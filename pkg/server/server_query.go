package server

import (
	"context"
	"encoding/json"

	"github.com/adrianliechti/wingman-index/pkg/index"
	"github.com/adrianliechti/wingman-index/pkg/to"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var QueryTool = &mcp.Tool{
	Name:        "query_knowledge_database",
	Title:       "Query Knowledge Database",
	Description: "The search query or question to find relevant information in the knowledge database. Use natural language and be specific about what information you're looking for.",
}

type QueryParams struct {
	Query string `json:"query" jsonschema:"The search query or question to find relevant information in the knowledge database. Use natural language and be specific about what information you're looking for."`
}

type QueryResult struct {
	Title   string `json:"title,omitempty"`
	Source  string `json:"source,omitempty"`
	Content string `json:"content,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`
}

func (s *Server) Query(ctx context.Context, ss *mcp.ServerSession, req *mcp.CallToolParamsFor[QueryParams]) (*mcp.CallToolResultFor[any], error) {
	query := req.Arguments.Query

	opts := &index.QueryOptions{
		Limit: to.Ptr(10),
	}

	results, err := s.Index.Query(ctx, query, opts)

	if err != nil {
		return nil, err
	}

	result := []QueryResult{}

	for _, r := range results {
		result = append(result, QueryResult{
			Title:   r.Title,
			Source:  r.Source,
			Content: r.Content,

			Metadata: r.Metadata,
		})
	}

	data, _ := json.Marshal(result)

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: string(data),
			},
		},
	}, nil
}
