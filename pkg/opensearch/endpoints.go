package opensearch

type endpoints string

const (
	// URI for creating new users: CreateNewIdexerUserURI:<new user>
	CreateNewIdexerUserURI endpoints = "_plugins/_security/api/internalusers"
)
