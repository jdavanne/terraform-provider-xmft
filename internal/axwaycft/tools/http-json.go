package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type Set interface {
	Set(name string, value string)
}

type HttpJSONOptions struct {
	ServerName string
	Cookie     *http.Cookie
	Headers    map[string]string
}

func HttpChildCorrelationID(id, childId string) string {
	if id == "" {
		if childId == "" {
			return GenerateShortID()
		}
		return childId
	}
	if childId == "" {
		childId = GenerateShortIDn(8)
	}
	return id + "." + childId
}

func HttpGenerateCorrelationID() string {
	return GenerateShortID()
}

func httpGetCorrelationId(h *http.Header) string {
	id := h.Get("X-TRACE-ID")
	if id == "" {
		id = h.Get("X-Correlation-Id")
	}
	if id == "" {
		id = GenerateShortID()
	}
	return id
}

func httpSetCorrelationIdm(h *map[string]string, id string) string {
	if id == "" {
		id = GenerateShortID()
	}
	(*h)["X-TRACE-ID"] = id
	(*h)["X-Correlation-Id"] = id
	return id
}

func httpSetCorrelationId(h Set, id string) string {
	if id == "" {
		id = GenerateShortID()
	}
	h.Set("X-TRACE-ID", id)
	h.Set("X-Correlation-Id", id)
	return id
}

func HttpJSONRequest(ctx context.Context, xTraceID string, client *http.Client, method, urlbase, uri string, options *HttpJSONOptions, data string, dataOut interface{}) ([]byte, error) {
	// FIXME: ctx = jlog.WithValue(ctx, "method", method)
	url := urlbase + uri
	// FIXME: ctx = jlog.WithValue(ctx, "url", url)

	var reader io.Reader = nil
	if data != "" {
		slog.DebugContext(ctx, "http: request body", "dataIn", data)
		reader = bytes.NewBuffer([]byte(data))
	}

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		slog.ErrorContext(ctx, "http: bad request", "error", err)
		return nil, err
	}

	if options != nil && options.ServerName != "" {
		slog.DebugContext(ctx, "http: request serverName", "serverName", options.ServerName)
		req.Host = options.ServerName
		t, v := client.Transport.(*http.Transport)
		if v {
			slog.DebugContext(ctx, "http: request serverName (TLS)", "serverName", options.ServerName)
			t.TLSClientConfig.ServerName = options.ServerName
		}
	}

	if options != nil {
		for k, v := range options.Headers {
			req.Header.Add(k, v)
		}
	}

	if req.Header.Get("content-type") == "" {
		req.Header.Add("content-type", "application/json")
		// req.Header.Add("accept", "application/json")
		req.Header.Add("accept", "*/*")
	}

	httpSetCorrelationId(req.Header, HttpChildCorrelationID(xTraceID, ""))

	if options != nil && options.Cookie != nil {
		req.AddCookie(options.Cookie)
	}

	resp, err := client.Do(req)
	if err != nil {
		slog.ErrorContext(ctx, "http: request error", "error", err)
		return nil, err
	}

	defer resp.Body.Close()

	debugHeader := false
	if debugHeader {
		ignoreHeaders := map[string]bool{
			"Cache-Control":             true,
			"Connection":                true,
			"Content-Length":            true,
			"Content-Type":              true,
			"Date":                      true,
			"Etag":                      true,
			"Keep-Alive":                true,
			"Pragma":                    true,
			"Strict-Transport-Security": true,
			"Vary":                      true,
			"X-Content-Type-Options":    true,
			"X-Powered-By":              true,
			"X-Xss-Protection":          true,
		}

		for k, v := range resp.Header {
			if !ignoreHeaders[k] {
				slog.DebugContext(ctx, "http: response header", "name", k, "value", v)
			}
		}
	}
	debug := true
	slog.InfoContext(ctx, "http: response", "statusCode", resp.StatusCode)
	if debug && dataOut != nil && resp.StatusCode < 300 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		slog.InfoContext(ctx, "http: response", "body", string(body))
		if err := json.NewDecoder(bytes.NewBuffer(body)).Decode(dataOut); err != nil {
			return nil, err
		}
		slog.DebugContext(ctx, "http: response", "body", dataOut)
		return nil, nil
	} else if !debug && dataOut != nil && resp.StatusCode < 300 {
		if err := json.NewDecoder(resp.Body).Decode(dataOut); err != nil {
			return nil, err
		}
		slog.DebugContext(ctx, "http: response", "body", dataOut)
		return nil, nil
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode >= 300 {
			slog.ErrorContext(ctx, "http: response", "body", string(body))
			return nil, errors.New(fmt.Sprint("unsuccessful statusCode (", resp.StatusCode, ")"))
		} else {
			slog.DebugContext(ctx, "http: response", "body", string(body))
		}
		return body, nil
	}
}

func HttpJSONGetWithOptions(ctx context.Context, xTraceId string, client *http.Client, url, uri string, options *HttpJSONOptions, dataOut interface{}) ([]byte, error) {
	return HttpJSONRequest(ctx, xTraceId, client, "GET", url, uri, options, "", dataOut)
}

func HttpJSONGet(ctx context.Context, xTraceId string, client *http.Client, url, uri string, dataOut interface{}) ([]byte, error) {
	return HttpJSONRequest(ctx, xTraceId, client, "GET", url, uri, nil, "", dataOut)
}

func HttpJSONPost(ctx context.Context, xTraceId string, client *http.Client, url, uri, dataIn string, dataOut interface{}) ([]byte, error) {
	return HttpJSONRequest(ctx, xTraceId, client, "POST", url, uri, nil, dataIn, dataOut)
}

func HttpJSONPut(ctx context.Context, xTraceId string, client *http.Client, url, uri, dataIn string, dataOut interface{}) ([]byte, error) {
	return HttpJSONRequest(ctx, xTraceId, client, "PUT", url, uri, nil, dataIn, dataOut)
}

func HttpJSONDelete(ctx context.Context, xTraceId string, client *http.Client, url, uri, dataIn string, dataOut interface{}) ([]byte, error) {
	return HttpJSONRequest(ctx, xTraceId, client, "DELETE", url, uri, nil, dataIn, dataOut)
}
