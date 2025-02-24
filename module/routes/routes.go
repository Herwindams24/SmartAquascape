package routes

import (
	"SmartAquascape/app/auth"
	"SmartAquascape/app/middlewares"
	"SmartAquascape/module/handler"
	"SmartAquascape/module/utilities/device"
	"SmartAquascape/module/utilities/user"
	"SmartAquascape/module/view"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *gin.Engine {
	// Load Repository
	userRepository := user.NewRepository(db)
	deviceRepository := device.NewRepository(db)
	//Load Service
	userService := user.NewService(userRepository)
	deviceService := device.NewService(deviceRepository)
	//Load Handler
	userHandler := handler.NewUserHandler(userService)
	//deviceHandler := handler.NewDeviceHandler(deviceService)
	//Load View
	deviceView := view.NewDeviceView(deviceService)
	userView := view.NewUserView(userService)

	router := gin.Default()
	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))
	router.Use(sessions.Sessions("tamaskapju", cookieStore))

	router.HTMLRender = ManualRender("./public/template/")

	// Routing Website Service
	// User
	
	// Device


	
	// Khusus Website
	// router.GET("/dashboard", middlewares.AllAkses(), userView.Dashboard)
	router.GET("/login", userView.Login)
	router.GET("/dashboard", userView.Dashboard)
	router.GET("/device", deviceView.Index)
	router.GET("/controldevice", middlewares.AllAkses(), deviceView.ControlDevice)
	router.GET("/controllingdevice", middlewares.AllAkses(), deviceView.ControllingDevice)
	router.GET("/getalldata", deviceView.MonitoringDevice)
	
	// Routing API Service
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	router.POST("/login", userHandler.Login)
	router.GET("/logout", userHandler.Logout)

	return router
}

func ManualRender(tmpDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("login", tmpDir+"home/home_login.html", tmpDir+"layouts/header.html", tmpDir+"layouts/footer.html")
	r.AddFromFiles("dashboard", tmpDir+"home/home_dashboard.html", tmpDir+"layouts/header.html", tmpDir+"layouts/menu.html", tmpDir+"layouts/navbar.html", tmpDir+"layouts/footer.html")
	r.AddFromFiles("controldevice", tmpDir+"device/device_control.html", tmpDir+"layouts/header.html", tmpDir+"layouts/menu.html", tmpDir+"layouts/navbar.html", tmpDir+"layouts/footer.html")
	return r
}
