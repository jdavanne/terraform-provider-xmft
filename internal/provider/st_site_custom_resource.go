package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
  "name": "string",
  "type": "custom",
  "protocol": "string",
  "transferType": "internal",
  "maxConcurrentConnection": 0,
  "default": false,
  "accessLevel": "PRIVATE",
  "account": "string",
  "additionalAttributes": {
    "additionalProp1": "string",
    "additionalProp2": "string",
    "additionalProp3": "string"
  },
  "clientCertificate": "string",
  "customProperties": {
    "property_1": "string",
    "property_2": "string"
  }
}
  {
      "type": "ExternalPersistedCustomSite",
      "id": "8a00b9468f148406018f14a695740003",
      "name": "S3_bucket",
      "account": "alice",
      "protocol": "s3",
      "transferType": "unspecified",
      "maxConcurrentConnection": 0,
      "default": false,
      "accessLevel": "PRIVATE",
      "clientCertificate": null,
      "customProperties": {
        "s3PatternType": "glob",
        "s3RequestTimeout": "25",
        "s3ReceivePtaFailureMoveRenameFileExpr": "",
        "s3Region": "eu-west-3",
        "s3ReceiveFileAsExpr": "",
        "s3FipsModeEnabled": "false",
        "s3ConnectionSecureEnabled": "true",
        "s3FailoverRegion": "us-standard",
        "s3UploadDestination": "",
        "s3ClientPoolMaxTotalInstances": "50",
        "s3ReceivePtaSuccess": "noAction",
        "s3SendFileAsExpr": "",
        "s3SslCipherSuites": "TLS_AES_256_GCM_SHA384, TLS_AES_128_GCM_SHA256, TLS_CHACHA20_POLY1305_SHA256, TLS_AES_128_CCM_SHA256, TLS_AES_128_CCM_8_SHA256, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384, TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256, TLS_DHE_RSA_WITH_AES_256_GCM_SHA384, TLS_DHE_DSS_WITH_AES_256_GCM_SHA384, TLS_DHE_DSS_WITH_AES_256_CBC_SHA256, TLS_DHE_RSA_WITH_AES_256_CBC_SHA256, TLS_DHE_RSA_WITH_AES_128_GCM_SHA256, TLS_DHE_DSS_WITH_AES_128_GCM_SHA256, TLS_DHE_DSS_WITH_AES_128_CBC_SHA256, TLS_DHE_RSA_WITH_AES_128_CBC_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256",
        "s3SystemMetadata": "",
        "s3AllowOverwrite": "false",
        "s3FailoverNetworkZone": "none",
        "s3JcloudsLoggingEnabled": "true",
        "s3ClientPoolEnabled": "true",
        "s3AutoCreateBucketCheck": "true",
        "s3FailoverSettingsEnabled": "false",
        "s3VerifyCertificateEnabled": "true",
        "s3AccessKey": "zouzou",
        "s3DownloadObjectKey": "/s3_download",
        "s3UploadMode": "auto",
        "s3DownloadRecursively": "true",
        "s3SslProtocols": "TLSv1.2, TLSv1.3",
        "s3ClientPoolMinEvictableIdleTime": "3600",
        "s3Bucket": "storage1-s3bucket",
        "s3ReceivePtaFailure": "noAction",
        "s3FlatDirectoryStructure": "false",
        "s3SecretKey": "zouzou",
        "s3UserMetadata": "",
        "s3ReceivePtaSuccessMoveRenameFileExpr": "",
        "s3NetworkZone": "none",
        "s3UserDownloadMetadata": "",
        "s3SoTimeout": "25",
        "s3DownloadPattern": "",
        "s3PrefixTimeout": "25",
        "s3VirtualHostedStyleEnabled": "true",
        "s3ConnectionTimeout": "25"
      },

 {
      "type": "ExternalPersistedCustomSite",
      "id": "8a00b9468f4d5acf018f9b54d4840003",
      "name": "AzureBlob",
      "account": "alice",
      "protocol": "azure-blob",
      "transferType": "unspecified",
      "maxConcurrentConnection": 0,
      "default": false,
      "accessLevel": "PUBLIC",
      "clientCertificate": null,
      "customProperties": {
        "azureblobReceiveTimeout": "25",
        "azureblobShowAdvancedSettingsEnabled": "false",
        "azureblobSendTimeout": "25",
        "azureblobUploadFolder": "folder_${date('yyyyMMddHHmmss')}",
        "azureblobAccountName": "hello",
        "azureblobSslProtocols": "TLSv1.2,TLSv1.3",
        "azureblobBlobType": "BLOCK_BLOB",
        "azureblobPtaReceiveOnSuccess": "noaction",
        "azureblobDownloadRecursively": "true",
        "azureblobPreserveFolderStructure": "true",
        "azureblobBlockSize": "4",
        "azureblobUseHttpsEnabled": "true",
        "azureblobDownloadMetadata": "",
        "azureblobPtaSendOnSuccessOverwriteExisting": "true",
        "azureblobAllowOverwriteExisting": "true",
        "azureblobPtaSendOnSuccessDeleteSnapshotsEnabled": "false",
        "azureblobStorageType": "AUTO",
        "azureblobAccountKey": "zouzou",
        "azureblobDownloadPattern": "",
        "azureblobSpSecretAuth": "false",
        "azureblobDownloadFolder": "downloadfolder_${date(\"dd.MM.yyyy\")}",
        "azureblobBlobContainer": "stcontainer",
        "azureblobMaxConcurrency": "8",
        "azureblobPtaReceiveOnSuccessOverwriteExisting": "true",
        "azureblobMaxSingleUploadSize": "4",
        "azureblobUploadContentType": "",
        "azureblobNetworkZone": "none",
        "azureblobAllowOverwriteUploadFolder": "false",
        "azureblobEndpointSuffix": "core.windows.net",
        "azureblobVerifyCertificateEnabled": "true",
        "azureblobConnectTimeout": "25",
        "azureblobConnectionType": "cs",
        "azureblobPtaSendFileAs": "",
        "azureblobPtaReceiveOnSuccessDeleteSnapshotsEnabled": "false",
        "azureblobPtaSendOnSuccess": "noaction",
        "azureblobDownloadPatternType": "glob",
        "azureblobBlobContainerCreateEnabled": "true",
        "azureblobUploadMetadata": "",
        "azureblobPageBlobAddZeroPadEnabled": "false",
        "azureblobSslCipherSuites": "TLS_AES_128_CCM_8_SHA256,TLS_AES_128_CCM_SHA256,TLS_AES_128_GCM_SHA256,TLS_AES_256_GCM_SHA384,TLS_DHE_DSS_WITH_AES_128_CBC_SHA,TLS_DHE_DSS_WITH_AES_128_CBC_SHA256,TLS_DHE_DSS_WITH_AES_128_GCM_SHA256,TLS_DHE_DSS_WITH_AES_256_CBC_SHA,TLS_DHE_DSS_WITH_AES_256_CBC_SHA256,TLS_DHE_DSS_WITH_AES_256_GCM_SHA384,TLS_DHE_RSA_WITH_AES_128_CBC_SHA,TLS_DHE_RSA_WITH_AES_128_CBC_SHA256,TLS_DHE_RSA_WITH_AES_128_CCM,TLS_DHE_RSA_WITH_AES_128_CCM_8,TLS_DHE_RSA_WITH_AES_128_GCM_SHA256,TLS_DHE_RSA_WITH_AES_256_CBC_SHA,TLS_DHE_RSA_WITH_AES_256_CBC_SHA256,TLS_DHE_RSA_WITH_AES_256_CCM,TLS_DHE_RSA_WITH_AES_256_CCM_8,TLS_DHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_ECDSA_WITH_AES_128_CCM,TLS_ECDHE_ECDSA_WITH_AES_128_CCM_8,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_ECDSA_WITH_AES_256_CCM,TLS_ECDHE_ECDSA_WITH_AES_256_CCM_8,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_AES_128_CBC_SHA,TLS_RSA_WITH_AES_128_CBC_SHA256,TLS_RSA_WITH_AES_128_CCM,TLS_RSA_WITH_AES_128_CCM_8,TLS_RSA_WITH_AES_128_GCM_SHA256,TLS_RSA_WITH_AES_256_CBC_SHA,TLS_RSA_WITH_AES_256_CBC_SHA256,TLS_RSA_WITH_AES_256_CCM,TLS_RSA_WITH_AES_256_CCM_8,TLS_RSA_WITH_AES_256_GCM_SHA384"
      },

*/

type stTransferSiteCustomModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                    types.String `tfsdk:"type" helper:",default:ExternalPersistedCustomSite"`
	Protocol                types.String `tfsdk:"protocol" helper:",required"`
	TransferType            types.String `tfsdk:"transfer_type" helper:"transferType,default:internal"`
	MaxConcurrentConnection types.Int64  `tfsdk:"max_concurrent_connection" helper:"maxConcurrentConnection,default:0"`
	Default                 types.Bool   `tfsdk:"default" helper:",default:false"`
	AccessLevel             types.String `tfsdk:"access_level" helper:"accessLevel,default:PRIVATE"`
	Account                 types.String `tfsdk:"account" helper:",required"`
	ClientCertificate       types.String `tfsdk:"client_certificate" helper:"clientCertificate,optional"`

	CustomProperties    types.Map `tfsdk:"custom_properties" helper:"customProperties,elementtype:string,optional,noread"`
	CustomPropertiesAll types.Map `tfsdk:"custom_properties_all" helper:"customProperties,elementtype:string,computed,nowrite,fieldMapOnRead:customProperties"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTTransferSiteCustomModelResource() resource.Resource {
	return NewSTResource(&stTransferSiteCustomModel{}, "st_site_custom", "", "/api/v2.0/sites", "/api/v2.0/sites/{id}").AddDiscriminator("[type=custom]")
}

func init() {
	registerResource(NewSTTransferSiteCustomModelResource)
}
