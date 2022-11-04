package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewEx2Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ex2",
		Short: "os.Openした後、ReadConfig()を実行する",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			viper.SetConfigType("yaml")
			f, err := os.Open(cfgFile)
			if err != nil {
				log.Error().Err(err).Send()
				return
			}
			defer f.Close()
			if err := viper.ReadConfig(f); err != nil {
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
	ex2 := NewEx2Command()
	rootCmd.AddCommand(ex2)
}
