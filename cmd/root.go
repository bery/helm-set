package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "set",
	Short: "Use environment variables as helm values",
	Long:  `Use environment variables as helm values by taking all variables starting with HELM_VAR_ to --set values.`,
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
}

var Mode string
var DryRun bool

var Debug bool
var Verbose bool
var Prefix string = "HELM_VAR_"

func main() {
	Execute()
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "enable verbose output")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "enable verbose output")
	rootCmd.PersistentFlags().BoolVar(&DryRun, "dry-run", false, "parameter parsing mode (rename, copy)")

	log.SetFormatter(&log.TextFormatter{DisableTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	if Debug {
		log.SetLevel(log.DebugLevel)
	} else if Verbose {
		log.SetLevel(log.InfoLevel)
	}

	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(upgradeCmd)

	rootCmd.Execute()
}
