package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	URL string
}

func NewClient(url string) Client {
	return Client{URL: url}
}

func (c *Client) makeURL(path string) string {
	return fmt.Sprintf("%s/%s", c.URL, path)
}

func (c *Client) GetData(path string, data string, method string) (string, error) {
	req, err := http.NewRequestWithContext(
		context.Background(),
		strings.ToUpper(method),
		c.makeURL(path),
		bytes.NewBuffer([]byte(data)),
	)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return PrintResponse(resp)
}

func PrintResponse(resp *http.Response) (string, error) {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", resp.Status, string(bodyBytes)), nil
}
