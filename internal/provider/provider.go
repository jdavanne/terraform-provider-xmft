// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"
	"slices"

	"terraform-provider-xmft/internal/cftapi"
	"terraform-provider-xmft/internal/stapi"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var (
	_ provider.Provider              = &xmftProvider{}
	_ provider.ProviderWithFunctions = &xmftProvider{}
)

// xmftProvider defines the provider implementation.
type xmftProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// xmftProviderModel describes the provider data model.
type xmftProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	Product  types.String `tfsdk:"product"`
}

func (p *xmftProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "xmft"
	resp.Version = p.version
}

func (p *xmftProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Example provider attribute",
				// Optional:            true,
			},
			"username": schema.StringAttribute{
				Optional: true,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"product": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *xmftProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring xmft client")
	var config xmftProviderModel

	diags := req.Config.Get(ctx, &config)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown xmft API Host",
			"The provider cannot create the xmft API client as there is an unknown configuration value for the xmft API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the XCO_MFT_HOST environment variable.",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown xmft API Username",
			"The provider cannot create the HashiCups API client as there is an unknown configuration value for the xmft API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the XCO_MFT_USERNAME environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown xmft API Password",
			"The provider cannot create the HashiCups API client as there is an unknown configuration value for the xmft API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the XCO_MFT_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("XMFT_HOST")
	username := os.Getenv("XMFT_USERNAME")
	password := os.Getenv("XMFT_PASSWORD")
	product := os.Getenv("XMFT_PRODUCT")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	if !config.Product.IsNull() {
		product = config.Product.ValueString()
	}
	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	list := []string{"cft", "st"}
	if !slices.Contains(list, product) {
		resp.Diagnostics.AddAttributeError(
			path.Root("type"),
			"Invalid XMFT Type",
			"The provider cannot create the XCO MFT API client as there is an invalid value for the XCO MFT type. "+
				"Set the type value in the configuration or use the XMFT_TYPE environment variable. "+
				"If either is already set, ensure the value is not empty and is one of the following: cft, st.",
		)
	}

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing xmft API Host",
			"The provider cannot create the xmft API client as there is a missing or empty value for the xmft API host. "+
				"Set the host value in the configuration or use the XCO_MFT_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing xmft API Username",
			"The provider cannot create the xmft API client as there is a missing or empty value for the xmft API username. "+
				"Set the username value in the configuration or use the XCO_MFT_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing xmft API Password",
			"The provider cannot create the xmft API client as there is a missing or empty value for the xmft API password. "+
				"Set the password value in the configuration or use the XMFT_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}
	ctx = tflog.SetField(ctx, "xmft_product", product)
	ctx = tflog.SetField(ctx, "xmft_host", host)
	ctx = tflog.SetField(ctx, "xmft_username", username)
	ctx = tflog.SetField(ctx, "xmft_password", password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "xmft_password")

	tflog.Debug(ctx, "Creating xmft client")

	if product == "cft" {
		// Create a new HashiCups client using the configuration values
		client, err := cftapi.NewClient(&host, &username, &password)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Create xmft API Client",
				"An unexpected error occurred when creating the xmft API client. "+
					"If the error is not clear, please contact the provider developers.\n\n"+
					"xmft Client Error: "+err.Error(),
			)
			return
		}
		resp.DataSourceData = client
		resp.ResourceData = client
		tflog.Info(ctx, "Configured XCO MFT CFT client", map[string]any{"success": true})
	} else if product == "st" {
		client, err := stapi.NewClient(&host, &username, &password)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Create XCO MFT CFT API Client",
				"An unexpected error occurred when creating the xmft API client. "+
					"If the error is not clear, please contact the provider developers.\n\n"+
					"xmft Client Error: "+err.Error(),
			)
			return
		}
		resp.DataSourceData = client
		resp.ResourceData = client
		tflog.Info(ctx, "Configured xmft client", map[string]any{"success": true})
	}
}

func (p *xmftProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewCFTSendResource,
		NewCFTRecvResource,
		NewCFTPartResource,
		NewCFTTcpResource,

		NewSTAccountResource,
		NewSTAdvancedRoutingApplicationResource,
		NewSTRouteTemplateResource,
		NewSTRouteCompositeResource,
		NewSTRouteSimpleResource,
		NewSTTransferSiteSSHModelResource,
		NewSTSubscriptionARModelResource,
	}
}

func (p *xmftProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAboutDataSource,

		NewSTVersionDataSource,
	}
}

func (p *xmftProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		// NewExampleFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &xmftProvider{
			version: version,
		}
	}
}
