# common-code

いろんな言語で共通の問題を解決する

## 目的

1. どんな言語でも書き方はだいたい同じであることを知る。
1. 次新しい言語を学ぶ時に、この処理の実装の仕方がわからない、とかを未然に防ぐ

## 問題

1. コマンドライン引数
1. ファイル読み込み
1. ファイル読み込み(1行ずつ)
1. ファイル出力
1. 日付
1. 文字コード処理(cp932←→utf-8)
1. エラーハンドリング
1. 正規表現
1. 集計
1. テストコード
1. HTTP Request
1. HTML処理
1. 外部ライブラリ
1. JSON処理
1. 画像処理

### コマンドライン引数

#### 要件

1. コマンドライン引数から数値を２つ受け取って、その和を標準出力する
1. 引数が渡されなかったときの対策をする

#### 理由

1. 実行時に値を切り替えたいことはよくある
1. 引数の扱いは言語によって異なる
1. 型変換を知る
1. 不測の事態に備えるのはアプリ実装者の領域であることを知る

### ファイル読み込み

#### 要件

1. コマンドライン引数からファイル名を受け取る
1. 受け取ったファイル名のファイルを読み込む
1. **読み込む際、一気にファイル内の全データを読み込む**
1. ファイルはutf-8エンコーディング、改行コードはLF
1. 行頭に行番号を0埋め3桁で出力する(001:foobar)
1. 空白のみの行は出力しない

#### 理由

1. とりあえず、全部のデータを読み込んで出力する方法を知る
1. 単純な実装で済むため、実装コストが低い
1. 扱うファイルサイズが小さい場合はこっちのが実装が楽

### ファイル読み込み(1行ずつ)

#### 要件

1. 前述の「ファイル読み込み」の要件のうち、3番目以外すべて満たす
1. **ファイルを1行ずつ読み込む**

#### 理由

1. 巨大なファイルを処理するとき、全データを読み込もうとすると激遅になるので1行ず
   つ処理する方法を知る
1. 処理対象のファイルサイズが巨大か否かで実装の仕方を変える

### ファイル出力

1. 何らかのデータをファイルに保存する

### 日付

1. ファイルを作るときに日付でファイル名を変えたいときがあるから

### 文字コード

1. Windowsで仕事してると度々文字コードが仕事の邪魔をするため