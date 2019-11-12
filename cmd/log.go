package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"strings"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "log at 10000",
	Run:   logRun,
}

func logRun(cmd *cobra.Command, args []string) {
	fmt.Println("local log running")
	fmt.Println(strings.TrimRight("/", "/"))
	http.Handle("/", http.FileServer(http.Dir("/data/logs")))
	http.ListenAndServe(":10000", nil)
}
