package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is the ccpctl version
const Version = "v0.1"

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "version",
		Short: "prints the ccpctl version",
		Long:  "prints the ccpctl version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(Version)
			return nil
		},
	}
	return cmd
}
