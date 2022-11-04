package cmd

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewEx1Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ex1",
		Short: "SetConfigFile()にfile pathを直接渡す",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			viper.SetConfigFile(cfgFile)
			if err := viper.ReadInConfig(); err != nil {
				log.Error().Err(err).Send()
				return
			}
			fmt.Println(viper.Get("env"))
			fmt.Println(viper.Get("data"))
		},
	}
	return cmd
}

func init() {
	ex1 := NewEx1Command()
	rootCmd.AddCommand(ex1)
}
