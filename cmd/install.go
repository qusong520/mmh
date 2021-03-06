package cmd

import (
	"github.com/mritd/mmh/utils"
	"github.com/spf13/cobra"
)

var installDir string
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install",
	Long: `
install.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Install(installDir)
	},
}

func init() {
	installCmd.PersistentFlags().StringVar(&installDir, "dir", "/usr/local/bin", "install dir")
	RootCmd.AddCommand(installCmd)
}
