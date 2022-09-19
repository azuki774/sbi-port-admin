package model

import "testing"

func TestValidateDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{date: "20060102"},
			wantErr: false,
		},
		{
			name:    "not existed date",
			args:    args{date: "20060431"},
			wantErr: true,
		},
		{
			name:    "not suitable format",
			args:    args{date: "2006-01-02"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateDate(tt.args.date); (err != nil) != tt.wantErr {
				t.Errorf("ValidateDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
