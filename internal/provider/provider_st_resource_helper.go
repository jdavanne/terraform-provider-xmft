package provider

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"
	"reflect"
	"time"

	"terraform-provider-xmft/internal/stapi"
	"terraform-provider-xmft/internal/tfhelper"
	"terraform-provider-xmft/internal/tools"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &stResource{}
	_ resource.ResourceWithConfigure = &stResource{}
)

type stProviderData struct {
	client               *stapi.Client
	additionalAttributes types.Map
}

type stResource struct {
	name         string
	kind         string
	uriCreate    string
	uriReplace   string
	obj          interface{}
	providerData *stProviderData
	// client       *stapi.Client

	ignoreDeleteNotFoundErrors bool
	onlyReplace                bool
	alwaysRecreate             bool

	swaggerDiscriminator string // for swagger comparison
	swaggerUri           string // for swagger comparison
}

func NewSTResource(obj interface{}, name string, kind string, uriCreate, uriReplace string) *stResource {
	/*if !strings.Contains(uriReplace, "{name}") {
		panic("no {name} in uri in :" + kind)
	}*/
	return &stResource{name: name, kind: kind, uriCreate: uriCreate, uriReplace: uriReplace, obj: obj}
}

func GetAttribute(state interface{}, attrname string) string {
	field := reflect.ValueOf(state).Elem().FieldByName(attrname)
	if field.IsValid() {
		value := field.Interface()
		switch value := value.(type) {
		case basetypes.StringValue:
			return value.ValueString()
		default:
			panic("unsupported type for: " + attrname + ":" + reflect.TypeOf(value).String())
		}
	}
	return ""
}

func GetRef(state interface{}) string {
	_name := GetAttribute(state, "Name")
	_id := GetAttribute(state, "Id")

	if _name == "" && _id == "" {
		return ""
	} else if _name == "" {
		return _id
	} else if _id == "" {
		return _name
	}
	return _name + "/" + _id
}

// Metadata returns the resource type name.
func (r *stResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.name
}

func (r *stResource) IgnoreDeleteNotFoundError() *stResource {
	r.ignoreDeleteNotFoundErrors = true
	return r
}

func (r *stResource) OnlyReplace() *stResource {
	r.onlyReplace = true
	return r
}

func (r *stResource) AlwaysRecreate() *stResource {
	r.alwaysRecreate = true
	return r
}

func (r *stResource) AddDiscriminator(disc string) *stResource {
	r.swaggerDiscriminator += disc
	return r
}

func (r *stResource) UseSwaggerUri(uri string) *stResource {
	r.swaggerUri = uri
	return r
}

//go:embed provider_st_field_descriptions.json
var providerStFieldDescriptions []byte
var providerStFieldDescriptionsMap map[string]string

func init() {
	err := json.Unmarshal(providerStFieldDescriptions, &providerStFieldDescriptionsMap)
	if err != nil {
		panic(err)
	}
}

// Schema defines the schema for the resource.
func (r *stResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	uri := r.uriReplace
	if r.swaggerUri != "" {
		uri = r.swaggerUri
	}
	ctx2 := context.WithValue(ctx, "fieldDescription", providerStFieldDescriptionsMap)
	resp.Schema = tfhelper.ModelToSchema(ctx2, r.name, uri[9:]+r.swaggerDiscriminator, r.obj)
}

func (r *stResource) createOrUpdate(ctx context.Context, pland *tfsdk.Plan, reqState *tfsdk.State, respState *tfsdk.State, odiags *diag.Diagnostics, create bool) {
	uri := ""
	// Retrieve values from plan
	plan := reflect.New(reflect.TypeOf(r.obj).Elem()).Interface()
	// plan := r.obj
	prev := make(map[string]interface{})
	if reqState != nil {
		diags := reqState.Get(ctx, plan)
		odiags.Append(diags...)
		if odiags.HasError() {
			return
		}
		slog.DebugContext(ctx, "CreateOrUpdate XCO ST "+r.name+" requestState", "data", fmt.Sprintf("%+v", plan))

		// Useful to support renames
		tfhelper.ResourceToAttributes(ctx, r.name, plan, prev)
		uri = tfhelper.ResolveURI(r.uriReplace, prev)
	}
	if reqState != nil && r.alwaysRecreate {
		ref := GetRef(plan)

		err := r.providerData.client.DeleteObject(ctx, uri)
		if err != nil {
			if e, ok := err.(*tools.HttpError); !(r.ignoreDeleteNotFoundErrors && ok && e.StatusCode == 404) {
				odiags.AddError(
					"Error Deleting XCO ST "+r.name+" ref="+ref+fmt.Sprint(" ignore=", r.ignoreDeleteNotFoundErrors, " ok=", ok, " statusCode=", e.StatusCode),
					"Could not delete "+r.name+"ref="+ref+" - unexpected error: "+err.Error(),
				)
				return
			}
		}
		create = true
	}

	diags := pland.Get(ctx, plan)
	odiags.Append(diags...)
	if odiags.HasError() {
		return
	}
	slog.DebugContext(ctx, "CreateOrUpdate XCO ST "+r.name+" planState", "data", fmt.Sprintf("%+v", plan))

	// if provider has AddtionAttributes, merge them into plan
	// name := plan.Name.ValueString()
	m := make(map[string]interface{})
	tfhelper.ResourceToAttributes(ctx, r.name, plan, m)

	var err error
	var stobject stapi.StObject
	ref := GetRef(plan)

	if create {
		uri := tfhelper.ResolveURI(r.uriCreate, m)
		stobject, err = r.providerData.client.CreateObject(ctx, uri, r.kind, m)
		if err != nil {
			odiags.AddError(
				"Error creating "+r.name+" Ref:"+ref,
				"Could not create "+r.name+" Ref:"+ref+", unexpected error: "+err.Error(),
			)
			return
		}
	} else {
		if ref == "" {
			odiags.AddError(
				"Error updating "+r.name,
				"Could not find "+r.name+" Name ",
			)
			return
		}

		// Useful to support renames
		uri2 := tfhelper.ResolveURI(r.uriReplace, m)
		if uri == "" {
			uri = uri2 // except when create is actually a PUT
		}
		stobject, err = r.providerData.client.ReplaceObject(ctx, uri, uri2, r.kind, m)
		if err != nil {
			odiags.AddError(
				"Error updating "+r.name+" Ref:"+ref,
				"Could not update "+r.name+" Ref:"+ref+", unexpected error: "+err.Error(),
			)
			return
		}
	}

	tfhelper.AttributesToResource(ctx, r.name, stobject, plan)
	// reflect.ValueOf(plan).Elem().FieldByName("Id").Set(reflect.ValueOf(types.StringValue(str(stobject["id"]))))
	reflect.ValueOf(plan).Elem().FieldByName("LastUpdated").Set(reflect.ValueOf(types.StringValue(time.Now().Format(time.RFC850))))

	slog.DebugContext(ctx, "CreateOrUpdate XCO ST "+r.name+" Ref="+ref, "data", fmt.Sprintf("%+v", plan), "stobject", fmt.Sprintf("%+v", stobject))

	// Set state to fully populated data
	diags = respState.Set(ctx, plan)
	odiags.Append(diags...)
	if odiags.HasError() {
		return
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *stResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.createOrUpdate(ctx, &req.Plan, nil, &resp.State, &resp.Diagnostics, !r.onlyReplace)
}

// Read refreshes the Terraform state with the latest data.
func (r *stResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	state := reflect.New(reflect.TypeOf(r.obj).Elem()).Interface()
	diags := req.State.Get(ctx, state)

	slog.DebugContext(ctx, "Read XCO ST "+r.name+" requestState", "data", fmt.Sprintf("%+v", state))

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from HashiCups
	ref := GetRef(state)
	if ref == "" {
		resp.Diagnostics.AddError(
			"Error Reading XCO ST "+r.name,
			"Could not find XCO ST "+r.name+" ref="+ref,
		)
		return
	}

	m := make(map[string]interface{})
	tfhelper.ResourceToAttributes(ctx, r.name, state, m)
	uri := tfhelper.ResolveURI(r.uriReplace, m)
	stobject, err := r.providerData.client.ReadObject(ctx, uri)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading XCO ST "+r.name,
			"Could not read XCO ST "+r.name+" ref="+ref+" : "+err.Error(),
		)
		return
	}

	state2 := state
	tfhelper.AttributesToResource(ctx, r.name, stobject, state2)

	slog.DebugContext(ctx, "Read XCO ST "+r.name+" ref="+ref, "data", fmt.Sprintf("%+v", state2))
	// Set refreshed state
	diags = resp.State.Set(ctx, state2)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *stResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.createOrUpdate(ctx, &req.Plan, &req.State, &resp.State, &resp.Diagnostics, false)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *stResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	if r.onlyReplace {
		// no Delete
		return
	}
	// Retrieve values from state
	state := reflect.New(reflect.TypeOf(r.obj).Elem()).Interface()
	diags := req.State.Get(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing order
	ref := GetRef(state)

	if ref == "" {
		resp.Diagnostics.AddError(
			"Error Deleting XCO ST "+r.name,
			"Could not find XCO "+r.name+" Reference ",
		)
		return
	}

	m := make(map[string]interface{})
	tfhelper.ResourceToAttributes(ctx, r.name, state, m)
	uri := tfhelper.ResolveURI(r.uriReplace, m)
	err := r.providerData.client.DeleteObject(ctx, uri)
	if err != nil {
		if e, ok := err.(*tools.HttpError); !(r.ignoreDeleteNotFoundErrors && ok && e.StatusCode == 404) {
			resp.Diagnostics.AddError(
				"Error Deleting XCO ST "+r.name+" ref="+ref+fmt.Sprint(" ignore=", r.ignoreDeleteNotFoundErrors, " ok=", ok, " statusCode=", e.StatusCode),
				"Could not delete "+r.name+"ref="+ref+" - unexpected error: "+err.Error(),
			)
			return
		}
	}
}

func (r *stResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.providerData = providerData
}
