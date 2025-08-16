package server

import (
	"net/http"

	"github.com/adrianliechti/wingman-index/pkg/config"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Server struct {
	*config.Config
	server *mcp.Server
}

func New(cfg *config.Config) *Server {
	impl := &mcp.Implementation{
		Name:    "index",
		Version: "v1.0.0",
	}

	s := &Server{
		Config: cfg,
		server: mcp.NewServer(impl, nil),
	}

	mcp.AddTool(s.server, QueryTool, s.Query)

	return s
}

func (s *Server) ListenAndServe() error {
	handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
		return s.server
	}, nil)

	server := &http.Server{
		Addr:    s.Address,
		Handler: handler,
	}

	return server.ListenAndServe()
}
