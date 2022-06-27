package cluster

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/zeet-dev/cli/pkg/api"
	"github.com/zeet-dev/cli/pkg/cmdutil"
	"github.com/zeet-dev/cli/pkg/iostreams"
)

type KubeconfigSetOptions struct {
	IO        *iostreams.IOStreams
	ApiClient func() (*api.Client, error)

	File      string
	ClusterID uuid.UUID
}

func NewKubeconfigSetCmd(f *cmdutil.Factory) *cobra.Command {
	var opts = &KubeconfigSetOptions{}
	opts.IO = f.IOStreams
	opts.ApiClient = f.ApiClient

	cmd := &cobra.Command{
		Use:   "kubeconfig:set [kubeconfig location] [cluster id]",
		Short: "Uploads a kubeconfig.yaml to Zeet",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.File = args[1]
			opts.ClusterID = uuid.MustParse(args[0])

			return runKubeconfigSet(opts)
		},
	}

	return cmd
}

func runKubeconfigSet(opts *KubeconfigSetOptions) error {
	client, err := opts.ApiClient()
	if err != nil {
		return err
	}

	dat, err := os.ReadFile(opts.File)
	if err != nil {
		return err
	}

	_, err = client.UpdateClusterKubeconfig(context.Background(), opts.ClusterID, dat)
	if err != nil {
		return err
	}

	fmt.Fprintln(opts.IO.Out, color.GreenString("Cluster updated"))
	return nil
}
