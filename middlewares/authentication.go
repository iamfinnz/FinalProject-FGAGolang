package middlewares

// Import library
import (
	"mygram/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Function Authentication untuk memeriksa validitas token JWT
func Authentication() gin.HandlerFunc {
	// Get context
	return func(c *gin.Context) {
		// Panggil verify token untuk verifikasi token JWT dari request
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			// Jika error hentikan eksekusi dan kirim response HTTP Unauthorizrd
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		// Jika verif berhasil simpan nilai ke context
		c.Set("userData", verifyToken)
		c.Next()
	}
}