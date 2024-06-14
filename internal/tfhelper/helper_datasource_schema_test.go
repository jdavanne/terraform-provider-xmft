package tfhelper

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToDataSourceSchema(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// Define the input data for the GenerateToSchema function
	var v TestResourceModel

	// Call the GenerateToSchema function
	generated := DataSourceModelToSchema(ctx, "testresource", &v)

	expected := schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{Required: true},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"str": schema.StringAttribute{
				Optional: true,
			},
			"str_default":        schema.StringAttribute{Optional: true},
			"str_co":             schema.StringAttribute{Optional: true, Computed: true},
			"str_no_read":        schema.StringAttribute{Optional: true},
			"str_no_write":       schema.StringAttribute{Optional: true},
			"str_optional":       schema.StringAttribute{Optional: true},
			"bool_false":         schema.BoolAttribute{Optional: true},
			"bool_true":          schema.BoolAttribute{Optional: true},
			"bool":               schema.BoolAttribute{Optional: true},
			"bool_default":       schema.BoolAttribute{Optional: true},
			"bool_default_true":  schema.BoolAttribute{Optional: true},
			"bool_default_false": schema.BoolAttribute{Optional: true},
			"int64":              schema.Int64Attribute{Optional: true},
			"int64_default":      schema.Int64Attribute{Optional: true},
			"int64_default_val":  schema.Int64Attribute{Optional: true},

			"str_list": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"sub_list": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id":      schema.StringAttribute{Required: true},
						"sub_str": schema.StringAttribute{Optional: true},
						"sub_str_co": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"sub_optional": schema.StringAttribute{Optional: true},
						"sub_sub_struct": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"sub_sub_str":  schema.StringAttribute{Optional: true},
								"sub_sub_bool": schema.BoolAttribute{Optional: true},
							},
							Optional: true,
						},
					},
				},
				Required:  false,
				Optional:  true,
				Computed:  false,
				Sensitive: false,
			},
			"sub_struct": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"id":           schema.StringAttribute{Required: true},
					"sub_str":      schema.StringAttribute{Optional: true},
					"sub_str_co":   schema.StringAttribute{Optional: true, Computed: true},
					"sub_optional": schema.StringAttribute{Optional: true},
					"sub_sub_struct": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"sub_sub_str":  schema.StringAttribute{Optional: true},
							"sub_sub_bool": schema.BoolAttribute{Optional: true},
						},
						Optional: true,
					},
				},
				Optional: true,
			},
			"sub_struct_p": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"id":           schema.StringAttribute{Required: true},
					"sub_str":      schema.StringAttribute{Optional: true},
					"sub_str_co":   schema.StringAttribute{Optional: true, Computed: true},
					"sub_optional": schema.StringAttribute{Optional: true},
					"sub_sub_struct": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"sub_sub_str":  schema.StringAttribute{Optional: true},
							"sub_sub_bool": schema.BoolAttribute{Optional: true},
						},
						Optional: true,
					},
				},
				Optional: true,
			},
			"sub_list_string": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"sub_list_object": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"sub_sub_str":  schema.StringAttribute{Optional: true},
						"sub_sub_bool": schema.BoolAttribute{Optional: true},
					},
				},
				Optional: true,
			},
			"sub_map_string": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"object": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"sub_sub_str":  schema.StringAttribute{Optional: true},
					"sub_sub_bool": schema.BoolAttribute{Optional: true},
				},
				Optional: true,
			},
			"poly_ptr": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"t1": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"sub_sub_str":  schema.StringAttribute{Optional: true},
							"sub_sub_bool": schema.BoolAttribute{Optional: true},
							"sub_sub_int":  schema.Int64Attribute{Optional: true},
						},
						Optional: true,
					},
					"t2": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"sub_sub_str":  schema.StringAttribute{Optional: true},
							"sub_sub_bool": schema.BoolAttribute{Optional: true},
						},
						Optional: true,
					},
					"t3": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"sub_sub_str":  schema.StringAttribute{Optional: true},
							"sub_sub_bool": schema.BoolAttribute{Optional: true},
						},
						Optional: true,
					},
				},
				Optional: true,
			},
			"list_poly_ptr": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"t1": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"sub_sub_str":  schema.StringAttribute{Optional: true},
								"sub_sub_bool": schema.BoolAttribute{Optional: true},
								"sub_sub_int":  schema.Int64Attribute{Optional: true},
							},
							Optional: true,
						},
						"t2": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"sub_sub_str":  schema.StringAttribute{Optional: true},
								"sub_sub_bool": schema.BoolAttribute{Optional: true},
							},
							Optional: true,
						},
						"t3": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"sub_sub_str":  schema.StringAttribute{Optional: true},
								"sub_sub_bool": schema.BoolAttribute{Optional: true},
							},
							Optional: true,
						},
					},
				},
				Optional: true,
			},
		},
	}
	// Check if the result matches the expected output
	assert.Equal(t, expected, generated)
}
