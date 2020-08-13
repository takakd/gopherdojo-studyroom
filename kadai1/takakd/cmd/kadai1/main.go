// 画像変換コマンド
package main

import (
	"flag"
	"fmt"
	"github.com/takakd/gopherdojo-studyroom/kadai1/takakd/internal/imgconv"
	"log"
	"os"
)

var convertFrom, convertTo *string
var dirname string

func init() {
	convertFrom = flag.String("srcfmt", "jpg", `output image format, suport "jpg", "png", and "gif"`)
	convertTo = flag.String("dstfmt", "png", `input image format, support "jpg", "png", and "gif"`)
	flag.Parse()

	if len(os.Args) < 2 {
		usage := `
Usage: kadai1 [-srcfmt format] [-dstfmt format] directory_path
`
		fmt.Println(usage)
		flag.PrintDefaults()
		// Add a space for easy to see.
		fmt.Println("")
		os.Exit(0)
	}

	dirname = os.Args[1]
}

func main() {
	opt := &imgconv.ConvertImageOption{
		DstDirPath: dirname,
	}
	opt.SrcFormat = imgconv.ImageFormat(*convertFrom)
	opt.DstFormat = imgconv.ImageFormat(*convertTo)

	if err := imgconv.ConvertImageInDirectory(dirname, opt); err != nil {
		log.Println(err)
	}

	log.Println("done.")
}
