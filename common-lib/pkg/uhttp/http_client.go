package uhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Get[T any](url string, query url.Values) (*T, error) {
	if query != nil {
		url += "?" + query.Encode()
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return ReadResponse[T](resp)
}

func Post[T any](url string, body any) (*T, error) {
	req, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, MIMEJSON, bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}
	return ReadResponse[T](resp)
}

func ReadResponse[T any](resp *http.Response) (*T, error) {
	if resp == nil {
		return nil, errors.New("nil response")
	}
	defer GracefulClose(resp)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := new(T)
	ct := resp.Header.Get(HeaderContentType)
	switch {
	case strings.Contains(ct, MIMEJSON):
		if err = json.Unmarshal(res, data); err != nil {
			return nil, err
		}
	case strings.Contains(ct, MIMEXML):
		if err = xml.Unmarshal(res, data); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported content type")
	}
	return data, nil
}

func GracefulClose(resp *http.Response) {
	if resp == nil || resp.Body == nil {
		return
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
}

type HttpClient[T any] struct {
	Endpoint string
	Method   string
	Header   http.Header
	Query    url.Values
	Body     []byte
	Timeout  time.Duration
}

func NewHttpClient[T any](endpoint string, method string, options ...HttpClientOption[T]) *HttpClient[T] {
	return &HttpClient[T]{
		Endpoint: endpoint,
		Method:   method,
		Header:   make(http.Header),
		Query:    make(url.Values),
		Body:     []byte{},
		Timeout:  10 * time.Second,
	}
}

func (clt *HttpClient[T]) Do(ctx context.Context) (*T, error) {
	if len(clt.Query) > 0 {
		clt.Endpoint += "?" + clt.Query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, clt.Method, clt.Endpoint, bytes.NewBuffer(clt.Body))
	if err != nil {
		return nil, err
	}
	if len(clt.Header) > 0 {
		req.Header = clt.Header
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return ReadResponse[T](resp)
}

type HttpClientOption[T any] func(c *HttpClient[T])

func WithTimeout[T any](timeout time.Duration) HttpClientOption[T] {
	return func(c *HttpClient[T]) {
		c.Timeout = timeout
	}
}

func WithBody[T any](data any) HttpClientOption[T] {
	return func(c *HttpClient[T]) {

	}
}
