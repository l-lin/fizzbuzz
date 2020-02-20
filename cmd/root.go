package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fizzbuzz",
	Short: "Fizz-buzz exercise",
	Long:  `The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz". The output would look like this: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...".`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version, buildDate string) {
	rootCmd.Version = func(version, buildDate string) string {
		res, err := json.Marshal(cliBuild{Version: version, BuildDate: buildDate})
		if err != nil {
			log.Fatal(err)
		}
		return string(res)
	}(version, buildDate)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(`{{printf "%s" .Version}}`)
}

type cliBuild struct {
	Version   string `json:"version"`
	BuildDate string `json:"buildDate"`
}
