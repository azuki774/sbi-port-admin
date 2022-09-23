package usecase

import (
	"context"
	"errors"
	"testing"

	"go.uber.org/zap"
)

var l *zap.Logger

func TestMain(m *testing.M) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	l, _ = config.Build()

	l.WithOptions(zap.AddStacktrace(zap.ErrorLevel))

	m.Run()
}

func TestRegistFileInformation_FillDateByFileName(t *testing.T) {
	type fields struct {
		FilePath string
		Date     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		want    string
	}{
		{
			name:    "/usr/tmp/20010102.csv",
			fields:  fields{FilePath: "/usr/tmp/20010102.csv"},
			wantErr: false,
			want:    "20010102",
		},
		{
			name:    "/usr/tmp/2001-01-02.csv",
			fields:  fields{FilePath: "/usr/tmp/2001-01-02.csv"},
			wantErr: false,
			want:    "20010102",
		},
		{
			name:    "/usr/tmp/2001-01-02",
			fields:  fields{FilePath: "/usr/tmp/2001-01-02"},
			wantErr: false,
			want:    "20010102",
		},
		{
			name:    "2001-01-02.csv",
			fields:  fields{FilePath: "2001-01-02.csv"},
			wantErr: false,
			want:    "20010102",
		},
		{
			name:    "20010102.csv",
			fields:  fields{FilePath: "2001-01-02.csv"},
			wantErr: false,
			want:    "20010102",
		},
		{
			name:    "",
			fields:  fields{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RegistFileInformation{
				FilePath: tt.fields.FilePath,
				Date:     tt.fields.Date,
			}
			if err := r.FillDateByFileName(); (err != nil) != tt.wantErr {
				t.Errorf("RegistFileInformation.FillDateByFileName() error = %v, wantErr %v", err, tt.wantErr)
			}

			if (!tt.wantErr) && (tt.want != r.Date) {
				t.Errorf("RegistFileInformation.FillDateByFileName(): date tt.fields.Date = %v, wantErr %v", r.Date, tt.want)
			}
		})
	}
}

func TestUsecaseClient_RegistJob(t *testing.T) {
	type fields struct {
		Logger     *zap.Logger
		HTTPClient HTTPClient
	}
	type args struct {
		ctx      context.Context
		filePath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Logger: l,
				HTTPClient: &MockHTTPClient{
					ResBody:    []byte(`{"created_number":5}`),
					StatusCode: 200,
					Err:        nil,
				},
			},
			args: args{
				ctx:      context.Background(),
				filePath: "/usr/date/20060102.csv",
			},
			wantErr: false,
		},
		{
			name: "status code 500",
			fields: fields{
				Logger: l,
				HTTPClient: &MockHTTPClient{
					ResBody:    []byte(``),
					StatusCode: 500,
					Err:        nil,
				},
			},
			args: args{
				ctx:      context.Background(),
				filePath: "/usr/date/20060102.csv",
			},
			wantErr: true,
		},
		{
			name: "send error",
			fields: fields{
				Logger: l,
				HTTPClient: &MockHTTPClient{
					Err: errors.New("error"),
				},
			},
			args: args{
				ctx:      context.Background(),
				filePath: "/usr/date/20060102.csv",
			},
			wantErr: true,
		},
		{
			name: "unknown file name",
			fields: fields{
				Logger: l,
				HTTPClient: &MockHTTPClient{
					ResBody:    []byte(`{"created_number":5}`),
					StatusCode: 200,
					Err:        nil,
				},
			},
			args: args{
				ctx:      context.Background(),
				filePath: "/usr/date/200601023.csv",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsecaseClient{
				Logger:     tt.fields.Logger,
				HTTPClient: tt.fields.HTTPClient,
			}
			if err := u.RegistJob(tt.args.ctx, tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("UsecaseClient.RegistJob() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
