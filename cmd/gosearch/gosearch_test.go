package main

import (
	"GoSearch/pkg/crawler"
	"testing"
)

func Test_printUrls(t *testing.T) {
	type args struct {
		urls     []crawler.Document
		flagname string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printUrls(tt.args.urls, tt.args.flagname)
		})
	}
}
