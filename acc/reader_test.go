package main

import (
	"encoding/csv"
	"reflect"
	"strings"
	"testing"
)

func Test_readCSV(t *testing.T) {
	accs := `
	1,1000
	2,2000
	`
	type args struct {
		r *csv.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantRec [][]string
		wantErr bool
	}{
		{
			name: "reads simple fields",
			args: args{
				r: csv.NewReader(strings.NewReader(accs)),
			},
			wantRec: [][]string{
				[]string{"1", "1000"},
				[]string{"2", "2000"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRec, err := readCSV(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRec, tt.wantRec) {
				t.Errorf("readCSV() = %v, want %v", gotRec, tt.wantRec)
			}
		})
	}
}
