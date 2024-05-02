package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
	{
	      "type": "cftsendno",
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
	        "prestate": " ",
	        "pretimeout": "0",
	        "premindate": "10000101",
	        "premintime": "00000000",
	        "ackstate": "IGNORE",
	        "acktimeout": "0",
	        "cycle": "0",
	        "delete": "NO",
	        "execsub": "LIST",
	        "execsuba": "SUBF",
	        "execsubpre": "LIST",
	        "faction": "NONE",
	        "fblksize": "0",
	        "fcode": "ASCII",
	        "fdisp": "SHR",
	        "filenotfound": "ABORT",
	        "filtertype": "STRJCMP",
	        "fkeylen": "0",
	        "fkeypos": "0",
	        "flrecl": "0",
	        "fname": "pub/FTEST",
	        "force": "NO",
	        "forg": " ",
	        "frecfm": " ",
	        "fspace": "0",
	        "ftype": "T",
	        "nblksize": "0",
	        "ncode": " ",
	        "ncomp": "15",
	        "netband": "1",
	        "nkeylen": "0",
	        "nkeypos": "0",
	        "nlrecl": "0",
	        "nrecfm": " ",
	        "nspace": "0",
	        "opermsg": "0",
	        "origin": "CFTUTIL",
	        "poststate": "DISP",
	        "posttimeout": "0",
	        "pri": "128",
	        "serial": " ",
	        "state": "DISP",
	        "tcycle": "DAY",
	        "trk": "UNDEFINED",
	        "idf": []
	      }
*/
type cftSendResourceModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required,noread,nowrite"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Fcode        types.String `tfsdk:"fcode" helper:",computed,optional"`
	Fname        types.String `tfsdk:"fname"`
	Faction      types.String `tfsdk:"faction" helper:",computed,optional"`
	Exec         types.String `tfsdk:"exec" helper:",computed,optional"`
	Parm         types.String `tfsdk:"parm" helper:",computed,optional"`
	Preexec      types.String `tfsdk:"preexec" helper:",computed,optional"`
	Cycdate      types.String `tfsdk:"cycdate" helper:",computed,optional"`
	Cyctime      types.String `tfsdk:"cyctime" helper:",computed,optional"`
	Maxdate      types.String `tfsdk:"maxdate" helper:",computed,optional"`
	Maxtime      types.String `tfsdk:"maxtime" helper:",computed,optional"`
	Mindate      types.String `tfsdk:"mindate" helper:",computed,optional"`
	MinTime      types.String `tfsdk:"mintime" helper:",computed,optional"`
	Maxduration  types.String `tfsdk:"maxduration" helper:",computed,optional"`
	Ackmindate   types.String `tfsdk:"ackmindate" helper:",computed,optional"`
	Ackmintime   types.String `tfsdk:"ackmintime" helper:",computed,optional"`
	Postmindate  types.String `tfsdk:"postmindate" helper:",computed,optional"`
	Postmintime  types.String `tfsdk:"postmintime" helper:",computed,optional"`
	Prestate     types.String `tfsdk:"prestate" helper:",computed,optional"`
	Pretimeout   types.String `tfsdk:"pretimeout" helper:",computed,optional"`
	Poststate    types.String `tfsdk:"poststate" helper:",computed,optional"`
	Posttimeout  types.String `tfsdk:"posttimeout" helper:",computed,optional"`
	Premindate   types.String `tfsdk:"premindate" helper:",computed,optional"`
	Premintime   types.String `tfsdk:"premintime" helper:",computed,optional"`
	Ackstate     types.String `tfsdk:"ackstate" helper:",computed,optional"`
	Acktimeout   types.String `tfsdk:"acktimeout" helper:",computed,optional"`
	Cycle        types.String `tfsdk:"cycle" helper:",computed,optional"`
	Delete       types.String `tfsdk:"delete" helper:",computed,optional"`
	Execsub      types.String `tfsdk:"execsub" helper:",computed,optional"`
	Execsuba     types.String `tfsdk:"execsuba" helper:",computed,optional"`
	Execsubpre   types.String `tfsdk:"execsubpre" helper:",computed,optional"`
	Fblksize     types.String `tfsdk:"fblksize" helper:",computed,optional"`
	Fdisp        types.String `tfsdk:"fdisp" helper:",computed,optional"`
	Filenotfound types.String `tfsdk:"filenotfound" helper:",computed,optional"`
	Filtertype   types.String `tfsdk:"filtertype" helper:",computed,optional"`
	Fkeylen      types.String `tfsdk:"fkeylen" helper:",computed,optional"`
	Fkeypos      types.String `tfsdk:"fkeypos" helper:",computed,optional"`
	Flrecl       types.String `tfsdk:"flrecl" helper:",computed,optional"`
	Force        types.String `tfsdk:"force" helper:",computed,optional"`
	Forg         types.String `tfsdk:"forg" helper:",computed,optional"`
	Frecfm       types.String `tfsdk:"frecfm" helper:",computed,optional"`
	Fspace       types.String `tfsdk:"fspace" helper:",computed,optional"`
	Ftype        types.String `tfsdk:"ftype" helper:",computed,optional"`
	Nblksize     types.String `tfsdk:"nblksize" helper:",computed,optional"`
	Ncode        types.String `tfsdk:"ncode" helper:",computed,optional"`
	Ncomp        types.String `tfsdk:"ncomp" helper:",computed,optional"`
	Netband      types.String `tfsdk:"netband" helper:",computed,optional"`
	Nkeylen      types.String `tfsdk:"nkeylen" helper:",computed,optional"`
	Nkeypos      types.String `tfsdk:"nkeypos" helper:",computed,optional"`
	Nlrecl       types.String `tfsdk:"nlrecl" helper:",computed,optional"`
	Nrecfm       types.String `tfsdk:"nrecfm" helper:",computed,optional"`
	Nspace       types.String `tfsdk:"nspace" helper:",computed,optional"`
	Opermsg      types.String `tfsdk:"opermsg" helper:",computed,optional"`
	Pri          types.String `tfsdk:"pri" helper:",computed,optional"`
	Serial       types.String `tfsdk:"serial" helper:",computed,optional"`
	State        types.String `tfsdk:"state" helper:",computed,optional"`
	Tcycle       types.String `tfsdk:"tcycle" helper:",computed,optional"`
	Trk          types.String `tfsdk:"trk" helper:",computed,optional"`
	Idf          types.String `tfsdk:"idf" helper:",computed,optional"`
}

// NewCFTSendResource is a helper function to simplify the provider implementation.
func NewCFTSendResource() resource.Resource {
	return NewCFTResource(&cftSendResourceModel{}, "cftsend", "cftsendno", "/cft/api/v1/objects/cftsend/implno", "/cft/api/v1/objects/cftsend/implno/{name}")
	// return &cftSendResource{}
}
