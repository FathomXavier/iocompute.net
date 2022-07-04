package cmd

import (
	"github.com/spf13/cobra"

	migratecmd "github.com/ovrclk/provider-services/cmd/provider-services/cmd/migrate"
)

func migrate() *cobra.Command {
	cmd := &cobra.Command{
		Use: "migrate",
	}

	cmd.AddCommand(migratecmd.V0_14ToV0_16())

	return cmd
}
