package secret

import (
	"github.com/pkg/errors"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/prerun"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/spf13/cobra"
)

const EmptyCreateError = "please provide a command for the type of secret"

func CreateCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:     constants.SECRET_COMMAND.Use,
		Aliases: constants.SECRET_COMMAND.Aliases,
		Short:   "Create a secret",
		Long:    "Create a secret",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := prerun.CallParentPrerun(cmd, args); err != nil {
				return err
			}
			if err := prerun.EnableVaultClients(opts.Create.Vault); err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.Errorf(EmptyCreateError)
		},
	}
	cmd.AddCommand(awsCmd(opts))
	cmd.AddCommand(azureCmd(opts))
	cmd.AddCommand(tlsCmd(opts))
	flagutils.AddVaultSecretFlags(cmd.PersistentFlags(), &opts.Create.Vault)
	cliutils.ApplyOptions(cmd, optionsFunc)
	return cmd
}
