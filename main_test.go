package main

import (
	"fmt"
	"golang.org/x/tools/go/analysis"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		t *testing.T
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Test(tt.args.t)
		})
	}
}

func TestRun(t *testing.T) {
	type args struct {
		pass *analysis.Pass
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name:    "TestRun(t *testing.T)",
			args:    args{},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Print("Running TestRun(t *testing.T)")
			//testdata := analysistest.TestData()
			//analysistest.Run(t, testdata, a, "a")
			AnalyzerTest(t)

			got, err := Run(tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}
