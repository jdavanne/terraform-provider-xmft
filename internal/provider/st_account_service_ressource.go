package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stAccountServiceResourceModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                  types.String `tfsdk:"type" helper:",default:service"`
	Uid                   types.String `tfsdk:"uid" helper:",default:10000"`
	Gid                   types.String `tfsdk:"gid" helper:",default:10000"`
	HomeFolder            types.String `tfsdk:"home_folder" helper:"homeFolder,required"`
	HomeFolderAccessLevel types.String `tfsdk:"home_folder_access_level" helper:"homeFolderAccessLevel,enum:/PRIVATE/PUBLIC/BUSINESS_UNIT,default:PRIVATE"`
	LastModified          types.String `tfsdk:"last_modified" helper:"lastModified,computed,nowrite"`

	User struct {
		Name types.String `tfsdk:"name" helper:",required"`

		AuthExternal             types.Bool   `tfsdk:"auth_external" helper:"authExternal,computed,nowrite"`
		FailedAuthAttempts       types.Int64  `tfsdk:"failed_auth_attempts" helper:"failedAuthAttempts,computed,nowrite"`
		FailedAuthMaximum        types.Int64  `tfsdk:"failed_auth_maximum" helper:"failedAuthMaximum,default:3"`
		FailedSshKeyAuthAttempts types.Int64  `tfsdk:"failed_ssh_key_auth_attempts" helper:"failedSshKeyAuthAttempts,computed,nowrite"`
		FailedSshKeyAuthMaximum  types.Int64  `tfsdk:"failed_ssh_key_auth_maximum" helper:"failedSshKeyAuthMaximum,default:3"`
		LastFailedAuth           types.String `tfsdk:"last_failed_auth" helper:"lastFailedAuth,computed,nowrite"`
		LastFailedSshKeyAuth     types.String `tfsdk:"last_failed_ssh_key_auth" helper:"lastFailedSshKeyAuth,computed,nowrite"`
		LastLogin                types.String `tfsdk:"last_login" helper:"lastLogin,computed,nowrite"`
		Locked                   types.Bool   `tfsdk:"locked" helper:"locked,computed,nowrite"`
		SuccessfulAuthMaximum    types.Int64  `tfsdk:"successful_auth_maximum" helper:"successfulAuthMaximum,emptyIsNull,optional"`
		SuccessfulLogins         types.Int64  `tfsdk:"successful_logins" helper:"successfulLogins,computed,nowrite"`
		SecretQuestion           struct {
			ForceSecretQuestionChange types.Bool   `tfsdk:"force_secret_question_change" helper:"forceSecretQuestionChange,computed,nowrite"`
			SecretAnswerGuessFailures types.Int64  `tfsdk:"secret_answer_guess_failures" helper:"secretAnswerGuessFailures,computed,nowrite"`
			SecretQuestion            types.String `tfsdk:"secret_question" helper:"secretQuestion,computed,nowrite"`
		} `tfsdk:"secret_question" helper:"secretQuestion,default:"`
		PasswordCredentials struct {
			ForcePasswordChange    types.Bool   `tfsdk:"force_password_change" helper:"forcePasswordChange,computed,optional"`
			Password               types.String `tfsdk:"password" helper:",noread"`
			PasswordDigest         types.String `tfsdk:"password_digest" helper:"passwordDigest,computed,nowrite,optional"`
			PasswordExpiryInterval types.Int64  `tfsdk:"password_expiry_interval" helper:"passwordExpiryInterval,emptyIsNull,optional"`
			LastOwnPasswordChange  types.String `tfsdk:"last_own_password_change" helper:"lastOwnPasswordChange,computed,nowrite"`
			LastPasswordChange     types.String `tfsdk:"last_password_change" helper:"lastPasswordChange,computed,nowrite"`
		} `tfsdk:"password_credentials" helper:"passwordCredentials,state,required"`
	} `tfsdk:"user" helper:",required"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	Disabled types.Bool `tfsdk:"disabled" helper:"disabled,default:false"`

	FileArchivingPolicy types.String `tfsdk:"file_archiving_policy" helper:"fileArchivingPolicy,enum:/DEFAULT/DISABLED/ENABLED,default:DEFAULT"`

	AuthByEmail                 types.Bool   `tfsdk:"auth_by_email" helper:"authByEmail,default:false"`
	AccountCreationDate         types.Int64  `tfsdk:"account_creation_date" helper:"accountCreationDate,computed,nowrite"`
	AccountEncryptMode          types.String `tfsdk:"account_encrypt_mode" helper:"accountEncryptMode,enum:/UNSPECIFIED/ENABLED,default:UNSPECIFIED"`
	AccountSubmitForApprove     types.Bool   `tfsdk:"account_submit_for_approve" helper:"accountSubmitForApprove,computed,nowrite"`
	AccountVerificationStatus   types.String `tfsdk:"account_verification_status" helper:"accountVerificationStatus,computed,nowrite"`
	BusinessUnit                types.String `tfsdk:"business_unit" helper:"businessUnit,emptyIsNull,default:"`
	LoginRestrictionPolicy      types.String `tfsdk:"login_restriction_policy" helper:"loginRestrictionPolicy,default:"`
	ManagedByCG                 types.Bool   `tfsdk:"managed_by_cg" helper:"managedByCG,default:"`
	MappedUser                  types.String `tfsdk:"mapped_user" helper:"mappedUser,emptyIsNull,default:"`
	Notes                       types.String `tfsdk:"notes" helper:"notes,emptyIsNull,default:"`
	PesitId                     types.String `tfsdk:"pesit_id" helper:"pesitId,emptyIsNull,default:"`
	RejectReason                types.String `tfsdk:"reject_reason" helper:"rejectReason,computed,nowrite"`
	RoutingMode                 types.String `tfsdk:"routing_mode" helper:"routingMode,enum:/accept/reject/ignore,default:reject"`
	Skin                        types.String `tfsdk:"skin" helper:"skin,default:Default HTML Template"`
	SubscriptionFolderDiscovery types.String `tfsdk:"subscription_folder_discovery" helper:"subscriptionFolderDiscovery,enum:/ITERATIVE/RECURSIVE,default:ITERATIVE"`
	TransferType                types.String `tfsdk:"transfer_type" helper:"transferType,enum:/false/I/E,default:false"`
	TransfersWebServiceAllowed  types.Bool   `tfsdk:"transfers_web_service_allowed" helper:"transfersWebServiceAllowed,default:false"`

	AdhocSettings struct {
		DeliveryMethod         types.String   `tfsdk:"delivery_method" helper:"deliveryMethod,enum:/DEFAULT/DISABLED/ANONYMOUS/ACCOUNT_WITHOUT_ENROLLMENT/ACCOUNT_WITH_ENROLLMENT/CUSTOM,emptyIsNull,default:DEFAULT"`
		EnrollmentTypes        []types.String `tfsdk:"enrollment_types" helper:"enrollmentTypes,elementtype:string,emptyIsNull,default:"`
		ImplicitEnrollmentType types.String   `tfsdk:"implicit_enrollment_type" helper:"implicitEnrollmentType,enum:/ANONYMOUS_LINK/CHALLENGED_LINK/EXISTING_ACCOUNT/ENROLL_UNLICENSED/ENROLL_LICENSED,emptyIsNull,default:"`
	} `tfsdk:"adhoc_settings" helper:"adhocSettings,default:"`

	Contact struct {
		Email types.String `tfsdk:"email" helper:"email,emptyIsNull,default:"`
		Phone types.String `tfsdk:"phone" helper:"phone,emptyIsNull,default:"`
	} `tfsdk:"contact" helper:"contact,default:"`
}

func NewSTAccountServiceResource() resource.Resource {
	return NewSTResource(&stAccountServiceResourceModel{}, "st_account_service", "", "/api/v2.0/accounts", "/api/v2.0/accounts/{name}").AddDiscriminator("[type=service]")
}

func init() {
	registerResource(NewSTAccountServiceResource)
}
