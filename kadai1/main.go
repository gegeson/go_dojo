package main

import (
	"flag"
	"fmt"
	"os"

	"convert"
)

func main() {
    from := flag.String("f", ".jpg", "Extension before conversion")
    to := flag.String("t", ".png", "Extension after conversion")
    // コマンドライン引数がパースされ、ポインタの指す先に値が設定される。
    flag.Parse()
    src := flag.Arg(0)

    
    fmt.Fprintln(os.Stdout, "now converting!!")
    
    err := convert.Convert(src, dst)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
    }
    fmt.Fprintln(os.Stdout, "convert done!!")
}