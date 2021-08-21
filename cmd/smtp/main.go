package main

import (
	"fmt"
	"github.com/emersion/go-smtp"
	"os"
	"time"

	"github.com/shauncampbell/unified-smtp/internal/authenticator"
	"github.com/shauncampbell/unified-smtp/internal/config"
	"github.com/shauncampbell/unified-smtp/internal/unified"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "smtp",
	RunE: runApplication,
}

const configurationFileFlag string = "configuration-file"

func runApplication(cmd *cobra.Command, args []string) error {
	// Figure out the config file path
	cfgFile, err := cmd.PersistentFlags().GetString(configurationFileFlag)
	if err != nil {
		return fmt.Errorf("must provide configuration file: %w", err)
	}
	// Read in the config file
	cfg, err := config.ReadConfigFromFile(cfgFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Create a new memory authenticator
	auth, err := authenticator.ForConfiguration(&cfg.Authenticator)
	if err != nil {
		return fmt.Errorf("failed to load authenticator: %s", err.Error())
	}
	// Create a memory backend
	be := unified.New(auth)

	// Create a new server
	s := smtp.NewServer(be)
	s.Addr = ":1125"
	s.Domain = "localhost"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Info().Msgf("Starting SMTP server at localhost:1125")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal().Msgf("%s", err.Error())
	}
	return nil
}

func main() {
	rootCmd.PersistentFlags().StringP(configurationFileFlag, "f", "", "path to a configuration file")
	err := rootCmd.MarkPersistentFlagRequired(configurationFileFlag)
	if err != nil {
		log.Error().Msgf("failed to set up configuration file flag")
		os.Exit(1)
	}

	rootCmd.Flags().IntP("port", "p", 1143, "port to run smtp server on")

	if err = rootCmd.Execute(); err != nil {
		log.Error().Msgf("unable to run application: %s", err.Error())
		if _, e := fmt.Fprintf(os.Stderr, "unable to run application: %s\n", err.Error()); e != nil {
			log.Error().Msgf("unable to write to stderr: %s", e.Error())
			log.Error().Msgf("unable to run application: %s", err.Error())
		}
		os.Exit(1)
	}
	os.Exit(0)
}
