package repository

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

type Client struct {
	Scheme string // http, https
	Host   string
	Port   string
}

func (c *Client) PostFile(ctx context.Context, endPoint string, filePath string) (resBody string, statusCode int, err error) {
	url := c.Scheme + "://" + net.JoinHostPort(c.Host, c.Port) + endPoint
	content, err := os.Open(filePath)
	if err != nil {
		return "", 0, err
	}

	req, err := http.NewRequest("POST", url, content)
	if err != nil {
		return "", 0, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer res.Body.Close()

	resBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", 0, err
	}

	return string(resBodyBytes), res.StatusCode, nil
}
