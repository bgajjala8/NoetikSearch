package domain 

//UserDB represents a user database with file pointers and segments based on search results
type UserDB struct {
	//Document DB pointer
    FilePointers [] string

	//Vector search segments
	Segments [] string
}

func NewUserDB(filePointers [] string, segments [] string) *UserDB {
	return &UserDB{
		FilePointers: filePointers,
		Segments: segments,
	}
}