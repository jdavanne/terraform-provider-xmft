package stapi

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/http"

	"terraform-provider-xmft/internal/tools"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type Client struct {
	client   *http.Client
	Url      string
	Login    string
	Password string
}

type StObject map[string]interface{}

type StObjectResp struct {
	Data StObject `json:"data"`
}

var client = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

func NewClient(url *string, username *string, password *string) (*Client, error) {
	return &Client{
		Url:      *url,
		Login:    *username,
		Password: *password,
		client:   client,
	}, nil
}

func (c *Client) Call(ctx context.Context, method, uri string, dataIn interface{}, dataOut interface{}) ([]byte, http.Header, error) {
	data := ""
	if dataIn != nil {
		d, err := tools.JSONMarchal(dataIn)
		if err != nil {
			return nil, nil, err
		}
		data = string(d)

	}

	// tflog.Info(ctx, "Call "+method+" uri="+uri+" data="+data)
	options := &tools.HttpJSONOptions{
		Headers: map[string]string{
			"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte((c.Login + ":" + c.Password))),
		},
	}
	rawdata, header, err := tools.HttpJSONRequest(ctx, "xxxxx", client, method, c.Url, uri, options, data, dataOut)
	tflog.Info(ctx, "xxxCall "+method+" uri="+uri+" data="+data)
	return rawdata, header, err
}

func (c *Client) ReadObject(ctx context.Context, uri string) (StObject, error) {
	var data StObject
	_, _, err := c.Call(ctx, "GET", uri, nil, &data)
	// stObjectFlatten(ctx, data.Data)
	return data, err
}

func (c *Client) CreateObject(ctx context.Context, uri string, typ string, dataIn map[string]interface{}) (StObject, error) {
	var data StObject
	// obj.Attributes["origin"] = "terraform"
	_, header, err := c.Call(ctx, "POST", uri, dataIn, nil)
	if err != nil {
		return data, err
	}

	location := header.Get("Location")
	_, _, err = c.Call(ctx, "GET", location[len(c.Url):], nil, &data)

	tflog.Info(ctx, "CreateObject :"+fmt.Sprint(data))
	// stObjectFlatten(ctx, data.Data)
	return data, err
}

func (c *Client) ReplaceObject(ctx context.Context, uri string, typ string, dataIn map[string]interface{}) (StObject, error) {
	var data StObject

	// obj.Attributes["origin"] = "terraform"
	_, _, err := c.Call(ctx, "PUT", uri, dataIn, nil)
	if err != nil {
		return data, err
	}
	_, _, err = c.Call(ctx, "GET", uri, nil, &data)
	tflog.Info(ctx, "UpdateObject :"+fmt.Sprint(data))

	// stObjectFlatten(ctx, data.Data)
	return data, err
}

func (c *Client) DeleteObject(ctx context.Context, uri string) error {
	_, _, err := c.Call(ctx, "DELETE", uri, nil, nil)
	return err
}
