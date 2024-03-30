package grpc

import (
	"fmt"
	di2 "github.com/boon-neko/golang-grpc-backend-example/internal/infrastructure/di"
	"github.com/boon-neko/golang-grpc-backend-example/internal/infrastructure/grpc"
	"github.com/spf13/cobra"
	"log"
	"net"
	"os"
	"os/signal"
)

func NewServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "serve grpc server",
		Run: func(cmd *cobra.Command, args []string) {
			port := "8080"
			listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
			if err != nil {
				panic(err)
			}
			di := di2.NewDIBox()
			s := grpc.NewServer(di)
			go func() {
				fmt.Printf("serve called. port:%s", port) //nolint:forbidigo
				if err := s.Serve(listener); err != nil {
					panic(err)
				}
			}()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt)
			<-quit
			log.Println("stopping gRPC server...")
			s.GracefulStop()
		},
	}
}
