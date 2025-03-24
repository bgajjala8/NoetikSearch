package service

import (
	"log/slog"
	"os"

	"github.com/bgajjala8/GoRagSearch/domain"
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

func Search(req domain.Search) *domain.UserDB {
	//Call repository and build
	logger.Info("Search request received", req)
}
