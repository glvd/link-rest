package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/glvd/link-rest/db"
	"github.com/glvd/link-rest/scrape"
	"github.com/goextension/log"
	"github.com/goextension/log/zap"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func subScrape() *cobra.Command {
	var path string
	cmd := &cobra.Command{
		Use: "scrape",
		Run: func(cmd *cobra.Command, args []string) {
			zap.InitZapSugar()
			ctx, cf := context.WithCancel(context.TODO())
			parseDone := make(chan bool)
			go func(ctx context.Context, done chan<- bool) {
				defer close(done)
				openedFile, err := os.Open(path)
				if err != nil {
					return
				}
				cfg := db.ParseFromMap(nil)
				engine, err := db.New(cfg)
				if err != nil {
					panic(err)
				}

				api, err := httpapi.NewLocalApi()
				if err != nil {
					return
				}
				s := scrape.NewScraper(api, engine)
				r := bufio.NewReader(openedFile)
			Loop:
				for {
					select {
					case <-ctx.Done():
						return
					default:
						line, _, err := r.ReadLine()
						if err != nil {
							log.Error(err)
							break Loop
						}
						err = s.ParseHash(ctx, string(line))
						if err != nil {
							log.Error(err)
							continue
						}
					}
				}
			}(ctx, parseDone)

			interrupts := make(chan os.Signal)
			signal.Notify(interrupts, os.Interrupt, syscall.SIGTERM)
			select {
			case <-interrupts:
				cf()
				fmt.Println("system exit with system interrupt")
			case <-parseDone:
				fmt.Println("system exit with parse hash done")
			}
		},
	}
	cmd.Flags().StringVarP(&path, "path", "p", "scrape.txt", "scrape hash from list and input into database")
	return cmd
}
