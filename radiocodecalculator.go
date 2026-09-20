/******************************************************************************
 * Radio Code Calculator API - WebApi interface
 *
 * Version      : v1.1.6
 * Language     : Go
 * Author       : Bartosz Wójcik (support@pelock.com)
 * Homepage     : https://www.pelock.com
 *
 *****************************************************************************/

package radiocodecalculator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

const APIURL = "https://www.pelock.com/api/radio-code-calculator/v1"

// Client is the Radio Code Calculator Web API client.
type Client struct {
	APIKey     string
	APIURL     string
	UserAgent  string
	HTTPClient *http.Client
}

// New creates a client. The activation key cannot be empty for a full license.
func New(apiKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		APIURL:     APIURL,
		UserAgent:  "PELock Radio Code Calculator",
		HTTPClient: &http.Client{Timeout: 120 * time.Second},
	}
}

// Login returns (errorCode, raw JSON map).
func (c *Client) Login(ctx context.Context) (int, map[string]any, error) {
	result, err := c.postRequest(ctx, map[string]string{"command": "login"})
	if err != nil {
		return ErrorConnection, result, err
	}
	return jsonError(result), result, nil
}

// Calc generates a radio code. radioModel may be a string or *RadioModel.
func (c *Client) Calc(ctx context.Context, radioModel any, serial, extra string) (int, map[string]any, error) {
	result, err := c.postRequest(ctx, map[string]string{
		"command":     "calc",
		"radio_model": modelName(radioModel),
		"serial":      serial,
		"extra":       extra,
	})
	if err != nil {
		return ErrorConnection, result, err
	}
	return jsonError(result), result, nil
}

// Info returns calculator parameters for one model.
func (c *Client) Info(ctx context.Context, radioModel any) (int, *RadioModel, error) {
	name := modelName(radioModel)
	result, err := c.postRequest(ctx, map[string]string{
		"command":     "info",
		"radio_model": name,
	})
	if err != nil {
		return ErrorConnection, nil, err
	}
	code := jsonError(result)
	if code != ErrorSuccess {
		return code, nil, nil
	}
	return code, modelFromResult(name, result), nil
}

// List returns every supported calculator.
func (c *Client) List(ctx context.Context) (int, []*RadioModel, error) {
	result, err := c.postRequest(ctx, map[string]string{"command": "list"})
	if err != nil {
		return ErrorConnection, nil, err
	}
	code := jsonError(result)
	if code != ErrorSuccess {
		return code, nil, nil
	}
	raw, _ := result["supportedRadioModels"].(map[string]any)
	var models []*RadioModel
	for name, item := range raw {
		entry, _ := item.(map[string]any)
		models = append(models, modelFromEntry(name, entry))
	}
	return code, models, nil
}

func modelFromResult(name string, result map[string]any) *RadioModel {
	return NewRadioModel(
		name,
		jsonInt(result["serialMaxLen"]),
		result["serialRegexPattern"],
		jsonInt(result["extraMaxLen"]),
		result["extraRegexPattern"],
	)
}

func modelFromEntry(name string, entry map[string]any) *RadioModel {
	if entry == nil {
		return NewRadioModel(name, 0, "", 0, nil)
	}
	return NewRadioModel(
		name,
		jsonInt(entry["serialMaxLen"]),
		entry["serialRegexPattern"],
		jsonInt(entry["extraMaxLen"]),
		entry["extraRegexPattern"],
	)
}

func (c *Client) postRequest(ctx context.Context, params map[string]string) (map[string]any, error) {
	if c.APIKey != "" {
		params["key"] = c.APIKey
	}

	body, err := c.postMultipart(ctx, params)
	if err != nil {
		return map[string]any{"error": float64(ErrorConnection)}, err
	}
	if len(body) == 0 {
		return map[string]any{"error": float64(ErrorConnection)}, fmt.Errorf("empty API response")
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil || result == nil {
		return map[string]any{"error": float64(ErrorConnection)}, err
	}
	return result, nil
}

func (c *Client) postMultipart(ctx context.Context, fields map[string]string) ([]byte, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	for k, v := range fields {
		if err := w.WriteField(k, v); err != nil {
			return nil, err
		}
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	url := c.APIURL
	if url == "" {
		url = APIURL
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func jsonError(result map[string]any) int {
	if result == nil {
		return ErrorConnection
	}
	return jsonInt(result["error"])
}

func jsonInt(v any) int {
	switch n := v.(type) {
	case float64:
		return int(n)
	case int:
		return n
	case json.Number:
		i, _ := n.Int64()
		return int(i)
	default:
		return 0
	}
}
