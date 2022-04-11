package context

type UserContext struct {
	Id           int
	Username     string
	Name         string
	Npm          string
	ClientId     string
	AccessToken  string
	RefreshToken string
}
