package utils

import (
	"context"

	"github.com/adrianliechti/wingman/pkg/client"
	"github.com/adrianliechti/wingman/pkg/provider"
)

type ClientEmbedder struct {
	client *client.Client
	model  string
}

func NewClientEmbedder(client *client.Client, model string) *ClientEmbedder {
	return &ClientEmbedder{
		client: client,
		model:  model,
	}
}

func (e *ClientEmbedder) Embed(ctx context.Context, texts []string) (*provider.Embedding, error) {
	result, err := e.client.Embeddings.New(ctx, client.EmbeddingsRequest{
		Model: e.model,
		Texts: texts,
	})

	if err != nil {
		return nil, err
	}

	return &provider.Embedding{
		Model: result.Model,

		Embeddings: result.Embeddings,
	}, nil
}
