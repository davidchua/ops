package main

import (
	"fmt"
	"log"
	"os"

	"embed"

	"github.com/davidchua/ops/cmd"
	"github.com/spf13/viper"
)

//go:embed templates
var tmpl embed.FS

func main() {
	viper.SetConfigName(".ops")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error establishing user's home directory: %#v", err)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			_, err := os.Create(fmt.Sprintf("%s/.ops.yaml", dirname))
			if err != nil {
				log.Fatalf("error creating peeaao config file: %#v", err)
			}
			viper.ReadInConfig()
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}
	cmd.Execute(tmpl)
}
