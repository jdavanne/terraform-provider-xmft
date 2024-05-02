package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
	{
	      "type": "cftrecv",
	      "id": "TXT",
	      "attributes": {
	        "cycdate": "00000000",
	        "cyctime": "00000000",
	        "maxdate": "99991231",
	        "maxtime": "23595999",
	        "mindate": "10000101",
	        "mintime": "00000000",
	        "maxduration": "0",
	        "ackmindate": "10000101",
	        "ackmintime": "00000000",
	        "postmindate": "10000101",
	        "postmintime": "00000000",
	        "ackstate": "IGNORE",
	        "acktimeout": "0",
	        "delete": "NO",
	        "dirnb": "0",
	        "execrall": "ALL",
	        "faction": "DELETE",
	        "fblksize": "0",
	        "fcheck": "NO",
	        "fcode": "ASCII",
	        "fdisp": "BOTH",
	        "filenotfound": "ABORT",
	        "fkeylen": "0",
	        "fkeypos": "0",
	        "flrecl": "0",
	        "fname": "pub/&IDT_&IDTU.RCV",
	        "force": "NO",
	        "forg": "SEQ",
	        "fspace": "0",
	        "ftype": "T",
	        "maction": " ",
	        "ncode": " ",
	        "ncomp": "15",
	        "netband": "1",
	        "opermsg": "0",
	        "origin": "CFTUTIL",
	        "poststate": "DISP",
	        "posttimeout": "0",
	        "pri": "128",
	        "serial": " ",
	        "state": "DISP",
	        "trk": "UNDEFINED",
	        "naction": "NONE",
	        "idf": []
	      },
*/

type cftRecvResourceModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required,noread,nowrite"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`
	Fname       types.String `tfsdk:"fname"`
	Exec        types.String `tfsdk:"exec"`

	Cycdate      types.String `tfsdk:"cycdate" helper:",computed,optional"`
	Cyctime      types.String `tfsdk:"cyctime" helper:",computed,optional"`
	Maxdate      types.String `tfsdk:"maxdate" helper:",computed,optional"`
	Maxtime      types.String `tfsdk:"maxtime" helper:",computed,optional"`
	Mindate      types.String `tfsdk:"mindate" helper:",computed,optional"`
	Mintime      types.String `tfsdk:"mintime" helper:",computed,optional"`
	Maxduration  types.String `tfsdk:"maxduration" helper:",computed,optional"`
	Ackmindate   types.String `tfsdk:"ackmindate" helper:",computed,optional"`
	Ackmintime   types.String `tfsdk:"ackmintime" helper:",computed,optional"`
	Postmindate  types.String `tfsdk:"postmindate" helper:",computed,optional"`
	Postmintime  types.String `tfsdk:"postmintime" helper:",computed,optional"`
	Ackstate     types.String `tfsdk:"ackstate" helper:",computed,optional"`
	Acktimeout   types.String `tfsdk:"acktimeout" helper:",computed,optional"`
	Delete       types.String `tfsdk:"delete" helper:",computed,optional"`
	Dirnb        types.String `tfsdk:"dirnb" helper:",computed,optional"`
	Execrall     types.String `tfsdk:"execrall" helper:",computed,optional"`
	Faction      types.String `tfsdk:"faction" helper:",computed,optional"`
	Fblksize     types.String `tfsdk:"fblksize" helper:",computed,optional"`
	Fcheck       types.String `tfsdk:"fcheck" helper:",computed,optional"`
	Fcode        types.String `tfsdk:"fcode" helper:",computed,optional"`
	Fdisp        types.String `tfsdk:"fdisp" helper:",computed,optional"`
	Filenotfound types.String `tfsdk:"filenotfound" helper:",computed,optional"`
	Fkeylen      types.String `tfsdk:"fkeylen" helper:",computed,optional"`
	Fkeypos      types.String `tfsdk:"fkeypos" helper:",computed,optional"`
	Flrecl       types.String `tfsdk:"flrecl" helper:",computed,optional"`
	Force        types.String `tfsdk:"force" helper:",computed,optional"`
	Forg         types.String `tfsdk:"forg" helper:",computed,optional"`
	Fspace       types.String `tfsdk:"fspace" helper:",computed,optional"`
	Ftype        types.String `tfsdk:"ftype" helper:",computed,optional"`
	Maction      types.String `tfsdk:"maction" helper:",computed,optional"`
	Ncode        types.String `tfsdk:"ncode" helper:",computed,optional"`
	Ncomp        types.String `tfsdk:"ncomp" helper:",computed,optional"`
	Netband      types.String `tfsdk:"netband" helper:",computed,optional"`
	Opermsg      types.String `tfsdk:"opermsg" helper:",computed,optional"`
	Poststate    types.String `tfsdk:"poststate" helper:",computed,optional"`
	Posttimeout  types.String `tfsdk:"posttimeout" helper:",computed,optional"`
	Pri          types.String `tfsdk:"pri" helper:",computed,optional"`
	Serial       types.String `tfsdk:"serial" helper:",computed,optional"`
	State        types.String `tfsdk:"state" helper:",computed,optional"`
	Trk          types.String `tfsdk:"trk" helper:",computed,optional"`
	Naction      types.String `tfsdk:"naction" helper:",computed,optional"`
	Idf          types.String `tfsdk:"idf" helper:",computed,optional"`
}

func NewCFTRecvResource() resource.Resource {
	return NewCFTResource(&cftRecvResourceModel{}, "cftrecv", "cftrecv", "/cft/api/v1/objects/cftrecv", "/cft/api/v1/objects/cftrecv/{name}")
}
