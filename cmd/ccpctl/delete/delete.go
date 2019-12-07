package delete

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swiftdiaries/ccp-clientlibrary-go/v3/pkg/delete"
)

type flagpole struct {
	Name string
}

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	flags := &flagpole{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "delete",
		Short: "delete a CCP cluster with name",
		Long:  "delete a CCP cluster with name",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags, cmd, args)
		},
	}
	cmd.Flags().StringVarP(&flags.Name, "name", "n", "", "cluster name for deletion")
	return cmd
}

func runE(flags *flagpole, cmd *cobra.Command, args []string) error {
	log.Infof("delete called..")
	err := delete.Cluster(flags.Name)
	if err != nil {
		log.Errorf("error creating cluster: %v", err)
		return err
	}
	return nil
}
