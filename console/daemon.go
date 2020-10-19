package main

import (
	service "github.com/glvd/link-rest/restapi"
	"github.com/goextension/log/zap"
	"github.com/spf13/cobra"
)

func subDaemon() *cobra.Command {
	var port int
	cmd := &cobra.Command{
		Use: "daemon",
		Run: func(cmd *cobra.Command, args []string) {
			zap.InitZapSugar()

			rest, err := service.New(port)
			if err != nil {
				panic(err)
			}
			if err := rest.Start(); err != nil {
				return
			}
		},
	}
	cmd.Flags().IntVarP(&port, "port", "p", 18080, "set the handle port")
	return cmd
}
