// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	"terraform-provider-axway-cft/internal/axwaycft"

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
	_ provider.Provider              = &AxwayCFTProvider{}
	_ provider.ProviderWithFunctions = &AxwayCFTProvider{}
)

// AxwayCFTProvider defines the provider implementation.
type AxwayCFTProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// AxwayCFTProviderModel describes the provider data model.
type AxwayCFTProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (p *AxwayCFTProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "axwaycft"
	resp.Version = p.version
}

func (p *AxwayCFTProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				MarkdownDescription: "Example provider attribute",
				Optional:            true,
			},
			"username": schema.StringAttribute{
				Optional: true,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *AxwayCFTProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring AxwayCFT client")
	var config AxwayCFTProviderModel

	diags := req.Config.Get(ctx, &config)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown AxwayCFT API Host",
			"The provider cannot create the AxwayCFT API client as there is an unknown configuration value for the AxwayCFT API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the AXWAY_CFT_HOST environment variable.",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown AxwayCFT API Username",
			"The provider cannot create the HashiCups API client as there is an unknown configuration value for the AxwayCFT API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the AXWAY_CFT_USERNAME environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown AxwayCFT API Password",
			"The provider cannot create the HashiCups API client as there is an unknown configuration value for the AxwayCFT API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the AXWAY_CFT_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("AXWAY_CFT_HOST")
	username := os.Getenv("AXWAY_CFT_USERNAME")
	password := os.Getenv("AXWAY_CFT_PASSWORD")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing AxwayCFT API Host",
			"The provider cannot create the AxwayCFT API client as there is a missing or empty value for the AxwayCFT API host. "+
				"Set the host value in the configuration or use the AXWAY_CFT_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing AxwayCFT API Username",
			"The provider cannot create the AxwayCFT API client as there is a missing or empty value for the AxwayCFT API username. "+
				"Set the username value in the configuration or use the AXWAY_CFT_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing AxwayCFT API Password",
			"The provider cannot create the AxwayCFT API client as there is a missing or empty value for the AxwayCFT API password. "+
				"Set the password value in the configuration or use the AXWAY_CFT_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "axwaycft_host", host)
	ctx = tflog.SetField(ctx, "axwaycft_username", username)
	ctx = tflog.SetField(ctx, "axwaycft_password", password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "axwaycft_password")

	tflog.Debug(ctx, "Creating AxwayCFT client")

	// Create a new HashiCups client using the configuration values
	client, err := axwaycft.NewClient(&host, &username, &password)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create AxwayCFT API Client",
			"An unexpected error occurred when creating the AxwayCFT API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"AxwayCFT Client Error: "+err.Error(),
		)
		return
	}

	// Example client configuration for data sources and resources
	resp.DataSourceData = client
	resp.ResourceData = client
	tflog.Info(ctx, "Configured AxwayCFT client", map[string]any{"success": true})
}

func (p *AxwayCFTProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewCFTSendResource,
	}
}

func (p *AxwayCFTProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAboutDataSource,
	}
}

func (p *AxwayCFTProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		// NewExampleFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AxwayCFTProvider{
			version: version,
		}
	}
}
