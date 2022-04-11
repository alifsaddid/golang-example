package response

type ResourceResponse struct {
	AccessToken  string `json:"access_token"`
	ClientId     string `json:"client_id"`
	UserId       int    `json:"user_id"`
	FullName     string `json:"full_name"`
	NPM          string `json:"npm"`
	Expires      string `json:"expires"`
	RefreshToken string `json:"refresh_token"`
}
