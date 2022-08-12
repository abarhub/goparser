package main

import (
	"testing"
)

func TestParserOK(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test1.txt", args: args{filename: "../test_data/ok/test1.txt"}, wantErr: false},
		{name: "test2.txt", args: args{filename: "../test_data/ok/test2.txt"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var listener calcListener
			if err := Parser(tt.args.filename, listener); (err != nil) != tt.wantErr {
				t.Errorf("Parser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParserKO(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name     string
		args     args
		msgError string
	}{
		{name: "test1.txt", args: args{filename: "../test_data/ko/test1.txt"},
			msgError: "error: line 2:6 token recognition error at: '!'"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var listener calcListener
			if err := Parser(tt.args.filename, listener); err == nil || err.Error() != tt.msgError {
				t.Errorf("Parser() error = %v, msgError %v!= error %v", err, tt.msgError, err.Error())
			}
		})
	}
}