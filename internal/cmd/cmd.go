package cmd

import (
	"github.com/NasSilverBullet/autospirit/internal/autospirit"
	"github.com/spf13/cobra"
)

func NewAutoSpiritCommand() *cobra.Command {
	cobra.OnInitialize()
	cmd := autospirit.Exec()
	cmd.AddCommand(autospirit.Go())
	return cmd
}
