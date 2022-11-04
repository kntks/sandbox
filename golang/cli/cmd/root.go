/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/rs/zerolog/log"

	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Caller().Logger()
}

var cfgFile string

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "A brief description of your application",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "", "config file")
	return cmd
}

var rootCmd = NewRootCommand()

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
