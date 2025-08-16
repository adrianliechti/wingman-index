package config

import (
	"errors"
	"os"
	"strings"

	"github.com/adrianliechti/wingman-index/pkg/index"
	"github.com/adrianliechti/wingman-index/pkg/index/azure"
	"github.com/adrianliechti/wingman-index/pkg/index/chroma"
	"github.com/adrianliechti/wingman-index/pkg/index/elasticsearch"
	"github.com/adrianliechti/wingman-index/pkg/index/memory"
	"github.com/adrianliechti/wingman-index/pkg/index/qdrant"
	"github.com/adrianliechti/wingman-index/pkg/index/weaviate"
	"github.com/adrianliechti/wingman-index/pkg/utils"
	"github.com/adrianliechti/wingman/pkg/client"
)

type Config struct {
	Address string

	Client *client.Client

	Index    index.Provider
	Embedder index.Embedder
}

func FromEnvironment() (*Config, error) {
	addr := ":3001"

	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	client, err := clientFromEnvironment()

	if err != nil {
		return nil, err
	}

	embedder, err := embedderFromEnvironment(client)

	if err != nil {
		return nil, err
	}

	index, err := indexFromEnvironment(embedder)

	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Address: addr,

		Client: client,

		Index:    index,
		Embedder: embedder,
	}

	return cfg, nil
}

func clientFromEnvironment() (*client.Client, error) {
	url := os.Getenv("WINGMAN_URL")
	token := os.Getenv("WINGMAN_TOKEN")

	if url == "" {
		url = "http://localhost:8080"
	}

	opts := []client.RequestOption{}

	if token != "" {
		opts = append(opts, client.WithToken(token))
	}

	client := client.New(url, opts...)

	return client, nil
}

func embedderFromEnvironment(client *client.Client) (index.Embedder, error) {
	embeddingModel := os.Getenv("WINGMAN_EMBEDDER")

	if embeddingModel == "" {
		embeddingModel = "text-embedding-3-small"
	}

	embedder := utils.NewClientEmbedder(client, embeddingModel)

	return embedder, nil
}

func indexFromEnvironment(embedder index.Embedder) (index.Provider, error) {
	switch strings.ToLower(os.Getenv("INDEX_TYPE")) {
	case "azure":
		return azureFromEnvironment()
	case "chroma":
		return chromaFromEnvironment(embedder)
	case "elasticsearch":
		return elasticsearchFromEnvironment()
	case "memory":
		return memoryFromEnvironment(embedder)
	case "qdrant":
		return qdrantFromEnvironment(embedder)
	case "weaviate":
		return weaviateFromEnvironment(embedder)
	default:
		return nil, errors.New("invalid index type, expected one of: azure, chroma, elasticsearch, memory, qdrant, weaviate")
	}
}

func azureFromEnvironment() (index.Provider, error) {
	url := os.Getenv("INDEX_URL")

	if url == "" {
		return nil, errors.New("INDEX_URL environment variable is required for Azure index")
	}

	token := os.Getenv("INDEX_TOKEN")

	if token == "" {
		return nil, errors.New("INDEX_TOKEN environment variable is required for Azure index")
	}

	namespace := os.Getenv("INDEX_NAMESPACE")

	if namespace == "" {
		namespace = "default"
	}

	return azure.New(url, namespace, token)
}

func chromaFromEnvironment(embedder index.Embedder) (index.Provider, error) {
	url := os.Getenv("INDEX_URL")

	if url == "" {
		url = "http://localhost:8000"
	}

	namespace := os.Getenv("INDEX_NAMESPACE")

	if namespace == "" {
		namespace = "default"
	}

	return chroma.New(url, namespace, chroma.WithEmbedder(embedder))
}

func elasticsearchFromEnvironment() (index.Provider, error) {
	url := os.Getenv("INDEX_URL")

	if url == "" {
		url = "http://localhost:6333"
	}

	namespace := os.Getenv("INDEX_NAMESPACE")

	if namespace == "" {
		namespace = "default"
	}

	return elasticsearch.New(url, namespace)
}

func memoryFromEnvironment(embedder index.Embedder) (index.Provider, error) {
	return memory.New(memory.WithEmbedder(embedder))
}

func qdrantFromEnvironment(embedder index.Embedder) (index.Provider, error) {
	url := os.Getenv("INDEX_URL")

	if url == "" {
		url = "http://localhost:6333"
	}

	namespace := os.Getenv("INDEX_NAMESPACE")

	if namespace == "" {
		namespace = "default"
	}

	return qdrant.New(url, namespace, qdrant.WithEmbedder(embedder))
}

func weaviateFromEnvironment(embedder index.Embedder) (index.Provider, error) {
	url := os.Getenv("INDEX_URL")

	if url == "" {
		url = "http://localhost:8080"
	}

	namespace := os.Getenv("INDEX_NAMESPACE")

	if namespace == "" {
		namespace = "default"
	}

	return weaviate.New(url, namespace, weaviate.WithEmbedder(embedder))
}
