// go build
// ./main img/ png
package main

import (
	"flag"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

var (
	flagDir              = flag.String("directory", "img", "ディレクトリ")
	flagOriginalImgType  = flag.String("original image type", "jpg", "変換前画像タイプ")
	flagConvertedImgType = flag.String("converted image type", "png", "変換後画像タイプ")
)

func main() {
	flag.Parse()
	args := flag.Args()

	switch {
	case len(args) == 2:
		*flagOriginalImgType = args[1]
	case len(args) == 3:
		*flagConvertedImgType = args[2]
	default:
		*flagDir = args[0]
	}

	if err := convertImg(*flagDir, *flagOriginalImgType, *flagConvertedImgType); err != nil {
		log.Fatal(err)
	}
}

func convertImg(d string, ft string, fd string) error {
	//ディレクトリを指定する
	files, err := os.ReadDir(d)
	if err != nil {
		log.Fatal(err)
	}

	//ディレクトリ以下は再帰的に処理する
	for _, file := range files {
		f, err := os.Open(d + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		name := (strings.Split(file.Name(), "."))[0]

		var img image.Image

		//指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
		//変換前と変換後の画像形式を指定できる（オプション）
		switch ft {
		case "png":
			img, err = png.Decode(f)
			if err != nil {
				log.Fatal(err)
			}
		case "jpeg, jpg":
			img, err = jpeg.Decode(f)
			if err != nil {
				log.Fatal(err)
			}
		case "gif":
			img, err = gif.Decode(f)
			if err != nil {
				log.Fatal(err)
			}
		default:
			img, err = jpeg.Decode(f)
			if err != nil {
				log.Fatal(err)
			}
		}

		switch fd {
		case "png":
			fso, err := os.Create("./img/" + name + ".png")
			if err != nil {
				log.Fatal(err)
			}
			png.Encode(fso, img)
			f.Close()
		case "jpeg, jpg":
			fso, err := os.Create("./img/" + name + ".jpg")
			if err != nil {
				log.Fatal(err)
			}
			jpeg.Encode(fso, img, nil)
			f.Close()
		case "gif":
			fso, err := os.Create("./img/" + name + ".jpg")
			if err != nil {
				log.Fatal(err)
			}
			gif.Encode(fso, img, nil)
			f.Close()
		default:
			fso, err := os.Create("./img/" + name + ".png")
			if err != nil {
				log.Fatal(err)
			}
			png.Encode(fso, img)
			f.Close()
		}
	}
	return err
}
