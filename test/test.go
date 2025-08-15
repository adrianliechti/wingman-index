package test

import (
	"context"
	"testing"

	"github.com/adrianliechti/wingman-index/pkg/index"
	"github.com/adrianliechti/wingman/pkg/provider"
)

type TestContext struct {
	Context context.Context

	Embedder provider.Embedder
}

func NewContext() *TestContext {
	return &TestContext{
		Context:  context.Background(),
		Embedder: NewMockEmbedder(),
	}
}

func TestIndex(t *testing.T, ctx *TestContext, provider index.Provider) {
}
