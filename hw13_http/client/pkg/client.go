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
	URL  string
	Path string
}

func NewClient(url string, path string) Client {
	return Client{URL: url, Path: path}
}

func (c *Client) makeURL() string {
	return fmt.Sprintf("%s/%s", c.URL, c.Path)
}

func (c *Client) GetData() (string, error) {
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		c.makeURL(),
		strings.NewReader(""))
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

func (c *Client) PostData(msg string) (string, error) {
	data := []byte(msg)
	body := bytes.NewReader(data)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, c.makeURL(), body)
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
