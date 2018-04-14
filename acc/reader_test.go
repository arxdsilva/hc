package main

import (
	"reflect"
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
