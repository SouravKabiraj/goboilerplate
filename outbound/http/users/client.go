package users

import (
	"context"
	"fmt"
	"net/http"
	httpclient "sourav.kabiraj/goboilerplate/outbound/http"
)

type Client struct {
	baseURL    string
	httpClient httpclient.Client
}

func NewClient(client httpclient.Client) Client {
	return Client{
		baseURL:    "https://api.kabiraj.com/users",
		httpClient: client,
	}
}

func (c Client) GetUserByID(ctx context.Context, id string) string {
	var resp *http.Response
	var err error
	var headers http.Header
	resp, err = c.httpClient.Get(ctx, "", headers)
	fmt.Println(resp, err)
	return ""
}
