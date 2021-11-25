package main

import (
	"convert"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
    if err := run(os.Args); err != nil {
        fmt.Fprintln(os.Stderr, "Error: ", err)
        os.Exit(1)
    }
}

func run(args []string) error {
    // ディレクトリ名が引数にあるかどうか
    if len(args) <= 1 {
        return errors.New("ディレクトリを指定してください")
    }

    err := filepath.WalkDir(os.Args[1], func(path string, info fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        // ディレクトリは除く
        if info.IsDir() || filepath.Ext(path) != ".png" {
            return nil
        }
        
        fmt.Println(path, info)
        if err := convert.Convert(path); err != nil {
            return err
        }
        return nil
    })

    if err != nil {
        return err
    }

    return nil
}