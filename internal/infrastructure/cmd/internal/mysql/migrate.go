package mysql

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "exec migration for mysql ",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("serve called")
		},
	}
}
