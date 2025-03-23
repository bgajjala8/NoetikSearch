package service

import (
	"github.com/bgajjala8/GoRagSearch/domain/search"
	"github.com/bgajjala8/GoRagSearch/domain/userDB"
	"io"
	"log"
	"net/http"
)

const (
	vertexUrl = "https://google.com/"
)

type SearchRequest struct {
	Prompt string
}

func Search(req search.Search) *userDB.UserDB {
	//Call repository and build
}
