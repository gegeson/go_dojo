package convert

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

const (
	JPG = "jpg"
	JPEG = "jpeg"
	PNG = "png"
	GIF = "gif"
)

func ExtensionCheck(ext string) error {
	switch ext {
	case JPG, JPEG, PNG, GIF:
		return nil
	default:
		return errors
	}
}


// 画像の拡張子を変換する
func Convert(src, dst string) error {
	fmt.Fprintln(os.Stdout, "start convert")

    // 変換前のファイルを読み取り専用で開く
    sf, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sf.Close()

    // image.Imageへとデコード
    img, _, err := image.Decode(sf)
    if err != nil {
        return err
    }

    // 読み書き用ファイルを作成
    df, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer df.Close()

    // .以降を取り出して条件分岐 その後image.Imageからエンコード
    switch filepath.Ext(dst) {
    case ".png":
        err = png.Encode(df, img)
    case ".jpg", ".jpeg":
       // 第三引数は画質で1~100まで。（今回は定数を使用:75)
        err = jpeg.Encode(df, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
    case ".gif":
        err = gif.Encode(df, img, nil)
    }

    if err != nil {
        return err
    }

    return nil
}