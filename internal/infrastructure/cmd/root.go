package cmd

import (
	"github.com/boon-neko/golang-grpc-backend-example/internal/infrastructure/cmd/internal/grpc"
	"github.com/boon-neko/golang-grpc-backend-example/internal/infrastructure/cmd/internal/mysql"

	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "golang-grpc-backend-example",
		Short: "golang-grpc-backend-example",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				err := cmd.Help()
				if err != nil {
					return
				}
			}
		},
	}
	cmd.AddCommand(grpc.NewServeCmd())
	cmd.AddCommand(mysql.NewMigrateCmd())
	cmd.AddCommand(mysql.NewSeedCmd())

	return cmd
}
