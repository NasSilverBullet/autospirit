package cmd

import (
	"github.com/NasSilverBullet/autospirit/internal/autospirit"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cobra.OnInitialize()
	cmd := autospirit.Exec()
	cmd.AddCommand(autospirit.Hello())
	cmd.AddCommand(autospirit.Bye())
	cmd.AddCommand(autospirit.Test())
	return cmd
}
