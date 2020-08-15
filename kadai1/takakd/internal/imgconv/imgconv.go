// 画像変換モジュール
package imgconv

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Supported image type
type ImageFormat string

const (
	PNG ImageFormat = "png"
	GIF ImageFormat = "gif"
	JPG ImageFormat = "jpg"
)

// Get file extension.
func (f *ImageFormat) ext() string {
	ext := ""
	if *f == PNG {
		ext = "png"
	} else if *f == GIF {
		ext = "gif"
	} else if *f == JPG {
		ext = "jpg"
	}
	return ext
}

func (f *ImageFormat) isCorrectExt(srcFilePath string) bool {
	ext := filepath.Ext(srcFilePath)
	correct := strings.ToLower(ext) == "."+f.ext()
	return correct
}

// Convert image file.
func ConvertImage(srcPath, dstDirPath string, srcFormat ImageFormat, dstFormat ImageFormat) error {
	if srcFormat == dstFormat {
		log.Println(fmt.Sprintf("not need to convert because format is same. format=%s", srcFormat))
		return nil
	}

	if !srcFormat.isCorrectExt(srcPath) {
		return fmt.Errorf("file does not match format. path=%s, format=%s", srcPath, srcFormat)
	}

	img, err := decodeImage(srcPath, srcFormat)
	if err != nil {
		return err
	}

	dstPath := buildDstFilepath(srcPath, dstDirPath, dstFormat.ext())
	if err := os.MkdirAll(filepath.Dir(dstPath), 0700); err != nil {
		return errors.New(fmt.Sprintf("failed to create directory. path=%s", filepath.Dir(srcPath)))
	}

	if err := encodeImage(dstPath, img, dstFormat); err != nil {
		return err
	}

	return nil
}

func buildDstFilepath(srcFilePath, dstDirPath, dstFileExt string) string {
	filename := filepath.Base(srcFilePath)
	fileExt := filepath.Ext(filename)

	// e.g. def.png -> def.png
	dstFilename := fmt.Sprintf("%s.%s", strings.TrimSuffix(filename, fileExt), dstFileExt)

	return filepath.Join(dstDirPath, dstFilename)
}

func decodeImage(filename string, format ImageFormat) (image.Image, error) {
	var srcFile *os.File

	defer func() {
		if srcFile != nil {
			if err := srcFile.Close(); err != nil {
				log.Printf("could not close file. name=%s\n", srcFile.Name())
			}
		}
	}()

	srcFile, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	var img image.Image
	switch format {
	case PNG:
		img, err = png.Decode(srcFile)
	case GIF:
		img, err = gif.Decode(srcFile)
	case JPG:
		img, err = jpeg.Decode(srcFile)
	}
	if err != nil {
		return nil, err
	}

	return img, nil
}

func encodeImage(filename string, img image.Image, format ImageFormat) error {
	dstFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func() {
		if closeErr := dstFile.Close(); closeErr != nil {
			log.Printf("could not close file. name=%s\n", dstFile.Name())
		}
	}()

	switch format {
	case PNG:
		err = png.Encode(dstFile, img)
	case GIF:
		err = gif.Encode(dstFile, img, nil)
	case JPG:
		err = jpeg.Encode(dstFile, img, nil)
	}
	if err != nil {
		return err
	}

	return nil
}
