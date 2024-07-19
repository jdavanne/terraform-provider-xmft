package stapi

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"

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

func (c *Client) Call(ctx context.Context, method, uri string, dataIn map[string]interface{}, dataOut interface{}) ([]byte, http.Header, error) {
	data := ""
	contentType := ""

	if strings.HasPrefix(uri, "/api/v2.0/configurations/options/") && method == "PUT" {
		method = "PATCH"
		data = `[{ 
			"op" :"replace",
			"path": "/values",
			"value": [ "` + dataIn["value"].(string) + `"]
		}]
		`
	} else if uri == "/api/v2.0/certificates" && method == "POST" {
		content, ok := dataIn["content"].(string)
		if !ok {
			return nil, nil, fmt.Errorf("/api/v2.0/certificates : content is not a string")
		}
		delete(dataIn, "content")
		jsonStr, err := tools.JSONMarchal(dataIn)
		if err != nil {
			return nil, nil, err
		}

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, _ := writer.CreatePart(textproto.MIMEHeader{"Content-Type": {"application/json"}})
		part.Write(jsonStr)
		part2, _ := writer.CreatePart(textproto.MIMEHeader{
			"Content-Type":        {"application/octet-stream"},
			"Content-Disposition": {`attachment; filename="cert.crt"`},
		})

		part2.Write([]byte(content))

		writer.Close()

		data = body.String()
		contentType = "multipart/mixed; boundary=" + writer.Boundary()
	}

	if data == "" && dataIn != nil {
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
	if contentType != "" {
		options.Headers["Content-Type"] = contentType
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

	if uri == "/api/v2.0/certificates" {
		_, _, err := c.Call(ctx, "POST", uri, dataIn, &data)
		if err != nil {
			return data, err
		}
	} else {
		_, header, err := c.Call(ctx, "POST", uri, dataIn, nil)
		if err != nil {
			return data, err
		}

		location := header.Get("Location")
		_, _, err = c.Call(ctx, "GET", location[len(c.Url):], nil, &data)
		if err != nil {
			return data, err
		}
	}

	tflog.Info(ctx, "CreateObject :"+fmt.Sprint(data))
	// stObjectFlatten(ctx, data.Data)
	return data, nil
}

func (c *Client) ReplaceObject(ctx context.Context, uri string, uri2 string, typ string, dataIn map[string]interface{}) (StObject, error) {
	var data StObject

	// obj.Attributes["origin"] = "terraform"
	_, _, err := c.Call(ctx, "PUT", uri, dataIn, nil)
	if err != nil {
		return data, err
	}
	_, _, err = c.Call(ctx, "GET", uri2, nil, &data)
	tflog.Info(ctx, "UpdateObject :"+fmt.Sprint(data))

	// stObjectFlatten(ctx, data.Data)
	return data, err
}

func (c *Client) DeleteObject(ctx context.Context, uri string) error {
	_, _, err := c.Call(ctx, "DELETE", uri, nil, nil)
	return err
}
