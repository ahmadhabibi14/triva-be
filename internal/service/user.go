package service

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"triva/configs"
	"triva/helper"
	"triva/internal/bootstrap/database"
	"triva/internal/bootstrap/logger"
	"triva/internal/repository/users"

	"github.com/nfnt/resize"
)

type UserService struct {
	Db *database.Database
}

func NewUserService(Db *database.Database) *UserService {
	return &UserService{Db: Db}
}

func (us *UserService) UpdateAvatar(img *multipart.FileHeader) (user users.User, err error) {
	buff := make([]byte, 512)
	reader, err := img.Open()
	if err != nil {
		return
	}

	_, err = reader.Read(buff)
	if err != nil {
		reader.Close()
		return
	}

	_, err = reader.Seek(0, io.SeekStart)
	if err != nil {
		return
	}

	contentType := http.DetectContentType(buff)

	extension := strings.ToLower(filepath.Ext(img.Filename))
	imgFilename := configs.PATH_AVATAR_IMAGE + helper.RandString(40) + extension

	writer, err := os.Create(imgFilename)
	if err != nil {
		return
	}

	err = resizeImage(writer, reader, contentType)
	if err != nil {
		logger.Log.Err(err)
		return
	}

	return
}

func resizeImage(writer io.Writer, reader io.Reader, contentType string) error {
	var img image.Image
	var err error
	
	switch contentType {
	case `image/png`:
		img, err = png.Decode(reader)
	case `image/jpeg`:
		img, err = jpeg.Decode(reader)
	default:
		img, _, err = image.Decode(reader)
	}

	if err != nil {
		return err
	}

	rect := img.Bounds()
	point := rect.Size()

	log.Println(`Point:`,point)

	nw, nh := uint(point.X), uint(point.Y)
	var maxWidthHeight uint = 256
	if nw > nh {
		if nw > maxWidthHeight {
			nh = nh * maxWidthHeight / nw
			nw = maxWidthHeight
		}
	} else {
		if nh > maxWidthHeight {
			nw = nw * maxWidthHeight / nh
			nh = maxWidthHeight
		}
	}

	img = resize.Resize(nw, nh, img, resize.Lanczos3)

	return jpeg.Encode(writer, img, &jpeg.Options{Quality: 100})
}

const (
	minPixelWidth  = 100
	minPixelHeight = 100
	maxPixelWidth  = 4000
	maxPixelHeight = 4000
	maxFileSize    = 1024 * 1024 * 5 // 5 MB
)

func ValidateImage(img *multipart.FileHeader) error {
	file, err := img.Open()
	if err != nil {
		return errors.New(`failed to open uploaded file`)
	}

	defer file.Close()

	imgConf, format, err := image.DecodeConfig(file)
	if err != nil {
		return errors.New(`uploaded file is not a valid image`)
	}

	switch format {
	case `jpeg`, `jpg`, `png`:
		break
	default:
		return errors.New(`invalid image extension, must be jpeg, jpg, and png`)
	}

	if imgConf.Width != imgConf.Height {
		return errors.New("uploaded image does not have the same dimension, recommend to 500x500")
	}

	if (imgConf.Width < minPixelWidth) || (imgConf.Height < minPixelHeight) {
		return fmt.Errorf("uploaded image does not meet minimum pixel dimensions of %dx%d", minPixelWidth, minPixelHeight)
	}

	if (imgConf.Width > maxPixelWidth) || (imgConf.Height > maxPixelHeight) {
		return fmt.Errorf("uploaded image exceeds maximum pixel dimensions of %dx%d", maxPixelWidth, maxPixelHeight)
	}

	max := float64(maxFileSize) / 1000000
	if img.Size > maxFileSize {
		return fmt.Errorf("uploaded file exceeds maximum size of %.2f MB", max)
	}

	return nil
}