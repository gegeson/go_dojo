package main

import (
	"convert"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
    from := flag.String("f", ".jpg", "extention before convert")
	to := flag.String("t", ".png", "extention after convert")
	dir := flag.String("d", ".", "target dir")
	flag.Parse()

    err := filepath.WalkDir(*dir, func(path string, info fs.DirEntry, err error) error {
        if err != nil {
            return err
        }

        // ディレクトリ・対象ファイルがない場合は除く
        if info.IsDir() || filepath.Ext(path) != *from {
            return nil
        }
        
        // fmt.Println(path, info)
        if err := convert.Convert(path, *from, *to); err != nil {
            return err
        }
        return nil
    })

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error: ", err)
        os.Exit(1)
    }
    
}
