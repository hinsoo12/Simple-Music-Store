package controllers

import (
	"net/http"

	"simple-music-store/models"

	"github.com/gin-gonic/gin"
)
//
type CreateMusicFields struct {
	Title    string `json:"title" binding:"required"`
	Singer   string `json:"singer" binding:"required"`
	Writer   string `json:"writer" binding:"required"`
	Director string `json:"director" binding:"required"`
}
//
type UpdateMusicField struct {
	Title    string `json:"title"`
	Singer   string `json:"singer" `
	Writer   string `json:"writer" `
	Director string `json:"director"`
}
//
var songs = []models.Song{
	{ID: 1, Title: "Hager Ethiopia", Singer: "Jano Band", Writer: "Jano Band", Director: "Jano Band"},
	{ID: 1, Title: "Hager Ethiopia", Singer: "Jano Band", Writer: "Jano Band", Director: "Jano Band"},
	{ID: 3, Title: "Hager Ethiopia", Singer: "Jano Band", Writer: "Jano Band", Director: "Jano Band"},
}

// GET /songs
func FindSongs(ctx *gin.Context) {
	var songs []models.Song
	models.DB.Find(&songs)

	ctx.JSON(http.StatusOK, gin.H{"result": songs})
}

// GET /songs/:id
func FindSong(ctx *gin.Context) {
	// Get model if exist
	var song models.Song
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&song).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": song})
}

// POST /songs

func AddSong(ctx *gin.Context) {
	// Validate input
	var input CreateMusicFields
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add song
	song := models.Song{Title: input.Title, Singer: input.Singer, Writer: input.Writer, Director: input.Director}
	models.DB.Create(&song)

	ctx.JSON(http.StatusOK, gin.H{"data": song})
}

// PATCH /song/:id

func UpdateSong(ctx *gin.Context) {
	// Get model if exist
	var song models.Song
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&song).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateMusicField
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&song).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{"data": song})
}

// DELETE /songs/:id

func DeleteSong(ctx *gin.Context) {
	// Get model if exist
	var song models.Song
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&song).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&song)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
