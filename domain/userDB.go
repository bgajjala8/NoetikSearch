package domain

// UserDB represents a user database with file pointers and segments based on search results
// This object is immutable once created
type UserDB struct {
	//Document DB pointer
	FilePointers []string

	//Vector search segments
	Segments []string
}

// NewUserDB creates a new user database
func NewUserDB(filePointers *[]string, segments *[]string) *UserDB {
	fpCopy := make([]string, len(*filePointers))
	copy(fpCopy, *filePointers)

	segsCopy := make([]string, len(*segments))
	copy(segsCopy, *segments)

	return &UserDB{
		FilePointers: filePointers,
		Segments:     segments,
	}
}
