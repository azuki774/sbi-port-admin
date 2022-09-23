package usecase

import "testing"

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
