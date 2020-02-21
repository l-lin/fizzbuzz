package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/l-lin/fizzbuzz/web"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Launch web server to perform fizz-buzz operation",
		Run:   runServe,
	}
	port int
)

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 3000, "the port to start the web server on")
}

func runServe(cmd *cobra.Command, args []string) {
	r := web.NewRouter()
	log.Printf("Server started on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
