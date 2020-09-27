package main

import (
	c "github.com/glvd/link-rest/controller"
	"github.com/goextension/log/zap"
	"github.com/spf13/cobra"
)

func subDaemon() *cobra.Command {
	var port int
	cmd := &cobra.Command{
		Use:                    "daemon",
		Aliases:                nil,
		SuggestFor:             nil,
		Short:                  "",
		Long:                   "",
		Example:                "",
		ValidArgs:              nil,
		Args:                   nil,
		ArgAliases:             nil,
		BashCompletionFunction: "",
		Deprecated:             "",
		Hidden:                 false,
		Annotations:            nil,
		Version:                "",
		PersistentPreRun:       nil,
		PersistentPreRunE:      nil,
		PreRun:                 nil,
		PreRunE:                nil,
		Run: func(cmd *cobra.Command, args []string) {
			zap.InitZapSugar()

			rest, err := c.New(port)
			if err != nil {
				panic(err)
			}
			rest.Start()
		},
		RunE:                       nil,
		PostRun:                    nil,
		PostRunE:                   nil,
		PersistentPostRun:          nil,
		PersistentPostRunE:         nil,
		SilenceErrors:              false,
		SilenceUsage:               false,
		DisableFlagParsing:         false,
		DisableAutoGenTag:          false,
		DisableFlagsInUseLine:      false,
		DisableSuggestions:         false,
		SuggestionsMinimumDistance: 0,
		TraverseChildren:           false,
		FParseErrWhitelist:         cobra.FParseErrWhitelist{},
	}
	cmd.Flags().IntVarP(&port, "port", "p", 18080, "set the handle port")
	return cmd
}
