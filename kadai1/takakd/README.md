# 課題1 - Gopher道場 自習室

[Gopher道場 自習室](https://gopherdojo.org/studyroom/)の`【TRY】画像変換コマンドを作ろう`の実装です。  
ディレクトリ配下の画像ファイルを変換します。

## 課題の回答について

要件 | 対応方法
--- | ---
mainパッケージ分離 | imgconvモジュールに変換処理を分離
自作・標準・準標準パッケージのみ使用 | その通りにした
ユーザー定義型を使う | imgconv.goで`ImageFormat`を使用
GoDoc生成 | `godoc -http=:8080` で`http://localhost:8080/pkg/?m=all`に列挙されていることを確認
Go Modulesを使ってみる | `go mod init`後、`go get ...`で`go.mod`, `go.sum`に列挙されることを確認（※）

※処理で利用するものが無かったので後に中身をクリア。

## Usage

```shell
./cmd/kadai1/kadai1 [-srcfmt format] [-dstfmt format] directory_path

  -dstfmt string
        入力画像のフォーマット, "jpg", "png", "gif"のいずれかを指定。
  -srcfmt string
        出力する画像のフォーマット, "jpg", "png", "gif"のいずれかを指定。
        
  directory_path
        変換する画像を含むディレクトリのフルパス
```
        
e.g. /somewhere/images内のpngファイルをjpgに変換

```shell
$ kadai1 -srcfmt png -dstfmt jpg /somewhere/images
```

## 環境

* go version go1.14.4 darwin/amd64
* macOS 10.15.5

## 動作確認方法

```shell
$ tree internal/imgconv/testdata 
internal/imgconv/testdata
├── gif
│   ├── sample.gif
│   └── sub
│       └── sample.gif
├── jpg
│   ├── sample.jpg
│   └── sub
│       └── sample.jpg
└── png
    ├── error.png
    ├── sample.png
    └── sub
        └── sample.png

6 directories, 7 files

$ file internal/imgconv/testdata/gif/sample.gif 
internal/imgconv/testdata/gif/sample.gif: GIF image data, version 87a, 400 x 400

$ make build
==> Formatting...
==> Building...
$ cmd/kadai1/kadai1 -srcfmt gif -dstfmt png internal/imgconv/testdata
2020/08/14 07:47:31 done.

$ tree internal/imgconv/testdata                                     
internal/imgconv/testdata
├── gif
│   ├── sample.gif
│   ├── sample.png
│   └── sub
│       ├── sample.gif
│       └── sample.png
├── jpg
│   ├── sample.jpg
│   └── sub
│       └── sample.jpg
└── png
    ├── error.png
    ├── sample.png
    └── sub
        └── sample.png

6 directories, 9 files

$ file internal/imgconv/testdata/gif/sample.png                     
internal/imgconv/testdata/gif/sample.png: PNG image data, 400 x 400, 8-bit colormap, non-interlaced
```

## 実行方法

ビルド

```shell
$ make build
# -> ./cmd/kadai1/kadai1に実行バイナリが生成されます。
```

テスト

```shell
$ make test
```

他

```shell
# コードフォーマット
$ make fmt

# godoc
$ make godoc

# gosec
$ make gosec
```

## ファイル構成

```shell
./
├── Makefile
├── README.md
├── cmd
│   └── kadai1          mainモジュール
│       ├── kadai1      実行ファイル
│       └── main.go     main
├── go.mod
└── internal
    ├── imgconv                 画像変換モジュール
    │   ├── fileutil.go         ディレクトリ単位の画像変換処理
    │   ├── fileutil_test.go
    │   ├── imgconv.go          画像変換処理
    │   └── imgconv_test.go
    └── testdata    テストデータ
        ├── gif
        │   ├── sample.gif
        │   └── ...
        ├── jpg
        │   ├── sample.jpg
        │   └── ...
        └── png
            ├── sample.png
            └── ...
```
