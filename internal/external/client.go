package external

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NoBypass/fds/internal/backend/tracing"
	"github.com/NoBypass/mincache"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go/ext"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type Client struct {
	tracing.Tracable

	cache         *mincache.Cache
	baseURL, name string
}

type Response struct {
	Body   any
	Header *http.Header
}

func NewClient(cache *mincache.Cache, baseURL, name string) *Client {
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}
	return &Client{
		cache:    cache,
		baseURL:  baseURL,
		name:     name,
		Tracable: tracing.NewTracable(name),
	}
}

func (c *Client) Request(ctx context.Context, url string, lifetime time.Duration, decode any) (*http.Header, error) {
	sp, ctx := c.StartSpan(ctx, url)
	defer sp.Finish()

	if reflect.TypeOf(decode).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("decode must be a pointer")
	}

	val, ok := c.cache.Get(fmt.Sprintf("%s:%s", c.baseURL, url))
	if ok {
		reflect.ValueOf(decode).Elem().Set(reflect.ValueOf(val.(*Response).Body).Elem())
		return val.(*Response).Header, nil
	}

	fullURL := fmt.Sprintf("%s%s", c.baseURL, url)
	sp.LogKV("url", fullURL)

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", c.name, err)
	}

	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusNotFound:
			return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s: not found", c.name))
		}

		var body string
		if resp.Body != nil {
			bodyBytes, _ := io.ReadAll(resp.Body)
			body = string(bodyBytes)
		}
		msg := fmt.Errorf("%s | body: %s", resp.Status, body)
		ext.LogError(sp, msg)
		return nil, msg
	}

	err = json.NewDecoder(resp.Body).Decode(decode)
	if err != nil {
		return nil, err
	}

	r := &Response{
		Body:   decode,
		Header: &resp.Header,
	}

	body, _ := json.Marshal(decode)
	header, _ := json.Marshal(r.Header)
	sp.LogKV("response", string(body))
	sp.LogKV("header", string(header))
	c.cache.Set(fmt.Sprintf("%s:%s", c.baseURL, url), r, lifetime)
	return r.Header, nil
}
