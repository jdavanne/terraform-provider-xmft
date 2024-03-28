package provider

import (
	"context"
	"fmt"
	"time"

	"terraform-provider-axway-cft/internal/axwaycft"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &cftSendResource{}
	_ resource.ResourceWithConfigure = &cftSendResource{}
)

// NewCFTSendResource is a helper function to simplify the provider implementation.
func NewCFTSendResource() resource.Resource {
	return &cftSendResource{}
}

// cftSendResource is the resource implementation.
type cftSendResource struct {
	client *axwaycft.Client
}

type cftSendResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Fcode       types.String `tfsdk:"fcode"`
	Fname       types.String `tfsdk:"fname"`
	Faction     types.String `tfsdk:"faction"`
	Exec        types.String `tfsdk:"exec"`
	Parm        types.String `tfsdk:"parm"`
	Preexec     types.String `tfsdk:"preexec"`
}

// Metadata returns the resource type name.
func (r *cftSendResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cftsend"
}

// Schema defines the schema for the resource.
func (r *cftSendResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{Required: true},
			"last_updated": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"fcode":   schema.StringAttribute{Optional: true},
			"fname":   schema.StringAttribute{Optional: true},
			"faction": schema.StringAttribute{Optional: true},
			"exec":    schema.StringAttribute{Optional: true},
			"parm":    schema.StringAttribute{Optional: true},
			"preexec": schema.StringAttribute{Optional: true},
		},
	}
}

func (r *cftSendResource) createOrUpdate(ctx context.Context, pland *tfsdk.Plan, state *tfsdk.State, odiags *diag.Diagnostics) {
	// Retrieve values from plan
	var plan cftSendResourceModel
	diags := pland.Get(ctx, &plan)
	odiags.Append(diags...)
	if odiags.HasError() {
		return
	}

	name := plan.Name.ValueString()

	m := make(map[string]interface{})
	m["exec"] = plan.Exec.ValueString()
	m["fcode"] = plan.Fcode.ValueString()
	m["fname"] = plan.Fname.ValueString()
	m["faction"] = plan.Faction.ValueString()
	m["parm"] = plan.Parm.ValueString()
	m["preexec"] = plan.Preexec.ValueString()

	// Create new order
	cftsend, err := r.client.CreateObject(ctx, "/cft/api/v1/objects/cftsend/implno/"+name, "cftsendno", name, m)
	if err != nil {
		odiags.AddError(
			"Error creating CFTSEND",
			"Could not create CFTSEND, unexpected error: "+err.Error(),
		)
		return
	}

	if cftsend.Id == "" {
		odiags.AddError(
			"Error creating Axway CFT CFTSEND",
			"Could not find CFT CFTSEND ID",
		)
		return
	}

	plan.Id = types.StringValue(str(cftsend.Id))
	plan.Name = types.StringValue(name)
	plan.Exec = types.StringValue(str(cftsend.Attributes["exec"]))
	plan.Fcode = types.StringValue(str(cftsend.Attributes["fcode"]))
	plan.Fname = types.StringValue(str(cftsend.Attributes["fname"]))
	plan.Faction = types.StringValue(str(cftsend.Attributes["faction"]))
	plan.Parm = types.StringValue(str(cftsend.Attributes["parm"]))
	plan.Preexec = types.StringValue(str(cftsend.Attributes["preexec"]))

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = state.Set(ctx, plan)
	odiags.Append(diags...)
	if odiags.HasError() {
		return
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *cftSendResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.createOrUpdate(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

// Read refreshes the Terraform state with the latest data.
func (r *cftSendResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state cftSendResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from HashiCups
	name := state.Id.ValueString()
	if name == "" {
		resp.Diagnostics.AddError(
			"Error Reading Axway CFT CFTSEND",
			"Could not find CFT CFTSEND ID "+state.Id.ValueString(),
		)
		return
	}

	cftsend, err := r.client.ReadObject(ctx, "/cft/api/v1/objects/cftsend/implno/"+name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Axway CFT CFTSEND",
			"Could not read Axway CFT CFTSEND ID "+state.Id.ValueString()+": "+err.Error(),
		)
		return
	}

	state.Id = types.StringValue(str(cftsend.Id))
	state.Exec = types.StringValue(str(cftsend.Attributes["exec"]))
	state.Fcode = types.StringValue(str(cftsend.Attributes["fcode"]))
	state.Fname = types.StringValue(str(cftsend.Attributes["fname"]))
	state.Faction = types.StringValue(str(cftsend.Attributes["faction"]))
	state.Parm = types.StringValue(str(cftsend.Attributes["parm"]))
	state.Preexec = types.StringValue(str(cftsend.Attributes["preexec"]))

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *cftSendResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.createOrUpdate(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *cftSendResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state cftSendResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing order
	name := state.Name.ValueString()
	if name == "" {
		resp.Diagnostics.AddError(
			"Error Reading Axway CFT CFTSEND",
			"Could not find CFT CFTSEND ID "+state.Id.ValueString(),
		)
		return
	}

	err := r.client.DeleteObject(ctx, "/cft/api/v1/objects/cftsend/implno/"+name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting AxwayCFT CFTSEND",
			"Could not delete CFTSEND, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *cftSendResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}
