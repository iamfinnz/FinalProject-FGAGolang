package middlewares

// Import library
import (
	"mygram/config"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Function untuk verifikasi akses pengguna ke data foto
func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get instance dari config
		db := config.GetDB()
		// Get ID foto dari parameter URL
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			// Jika error kirim response error Bad Request
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}
		// Get data pengguna dari konteks
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		photo := models.Photo{}

		// Cari data foto dari database berdasarkan ID foto
		err = db.Select("user_id").First(&photo, uint(photoId)).Error
		if err != nil {
			// Jika tidak ditemukan kirim response HTTP NotFound
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})
			return
		}

		if photo.UserId != userID {
			// Jika pengguna tidak punya hak akses, kirim response Unauthorized
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		// Jika verif selesai, lanjut eksekusi handler selanjutnya
		c.Next()
	}
}

// Function untuk verifikasi akses pengguna ke data comment
// Cara kerja sama dengan foto
func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		comment := models.Comment{}

		err = db.Select("user_id").First(&comment, uint(commentId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})
			return
		}

		if comment.UserId != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

// Function untuk verifikasi akses pengguna ke data social media
// Cara kerja sama dengan function sebelumnya
func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GetDB()
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		socialMedia := models.SocialMedia{}

		err = db.Select("user_id").First(&socialMedia, uint(socialMediaId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})
			return
		}

		if socialMedia.UserId != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}