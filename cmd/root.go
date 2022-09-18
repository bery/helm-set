package cmd

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	Execute()
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "set",
	Short: "Use envrironment variables as helm values",
	Long:  `Use envrironment variables as helm values by taking all variables starting with HELM_VAR_ to --set values.`,
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
}

// ④ Now use the FooMode enum flag. If you want a non-zero default, then
// simply set it here, such as in "foomode = Bar".
var Mode string

var Debug bool
var Verbose bool
var Prefix string = "HELM_VAR_"

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Cleanup OpenStack environment for compatibility with the golang client
	for _, v := range []string{
		"OS_IDENTITY_PROVIDER", "OS_AUTH_TYPE", "OS_MUTUAL_AUTH", "OS_PROTOCOL"} {
		os.Unsetenv(v)
	}
	os.Setenv("OS_AUTH_URL", strings.Replace(os.Getenv("OS_AUTH_URL"), "krb/", "", 1))
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "", false, "enable verbose output")
	RootCmd.PersistentFlags().StringVarP(&Mode, "mode", "", "rename", "enable verbose output")

	log.SetFormatter(&log.TextFormatter{DisableTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	if Debug {
		log.SetLevel(log.DebugLevel)
	} else if Verbose {
		log.SetLevel(log.InfoLevel)
	}

	RootCmd.AddCommand(installCmd)
	RootCmd.AddCommand(upgradeCmd)
}
