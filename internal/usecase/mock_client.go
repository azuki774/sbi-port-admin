package usecase

import "context"

type MockHTTPClient struct {
	ResBody    []byte
	StatusCode int
	Err        error
}

func (m *MockHTTPClient) PostFile(ctx context.Context, endPoint string, filePath string) (resBody []byte, statusCode int, err error) {
	return m.ResBody, m.StatusCode, m.Err
}
