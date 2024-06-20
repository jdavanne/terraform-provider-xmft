package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
  "id": "string",
  "name": "string",
  "default": false,
  "account": "string",
  "sendMapping": "/*",
  "receiveMapping": "/${pesit.fileLabel}",
  "sendingAcknowledgmentEnabled": false,
  "fileLabelOption": "DONT_SEND",
  "multiSelect": false,
  "transferMode": "BINARY",
  "recordFormat": "Variable",
  "recordLength": 2048,
  "paddingStripEnabled": false,
  "additionalAttributes": {
    "additionalProp1": "string",
    "additionalProp2": "string",
    "additionalProp3": "string"
  },
  "advancedSettings": {
    "enabled": true,
    "callerTranscoding": {
      "type": "ebcdic_predefined",
      "localDataCode": "EBCDIC",
      "networkDataCode": "EBCDIC",
      "transcoding": "predefined",
      "outputRecordLength": 2048,
      "outputRecordFormat": "VARIABLE",
      "paddingCharacter": "\\u0040",
      "sourceEncodingScheme": "IBM1047",
      "outputEncodingScheme": "IBM1047"
    },
    "receiverTranscoding": {
      "type": "predefined",
      "transcoding": "predefined",
      "outputRecordLength": 2048,
      "outputRecordFormat": "VARIABLE",
      "paddingCharacter": "\\u0020",
      "lineEndingFormat": "DEFAULT",
      "sourceEncodingScheme": "UTF-8",
      "outputEncodingScheme": "UTF-8"
    }
  }
}
*/

type stTransferProfileModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Default                      types.Bool   `tfsdk:"default" helper:",default:false"`
	Account                      types.String `tfsdk:"account" helper:",required"`
	SendMapping                  types.String `tfsdk:"send_mapping" helper:"sendMapping,default:/*"`
	ReceiveMapping               types.String `tfsdk:"receive_mapping" helper:"receiveMapping,default:/${pesit.fileLabel}"`
	SendingAcknowledgmentEnabled types.Bool   `tfsdk:"sending_acknowledgment_enabled" helper:"sendingAcknowledgmentEnabled,default:false"`
	FileLabelOption              types.String `tfsdk:"file_label_option" helper:"fileLabelOption,enum:/DONT_SEND/SEND_FILENAME/SEND_FILENAME_AND_PATH,default:DONT_SEND"`
	MultiSelect                  types.Bool   `tfsdk:"multi_select" helper:"multiSelect,default:false"`
	TransferMode                 types.String `tfsdk:"transfer_mode" helper:"transferMode,enum:/BINARY/ASCII/EBCDIC/EBCDIC_NATIVE,default:BINARY"`
	RecordFormat                 types.String `tfsdk:"record_format" helper:"recordFormat,enum:/Variable/Fixed,default:Variable"`
	RecordLength                 types.Int64  `tfsdk:"record_length" helper:"recordLength,max:32767,default:32767"`
	PaddingStripEnabled          types.Bool   `tfsdk:"padding_strip_enabled" helper:"paddingStripEnabled,default:false"`
	AdditionalAttributes         types.Map    `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
	/*AdvancedSettings             *struct {
		Enabled           types.Bool `tfsdk:"enabled" helper:"enabled,default:true"`
		CallerTranscoding *struct {
			Type                 types.String `tfsdk:"type" helper:"type,default:ebcdic_predefined"`
			LocalDataCode        types.String `tfsdk:"local_data_code" helper:"localDataCode,default:EBCDIC"`
			NetworkDataCode      types.String `tfsdk:"network_data_code" helper:"networkDataCode,default:EBCDIC"`
			Transcoding          types.String `tfsdk:"transcoding" helper:"transcoding,default:predefined"`
			OutputRecordLength   types.Int64  `tfsdk:"output_record_length" helper:"outputRecordLength,default:2048"`
			OutputRecordFormat   types.String `tfsdk:"output_record_format" helper:"outputRecordFormat,default:VARIABLE"`
			PaddingCharacter     types.String `tfsdk:"padding_character" helper:"paddingCharacter,default:\\u0040"`
			SourceEncodingScheme types.String `tfsdk:"source_encoding_scheme" helper:"sourceEncodingScheme,default:IBM1047"`
			OutputEncodingScheme types.String `tfsdk:"output_encoding_scheme" helper:"outputEncodingScheme,default:IBM1047"`
		} `tfsdk:"caller_transcoding" helper:"callerTranscoding,default:"`
		ReceiverTranscoding *struct {
			Type                 types.String `tfsdk:"type" helper:"type,default:predefined"`
			Transcoding          types.String `tfsdk:"transcoding" helper:"transcoding,default:predefined"`
			OutputRecordLength   types.Int64  `tfsdk:"output_record_length" helper:"outputRecordLength,default:2048"`
			OutputRecordFormat   types.String `tfsdk:"output_record_format" helper:"outputRecordFormat,default:VARIABLE"`
			PaddingCharacter     types.String `tfsdk:"padding_character" helper:"paddingCharacter,default:\\u0020"`
			LineEndingFormat     types.String `tfsdk:"line_ending_format" helper:"lineEndingFormat,default:DEFAULT"`
			SourceEncodingScheme types.String `tfsdk:"source_encoding_scheme" helper:"sourceEncodingScheme,default:UTF-8"`
			OutputEncodingScheme types.String `tfsdk:"output_encoding_scheme" helper:"outputEncodingScheme,default:UTF-8"`
		} `tfsdk:"receiver_transcoding" helper:"receiverTranscoding,default:"`
	} `tfsdk:"advanced_settings" helper:"advancedSettings,default:"`*/
}

func NewSTTransferProfileModelResource() resource.Resource {
	return NewSTResource(&stTransferProfileModel{}, "st_transfer_profile", "", "/api/v2.0/transferProfiles", "/api/v2.0/transferProfiles/{id}")
}

func init() {
	registerResource(NewSTTransferProfileModelResource)
}
