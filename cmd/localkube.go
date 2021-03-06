package cmd

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/e2e-harness/pkg/localkube"
	"github.com/giantswarm/e2e-harness/pkg/tasks"
	"github.com/giantswarm/microerror"
)

var (
	LocalkubeCmd = &cobra.Command{
		Use:   "localkube",
		Short: "setup localkube",
		RunE:  runLocalkube,
	}
)

var (
	minikubeVersion string
)

func init() {
	RootCmd.AddCommand(LocalkubeCmd)

	SetupCmd.Flags().StringVar(&minikubeVersion, "minikube-version", "v0.25.2", "Minikube version to use.")
}

func runLocalkube(cmd *cobra.Command, args []string) error {
	var err error

	var l *localkube.Localkube
	{
		c := localkube.Config{
			MinikubeVersion: minikubeVersion,
		}

		l, err = localkube.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	// tasks to run
	bundle := []tasks.Task{
		l.SetUp,
	}

	return tasks.Run(bundle)
}
