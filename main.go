package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/adrianliechti/wingman-index/pkg/config"
	"github.com/adrianliechti/wingman-index/pkg/indexer"
	"github.com/adrianliechti/wingman-index/pkg/server"
)

func main() {
	ctx := context.Background()

	root, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	cfg, err := config.FromEnvironment()

	if err != nil {
		panic(err)
	}

	indexer, err := indexer.New(cfg)

	if err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := indexer.IndexDir(ctx, root); err != nil {
				log.Printf("Error indexing directory: %v\n", err)
			}

			time.Sleep(5 * time.Minute)
		}
	}()

	server := server.New(cfg)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
