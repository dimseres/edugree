package controllers

import (
	"authorization/internal/helpers"
	"authorization/internal/transport/rest/middlewares"
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func InitUploadRoutes(app *echo.Group) {
	protected := app.Group("")
	protected.Use(middlewares.JwtProtect())
	protected.POST("/image", Upload)
}

func Upload(c echo.Context) error {
	_fileName, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error":   true,
			"message": err,
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	availableExtensions := []string{
		".jpg", ".png", ".jpeg",
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error":   true,
			"message": err,
		})
	}
	defer src.Close()

	fileName := md5.Sum([]byte(_fileName.String() + file.Filename))

	ext := filepath.Ext(file.Filename)
	allow := false
	for _, available := range availableExtensions {
		if available == ext {
			allow = true
			break
		}
	}

	if !allow {
		return c.JSON(http.StatusUnsupportedMediaType, echo.Map{
			"error":   true,
			"message": "unsupported media",
		})
	}

	err = helpers.CreateFolder(helpers.GetPublicUploadPath() + "images")
	if err != nil {
		return c.JSON(500, echo.Map{
			"error":   true,
			"message": err,
		})
	}

	filePath := "images/" + hex.EncodeToString(fileName[:]) + ext
	dest := helpers.GetPublicUploadPath() + filePath

	dst, err := os.Create(dest)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error":   true,
			"message": err,
		})
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(500, echo.Map{
			"error":   true,
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": false,
		"url":   helpers.MakeAbsoluteStaticUrl(filePath),
	})
}
