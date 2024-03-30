package main

import (
	"github.com/boon-neko/golang-grpc-backend-example/internal/infrastructure/cmd"
)

func main() {
	rootCmd := cmd.NewCmdRoot()
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
