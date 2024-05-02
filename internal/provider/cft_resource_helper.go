package provider

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"terraform-provider-xmft/internal/cftapi"
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &cftResource{}
	_ resource.ResourceWithConfigure = &cftResource{}
)

func str(v interface{}) string {
	r, t := v.(string)
	if !t {
		return ""
	}
	return r
}

func NewCFTResource(obj interface{}, name string, kind string, uriCreate, uriReplace string) resource.Resource {
	if !strings.Contains(uriReplace, "{name}") {
		panic("no {name} in uri in :" + kind)
	}
	return &cftResource{name: name, kind: kind, uriCreate: uriCreate, uriReplace: uriReplace, obj: obj}
}

type cftResource struct {
	name       string
	kind       string
	uriCreate  string
	uriReplace string
	obj        interface{}
	client     *cftapi.Client
}

// Metadata returns the resource type name.
func (r *cftResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.name
}

// Schema defines the schema for the resource.
func (r *cftResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = tfhelper.ModelToSchema(ctx, r.name, r.obj)
}

func (r *cftResource) resolveURI(uri, name string) string {
	return strings.Replace(uri, "{name}", name, -1)
}

func (r *cftResource) createOrUpdate(ctx context.Context, pland *tfsdk.Plan, state *tfsdk.State, odiags *diag.Diagnostics) {
	// Retrieve values from plan
	plan := r.obj
	diags := pland.Get(ctx, plan)

	odiags.Append(diags...)
	if odiags.HasError() {
		return
	}

	// name := plan.Name.ValueString()
	m := make(map[string]interface{})
	tfhelper.ResourceToAttributes(ctx, r.name, plan, m)

	id := ""
	if reflect.ValueOf(state).Elem().FieldByName("Id") != (reflect.Value{}) {
		id = reflect.ValueOf(state).Elem().FieldByName("Id").Interface().(basetypes.StringValue).ValueString()
	}
	name := reflect.ValueOf(plan).Elem().FieldByName("Name").Interface().(basetypes.StringValue).ValueString()

	// FIXME: hack for cftpart (part1) : should be moved in client
	var attr2l []interface{}
	if r.name == "cftpart" {
		attr2l = m["tcp"].([]interface{})
		delete(m, "tcp")
	}
	// FIXME: End of hack (part1)

	alwaysReplace := true
	var err error
	var cftobject cftapi.CftObject
	if !alwaysReplace && id == "" {
		uri := r.resolveURI(r.uriCreate, name)
		cftobject, err = r.client.CreateObject(ctx, uri, r.kind, name, m)
		if err != nil {
			odiags.AddError(
				"Error creating "+r.name+" Id:"+name,
				"Could not create "+r.name+" Id:"+name+", unexpected error: "+err.Error(),
			)
			return
		}
	} else {
		uri := r.resolveURI(r.uriReplace, name)
		cftobject, err = r.client.ReplaceObject(ctx, uri, r.kind, name, m)
		if err != nil {
			odiags.AddError(
				"Error updating "+r.name+" Id:"+name,
				"Could not update "+r.name+" Id:"+name+", unexpected error: "+err.Error(),
			)
			return
		}
	}

	// FIXME: hack for cftpart (part2) : should be moved in client
	if r.name == "cftpart" {
		var err error
		// var cftobject xmft.CftObject
		for i := 0; i < len(attr2l); i++ {
			attr2 := attr2l[i].(map[string]interface{})
			id := attr2["id"].(string)
			delete(attr2, "id")
			if !alwaysReplace && id == "" {
				uri := r.resolveURI(r.uriReplace+"/tcp", name)
				_, err = r.client.CreateObject(ctx, uri, "cfttcp", "1", attr2)
			} else {
				uri := r.resolveURI(r.uriReplace+"/tcp/"+id, name)
				_, err = r.client.ReplaceObject(ctx, uri, "cfttcp", "1", attr2)
			}
			if err != nil {
				odiags.AddError(
					"Error creating tcp "+r.name+" Id:"+name,
					"Could not create tcp "+r.name+" Id:"+name+", unexpected error: "+err.Error(),
				)
				return
			}
		}

		uri := r.resolveURI(r.uriReplace, name)
		cftobject, err = r.client.ReadObject(ctx, uri)
		if err != nil {
			odiags.AddError(
				"Error Reading XCO CFT "+r.name+" Id:"+name,
				"Could not read XCO CFT "+r.name+" Id:"+name+": "+err.Error(),
			)
			return
		}
	}
	// FIXME: End of hack (part2)

	if cftobject.Id == "" {
		odiags.AddError(
			"Error creating XCO CFT "+r.name+" Id:"+name,
			"Could not find CFT "+r.name+" Id:"+name,
		)
		return
	}

	tfhelper.AttributesToResource(ctx, r.name, cftobject.Attributes, plan)
	reflect.ValueOf(plan).Elem().FieldByName("Id").Set(reflect.ValueOf(types.StringValue(str(cftobject.Id))))
	reflect.ValueOf(plan).Elem().FieldByName("LastUpdated").Set(reflect.ValueOf(types.StringValue(time.Now().Format(time.RFC850))))
	// plan.Id = types.StringValue(str(cftsend.Id))
	// plan.LastUpdated =

	// Set state to fully populated data
	diags = state.Set(ctx, plan)
	odiags.Append(diags...)
	if odiags.HasError() {
		return
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *cftResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.createOrUpdate(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

// Read refreshes the Terraform state with the latest data.
func (r *cftResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	state := r.obj
	diags := req.State.Get(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from HashiCups
	// name := state.Id.ValueString()
	name := reflect.ValueOf(state).Elem().FieldByName("Id").Interface().(basetypes.StringValue).ValueString()

	uri := r.resolveURI(r.uriReplace, name)
	cftobject, err := r.client.ReadObject(ctx, uri)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading XCO CFT "+r.name,
			"Could not read XCO CFT "+r.name+" ID "+name+": "+err.Error(),
		)
		return
	}

	// state.Id = types.StringValue(str(cftsend.Id))
	tfhelper.AttributesToResource(ctx, r.name, cftobject.Attributes, state)
	reflect.ValueOf(state).Elem().FieldByName("Id").Set(reflect.ValueOf(types.StringValue(str(cftobject.Id))))

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *cftResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.createOrUpdate(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *cftResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	state := r.obj
	diags := req.State.Get(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing order
	name := reflect.ValueOf(state).Elem().FieldByName("Id").Interface().(basetypes.StringValue).ValueString()

	if name == "" {
		resp.Diagnostics.AddError(
			"Error Reading XCO CFT "+r.name,
			"Could not find CFT "+r.name+" ID ",
		)
		return
	}
	uri := r.resolveURI(r.uriReplace, name)
	err := r.client.DeleteObject(ctx, uri)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting xmft "+r.name,
			"Could not delete "+r.name+" unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *cftResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cftapi.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *xmft.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}
