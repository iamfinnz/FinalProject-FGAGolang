package router

// Import library
import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default() // Membuat instance baru dari engine Gin

	// Definisi rute-rute aplikasi
	userRouter := r.Group("/users")
	// Rute-rute terkait pengguna
	// Meliputi register, login, update, dan delete pengguna
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/", middlewares.Authentication(), controllers.UpdateUser)
		userRouter.DELETE("/", middlewares.Authentication(), controllers.DeleteUser)
	}

	// Definisi rute-rute terkait foto
	photoRouter := r.Group("/photos")
	// Middleware authentication digunakan untuk semua rute dalam group ini
	// Meliputi create, get, update, dan delete foto
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	// Definisi rute-rute terkait komentar
	commentRouter := r.Group("/comments")
	// Middleware authentication digunakan untuk semua rute dalam group ini
	// Meliputi create, get, update, dan delete komentar
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	// Definisi rute-rute terkait media sosial
	socialMediaRouter := r.Group("/socialmedias")
	// Middleware authentication digunakan untuk semua rute dalam group ini
	// Meliputi create, get, update, dan delete media sosial
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r // Mengembalikan instance dari engine Gin yang telah dikonfigurasi dengan rute-rute aplikasi
}
