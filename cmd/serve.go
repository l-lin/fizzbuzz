package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/l-lin/fizzbuzz/stats"
	"github.com/l-lin/fizzbuzz/stats/kafka"
	"github.com/l-lin/fizzbuzz/stats/memory"
	"github.com/l-lin/fizzbuzz/web"
	"github.com/rs/zerolog/log"
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
	brokers          string
	topic            string
)

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 3000, "the port to start the web server on")
	serveCmd.Flags().StringVarP(&statsStorageMode, "stats-storage-mode", "m", strings.ToLower(stats.Memory.String()), fmt.Sprintf("the mode to store request stats (available modes: %v)", stats.ModesToSlice()))
	serveCmd.Flags().StringVarP(&brokers, "brokers", "b", "localhost:19092,localhost:29092,localhost:39092", "bootstrap broker(s) (host:port separated by a comma)")
	serveCmd.Flags().StringVarP(&topic, "topic", "t", "fizzbuzz-request-stats", "topic to consume from and produce to")
}

func runServe(cmd *cobra.Command, args []string) {
	setupLogger()
	m := stats.GetMode(statsStorageMode)
	repo := getStatsRepository(m)
	defer repo.Close()
	r := web.NewRouter(repo)
	log.Info().Msgf("Server started on port %d", port)
	log.Fatal().Err(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func getStatsRepository(m stats.Mode) stats.Repository {
	if stats.Kafka == m {
		log.Debug().Str("brokers", brokers).Str("topic", topic).Msg("Connecting to kafka brokers")
		repo := kafka.NewRepository(brokers, topic)
		go repo.Listen()
		return repo
	}
	log.Warn().Msg("Do not use this mode if you are running the application in cluster")
	return memory.NewRepository()
}
