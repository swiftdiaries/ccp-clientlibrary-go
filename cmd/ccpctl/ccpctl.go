package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"sigs.k8s.io/kind/cmd/kind/version"
)

const defaultLevel = log.WarnLevel

// Flags for the kind command
type Flags struct {
	LogLevel string
}

// NewCommand creates the root cobra command
func NewCommand() *cobra.Command {
	flags := &Flags{LogLevel: "info"}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "ccpctl",
		Short: "ccpctl is a productivity tool for Cisco Container Platform",
		Long: `
		ccpctl is a productivity tool for Cisco Container Platform
	`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags, cmd, args)
		},
		SilenceUsage: true,
		Version:      version.Version,
	}
	cmd.AddCommand(version.NewCommand())
	return cmd
}

func runE(flags *Flags, cmd *cobra.Command, args []string) error {
	level := defaultLevel
	parsed, err := log.ParseLevel(flags.LogLevel)
	if err != nil {
		log.Warnf("Invalid log level '%s', defaulting to '%s'", flags.LogLevel, level)
	} else {
		level = parsed
	}
	log.SetLevel(level)
	return nil
}

// Run runs the `ccpctl` root command
func Run() error {
	return NewCommand().Execute()
}

func main() {
	log.SetOutput(os.Stdout)
	// this formatter is the default, but the timestamps output aren't
	// particularly useful, they're relative to the command start
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
		ForceColors:     true,
	})
	if err := Run(); err != nil {
		os.Exit(1)
	}
}
