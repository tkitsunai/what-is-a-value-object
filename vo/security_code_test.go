package vo

import (
	"reflect"
	"testing"
)

func TestNewSecurityCode(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name    string
		args    args
		want    SecurityCode
		wantErr bool
	}{
		{
			name:    "1300未満は証券コードではない",
			args:    args{value: 1000},
			want:    SecurityCode{value: 0},
			wantErr: true,
		},
		{
			name:    "1300は証券コードとして扱える",
			args:    args{value: 1300},
			want:    SecurityCode{value: 1300},
			wantErr: false,
		},
		{
			name:    "9999は証券コードとして扱える",
			args:    args{value: 9999},
			want:    SecurityCode{value: 9999},
			wantErr: false,
		},
		{
			name:    "10000は証券コードではない",
			args:    args{value: 10000},
			want:    SecurityCode{value: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSecurityCode(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSecurityCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSecurityCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
