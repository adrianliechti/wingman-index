package chroma_test

import (
	"testing"

	"github.com/adrianliechti/wingman-index/pkg/index/chroma"
	"github.com/adrianliechti/wingman-index/test"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestChroma(t *testing.T) {
	context := test.NewContext()

	server, err := testcontainers.GenericContainer(context.Context, testcontainers.GenericContainerRequest{
		Started: true,

		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "ghcr.io/chroma-core/chroma:1.0.17",
			ExposedPorts: []string{"8000/tcp"},
			WaitingFor:   wait.ForExposedPort(),
		},
	})

	require.NoError(t, err)

	url, err := server.Endpoint(context.Context, "")
	require.NoError(t, err)

	c, err := chroma.New("http://"+url, "test", chroma.WithEmbedder(context.Embedder))
	require.NoError(t, err)

	test.TestIndex(t, context, c)
}
