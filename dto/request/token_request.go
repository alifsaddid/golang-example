package request

type TokenRequest struct {
	Username      string `form:"username" binding:"required"`
	Password      string `form:"password" binding:"required"`
	Grant_type    string `form:"grant_type" binding:"required"`
	Client_id     string `form:"client_id" binding:"required"`
	Client_secret string `form:"client_secret" binding:"required"`
}
