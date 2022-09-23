package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "http://example.com:80/regist/20060102",
		httpmock.NewStringResponder(200, `{"created_number":5}`),
	)
	m.Run()
}

func TestClient_PostFile(t *testing.T) {
	type fields struct {
		Scheme string
		Host   string
		Port   string
	}
	type args struct {
		ctx      context.Context
		endPoint string
		filePath string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantResBody    []byte
		wantStatusCode int
		wantErr        bool
	}{
		{
			name:           "ok",
			fields:         fields{Scheme: "http", Host: "example.com", Port: "80"},
			args:           args{ctx: context.Background(), endPoint: "/regist/20060102", filePath: "../../test/20060102.csv"},
			wantResBody:    []byte(`{"created_number":5}`),
			wantStatusCode: 200,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Scheme: tt.fields.Scheme,
				Host:   tt.fields.Host,
				Port:   tt.fields.Port,
			}
			gotResBody, gotStatusCode, err := c.PostFile(tt.args.ctx, tt.args.endPoint, tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.PostFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResBody, tt.wantResBody) {
				t.Errorf("Client.PostFile() gotResBody = %v, want %v", gotResBody, tt.wantResBody)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("Client.PostFile() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
	}
}
