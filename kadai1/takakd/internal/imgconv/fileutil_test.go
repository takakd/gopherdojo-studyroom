package imgconv

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestConvertImageInDirectory(t *testing.T) {

	// Setup
	_, filename, _, _ := runtime.Caller(0)
	testdataDir, _ := filepath.Abs(filepath.Join(filepath.Dir(filename), "testdata"))
	outputTestDataDir := filepath.Join(testdataDir, "output")
	tests := []struct {
		name string
		opt  *ConvertImageOption
	}{
		{name: "png-jpg", opt: &ConvertImageOption{SrcFormat: "png", DstFormat: "jpg", DstDirPath: outputTestDataDir}},
	}
	if err := os.RemoveAll(outputTestDataDir); err != nil {
		t.Error("failed to cleanup")
	}

	// Run
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srcBaseDirPath := filepath.Join(testdataDir, "/", string(test.opt.SrcFormat))
			err := ConvertImageInDirectory(srcBaseDirPath, test.opt)
			if err != nil {
				t.Error(fmt.Sprintf("ConvertImageInDirectory should be succeed. err=%v", err))
				return
			}

			srcCount := 0
			filepath.Walk(srcBaseDirPath, func(path string, f os.FileInfo, err error) error {
				if !f.IsDir() {
					return nil
				}
				if test.opt.SrcFormat.isCorrectExt(path) && isValidFile(path, test.opt.SrcFormat) {
					srcCount++
				}
				return nil
			})

			dstCount := 0
			var dstFilePathList []string
			filepath.Walk(test.opt.DstDirPath, func(path string, f os.FileInfo, err error) error {
				if !f.IsDir() {
					return nil
				}
				if test.opt.DstFormat.isCorrectExt(path) && isValidFile(path, test.opt.DstFormat) {
					dstCount += 1
					dstFilePathList = append(dstFilePathList, path)
				}
				return nil
			})

			if srcCount != dstCount {
				t.Errorf(fmt.Sprintf("failed to convert files in directory. got=%d, want=%d", srcCount, dstCount))
			}

			// Clean up
			os.RemoveAll(test.opt.DstDirPath)
		})
	}
}
