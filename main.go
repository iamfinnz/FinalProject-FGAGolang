package main

// Import library
import (
	"mygram/config"
	"mygram/router"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	_ "mygram/docs"
)

// @title MyGram API
// @version 1.0
// @description This is a MyGram server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /mygram

// Function yang akan dijalankan
func main() {
	config.StartDB()  // Memulai koneksi ke database
	r := router.StartApp()  // Memulai aplikasi Echo dan mendapatkan router-nya
	err := r.Run(":8080")  // Menjalankan server HTTP pada port 8080
	if err != nil {
		panic(err)
	}

	e := echo.New()  // Membuat instance baru dari Echo framework
	e.GET("/mygram/*", echoSwagger.WrapHandler)  // Menambahkan handler untuk rute swagger
	e.Logger.Fatal(e.Start(":8080"))  // Menjalankan server HTTP pada port 8080
}

