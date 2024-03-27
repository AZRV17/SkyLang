package service

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/AZRV17/Skylang/internal/repository"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

const dir = "filestorage"

type ImageService struct {
	userRepo   repository.Users
	courseRepo repository.Courses
}

func NewImageService(userRepo repository.Users, courseRepo repository.Courses) *ImageService {
	return &ImageService{
		userRepo:   userRepo,
		courseRepo: courseRepo,
	}
}

func (i ImageService) SetCourseImage(id int, imgBase64 string) error {
	imageData, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return err
	}

	img, format, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		return err
	}

	// Генерировать случайное имя файла для изображения
	randBytes := make([]byte, 16)
	_, err = rand.Read(randBytes)
	if err != nil {
		return err
	}
	filename := hex.EncodeToString(randBytes)
	outputPath := filepath.Join(dir, fmt.Sprintf("%s.%s", filename, format))

	err = saveImage(img, format, outputPath)
	if err != nil {
		return err
	}

	absolutePath, err := filepath.Abs(outputPath)
	if err != nil {
		return err
	}

	return i.courseRepo.SetCourseIcon(id, absolutePath)
}

func (i ImageService) SetUserAvatar(id int, avatar string) error {
	imageData, err := base64.StdEncoding.DecodeString(avatar)
	if err != nil {
		return err
	}

	img, format, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		return err
	}

	// Генерировать случайное имя файла для изображения
	randBytes := make([]byte, 16)
	_, err = rand.Read(randBytes)
	if err != nil {
		return err
	}
	filename := hex.EncodeToString(randBytes)
	outputPath := filepath.Join(dir, fmt.Sprintf("%s.%s", filename, format))

	err = saveImage(img, format, outputPath)
	if err != nil {
		return err
	}

	absolutePath, err := filepath.Abs(outputPath)
	if err != nil {
		return err
	}

	_, err = i.userRepo.SetUserAvatar(id, absolutePath)

	return err
}

func saveImage(resizedImg image.Image, originalFormat string, outputPath string) error {
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(outFile)

	switch originalFormat {
	case "jpeg":
		err = jpeg.Encode(outFile, resizedImg, nil)
	case "png":
		err = png.Encode(outFile, resizedImg)
	default:
		return fmt.Errorf("unsupported format: %s", originalFormat)
	}

	return err
}
