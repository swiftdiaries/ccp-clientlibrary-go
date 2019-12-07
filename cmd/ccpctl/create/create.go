package create

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swiftdiaries/ccp-clientlibrary-go/v3/pkg/create"
)

type flagpole struct {
	Config string
}

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	flags := &flagpole{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "create",
		Short: "create a CCP cluster with a cluster YAML/JSON file",
		Long:  "create a CCP cluster with a cluster YAML/JSON file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags, cmd, args)
		},
	}
	cmd.Flags().StringVarP(&flags.Config, "file", "f", "", "config file for CCP cluster creation.")
	return cmd
}

func runE(flags *flagpole, cmd *cobra.Command, args []string) error {
	log.Infof("create called..")
	err := create.Cluster(flags.Config)
	if err != nil {
		log.Errorf("error creating cluster: %v", err)
		return err
	}
	return nil
}
