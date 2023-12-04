package tests

import "testing"

func chk(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
