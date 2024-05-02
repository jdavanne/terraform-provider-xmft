package cftapi

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"terraform-provider-xmft/internal/tools"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type Client struct {
	client   *http.Client
	Url      string
	Login    string
	Password string
}

type CftObject struct {
	Type       string                 `json:"type"`
	Id         string                 `json:"id"`
	Attributes map[string]interface{} `json:"attributes"`
}

type CftObjectResp struct {
	Data CftObject `json:"data"`
}

var client = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

func (c *Client) Call(ctx context.Context, method, uri string, dataIn interface{}, dataOut interface{}) ([]byte, error) {
	data := ""
	if dataIn != nil {
		d, err := tools.JSONMarchal(dataIn)
		if err != nil {
			return nil, err
		}
		data = string(d)

	}

	// tflog.Info(ctx, "Call "+method+" uri="+uri+" data="+data)
	options := &tools.HttpJSONOptions{
		Headers: map[string]string{
			"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte((c.Login + ":" + c.Password))),
		},
	}
	rawdata, _, err := tools.HttpJSONRequest(ctx, "xxxxx", client, method, c.Url, uri, options, data, dataOut)
	tflog.Info(ctx, "xxxCall "+method+" uri="+uri+" data="+data)
	return rawdata, err
}

func init() {
	programLevel := new(slog.LevelVar)
	programLevel.Set(slog.LevelDebug)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
}

func NewClient(url *string, username *string, password *string) (*Client, error) {
	return &Client{
		Url:      *url,
		Login:    *username,
		Password: *password,
		client:   client,
	}, nil
}

func (c *Client) About(ctx context.Context) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	_, err := c.Call(ctx, "GET", "/cft/api/v1/about", nil, &data)

	return data, err
}

func cftObjectFlatten(ctx context.Context, obj CftObject) {
	if obj.Attributes == nil {
		obj.Attributes = make(map[string]interface{})
	}
	// FIXME: huge hack to handle tcp particularity
	if obj.Type == "cftpart" {
		tcp := obj.Attributes["tcp"]
		if tcp != nil {
			if tcpl, ok := tcp.([]interface{}); ok {
				l := make([]interface{}, 0, len(tcpl))
				for i := 0; i < len(tcpl); i++ {
					if tcpItem, ok := tcpl[i].(map[string]interface{}); ok {
						// obj.Attributes["tcp"] = tcpItem["attributes"]
						attrs := tcpItem["attributes"].(map[string]interface{})
						attrs["id"] = tcpItem["id"]
						l = append(l, attrs)
					} else {
						panic("expecting a interface{}")
					}
				}
				obj.Attributes["tcp"] = l
			} else {
				panic("expecting a list")
			}
		}
		tflog.Info(ctx, "xxxCall flatten"+fmt.Sprint(obj, " ", tcp))
	}
	obj.Attributes["_id"] = obj.Id
	obj.Attributes["_type"] = obj.Type
}

func (c *Client) ReadObject(ctx context.Context, uri string) (CftObject, error) {
	var data CftObjectResp
	_, err := c.Call(ctx, "GET", uri, nil, &data)
	cftObjectFlatten(ctx, data.Data)
	return data.Data, err
}

func (c *Client) CreateObject(ctx context.Context, uri string, typ string, id string, dataIn map[string]interface{}) (CftObject, error) {
	var data CftObjectResp
	obj := CftObject{
		Id:         id,
		Type:       typ,
		Attributes: dataIn,
	}
	// obj.Attributes["origin"] = "terraform"
	_, err := c.Call(ctx, "POST", uri, obj, &data)
	tflog.Info(ctx, "CreateObject :"+fmt.Sprint(data.Data))
	cftObjectFlatten(ctx, data.Data)
	return data.Data, err
}

func (c *Client) ReplaceObject(ctx context.Context, uri string, typ string, id string, dataIn map[string]interface{}) (CftObject, error) {
	var data CftObjectResp
	obj := CftObject{
		Id:         id,
		Type:       typ,
		Attributes: dataIn,
	}
	// obj.Attributes["origin"] = "terraform"
	_, err := c.Call(ctx, "PUT", uri, obj, &data)
	tflog.Info(ctx, "CreateObject :"+fmt.Sprint(data.Data))
	cftObjectFlatten(ctx, data.Data)
	return data.Data, err
}

func (c *Client) DeleteObject(ctx context.Context, uri string) error {
	_, err := c.Call(ctx, "DELETE", uri, nil, nil)
	return err
}
