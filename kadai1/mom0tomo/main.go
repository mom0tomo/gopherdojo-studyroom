// go build
// ./main img/ png
package main

import (
	"flag"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

var (
	flagDir     = flag.String("directory", "img", "ディレクトリ")
	flagImgType = flag.String("image type", "png", "画像タイプ")
)

func main() {
	//引数を取得する
	flag.Parse()
	args := flag.Args()
	switch {
	case len(args) == 2:
		*flagDir = args[0]
		*flagImgType = args[1]
	default:
		*flagDir = args[0]
	}

	if err := convertFile(*flagDir, *flagImgType); err != nil {
		log.Fatal(err)
	}
}

func convertFile(d string, t string) error {
	//ディレクトリを指定する
	files, err := os.ReadDir(d)
	if err != nil {
		log.Fatal(err)
	}

	//ディレクトリ以下は再帰的に処理する
	for _, file := range files {
		//指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
		f, err := os.Open(d + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		img, err := jpeg.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		name := (strings.Split(file.Name(), "."))[0]

		fso, err := os.Create("./img/" + name + ".png")
		if err != nil {
			log.Fatal(err)
		}
		png.Encode(fso, img)
		f.Close()
	}

	// //変換前と変換後の画像形式を指定できる（オプション）
	// fileType := t

	// sc := bufio.NewScanner(f)
	// for sc.Scan() {
	// 	fmt.Println(sc.Text())
	// }
	return err
}
