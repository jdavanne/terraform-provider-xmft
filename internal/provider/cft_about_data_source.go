package provider

import (
	"context"
	"fmt"

	"terraform-provider-xmft/internal/cftapi"
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type cftAboutDataSource struct {
	client *cftapi.Client
}

type cftAboutDataSourceModel struct {
	InstanceId       types.String `tfsdk:"instance_id"`
	Version          types.String `tfsdk:"version"`
	Level            types.String `tfsdk:"level"`
	MultinodeEnabled types.Bool   `tfsdk:"multinode_enabled"`
	System           types.String `tfsdk:"system"`
}

var (
	_ datasource.DataSource              = &cftAboutDataSource{}
	_ datasource.DataSourceWithConfigure = &cftAboutDataSource{}
)

func NewAboutDataSource() datasource.DataSource {
	return &cftAboutDataSource{}
}

func (d *cftAboutDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cft_about"
}

func (d *cftAboutDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	var obj cftAboutDataSourceModel
	resp.Schema = tfhelper.DataSourceModelToSchema(ctx, "version", &obj)
}

func (d *cftAboutDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state cftAboutDataSourceModel
	about, err := d.client.About(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read XCO CFT About",
			err.Error(),
		)
		return
	}

	tfhelper.AttributesToResource(ctx, "About", about, &state)

	/*aboutState := aboutModel{
		InstanceId: types.StringValue(str(about["instance_id"])),
		Name:       types.StringValue(str(about["name"])),
		Version:    types.StringValue(str(about["version"])),
		Level:      types.StringValue(str(about["level"])),
		System:     types.StringValue(str(about["system"])),
	}

	state.About = aboutState*/

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *cftAboutDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cftapi.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *cftapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}
