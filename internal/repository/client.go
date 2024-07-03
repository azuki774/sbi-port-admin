package repository

import (
	"context"
	"io"
	"net"
	"net/http"
	"os"
)

type Client struct {
	Scheme string // http, https
	Host   string
	Port   string
}

func (c *Client) PostFile(ctx context.Context, endPoint string, filePath string) (resBody []byte, statusCode int, err error) {
	url := c.Scheme + "://" + net.JoinHostPort(c.Host, c.Port) + endPoint
	content, err := os.Open(filePath)
	if err != nil {
		return []byte{}, 0, err
	}

	req, err := http.NewRequest("POST", url, content)
	if err != nil {
		return []byte{}, 0, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, 0, err
	}
	defer res.Body.Close()

	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, 0, err
	}

	return resBody, res.StatusCode, nil
}
