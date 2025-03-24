package service

import (
	"log/slog"
	"os"

	"GoRagSearch/domain/search"
	"GoRagSearch/domain/userDB"
)

var (
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
)

const (
	vertexUrl = "https://google.com/"
)

type SearchRequest struct {
	Prompt string
}

// newSearchRequest creates a new search request (package use only)
func newSearchRequest(prompt string) *SearchRequest {
	return &SearchRequest{
		Prompt: prompt,
	}
}

func Search(req search.Search) *userDB.UserDB {
	//Call repository and build
	logger.Info("Search request received", req)
}
