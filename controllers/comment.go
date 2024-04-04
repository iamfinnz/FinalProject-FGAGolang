package controllers

// Import library
import (
	"encoding/json"
	"mygram/config"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	// Get instance database dari konfigurasi
	db := config.GetDB()
	// Get data pengguna dari konteks
	userData := c.MustGet("userData").(jwt.MapClaims)
	// Get content type dari request
	contentType := helpers.GetContentType(c)

	// Deklarasi struct untuk request pembuatan comment
	commentRequest := models.CreateCommentRequest{}
	// Get ID pengguna dari data pengguna
	userID := uint(userData["id"].(float64))

	// Periksa content type request
	if contentType == appJSON {
		// Bind JSON request ke struct commentRequest
		if err := c.ShouldBindJSON(&commentRequest); err != nil {
			// Jika error, kirim response Bad Request
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		// Bind JSON request ke struct commentRequest
		if err := c.ShouldBind(&commentRequest); err != nil {
			// Jika error, kirim response Bad Request
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	// Create objek comment
	comment := models.Comment{
		PhotoId: commentRequest.PhotoId,
		Message: commentRequest.Message,
		UserId:  userID,
	}

	// Create new comment ke dalam database
	err := db.Debug().Create(&comment).Error
	if err != nil {
		// Jika error kirim response error Bad Request
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Marshal objek comment ke JSON
	commentString, _ := json.Marshal(comment)
	// Deklarasi objek response untuk membuak komentar
	commentResponse := models.CreateCommentResponse{}
	// Unmarshal JSON ke response
	json.Unmarshal(commentString, &commentResponse)

	//Kirim response OK ke objek response
	c.JSON(http.StatusCreated, commentResponse)
}

func GetComment(c *gin.Context) {
	// Get instance database dari konfigurasi
	db := config.GetDB()

	// Deklarasi slice untuk simpan komentar
	comments := []models.Comment{}

	// Get semua komentar dari database
	err := db.Debug().Preload("User").Preload("Photo").Order("id asc").Find(&comments).Error
	if err != nil {
		// Jika error kirim response Bad Request
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Marshal komentar ke JSON
	commentsString, _ := json.Marshal(comments)
	// Deklarasi slice untuk menyimpan respons komentar
	commentsResponse := []models.CommentResponse{}
	// Unmarshal JSON ke slice respons
	json.Unmarshal(commentsString, &commentsResponse)

	// Kirim response OK dengan slice respons komentar
	c.JSON(http.StatusOK, commentsResponse)
}

func UpdateComment(c *gin.Context) {
	// Get instance database dari konfigurasi
	db := config.GetDB()
	// Get data pengguna dari konteks
	userData := c.MustGet("userData").(jwt.MapClaims)
	// Get content type dari request
	contentType := helpers.GetContentType(c)

	// Deklarasi struct untuk request update komentar
	commentRequest := models.UpdateCommentRequest{}
	// Get ID komentar dari parameter URL
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	// Get ID pengguna dari data pengguna
	userID := uint(userData["id"].(float64))

	// Periksa content type request
	if contentType == appJSON {
		// Bind JSON request ke struct commentRequest
		if err := c.ShouldBindJSON(&commentRequest); err != nil {
			// Jika error kirim response error Bad Request
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		// Bind request ke struct commentRequest
		if err := c.ShouldBind(&commentRequest); err != nil {
			// Jika error kirim response error Bad Request
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	// Deklarasi objek komentar yang akan diupdate
	comment := models.Comment{}
	comment.ID = uint(commentId)
	comment.UserId = userID

	// Marshal data update ke JSON
	updateString, _ := json.Marshal(commentRequest)
	// Deklarasikan objek data update komentar
	updateData := models.Comment{}
	// Unmarshal JSON ke objek data update komentar
	json.Unmarshal(updateString, &updateData)

	// Update komentar dalam database
	err := db.Model(&comment).Updates(updateData).Error
	if err != nil {
		// Jika error kirim response error Bad Request
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	// Get data komentar yang telah diupdate
	_ = db.First(&comment, comment.ID).Error

	// Marshal komentar ke JSON
	commentString, _ := json.Marshal(comment)
	// Deklarasi objek response untuk update komentar
	commentResponse := models.UpdateCommentResponse{}
	// Unmarshal JSON ke objek respons
	json.Unmarshal(commentString, &commentResponse)

	// Kirim response OK dengan objek respons
	c.JSON(http.StatusOK, commentResponse)
}

func DeleteComment(c *gin.Context) {
	// Get instance database dari konfigurasi
	db := config.GetDB()
	// Get data pengguna dari konteks
	userData := c.MustGet("userData").(jwt.MapClaims)

	// Get ID komentar dari parameter URL
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	// Get ID pengguna dari data pengguna
	userID := uint(userData["id"].(float64))

	// Deklarasikan objek komentar yang akan dihapus
	comment := models.Comment{}
	comment.ID = uint(commentId)
	comment.UserId = userID

	// Hapus komentar dari database
	err := db.Delete(&comment).Error
	if err != nil {
		// Jika error kirim response error Bad Request
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Kirim response OK dengan pesan berhasil delete
	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
