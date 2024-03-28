package provider

import (
	"context"
	"fmt"

	"terraform-provider-axway-cft/internal/axwaycft"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type cftAboutDataSource struct {
	client *axwaycft.Client
}

type cftAboutDataSourceModel struct {
	About aboutModel `tfsdk:"about"`
}

type aboutModel struct {
	InstanceId       types.String `tfsdk:"instance_id"`
	Name             types.String `tfsdk:"name"`
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
	resp.TypeName = req.ProviderTypeName + "_about"
}

func (d *cftAboutDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"about": schema.SingleNestedAttribute{
				Computed: true,
				// NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"instance_id":       schema.StringAttribute{Computed: true},
					"name":              schema.StringAttribute{Computed: true},
					"version":           schema.StringAttribute{Computed: true},
					"level":             schema.StringAttribute{Computed: true},
					"multinode_enabled": schema.BoolAttribute{Computed: true},
					"system":            schema.StringAttribute{Computed: true},
				},
				//},
				//		},
				//

			},
		},
	}
}

func str(v interface{}) string {
	r, t := v.(string)
	if !t {
		return ""
	}
	return r
}

func (d *cftAboutDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state cftAboutDataSourceModel
	about, err := d.client.About(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read AxwayCFT About",
			err.Error(),
		)
		return
	}

	aboutState := aboutModel{
		InstanceId: types.StringValue(str(about["instance_id"])),
		Name:       types.StringValue(str(about["name"])),
		Version:    types.StringValue(str(about["version"])),
		Level:      types.StringValue(str(about["level"])),
		System:     types.StringValue(str(about["system"])),
	}

	state.About = aboutState

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

	client, ok := req.ProviderData.(*axwaycft.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *axwaycft.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}
