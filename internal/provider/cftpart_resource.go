package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type cftTcpResourceModelSub struct {
	Id types.String `tfsdk:"id" helper:",required"`

	Cnxin    types.String `tfsdk:"cnxin" helper:",computed,optional"`
	Cnxout   types.String `tfsdk:"cnxout" helper:",computed,optional"`
	Cnxinout types.String `tfsdk:"cnxinout" helper:",computed,optional"`

	Omaxtime types.String `tfsdk:"omaxtime"  helper:",computed,optional"`
	Omintime types.String `tfsdk:"omintime"  helper:",computed,optional"`
	Imaxtime types.String `tfsdk:"imaxtime"  helper:",computed,optional"`
	Imintime types.String `tfsdk:"imintime"  helper:",computed,optional"`

	Retrym types.String `tfsdk:"retrym" helper:",computed,optional"`
	Retryn types.String `tfsdk:"retryn" helper:",computed,optional"`
	Retryw types.String `tfsdk:"retryw" helper:",computed,optional"`

	Host types.String `tfsdk:"host" helper:",optional"`
}

/*
	{
	      "type": "cftpart",
	      "id": "CFT2",
	      "attributes": {
	        "commut": "YES",
	        "ctrlpart": "IGNORE",
	        "imaxtime": "23595999",
	        "imintime": "00000000",
	        "nack": " ",
	        "nrpart": "CFT2",
	        "nrpassw": "<CFT_PASSWORD>e96t69lGgYEtYCrbEIwFbssdW7Q70Oe7",
	        "nspart": "CFT1",
	        "nspassw": "<CFT_PASSWORD>o+PWhYx5WJ8pIKDT2ORtJBTlVEcd6o9Z",
	        "omaxtime": "23595999",
	        "omintime": "00000000",
	        "origin": "DESIGNER",
	        "prot": [
	          "PESIT"
	        ],
	        "rauth": "*",
	        "sap": [
	          "1788"
	        ],
	        "sauth": "*",
	        "state": "ACTIVEBOTH",
	        "trk": "UNDEFINED",
	        "syst": "UNIX",
	        "tcp": [
	          {
	            "type": "cfttcp",
	            "id": "1",
	            "attributes": {
	              "cnxin": "2",
	              "cnxinout": "4",
	              "cnxout": "100",
	              "imaxtime": "23595999",
	              "imintime": "00000000",
	              "omaxtime": "23595999",
	              "omintime": "00000000",
	              "origin": "DESIGNER",
	              "retrym": "12",
	              "retryn": "4",
	              "retryw": "1",
	              "host": [
	                "cft2"
	              ],
	              "verify": "0"
	            }
	          }
	        ]
	      }
*/
type cftPartResourceModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state,nowrite"`
	Name        types.String `tfsdk:"name" helper:",required,noread,nowrite"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Nrpart  types.String `tfsdk:"nrpart" helper:",computed,optional"`
	Nrpassw types.String `tfsdk:"nrpassw" helper:",noread"`
	// Nrpasswe types.String `tfsdk:"nrpasswe" helper:"computed,nowrite,readmap"`
	Nspart  types.String `tfsdk:"nspart" helper:",computed,optional"`
	Nspassw types.String `tfsdk:"nspassw" helper:",noread"`
	// Nspasswe types.String `tfsdk:"nspasswe" helper:"computed,nowrite,readmap"`

	Omaxtime types.String `tfsdk:"omaxtime"  helper:",computed,optional"`
	Omintime types.String `tfsdk:"omintime"  helper:",computed,optional"`
	Imaxtime types.String `tfsdk:"imaxtime"  helper:",computed,optional"`
	Imintime types.String `tfsdk:"imintime"  helper:",computed,optional"`

	Commut types.String `tfsdk:"commut" helper:",computed,optional"`

	Prot types.String `tfsdk:"prot"` // (array)
	Sap  types.String `tfsdk:"sap"`  //(array)

	Sauth types.String `tfsdk:"sauth" helper:",computed,optional"`
	Rauth types.String `tfsdk:"rauth" helper:",computed,optional"`

	State types.String `tfsdk:"state"  helper:",computed,optional"`
	Trk   types.String `tfsdk:"trk"  helper:",computed,optional"`
	Syst  types.String `tfsdk:"syst"  helper:",computed,optional"`

	Comment types.String `tfsdk:"comment" helper:",optional,default"`

	Tcp []cftTcpResourceModelSub `tfsdk:"tcp" helper:",optional"`
}

func NewCFTPartResource() resource.Resource {
	return NewCFTResource(&cftPartResourceModel{}, "cftpart", "cftpart", "/cft/api/v1/objects/cftpart", "/cft/api/v1/objects/cftpart/{name}")
}

func init() {
	registerResource(NewCFTPartResource)
}
