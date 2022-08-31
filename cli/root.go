package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:     "cloud-station-cli",
		Long:    "cloud-station-cli 云中转站Long",
		Short:   "cloud-station-cli 云中转站Short",
		Example: "cloud-station-cli cmds",
		RunE: func(cmd *cobra.Command, args []string) error {
			if version {
				fmt.Println("cloud-station-cli v0.0.1")
			}
			return nil
		},
	}
	version = false
)

func init() {
	f := RootCmd.PersistentFlags()
	f.BoolVarP(&version, "version", "v", false, "当前CLI的版本信息")
}
