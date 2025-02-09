package http

import (
	"context"
	"io"
	"net/http"
)

type Client interface {
	Get(ctx context.Context, url string, header http.Header) (*http.Response, error)
	Post(ctx context.Context, url string, header http.Header, body io.Reader) (*http.Response, error)
}

type ClientImpl struct {
	client *http.Client
}

func (h *ClientImpl) Get(ctx context.Context, url string, header http.Header) (*http.Response, error) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	return h.client.Do(request)
}

func (h *ClientImpl) Post(ctx context.Context, url string, header http.Header, body io.Reader) (*http.Response, error) {
	request, _ := http.NewRequest(http.MethodPost, url, body)
	request.Header = header
	return h.client.Do(request)
}
