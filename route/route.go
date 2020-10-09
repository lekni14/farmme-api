package routes

import (
	"log"
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	auth "farmme-api/middleware"
	"farmme-api/repository"	
	jwt "github.com/appleboy/gin-jwt/v2"
	handle_user "farmme-api/api/v1/user"
	handle_farm "farmme-api/api/v1/farm"
	handle_cows "farmme-api/api/v1/cow"
	handle_gen "farmme-api/api/v1/generate"
)
// Route main 
func Route(route *gin.Engine, connectionDB *mongo.Database) {
	userRepository := repository.UserRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	userAPI := handle_user.UserAPI{
		UserRepository: &userRepository,
	}
	middleware := auth.Auth(userAPI)

	//route.StaticFS("/uploads/images/", http.Dir(upload.GetImagePath()))
	//route.Static("/upload", "./upload")
	api := route.Group("/api/v1")
	{

		api.POST("/user/ep", userAPI.AddEP)
		api.POST("/user/pd", userAPI.AddPD)
		api.POST("/user/login", middleware.LoginHandler)
		api.POST("/user/loginPD", middleware.LoginHandler)
		// api.POST("/user/forgotpass", userAPI.ForgotPassword)
		// api.POST("/user/forgotpassword", userAPI.ForgotPasswordMobile)
		// api.POST("/user/updatepass", userAPI.UpdatePassword)
		api.Use(middleware.MiddlewareFunc())
		{
		// 	api.POST("/uploads", uploads.Uploads)
		// 	api.POST("/uploadCover", uploads.UploadCover)
		// 	api.POST("/user/avatar", userAPI.UpdateAvatar)
		// 	api.POST("/user/address", userAPI.AddAdress)
		// 	api.POST("/uploadSlip", uploads.UploadSlip)
		// 	api.POST("/uploadWithFolder", uploads.UploadWithFolder)
		 	api.GET("/user", userAPI.Get)
		// 	api.PUT("/user", userAPI.Edit)
		// 	api.GET("/user/confirm", userAPI.Confirm)
		// 	api.POST("/user/changepass", userAPI.ChangePassword)
		// 	api.DELETE("/user/:id", userAPI.Delete)
			api.GET("/user/logout", func(c *gin.Context) {
				log.Println("logout")
				if token, err := middleware.CheckIfTokenExpire(c); err == nil {
					if err2 := token.Valid(); err2 == nil {
						log.Println("valid")
						middleware.DisabledAbort = true
						token["exp"] = time.Now().UTC().Unix()
						c.Abort()
						res := gin.H{"msg": "success"}
						c.JSON(http.StatusOK, res)
					}
				} else {
					log.Println(err)
				}

			})
		}
	}
	route.NoRoute(func(c *gin.Context) {
        c.AbortWithStatus(http.StatusNotFound)
	})
	FarmRoute(route, connectionDB, middleware)
	CowRoute(route, connectionDB, middleware)
	GenRoute(route, connectionDB, middleware)
}
// FarmRoute  path
func FarmRoute(route *gin.Engine, connectionDB *mongo.Database, middleware *jwt.GinJWTMiddleware) {
	farmRepository := repository.FarmRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	farmAPI := handle_farm.FarmAPI{
		FarmRepository: &farmRepository,
	}

	api := route.Group("/api/v1/farm")
	{
		// api.GET("/findByStatus/:status", farmAPI.GetByStatus)
		// api.GET("/eventInfo/:id", eventAPI.GetByID)
		// api.GET("/all", eventAPI.GetAll)
		

		api.Use(middleware.MiddlewareFunc())
		{

			api.GET("/myFarm", farmAPI.MyFarm)
			api.POST("", farmAPI.AddFarm)
			// api.GET("/myEvent", eventAPI.MyEvent)
			
		}
	}

}
// CowRoute in path
func CowRoute(route *gin.Engine, connectionDB *mongo.Database, middleware *jwt.GinJWTMiddleware) {
	cowRepository := repository.CowRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	cowAPI := handle_cows.CowAPI{
		CowRepository: &cowRepository,
	}

	api := route.Group("/api/v1/cow")
	{
		api.Use(middleware.MiddlewareFunc())
		{
			api.POST("", cowAPI.AddCow)			
		}
	}

}
// GenRoute path generate
func GenRoute(route *gin.Engine, connectionDB *mongo.Database, middleware *jwt.GinJWTMiddleware) {
	genRepository := repository.CowRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	genAPI := handle_gen.GenAPI{
		GenRepository: &genRepository,
	}

	api := route.Group("/api/v1/generate")
	{
		api.Use(middleware.MiddlewareFunc())
		{
			api.POST("", genAPI.AddGen)			
		}
	}

}