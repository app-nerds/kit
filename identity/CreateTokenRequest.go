package identity

/*
A CreateTokenRequest is used when creating a new JWT token.
It contians basic information about a user, and then allows
for additional data
*/
type CreateTokenRequest struct {
	UserID string
	UserName string
	AdditionalData map[string]interface{}
}