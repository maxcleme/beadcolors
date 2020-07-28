package main

import (
	"reflect"
	"testing"
)

func TestV3(t *testing.T) {
	type args struct {
		r     int
		g     int
		b     int
		index int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				r:     171,
				g:     38,
				b:     144,
				index: 0,
			},
			want:    []string{"!", "171", "38", "144", "312.18", "0.64", "0.41", "41.31", "62.39", "-27.65"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := V3(tt.args.r, tt.args.g, tt.args.b, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("V3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("V3() got = %v, want %v", got, tt.want)
			}
		})
	}
}
