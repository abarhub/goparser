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

func TestChecker(t *testing.T) {
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
