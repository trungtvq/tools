package test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("test main")
	os.Exit(m.Run())
}
