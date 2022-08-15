package main

import (
	"fmt"
	"testing"
)

func runChecker(filename string) error {
	functionList, err := Parser(filename)
	if err != nil {
		return err
	} else if len(functionList) != 1 {
		return fmt.Errorf("no function find in file")
	}
	return Checker(functionList)
}

func TestCheckerOK(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test1", args: args{filename: "../test_data/sementic/ok/test1.txt"}, wantErr: false},
		{name: "test2", args: args{filename: "../test_data/sementic/ok/test2.txt"}, wantErr: false},
		{name: "test3", args: args{filename: "../test_data/sementic/ok/test3.txt"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runChecker(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("Checker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckerKO(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		err  string
	}{
		{name: "test1", args: args{filename: "../test_data/sementic/ko/test1.txt"},
			err: "error: operator '2' need int parameters (line:2, column:7)"},
		{name: "test2", args: args{filename: "../test_data/sementic/ko/test2.txt"},
			err: "error: operator '2' need int parameters (line:3, column:7)"},
		{name: "test3", args: args{filename: "../test_data/sementic/ko/test3.txt"},
			err: "error: operator '3' need int parameters (line:4, column:11)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runChecker(tt.args.filename); err == nil || err.Error() != tt.err {
				t.Errorf("Checker() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
