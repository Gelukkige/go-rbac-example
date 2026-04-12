package initialize

import (
	"fmt"
	"go-rbac-example/internal/api"
	"go-rbac-example/internal/dao"
	"go-rbac-example/internal/global"
	"go-rbac-example/internal/service"

	"github.com/gin-gonic/gin"
)

func RouterInit() {
	serverConfig := global.Config.Server
	r := gin.Default()
	db := global.DB

	logInApi := api.NewLogInAPI()
	r.POST("/login", logInApi.LogIn)

	// USER
	userDao := dao.NewUserDao(db)
	userService := service.NewUserService(userDao)
	userAPI := api.NewUserAPI(userService)
	userGroup := r.Group("/user")
	{
		userGroup.POST("/create", userAPI.CreateUser)
		userGroup.DELETE("/delete", userAPI.DeleteUser)
		userGroup.PUT("/update", userAPI.UpdateUser)
		userGroup.POST("/list", userAPI.ListUsers)
	}

	// ROLE
	roleDao := dao.NewRoleDao(db)
	roleService := service.NewRoleService(roleDao)
	roleAPI := api.NewRoleAPI(roleService)
	roleGroup := r.Group("/role")
	{
		roleGroup.POST("/create", roleAPI.CreateRole)
		roleGroup.DELETE("/delete", roleAPI.DeleteRole)
		roleGroup.PUT("/update", roleAPI.UpdateRole)
		roleGroup.POST("/list", roleAPI.ListRoles)
	}

	// DATA
	dataDao := dao.NewDataDao(db)
	dataService := service.NewDataService(dataDao)
	dataAPI := api.NewDataAPI(dataService)
	dataGroup := r.Group("/data")
	{
		dataGroup.POST("/list", dataAPI.ListData)
	}

	addr := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	go r.Run(addr)
}
