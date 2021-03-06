# 課題 1

## 画像変換コマンドを作ろう

### 次の仕様を満たすコマンドを作って下さい

- ディレクトリを指定する
- 指定したディレクトリ以下の PNG ファイルを JPG に変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）

### 以下を満たすように開発してください

- main パッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
- 準標準パッケージ：golang.org/x 以下のパッケージ
- ユーザ定義型を作ってみる
- GoDoc を生成してみる
- Go Modules を使ってみる
- test も実装

### command

```
# 実行
go run main.go -d sample -f .png -t .jpg

# 作成した画像消去
sh remove.sh
```

### 備考

### 参考 code

- [[Go] 画像の拡張子の変更をする CLI コマンドを作る](https://github.com/marnysan111/gopherdojo-studyroom/tree/kadai1-marny/kadai1/marny)

## io.Reader と io.Writer について

io.Reader も io.Writer もインターフェースであり、汎用性が高くさまざまものに利用されている.

- [io.Reader をすこれ](https://qiita.com/ktnyt/items/8ede94469ba8b1399b12)
