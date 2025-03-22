package domain

//Search represents a search request
type Search struct {
	//Query Prompt
	Prompt string

	//User ID for logical search
	UserID string
}

func NewSearch(prompt string, userID string) *Search {
	return &Search{
		Prompt: prompt,
		UserID: userID,
	}
}