package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
  "name": "string",
  "type": "pesit",
  "protocol": "pesit",
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
  "pesitId": "string",
  "host": "string",
  "port": "string",
  "dmz": "none",
  "configurePreconnection": false,
  "preconnectionPartnerId": "none",
  "usePreconnectionPartnerPassword": false,
  "usePreconnectionPartnerPasswordExpr": false,
  "preconnectionPartnerPassword": "none",
  "preconnectionServerId": "none",
  "usePreconnectionServerPassword": false,
  "usePreconnectionServerPasswordExpr": false,
  "preconnectionServerPassword": "none",
  "useServerPassword": false,
  "useServerPasswordExpr": false,
  "serverPassword": "string",
  "usePartnerPassword": false,
  "usePartnerPasswordExpr": false,
  "partnerPassword": "string",
  "compression": "none",
  "resyncAllowed": false,
  "checkpointInterval": 1024,
  "checkpointWindow": 4,
  "connectionTimeout": 60,
  "bufferSize": 8192,
  "sendMessage": "string",
  "receiveMessage": "string",
  "storeAndForwardMode": "START_NEW",
  "originator": "string",
  "finalDestination": "string",
  "ptcpConnections": 1,
  "ptcpPacketSize": 3000,
  "socketSendReceiveBuffersize": 65536,
  "ptcpConnectRetryCount": 10,
  "isSecure": true,
  "verifyCert": false,
  "fipsMode": false,
  "cftCompatibleSslMode": false,
  "loginCertificate": "string",
  "partnerCertificate": "string",
  "cipherSuites": "string",
  "protocols": "string",
  "alternativeAddresses": [
    {
      "host": "host",
      "port": "string",
      "position": 1
    }
  ]
}
*/

type stTransferSitePesitModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                                types.String `tfsdk:"type" helper:",default:pesit"`
	Protocol                            types.String `tfsdk:"protocol" helper:",default:pesit"`
	TransferType                        types.String `tfsdk:"transfer_type" helper:"transferType,default:internal"`
	MaxConcurrentConnection             types.Int64  `tfsdk:"max_concurrent_connection" helper:"maxConcurrentConnection,default:0"`
	Default                             types.Bool   `tfsdk:"default" helper:",default:false"`
	AccessLevel                         types.String `tfsdk:"access_level" helper:"accessLevel,default:PRIVATE"`
	Account                             types.String `tfsdk:"account" helper:",required"`
	PesitId                             types.String `tfsdk:"pesit_id" helper:"pesitId,default:"`
	Host                                types.String `tfsdk:"host" helper:",required"`
	Port                                types.String `tfsdk:"port" helper:",default:1761"`
	Dmz                                 types.String `tfsdk:"dmz" helper:",emptyIsNull,default"`
	ConfigurePreconnection              types.Bool   `tfsdk:"configure_preconnection" helper:"configurePreconnection,default:false"`
	PreconnectionPartnerId              types.String `tfsdk:"preconnection_partner_id" helper:"preconnectionPartnerId,noread,default:"`
	UsePreconnectionPartnerPassword     types.Bool   `tfsdk:"use_preconnection_partner_password" helper:"usePreconnectionPartnerPassword,noread,default:false"`
	UsePreconnectionPartnerPasswordExpr types.Bool   `tfsdk:"use_preconnection_partner_password_expr" helper:"usePreconnectionPartnerPasswordExpr,default:false"`
	PreconnectionPartnerPassword        types.String `tfsdk:"preconnection_partner_password" helper:"preconnectionPartnerPassword,noread,,emptyIsNull,default"`
	PreconnectionServerId               types.String `tfsdk:"preconnection_server_id" helper:"preconnectionServerId,noread,default:"`
	UsePreconnectionServerPassword      types.Bool   `tfsdk:"use_preconnection_server_password" helper:"usePreconnectionServerPassword,default:false"`
	UsePreconnectionServerPasswordExpr  types.Bool   `tfsdk:"use_preconnection_server_password_expr" helper:"usePreconnectionServerPasswordExpr,default:false"`
	PreconnectionServerPassword         types.String `tfsdk:"preconnection_server_password" helper:"preconnectionServerPassword,default:"`
	UseServerPassword                   types.Bool   `tfsdk:"use_server_password" helper:"useServerPassword,default:false"`
	UseServerPasswordExpr               types.Bool   `tfsdk:"use_server_password_expr" helper:"useServerPasswordExpr,default:false"`
	ServerPassword                      types.String `tfsdk:"server_password" helper:"serverPassword,emptyIsNull,noread,default"`
	UsePartnerPassword                  types.Bool   `tfsdk:"use_partner_password" helper:"usePartnerPassword,default:false"`
	UsePartnerPasswordExpr              types.Bool   `tfsdk:"use_partner_password_expr" helper:"usePartnerPasswordExpr,default:false"`
	PartnerPassword                     types.String `tfsdk:"partner_password" helper:"partnerPassword,noread,emptyIsNull,default"`
	Compression                         types.String `tfsdk:"compression" helper:"compression,enum:/none/vertical/horizontal/both,default:none"`
	ResyncAllowed                       types.Bool   `tfsdk:"resync_allowed" helper:"resyncAllowed,default:false"`
	CheckpointInterval                  types.Int64  `tfsdk:"checkpoint_interval" helper:"checkpointInterval,default:1024"`
	CheckpointWindow                    types.Int64  `tfsdk:"checkpoint_window" helper:"checkpointWindow,default:4"`
	ConnectionTimeout                   types.Int64  `tfsdk:"connection_timeout" helper:"connectionTimeout,default:60"`
	BufferSize                          types.Int64  `tfsdk:"buffer_size" helper:"bufferSize,default:8192"`
	SendMessage                         types.String `tfsdk:"send_message" helper:"sendMessage,emptyIsNull,default"`
	ReceiveMessage                      types.String `tfsdk:"receive_message" helper:"receiveMessage,emptyIsNull,default"`
	StoreAndForwardMode                 types.String `tfsdk:"store_and_forward_mode" helper:"storeAndForwardMode,default:START_NEW"`
	Originator                          types.String `tfsdk:"originator" helper:"originator,emptyIsNull,default"`
	FinalDestination                    types.String `tfsdk:"final_destination" helper:"finalDestination,default"`
	PtcpConnections                     types.Int64  `tfsdk:"ptcp_connections" helper:"ptcpConnections,default:1"`
	PtcpPacketSize                      types.Int64  `tfsdk:"ptcp_packet_size" helper:"ptcpPacketSize,default:3000"`
	SocketSendReceiveBuffersize         types.Int64  `tfsdk:"socket_send_receive_buffersize" helper:"socketSendReceiveBuffersize,default:65536"`
	PtcpConnectRetryCount               types.Int64  `tfsdk:"ptcp_connect_retry_count" helper:"ptcpConnectRetryCount,default:10"`
	IsSecure                            types.Bool   `tfsdk:"is_secure" helper:"isSecure,default:false"`
	VerifyCert                          types.Bool   `tfsdk:"verify_cert" helper:"verifyCert,default:false"`
	CftCompatibleSslMode                types.Bool   `tfsdk:"cft_compatible_ssl_mode" helper:"cftCompatibleSslMode,default:false"`
	LoginCertificate                    types.String `tfsdk:"login_certificate" helper:"loginCertificate,emptyIsNull,default"`
	PartnerCertificate                  types.String `tfsdk:"partner_certificate" helper:"partnerCertificate,emptyIsNull,default"`
	CipherSuites                        types.String `tfsdk:"cipher_suites" helper:"cipherSuites,optional,computed"`
	Protocols                           types.String `tfsdk:"protocols" helper:"protocols,optional,computed"`

	AlternativeAddresses []struct {
		Host     types.String `tfsdk:"host" helper:",required"`
		Port     types.String `tfsdk:"port" helper:",required"`
		Position types.Int64  `tfsdk:"position" helper:",required"`
	} `tfsdk:"alternative_addresses" helper:"alternativeAddresses"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTTransferSitePesitModelResource() resource.Resource {
	return NewSTResource(&stTransferSitePesitModel{}, "st_site_pesit", "", "/api/v2.0/sites", "/api/v2.0/sites/{id}").AddDiscriminator("[type=pesit]")
}

func init() {
	registerResource(NewSTTransferSitePesitModelResource)
}
