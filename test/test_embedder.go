package test

import (
	"context"

	"github.com/adrianliechti/wingman/pkg/provider"
)

type MockEmbedder struct {
	dimension int
}

// NewMockEmbedder creates a new mock embedder with specified dimension
func NewMockEmbedder() *MockEmbedder {
	return &MockEmbedder{
		dimension: 10,
	}
}

func (m *MockEmbedder) Embed(ctx context.Context, texts []string) (*provider.Embedding, error) {
	vectors := make([][]float32, len(texts))

	for idx, text := range texts {
		embedding := make([]float32, m.dimension)

		// Generate deterministic embeddings based on text
		textLen := len(text)
		for i := 0; i < m.dimension; i++ {
			// Simple hash-like function for deterministic results
			embedding[i] = float32((textLen+i*7)%100) / 100.0
		}

		vectors[idx] = embedding
	}

	return &provider.Embedding{Embeddings: vectors}, nil
}
