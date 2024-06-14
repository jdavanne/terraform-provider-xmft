package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
  "name": "string",
  "baseFolder": "string",
  "parent": "string",
  "businessUnitHierarchy": "string",
  "baseFolderModifyingAllowed": false,
  "homeFolderModifyingAllowed": false,
  "dmz": "string",
  "managedByCG": false,
  "enabledIcapServers": [
    "string"
  ],
  "additionalAttributes": {
    "additionalProp1": "string",
    "additionalProp2": "string",
    "additionalProp3": "string"
  },
  "bandwidthLimits": {
    "policy": "default",
    "inboundLimit": 0,
    "outboundLimit": 0,
    "modifyLimitsAllowed": false
  },
  "accountMaintenanceSettings": {
    "criteria": {
      "daysAfterCreation": 1,
      "daysOfInactivity": 1,
      "specificDate": "03.31/2004"
    },
    "action": {
      "action": "DELETE",
      "daysDeleteDisabled": 1
    },
    "emailNotificationBeforeAction": {
      "emailTemplate": "AccountMaintenanceNotification.xhtml",
      "notifyDays": "string",
      "notifyAccount": false,
      "notifyEmails": "string"
    },
    "emailNotificationForUserPassword": {
      "emailTemplate": "AccountMaintenanceNotification.xhtml",
      "notifyDays": "string",
      "notifyAccount": false,
      "notifyEmails": "string"
    },
    "emailNotificationForUserCertificate": {
      "emailTemplate": "AccountMaintenanceNotification.xhtml",
      "notifyDays": "string",
      "notifyAccount": false,
      "notifyEmails": "string"
    },
    "policy": "default",
    "policyModifyingAllowed": false
  },
  "htmlTemplateSettings": {
    "htmlTemplateFolderPath": "Default HTML Template",
    "isAllowedForModifying": false
  },
  "transfersApiSettings": {
    "isWebServiceRightsModifyingAllowed": false,
    "transfersWebServiceAllowed": false
  },
  "adHocSettings": {
    "authByEmail": false,
    "isAuthByEmailModifyingAllowed": false,
    "isDeliveryMethodModifyingAllowed": false,
    "deliveryMethod": "DEFAULT",
    "enrollmentTypes": [
      "string"
    ],
    "enrollmentTemplate": "Default",
    "implicitEnrollmentType": "ANONYMOUS_LINK",
    "notificationTemplate": "string"
  },
  "fileArchivingSettings": {
    "policy": "default",
    "policyModifyingAllowed": false,
    "folderPolicy": "default",
    "customFolder": "string",
    "encryptionCertificatePolicy": "string",
    "customEncryptionCertificate": "string",
    "customFileSizePolicy": "default",
    "customFileSize": 0
  },
  "loginRestrictionSettings": {
    "isPolicyModifyingAllowed": false,
    "policy": "string"
  },
  "addressBookSettings": {
    "policy": "default",
    "nonAddressBookCollaborationAllowed": false,
    "sources": [
      {
        "id": "string",
        "name": "string",
        "type": "LOCAL",
        "parentGroup": "string",
        "enabled": false,
        "customProperties": {
          "additionalProp1": "string",
          "additionalProp2": "string",
          "additionalProp3": "string"
        }
      }
    ],
    "modifySourcesAllowed": false,
    "modifyCollaborationAllowed": false
  },
  "fileMaintenanceSettings": {
    "policy": "default",
    "deleteFilesDays": 0,
    "pattern": "string",
    "expirationPeriod": false,
    "removeFolders": false,
    "warningNotifications": false,
    "sendSentinelAlert": false,
    "warnNotifyAccount": false,
    "warningNotificationsTemplate": "FileMaintenanceNotification.xhtml",
    "notifyDays": "string",
    "deletionNotifications": false,
    "deletionNotifyAccount": false,
    "deletionNotificationsTemplate": "FileMaintenanceNotification.xhtml",
    "allowPolicyModifying": false,
    "warnNotifyEmails": "string",
    "deletionNotifyEmails": "string"
  },
  "sharedFoldersCollaborationAllowed": false
}
*/

type stBusinessUnitResourceModel struct {
	// Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	BaseFolder                 types.String   `tfsdk:"base_folder" helper:"baseFolder,required"`
	Parent                     types.String   `tfsdk:"parent" helper:",emptyIsNull,default:"`
	BusinessUnitHierarchy      types.String   `tfsdk:"business_unit_hierarchy" helper:"businessUnitHierarchy,computed"`
	BaseFolderModifyingAllowed types.Bool     `tfsdk:"base_folder_modifying_allowed" helper:"baseFolderModifyingAllowed,default:false"`
	HomeFolderModifyingAllowed types.Bool     `tfsdk:"home_folder_modifying_allowed" helper:"homeFolderModifyingAllowed,default:false"`
	Dmz                        types.String   `tfsdk:"dmz" helper:",emptyIsNull,default:"`
	ManagedByCG                types.Bool     `tfsdk:"managed_by_cg" helper:"managedByCG,default:false"`
	EnabledIcapServers         []types.String `tfsdk:"enabled_icap_servers" helper:"enabledIcapServers,default:"`

	SharedFoldersCollaborationAllowed types.Bool `tfsdk:"shared_folders_collaboration_allowed" helper:"sharedFoldersCollaborationAllowed,optional"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTBusinessUnitResource() resource.Resource {
	return NewSTResource(&stBusinessUnitResourceModel{}, "st_business_unit", "", "/api/v2.0/businessUnits", "/api/v2.0/businessUnits/{name}")
}

func init() {
	registerResource(NewSTBusinessUnitResource)
}
