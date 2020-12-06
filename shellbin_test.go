package main

import "testing"

func TestBinCreation(t *testing.T) {
	_, err := makeHTTPRequest("ShellBin Testing")
	if err != nil {
		t.Errorf("Error creating a bin: %v\n", err)
	}
}
