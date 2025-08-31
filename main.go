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

	// Read index directory from env var, default to current working directory
	root := os.Getenv("INDEXDIR")
	if root == "" {
		var err error
		root, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	}

	cfg, err := config.FromEnvironment()

	if err != nil {
		panic(err)
	}

	// Read indexing interval from env var, default to 5 minutes
	intervalStr := os.Getenv("INDEX_INTERVAL")
	if intervalStr == "" {
		intervalStr = "5m"
	}
	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		log.Printf("Invalid INDEX_INTERVAL, using default 5m: %v", err)
		interval = 5 * time.Minute
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

			time.Sleep(interval)
		}
	}()

	server := server.New(cfg)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
