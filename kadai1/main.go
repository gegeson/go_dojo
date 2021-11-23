package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"convert"
)

// 第一引数に名前、第二引数にデフォルト値、第三引数に使い方を渡す。戻り値はフラグの値が格納されるそれぞれの型の変数のアドレス。
var from = flag.String("f", ".jpeg", "Extension before conversion")
var to = flag.String("t", ".png", "Extension after conversion")
var rm = flag.Bool("r", false, "Remove file before conversion")

func main() {
    flag.Parse() // コマンドライン引数がパースされ、ポインタの指す先に値が設定される。
    src := flag.Arg(0)
    fmt.Println(*from, *to, *rm)
    // 第一引数の文字列のコピーに対し、第四引数個、第二引数の部分を第三引数に置換したものを返す。
    dst := strings.Replace(flag.Arg(0), *from, *to, 1)

    err := convert.Convert(src, dst, *rm)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
    }
    fmt.Fprintln(os.Stdout, "end image convert")
}