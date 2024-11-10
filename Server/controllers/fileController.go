package controllers

import (
	"net/http"

	"github.com/Raffique/JL_Adventure_Tours/Server/storage"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No file is uploaded"})
        return
    }

    fileURL, err := storage.UploadFile(c, file)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"url": fileURL})
}
