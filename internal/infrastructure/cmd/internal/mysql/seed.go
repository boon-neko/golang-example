/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package mysql

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewSeedCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "seed",
		Short: "seed data for mysql database;",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("serve called")
		},
	}
}
