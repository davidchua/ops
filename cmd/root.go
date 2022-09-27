package cmd

import (
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"embed"
)

var (
	tmpl embed.FS
	// Used for flags.
	clog     *logrus.Logger
	loglevel string

	rootCmd = &cobra.Command{
		Use:   "ops",
		Short: "David's Ops Toolkit",
		Long:  "David's Everything Ops Toolkit",
		//	Run: func(cmd *cobra.Command, args []string) {
		//		loglevelstr, err := logrus.ParseLevel(loglevel)
		//		if err != nil {
		//			clog.Fatalf("unable to parse loglevel: %#v", err)
		//		}

		//		clog.SetLevel(loglevelstr)

		//	},
	}
)

func Execute(templates embed.FS) error {
	tmpl = templates
	return rootCmd.Execute()
}

func init() {

	rootCmd.PersistentFlags().StringVar(&loglevel, "loglevel", "info", "log level, allowed [debug, info, warn, error, fatal]")

	// initiate logging
	clog = logrus.New()
	clog.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

}
