package facade

//SearchResultsResponse represents the response from the search results
type SearchResultsResponse struct {
	Request SearchRequest
	Results [] string
}