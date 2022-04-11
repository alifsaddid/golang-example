package controllers

import (
	"Oauth/dto/request"
	"Oauth/dto/response"
	"Oauth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController interface {
	Setup(router *gin.RouterGroup)
	CreateRole(c *gin.Context)
	GetAllRole(c *gin.Context)
	GetRoleById(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}

type roleController struct {
	roleService services.RoleService
}

func NewRoleController(roleService services.RoleService) RoleController {
	return &roleController{
		roleService: roleService,
	}
}

func (r roleController) Setup(router *gin.RouterGroup) {
	roles := router.Group("/roles")
	roles.POST("", r.CreateRole)
	roles.GET("", r.GetAllRole)
	roles.GET("/:id", r.GetRoleById)
	roles.PUT("/:id", r.UpdateRole)
	roles.DELETE("/:id", r.DeleteRole)
}

// @Summary 		Create Role
// @Description 	Endpoint for create a new role
// @Tags 			Roles
// @Param           Role			body	request.RoleRequest	true	"Request body"
// @Produce  		json
// @Success 		200
// @Router 			/roles [post]
func (r roleController) CreateRole(c *gin.Context) {
	var role request.RoleRequest
	err := c.ShouldBindJSON(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "Request format invalid",
		})
		return
	}

	res, err := r.roleService.Create(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		Message: "Create new role success",
		Data:    res,
	})
}

// @Summary 		Get All Role
// @Description 	Endpoint for get all roles
// @Tags 			Roles
// @Produce  		json
// @Success 		200
// @Router 			/roles [get]
func (r roleController) GetAllRole(c *gin.Context) {
	res, err := r.roleService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		Message: "Get all role success",
		Data:    res,
	})
}

// @Summary 		Get Role by Id
// @Description 	Endpoint for get role by id
// @Tags 			Roles
// @Param 			id			path	int			true	"Role ID"
// @Produce  		json
// @Success 		200
// @Router 			/roles/{id} [get]
func (r roleController) GetRoleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "Invalid role id",
		})
		return
	}

	res, err := r.roleService.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		Message: "Get role success",
		Data:    res,
	})
}

// @Summary 		Update Role
// @Description 	Endpoint for update a role
// @Tags 			Roles
// @Param           Role	body	request.RoleRequest	true	"Request body"
// @Param 			id		path	int					true	"Role ID"
// @Produce  		json
// @Success 		200
// @Router 			/roles/{id} [put]
func (r roleController) UpdateRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "invalid role id",
		})
		return
	}

	var role request.RoleRequest
	err = c.ShouldBindJSON(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "Request format invalid",
		})
		return
	}

	res, err := r.roleService.UpdateById(id, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		Message: "Update new role success",
		Data:    res,
	})
}

// @Summary 		Delete Role
// @Description 	Endpoint for delete a role
// @Tags 			Roles
// @Param 			id	path	int		true	"Role ID"
// @Produce  		json
// @Success 		200
func (r roleController) DeleteRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "invalid role id",
		})
		return
	}
	err = r.roleService.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		Message: "Delete role success",
	})
}
