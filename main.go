package main

import (
	"github.com/trungtvq/go-utils/cmd"
)

var revision string

func main() {
	cmd.SetRevision(revision)
	cmd.Execute()
}
