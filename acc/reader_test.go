package main

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func Test_mapAccs(t *testing.T) {
	type args struct {
		accs [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    map[int]int
		wantErr bool
	}{
		{
			name: "simple map from slice of slices",
			args: args{
				accs: [][]string{
					[]string{"1", "2"},
					[]string{"2", "2"},
				},
			},
			want: map[int]int{
				1: 2,
				2: 2,
			},
			wantErr: false,
		},
		{
			name: "simple map from slice of slices with negative numbers",
			args: args{
				accs: [][]string{
					[]string{"1", "-2"},
					[]string{"2", "2"},
				},
			},
			want: map[int]int{
				1: -2,
				2: 2,
			},
			wantErr: false,
		},
		{
			name: "simple map from slice of slices with parse error",
			args: args{
				accs: [][]string{[]string{"a", "-2"}},
			},
			want:    map[int]int{},
			wantErr: true,
		},
		{
			name: "simple map from slice of slices with parse error in second parameter",
			args: args{
				accs: [][]string{[]string{"1", "a"}},
			},
			want:    map[int]int{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapAccs(tt.args.accs)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapAccs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && len(got) > 0 {
				t.Errorf("mapAccs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateAccountsFromRaw(t *testing.T) {
	wd, _ := os.Getwd()
	type args struct {
		file string
	}
	tests := []struct {
		name     string
		args     args
		wantAccs map[int]int
		wantErr  bool
	}{
		{
			name: "sample file test",
			args: args{
				file: filepath.Join(wd, "contas.csv"),
			},
			wantAccs: map[int]int{
				1: 2,
				2: 3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAccs, err := generateAccountsFromRaw(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateAccountsFromRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAccs, tt.wantAccs) {
				t.Errorf("generateAccountsFromRaw() = %v, want %v", gotAccs, tt.wantAccs)
			}
		})
	}
}

func Test_calcBalancesAfterTransactions(t *testing.T) {
	type args struct {
		file string
		accs map[int]int
	}
	tests := []struct {
		name    string
		args    args
		want    map[int]int
		wantErr bool
	}{
		{
			name: "calculate positive deposits",
			args: args{
				file: "transacoes.csv",
				accs: map[int]int{
					1: 1000,
					2: 2000,
				},
			},
			want: map[int]int{
				1: 3000,
				2: 5000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calcBalancesAfterTransactions(tt.args.file, tt.args.accs)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcBalancesAfterTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcBalancesAfterTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateBalance(t *testing.T) {
	positiveTransacToAccs := `1,2000
2,3000`
	type args struct {
		accs map[int]int
		r    *csv.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantM   map[int]int
		wantErr bool
	}{
		{
			name: "positiveTransacToAccs",
			args: args{
				accs: map[int]int{
					1: 1000,
					2: 2000,
				},
				r: csv.NewReader(strings.NewReader(positiveTransacToAccs)),
			},
			wantM: map[int]int{
				1: 3000,
				2: 5000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, err := calculateBalance(tt.args.accs, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("calculateBalance() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}
