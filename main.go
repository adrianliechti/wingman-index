package main

import (
	"github.com/adrianliechti/wingman-index/pkg/index/azure"
	"github.com/adrianliechti/wingman-index/pkg/index/chroma"
	"github.com/adrianliechti/wingman-index/pkg/index/elasticsearch"
	"github.com/adrianliechti/wingman-index/pkg/index/memory"
	"github.com/adrianliechti/wingman-index/pkg/index/qdrant"
	"github.com/adrianliechti/wingman-index/pkg/index/weaviate"
)

func main() {
}

func indexFromEnv() {
	var url string
	var namespace string
	var token string

	azure.New(url, namespace, token)
	chroma.New(url, namespace)
	elasticsearch.New(url, namespace)
	memory.New()
	qdrant.New(url, namespace)
	weaviate.New(url, namespace)
}
