package bean

import (
	"Oauth/controllers"
	"Oauth/repositories"
	"Oauth/services"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type BeanFactory interface {
	// repositories
	GetClientRepository() repositories.ClientRepository
	GetUserRepository() repositories.UserRepository

	// services
	GetAuthService() services.AuthService
	GetClientService() services.ClientService
	GetJWTService() services.JWTService
	GetRedisService() services.RedisService
	GetUserService() services.UserService

	// controllers
	GetOauthController() controllers.OauthController
	GetRoleController() controllers.RoleController
}

type beanFactory struct {
	// db
	db          *gorm.DB
	redisClient *redis.Client

	// repositories
	clientRepository     repositories.ClientRepository
	userRepository       repositories.UserRepository
	permissionRepository repositories.PermissionRepository
	roleRepository       repositories.RoleRepository

	// services
	authService       services.AuthService
	clientService     services.ClientService
	jwtService        services.JWTService
	redisService      services.RedisService
	userService       services.UserService
	permissionService services.PermissionService
	roleService       services.RoleService

	// controllers
	oauthController controllers.OauthController
	roleController  controllers.RoleController
}

func NewFactory(db *gorm.DB, redisClient *redis.Client) BeanFactory {
	return &beanFactory{
		db:          db,
		redisClient: redisClient,
	}
}

/*
 * Repositories
 */
func (f beanFactory) GetClientRepository() repositories.ClientRepository {
	if f.clientRepository == nil {
		f.clientRepository = repositories.NewClientRepository(f.db)
	}

	return f.clientRepository
}

func (f beanFactory) GetUserRepository() repositories.UserRepository {
	if f.userRepository == nil {
		f.userRepository = repositories.NewUserRepository(f.db)
	}
	return f.userRepository
}

func (f beanFactory) GetPermissionRepository() repositories.PermissionRepository {
	if f.permissionRepository == nil {
		f.permissionRepository = repositories.NewPermissionRepository(f.db)
	}
	return f.permissionRepository
}

func (f beanFactory) GetRoleRepository() repositories.RoleRepository {
	if f.roleRepository == nil {
		f.roleRepository = repositories.NewRoleRepository(f.db)
	}
	return f.roleRepository
}

/*
 * Services
 */
func (f beanFactory) GetAuthService() services.AuthService {
	if f.authService == nil {
		f.authService = services.NewAuthService(
			f.GetJWTService(),
			f.GetUserService(),
			f.GetRedisService(),
			f.GetClientService(),
		)
	}
	return f.authService
}

func (f beanFactory) GetClientService() services.ClientService {
	if f.clientService == nil {
		f.clientService = services.NewClientService(
			f.GetClientRepository(),
		)
	}
	return f.clientService
}

func (f beanFactory) GetJWTService() services.JWTService {
	if f.jwtService == nil {
		f.jwtService = services.NewJwtService(
			f.GetRedisService(),
			f.GetUserService(),
			f.GetClientService(),
		)
	}
	return f.jwtService
}

func (f beanFactory) GetRedisService() services.RedisService {
	if f.redisService == nil {
		f.redisService = services.NewRedisService(f.redisClient)
	}
	return f.redisService
}

func (f beanFactory) GetUserService() services.UserService {
	if f.userService == nil {
		f.userService = services.NewUserService(
			f.GetUserRepository(),
		)
	}
	return f.userService
}

func (f beanFactory) GetPermissionService() services.PermissionService {
	if f.permissionService == nil {
		f.permissionService = services.NewPermissionService(
			f.GetPermissionRepository(),
		)
	}
	return f.permissionService
}

func (f beanFactory) GetRoleService() services.RoleService {
	if f.roleService == nil {
		f.roleService = services.NewRoleService(
			f.GetRoleRepository(),
			f.GetPermissionService(),
		)
	}
	return f.roleService
}

/*
 * Controllers
 */
func (f beanFactory) GetOauthController() controllers.OauthController {
	if f.oauthController == nil {
		f.oauthController = controllers.NewOauthController(
			f.GetAuthService(),
		)
	}
	return f.oauthController
}

func (f beanFactory) GetRoleController() controllers.RoleController {
	if f.roleController == nil {
		f.roleController = controllers.NewRoleController(
			f.GetRoleService(),
		)
	}
	return f.roleController
}
