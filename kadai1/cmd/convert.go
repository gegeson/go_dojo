package convert

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

const (
	JPG = "jpg"
	JPEG = "jpeg"
	PNG = "png"
	GIF = "gif"
)

// 画像の拡張子を変換する
func Convert(path string, from string, to string) error {
    src, err := os.Open(path)
    if err != nil {
        return err
    }
    defer src.Close()

    output := strings.Replace(path, from, "", 1)
    dest, err := os.Create(output + to)
    if err != nil {
        return err
    }
    defer dest.Close()

    img, _, err := image.Decode(src)
    if err != nil {
        return errors.New("Decode失敗")
    }


    // .以降を取り出して条件分岐 その後image.Imageからエンコード
    switch to {
    case ".png":
        err = png.Encode(dest, img)
    case ".jpg", ".jpeg":
       // 第三引数は画質で1~100まで。（今回は定数を使用:75)
        err = jpeg.Encode(dest, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
    case ".gif":
        err = gif.Encode(dest, img, nil)
    }

    if err != nil {
        return errors.New("Encode失敗")
    }

    return nil
}