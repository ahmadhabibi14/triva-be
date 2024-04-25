package users

type Session struct {
	UserID string `json:"user_id"`
	Username string `json:"username"`
	Authenticated bool `json:"authenticated"`
}

func NewSession(userId, username string, authenticated bool) *Session {
	return &Session{
		UserID: userId,
		Username: username,
		Authenticated: authenticated,
	}
}