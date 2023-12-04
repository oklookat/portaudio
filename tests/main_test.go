package tests

import (
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	cmd := exec.Command("python", "../bootstrap.py")
	if err := cmd.Run(); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}
