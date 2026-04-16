package main_test

import (
	_ "embed"
	"io/fs"
	"os"
	"testing"
)

//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
