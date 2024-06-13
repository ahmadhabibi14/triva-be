package service

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
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

func (us *UserService) UpdateAvatar(img *multipart.FileHeader, userId uint64) (user *users.User, err error) {
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

	imgFilename := helper.RandString(40) + extension
	fullImgFileName := filepath.Join(configs.OS_PATH_AVATAR_IMAGE, imgFilename)

	err = os.MkdirAll(configs.OS_PATH_AVATAR_IMAGE, os.ModePerm)
	if err != nil {
		logger.Log.Err(err).Msg(`error while creating directory`)
		return
	}

	writer, err := os.Create(fullImgFileName)
	if err != nil {
		logger.Log.Err(err).Msg(`error while creating file`)
		return
	}

	err = resizeImage(writer, reader, contentType, 256)
	if err != nil {
		logger.Log.Err(err).Msg(`error while resizing image`)
		return
	}

	user = users.NewUserMutator(us.Db)
	user.Id = userId
	user.AvatarURL = configs.WEB_PATH_AVATAR_IMAGE + imgFilename

	err = user.UpdateAvatarById()
	if err != nil {
		return
	}

	return
}

func resizeImage(writer io.Writer, reader io.Reader, contentType string, toWidthHeight uint) error {
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

	width, height := uint(point.X), uint(point.Y)
	if width != height {
		return errors.New("uploaded image does not have the same dimension, recommend to 500x500")
	}

	if (width < minPixelWidth) || (height < minPixelHeight) {
		return fmt.Errorf("uploaded image does not meet minimum pixel dimensions of %dx%d", minPixelWidth, minPixelHeight)
	}

	if (width > maxPixelWidth) || (height > maxPixelHeight) {
		return fmt.Errorf("uploaded image exceeds maximum pixel dimensions of %dx%d", maxPixelWidth, maxPixelHeight)
	}

	if width > toWidthHeight {
		width = toWidthHeight
		height = toWidthHeight
	}

	img = resize.Resize(width, height, img, resize.Lanczos3)

	return jpeg.Encode(writer, img, &jpeg.Options{Quality: 100})
}

const (
	minPixelWidth  uint = 150
	minPixelHeight uint = 150
	maxPixelWidth  uint = 4000
	maxPixelHeight uint = 4000
	maxFileSize    int = 1024 * 1024 * 5 // 5 MB
)

// func ValidateImage(img *multipart.FileHeader) error {
// 	file, err := img.Open()
// 	if err != nil {
// 		return errors.New(`failed to open uploaded file`)
// 	}

// 	defer file.Close()

// 	imgConf, format, err := image.DecodeConfig(file)
// 	if err != nil {
// 		return errors.New(`uploaded file is not a valid image`)
// 	}

// 	switch format {
// 	case `jpeg`, `jpg`, `png`:
// 		break
// 	default:
// 		return errors.New(`invalid image extension, must be jpeg, jpg, and png`)
// 	}

// 	if imgConf.Width != imgConf.Height {
// 		return errors.New("uploaded image does not have the same dimension, recommend to 500x500")
// 	}

// 	if (imgConf.Width < minPixelWidth) || (imgConf.Height < minPixelHeight) {
// 		return fmt.Errorf("uploaded image does not meet minimum pixel dimensions of %dx%d", minPixelWidth, minPixelHeight)
// 	}

// 	if (imgConf.Width > maxPixelWidth) || (imgConf.Height > maxPixelHeight) {
		
// 	}

// 	max := float64(maxFileSize) / 1000000
// 	if img.Size > maxFileSize {
// 		return fmt.Errorf("uploaded file exceeds maximum size of %.2f MB", max)
// 	}

// 	return nil
// }