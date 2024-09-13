package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type Set interface {
	Set(name string, value string)
}

type HttpError struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
	Method     string `json:"method"`
	URL        string `json:"url"`
}

func (h HttpError) Error() string {
	return fmt.Sprintf("HTTP error: %d %s (%s %s)", h.StatusCode, h.Message, h.Method, h.URL)
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

/*
func HttpGenerateCorrelationID() string {
	return GenerateShortID()
}
*/

/*
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
*/

/*
func httpSetCorrelationIdm(h *map[string]string, id string) string {
	if id == "" {
		id = GenerateShortID()
	}
	(*h)["X-TRACE-ID"] = id
	(*h)["X-Correlation-Id"] = id
	return id
}
*/

func httpSetCorrelationId(h Set, id string) string {
	if id == "" {
		id = GenerateShortID()
	}
	h.Set("X-TRACE-ID", id)
	h.Set("X-Correlation-Id", id)
	return id
}

func HttpJSONRequest(ctx context.Context, xTraceID string, client *http.Client, method, urlbase, uri string, options *HttpJSONOptions, dataIn string, dataOut interface{}) ([]byte, http.Header, error) {
	debug := true
	// FIXME: ctx = jlog.WithValue(ctx, "method", method)
	url := urlbase + uri
	// FIXME: ctx = jlog.WithValue(ctx, "url", url)
	var reader io.Reader = nil
	if dataIn != "" {
		var dataIn2 interface{}
		if err := json.NewDecoder(bytes.NewBuffer([]byte(dataIn))).Decode(&dataIn2); err != nil {
			slog.ErrorContext(ctx, "http: "+method+" "+url+" - error decoding generic json input", "err", err, "body", dataIn)
			// return nil, nil, err
		}
		slog.DebugContext(ctx, "http: "+method+" "+url+" - request body", "dataIn-decoded-raw", dataIn2)
		slog.DebugContext(ctx, "http: "+method+" "+url+" - request body", "dataIn-reader-json", dataIn)
		reader = bytes.NewBuffer([]byte(dataIn))
	}

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		slog.ErrorContext(ctx, "http: "+method+" "+url+" - bad request", "error", err)
		return nil, nil, err
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
	}

	if req.Header.Get("accept") == "" {
		req.Header.Add("accept", "application/json")
		// req.Header.Add("accept", "*/*")
	}

	httpSetCorrelationId(req.Header, HttpChildCorrelationID(xTraceID, ""))

	if options != nil && options.Cookie != nil {
		req.AddCookie(options.Cookie)
	}

	timeStart := time.Now()
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				slog.WarnContext(ctx, "http: "+method+" "+url+" - slow", "time", time.Since(timeStart).String())
			case <-done:
				return
			}
		}
	}()
	resp, err := client.Do(req)
	done <- struct{}{}
	if err != nil {
		slog.ErrorContext(ctx, "http: "+method+" "+url+" - request error", "error", err)
		return nil, nil, err
	}

	defer resp.Body.Close()
	timeEnd := time.Now()
	elapsed := timeEnd.Sub(timeStart)

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

	if !debug {
		slog.InfoContext(ctx, "http: "+method+" "+url+" - response", "statusCode", resp.StatusCode, "time", elapsed.String())
	}
	if debug && dataOut != nil && resp.StatusCode < 300 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			slog.ErrorContext(ctx, "http: "+method+" "+url+" - error reading body", "statusCode", resp.StatusCode, "err", err)
			return nil, nil, err
		}
		var dataOut2 interface{}
		if err := json.NewDecoder(bytes.NewBuffer(body)).Decode(&dataOut2); err != nil {
			slog.ErrorContext(ctx, "http: "+method+" "+url+" - error decoding genric json", "statusCode", resp.StatusCode, "err", err, "body", string(body))
			return nil, nil, err
		}
		// slog.InfoContext(ctx, "http: response", "statusCode", resp.StatusCode, "body", string(body))
		if err := json.NewDecoder(bytes.NewBuffer(body)).Decode(dataOut); err != nil {
			slog.ErrorContext(ctx, "http: "+method+" "+url+" - error decoding to object", "statusCode", resp.StatusCode, "err", err, "body", string(body))
			return nil, nil, err
		}
		slog.DebugContext(ctx, "http: "+method+" "+url+" - response", "statusCode", resp.StatusCode, "time", elapsed.String(), "body", dataOut2, "data", dataOut)
		return nil, resp.Header, nil
	} else if !debug && dataOut != nil && resp.StatusCode < 300 {
		if err := json.NewDecoder(resp.Body).Decode(dataOut); err != nil {
			slog.ErrorContext(ctx, "http: "+method+" "+url+" - error decoding", "statusCode", resp.StatusCode, "err", err)
			return nil, nil, err
		}
		slog.DebugContext(ctx, "http: "+method+" "+url+" - response", "statusCode", resp.StatusCode, "time", elapsed.String(), "body", dataOut)
		return nil, resp.Header, nil
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			slog.ErrorContext(ctx, "http: "+method+" "+url+" - error reading body", "statusCode", resp.StatusCode, "err", err)
			return nil, nil, err
		}

		if resp.StatusCode >= 300 {
			slog.ErrorContext(ctx, "http: "+method+" "+url+" - response", "statusCode", resp.StatusCode, "time", elapsed.String(), "body", string(body))
			err := &HttpError{resp.StatusCode, string(body), method, url}
			return nil, nil, err
		} else {
			slog.DebugContext(ctx, "http: "+method+" "+url+" - response", "statusCode", resp.StatusCode, "time", elapsed.String(), "body", string(body))
		}
		return body, resp.Header, nil
	}
}

/*

func HttpJSONGetWithOptions(ctx context.Context, xTraceId string, client *http.Client, url, uri string, options *HttpJSONOptions, dataOut interface{}) ([]byte, http.Header, error) {
	return HttpJSONRequest(ctx, xTraceId, client, "GET", url, uri, options, "", dataOut)
}

func HttpJSONGet(ctx context.Context, xTraceId string, client *http.Client, url, uri string, dataOut interface{}) ([]byte, http.Header, error) {
	return HttpJSONRequest(ctx, xTraceId, client, "GET", url, uri, nil, "", dataOut)
}

func HttpJSONPost(ctx context.Context, xTraceId string, client *http.Client, url, uri, dataIn string, dataOut interface{}) ([]byte, http.Header, error) {
	return HttpJSONRequest(ctx, xTraceId, client, "POST", url, uri, nil, dataIn, dataOut)
}

func HttpJSONPut(ctx context.Context, xTraceId string, client *http.Client, url, uri, dataIn string, dataOut interface{}) ([]byte, http.Header, error) {
	return HttpJSONRequest(ctx, xTraceId, client, "PUT", url, uri, nil, dataIn, dataOut)
}

func HttpJSONDelete(ctx context.Context, xTraceId string, client *http.Client, url, uri, dataIn string, dataOut interface{}) ([]byte, http.Header, error) {
	return HttpJSONRequest(ctx, xTraceId, client, "DELETE", url, uri, nil, dataIn, dataOut)
}
*/
