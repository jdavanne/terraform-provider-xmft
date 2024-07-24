package provider

import (
	"context"
	"fmt"

	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stVersionDataSource struct {
	providerData *stProviderData
}

/*
{
  "serverType" : "ST-Core-Server",
  "version" : "5.5-20240328",
  "build" : "${env.BUILD_NUMBER}",
  "os" : "Linux",
  "dockerMode" : true,
  "root" : false,
  "updateLevel" : null,
  "updateHistory" : [ ],
  "spiVersions" : [ {
    "name" : "AUTHENTICATION",
    "versions" : [ "1.0", "1.1", "1.2" ]
  }, {
    "name" : "ADVANCED_ROUTING",
    "versions" : [ "1.0", "1.1" ]
  }, {
    "name" : "TRANSFER_SITE",
    "versions" : [ "1.0", "1.1", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8" ]
  }, {
    "name" : "AUTHORIZATION",
    "versions" : [ "1.0", "1.1", "1.2", "1.3" ]
  } ],
  "other" : [ ]
*/

type stVersionDataSourceModel struct {
	ServerType types.String `tfsdk:"server_type" helper:"serverType,computed"`
	Version    types.String `tfsdk:"version" helper:",computed"`
	Build      types.String `tfsdk:"build" helper:",computed"`

	Os         types.String `tfsdk:"os" helper:",computed"`
	DockerMode types.Bool   `tfsdk:"docker_mode" helper:"dockerMode,computed"`
	Root       types.Bool   `tfsdk:"root" helper:",computed"`

	// UpdateLevel   types.String `tfsdk:"updateLevel"`
	// UpdateHistory types.List   `tfsdk:"updateHistory"`
	SpiVersions []struct {
		Name     types.String   `tfsdk:"name"  helper:",computed"`
		Versions []types.String `tfsdk:"version"  helper:",computed"`
	} `tfsdk:"spi_versions" helper:"spiVersion"`

	// Other []interface{}
}

var (
	_ datasource.DataSource              = &stVersionDataSource{}
	_ datasource.DataSourceWithConfigure = &stVersionDataSource{}
)

func NewSTVersionDataSource() datasource.DataSource {
	return &stVersionDataSource{}
}

func (d *stVersionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_st_version"
}

func (d *stVersionDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	var obj stVersionDataSourceModel
	resp.Schema = tfhelper.DataSourceModelToSchema(ctx, "version", "/version", &obj)
}

func (d *stVersionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state stVersionDataSourceModel
	stobject, err := d.providerData.client.ReadObject(ctx, "/api/v2.0/version")
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read xmft About",
			err.Error(),
		)
		return
	}

	tfhelper.AttributesToResource(ctx, "Version", stobject, &state)
	fmt.Printf("%+v\n", state)
	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *stVersionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*stProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *xmft.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.providerData = providerData
}

func init() {
	registerDataSource(NewSTVersionDataSource)
}
