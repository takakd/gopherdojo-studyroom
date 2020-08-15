package imgconv

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// `ConvertImageInDirectory` func option parameter
type ConvertImageOption struct {
	SrcFormat  ImageFormat
	DstFormat  ImageFormat
	DstDirPath string
}

// Check if values are valid.
//
// return true or false, and the detail of invalid if false.
func (c *ConvertImageOption) validate() (bool, string) {
	if c.SrcFormat != PNG && c.SrcFormat != GIF && c.SrcFormat != JPG {
		return false, "src format is not supported"
	}
	if c.DstFormat != PNG && c.DstFormat != GIF && c.DstFormat != JPG {
		return false, "dst format is not supported"
	}
	if c.DstDirPath == "" {
		return false, "dst directory path is empty"
	}
	return true, ""
}

// Convert images in the directory including lower directories.
func ConvertImageInDirectory(srcDirPath string, opt *ConvertImageOption) error {

	if opt.SrcFormat == opt.DstFormat {
		log.Println("not need to convert")
		return nil
	}

	if !fileExists(srcDirPath) {
		return errors.New(fmt.Sprintf("directory does not exist. path=%s", srcDirPath))
	}

	if !fileExists(opt.DstDirPath) {
		if err := os.MkdirAll(opt.DstDirPath, 0700); err != nil {
			return errors.New(fmt.Sprintf("failed to output directory. path=%s", opt.DstDirPath))
		}
	}

	err := filepath.Walk(srcDirPath,
		func(srcFilePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			isConvertTarget := opt.SrcFormat.isCorrectExt(srcFilePath)
			if !isConvertTarget {
				return nil
			}

			dstDirPath := strings.Replace(filepath.Dir(srcFilePath), srcDirPath, opt.DstDirPath, 1)
			if err := ConvertImage(srcFilePath, dstDirPath, opt.SrcFormat, opt.DstFormat); err != nil {
				log.Printf("failed to convert file. src=%s, dstDir=%s", srcFilePath, dstDirPath)
				return nil
			}

			return nil
		})
	if err != nil {
		return err
	}

	return nil
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
