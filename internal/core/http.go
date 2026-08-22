package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is a small HTTP wrapper with retries and timeouts.
type Client struct {
	base string
	token string
	timeout time.Duration
	retries int
	client *http.Client
}

// NewClient builds an HTTP client.
func NewClient(base, token string, timeout time.Duration, retries int) *Client {
	return &Client{
 base: base,
 token: token,
 timeout: timeout,
 retries: retries,
 client: &http.Client{Timeout: timeout},
	}
}

func (c *Client) request(ctx context.Context, method, path string, body io.Reader, out any) error {
	var lastErr error
	for attempt := 0; attempt <= c.retries; attempt++ {
 req, err := http.NewRequestWithContext(ctx, method, c.base+path, body)
 if err != nil {
 return err
 }
 req.Header.Set("Content-Type", "application/json")
 if c.token != "" {
 req.Header.Set("Authorization", "Bearer "+c.token)
 }
 resp, err := c.client.Do(req)
 if err != nil {
 lastErr = err
 time.Sleep(time.Duration(attempt+1) * 200 * time.Millisecond)
 continue
 }
 defer resp.Body.Close()
 if resp.StatusCode >= 400 {
 lastErr = fmt.Errorf("http %d", resp.StatusCode)
 time.Sleep(time.Duration(attempt+1) * 200 * time.Millisecond)
 continue
 }
 if out == nil {
 return nil
 }
 return json.NewDecoder(resp.Body).Decode(out)
	}
	return lastErr
}

func (c *Client) Get(ctx context.Context, path string, out any) error {
	return c.request(ctx, http.MethodGet, path, nil, out)
}

func (c *Client) Post(ctx context.Context, path string, payload, out any) error {
	body, err := json.Marshal(payload)
	if err != nil {
 return err
	}
	return c.request(ctx, http.MethodPost, path, bytesReader(body), out)
}