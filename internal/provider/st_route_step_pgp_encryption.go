package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
 {
	"type": "PgpEncryption",
	"status": "ENABLED",
	"precedingStep": "string",
	"conditionType": "ALWAYS",
	"condition": "string",
	"actionOnStepSuccess": "PROCEED",
	"actionOnStepFailure": "FAIL",
	"autostart": false,
	"fileFilterExpression": "string",
	"fileFilterExpressionType": "GLOB",
	"usePrecedingStepFiles": false,
	"encryptKeyOwnerExpression": "string",
	"encryptKeyExpression": "string",
	"encryptKeyOwnerExpressionType": "string",
	"encryptKeyExpressionType": "string",
	"signKeyOwnerExpression": "string",
	"signKeyExpression": "string",
	"signKeyOwnerExpressionType": "string",
	"signKeyExpressionType": "string",
	"useAsciiArmour": false,
	"compressionType": "string",
	"compressionLevel": "string",
	"postTransformationActionRenameAsExpression": "string"
},
*/

type stRouteStepPgpEncryptionResourceModel struct {
	Id     types.String `tfsdk:"id" helper:",computed,state"`
	Type   types.String `tfsdk:"type" helper:",default:PgpEncryption"`
	Status types.String `tfsdk:"status" helper:",default:ENABLED"`

	// PrecedingStep       types.String `tfsdk:"preceding_step" helper:"precedingStep,optional"`
	ConditionType       types.String `tfsdk:"condition_type" helper:"conditionType,enum:ALWAYS/EL,default:ALWAYS"`
	Condition           types.String `tfsdk:"condition" helper:",emptyIsNull,default:"`
	ActionOnStepSuccess types.String `tfsdk:"action_on_step_success" helper:"actionOnStepSuccess,enum:/PROCEED/STOP,default:PROCEED"`
	ActionOnStepFailure types.String `tfsdk:"action_on_step_failure" helper:"actionOnStepFailure,enum:/FAIL/PROCEED,default:FAIL"`
	// Autostart                types.Bool   `tfsdk:"autostart" helper:",default:false"`
	FileFilterExpression     types.String `tfsdk:"file_filter_expression" helper:"fileFilterExpression,default:*"`
	FileFilterExpressionType types.String `tfsdk:"file_filter_expression_type" helper:"fileFilterExpressionType,enum:/GLOB/REGEXP/TEXT_FILES,default:GLOB"`
	// UsePrecedingStepFiles                      types.Bool   `tfsdk:"use_preceding_step_files" helper:"usePrecedingStepFiles,default:false"`
	EncryptKeyOwnerExpression                  types.String `tfsdk:"encrypt_key_owner_expression" helper:"encryptKeyOwnerExpression,emptyIsNull,default:"`
	EncryptKeyOwnerExpressionType              types.String `tfsdk:"encrypt_key_owner_expression_type" helper:"encryptKeyOwnerExpressionType,emptyIsNull,enum:/NAME/EXPRESSION,default:"`
	EncryptKeyExpression                       types.String `tfsdk:"encrypt_key_expression" helper:"encryptKeyExpression,emptyIsNull,default:"`
	EncryptKeyExpressionType                   types.String `tfsdk:"encrypt_key_expression_type" helper:"encryptKeyExpressionType,enum:/ALIAS/EXPRESSION_WILDCARD,default:ALIAS"`
	SignKeyOwnerExpression                     types.String `tfsdk:"sign_key_owner_expression" helper:"signKeyOwnerExpression,emptyIsNull,default:"`
	SignKeyOwnerExpressionType                 types.String `tfsdk:"sign_key_owner_expression_type" helper:"signKeyOwnerExpressionType,emptyIsNull,enum:/NAME/EXPRESSION,default:"`
	SignKeyExpression                          types.String `tfsdk:"sign_key_expression" helper:"signKeyExpression,emptyIsNull,default:"`
	SignKeyExpressionType                      types.String `tfsdk:"sign_key_expression_type" helper:"signKeyExpressionType,emptyIsNull,enum:/ALIAS/EXPRESSION_WILDCARD,default:"`
	UseAsciiArmour                             types.Bool   `tfsdk:"use_ascii_armour" helper:"useAsciiArmour,default:false"`
	CompressionType                            types.String `tfsdk:"compression_type" helper:"compressionType,default:0"`
	CompressionLevel                           types.String `tfsdk:"compression_level" helper:"compressionLevel,default:2"`
	PostTransformationActionRenameAsExpression types.String `tfsdk:"post_transformation_action_rename_as_expression" helper:"postTransformationActionRenameAsExpression,emptyIsNull,default:"`
}

func init() {
	tfhelper.RegisterType("stRouteStepPgpEncryptionResourceModel", &stRouteStepPgpEncryptionResourceModel{})
}
