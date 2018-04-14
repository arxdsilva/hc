package main

import (
	"encoding/csv"
	"reflect"
	"testing"
)

func Test_readCSV(t *testing.T) {
	type args struct {
		r csv.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantRec [][]string
		wantErr bool
	}{
	// TODO: Add test cases.
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
