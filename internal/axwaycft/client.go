package axwaycft

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"terraform-provider-axway-cft/internal/axwaycft/tools"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type Client struct {
	client   *http.Client
	Url      string
	Login    string
	Password string
}

type cftobject struct {
	Type       string                 `json:"type"`
	Id         string                 `json:"id"`
	Attributes map[string]interface{} `json:"attributes"`
}

type cftobjectResp struct {
	Data cftobject `json:"data"`
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

	tflog.Info(ctx, "Call "+method+" uri="+uri+" data="+data)
	options := &tools.HttpJSONOptions{
		Headers: map[string]string{
			"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte((c.Login + ":" + c.Password))),
		},
	}
	rawdata, err := tools.HttpJSONRequest(ctx, "xxxxx", client, method, c.Url, uri, options, data, dataOut)
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
		client:   &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}},
	}, nil
}

func (c *Client) About(ctx context.Context) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	_, err := c.Call(ctx, "GET", "/cft/api/v1/about", nil, &data)

	return data, err
}

func (c *Client) ReadObject(ctx context.Context, uri string) (cftobject, error) {
	var data cftobjectResp
	_, err := c.Call(ctx, "GET", uri, nil, &data)
	return data.Data, err
}

func (c *Client) CreateObject(ctx context.Context, uri string, typ string, id string, dataIn map[string]interface{}) (cftobject, error) {
	var data cftobjectResp
	obj := cftobject{
		Id:         id,
		Type:       typ,
		Attributes: dataIn,
	}
	_, err := c.Call(ctx, "PUT", uri, obj, &data)
	tflog.Info(ctx, "CreateObject :"+fmt.Sprint(data.Data))
	return data.Data, err
}

func (c *Client) DeleteObject(ctx context.Context, uri string) error {
	_, err := c.Call(ctx, "DELETE", uri, nil, nil)
	return err
}
