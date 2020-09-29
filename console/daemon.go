package main

import (
	c "github.com/glvd/link-rest/controller"
	"github.com/goextension/log/zap"
	"github.com/spf13/cobra"
)

func subDaemon() *cobra.Command {
	var port int
	cmd := &cobra.Command{
		Use: "daemon",
		Run: func(cmd *cobra.Command, args []string) {
			zap.InitZapSugar()

			rest, err := c.New(port)
			if err != nil {
				panic(err)
			}
			rest.Start()
		},
	}
	cmd.Flags().IntVarP(&port, "port", "p", 18080, "set the handle port")
	return cmd
}
