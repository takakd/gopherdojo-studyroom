package imgconv

import (
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestConvertImage(t *testing.T) {

	// Setup
	_, filename, _, _ := runtime.Caller(0)
	testdataDir, _ := filepath.Abs(filepath.Join(filepath.Dir(filename), "testdata"))
	outputTestDataDir := filepath.Join(testdataDir, "output")
	tests := []struct {
		name    string
		opt     *ConvertImageOption
		srcPath string
		dstPath string
	}{
		{name: "png-jpg", opt: &ConvertImageOption{SrcFormat: "png", DstFormat: "jpg", DstDirPath: outputTestDataDir}},
		{name: "png-gif", opt: &ConvertImageOption{SrcFormat: "png", DstFormat: "gif", DstDirPath: outputTestDataDir}},
		{name: "gif-jpg", opt: &ConvertImageOption{SrcFormat: "gif", DstFormat: "jpg", DstDirPath: outputTestDataDir}},
		{name: "gif-png", opt: &ConvertImageOption{SrcFormat: "gif", DstFormat: "png", DstDirPath: outputTestDataDir}},
		{name: "jpg-gif", opt: &ConvertImageOption{SrcFormat: "jpg", DstFormat: "gif", DstDirPath: outputTestDataDir}},
		{name: "jpg-png", opt: &ConvertImageOption{SrcFormat: "jpg", DstFormat: "png", DstDirPath: outputTestDataDir}},
	}
	for i, _ := range tests {
		tests[i].srcPath = filepath.Join(testdataDir, "/", string(tests[i].opt.SrcFormat), "sample."+string(tests[i].opt.SrcFormat))
		tests[i].dstPath = buildDstFilepath(tests[i].srcPath, outputTestDataDir, string(tests[i].opt.DstFormat))
	}

	cleanup := func() {
		if removeErr := os.RemoveAll(outputTestDataDir); removeErr != nil {
			t.Errorf("failed to delete files. dir=%s", outputTestDataDir)
		}
	}
	cleanup()

	// Run
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := ConvertImage(test.srcPath, outputTestDataDir, test.opt.SrcFormat, test.opt.DstFormat); err != nil {
				t.Error(fmt.Sprintf("should be succeed. err=%v", err))
				return
			}

			if !fileExists(test.dstPath) {
				t.Error("dst file should exists")
				return
			}

			convertSucceeded := isValidFile(test.dstPath, test.opt.DstFormat)
			if !convertSucceeded {
				t.Error(fmt.Sprintf("failed to convert. path=%s", test.dstPath))
				return
			}

			os.Remove(test.dstPath)
		})
	}

	// Clean up
	cleanup()
}

func isValidFile(srcFilePath string, f ImageFormat) bool {
	srcFile, err := os.OpenFile(srcFilePath, os.O_RDONLY, 0600)
	if err != nil {
		return false
	}

	if f == PNG {
		_, err = png.Decode(srcFile)
	} else if f == GIF {
		_, err = gif.Decode(srcFile)
	} else if f == JPG {
		_, err = jpeg.Decode(srcFile)
	}

	valid := err == nil
	return valid
}
