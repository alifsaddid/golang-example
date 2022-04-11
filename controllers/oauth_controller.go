package controllers

import (
	"Oauth/dto/context"
	"Oauth/dto/request"
	"Oauth/dto/response"
	"Oauth/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OauthController interface {
	Token(c *gin.Context)
	Resource(c *gin.Context)
}

type oauthController struct {
	authService services.AuthService
}

func NewOauthController(authService services.AuthService) OauthController {
	return &oauthController{
		authService: authService,
	}
}

// @Summary 		Token
// @Description 	Login endpoint to obtain OAuth token
// @Tags 			Auth
// @Accept 			x-www-form-urlencoded
// @Param			tokenRequest 	formData 	request.TokenRequest	true	"Token request form"
// @Produce  		json
// @Success 		200 			{object} 	response.TokenResponse
// @Router 			/oauth/token [post]
func (o oauthController) Token(c *gin.Context) {
	var req request.TokenRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseErrorResponse{
			Error:            "Binding failed",
			ErrorDescription: err.Error(),
		})
	}

	access, refresh, err := o.authService.Login(req.Username, req.Password, req.Client_id, req.Client_secret)
	if err != nil {
		c.JSON(err.Status, response.BaseErrorResponse{
			Error:            err.Err.Error(),
			ErrorDescription: err.Err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.TokenResponse{
		AccessToken:  access.Token,
		RefreshToken: refresh.Token,
		ExpiresIn:    300,
		TokenType:    "Bearer",
	})
}

// @Summary 		Resource
// @Description 	Resource endpoint
// @Tags 			Auth
// @Param 			Authorization	header	string		true	"Access Token"
// @Produce  		json
// @Success 		200 			{object} 	response.TokenResponse
// @Router 			/oauth/resource [post]
func (o oauthController) Resource(c *gin.Context) {
	cur, _ := c.Get("currentUser")
	user := cur.(context.UserContext)

	c.JSON(http.StatusOK, response.ResourceResponse{
		AccessToken:  user.AccessToken,
		ClientId:     user.ClientId,
		FullName:     user.Name,
		NPM:          user.Npm,
		RefreshToken: user.RefreshToken,
		UserId:       user.Id,
	})
}
