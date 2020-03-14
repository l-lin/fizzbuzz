package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/l-lin/fizzbuzz/stats"
	"github.com/l-lin/fizzbuzz/web"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Launch web server to perform fizz-buzz operation",
		Run:   runServe,
	}
	port             int
	statsStorageMode string
)

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 3000, "the port to start the web server on")
	serveCmd.Flags().StringVarP(&statsStorageMode, "stats-storage-mode", "m", strings.ToLower(stats.Memory.String()), fmt.Sprintf("the mode to store request stats (available modes: %v)", stats.ModesToSlice()))
}

func runServe(cmd *cobra.Command, args []string) {
	r := web.NewRouter(statsStorageMode)
	log.Printf("Server started on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
