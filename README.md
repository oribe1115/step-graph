# step-graph

`make test`: ユニットテストの一括実行

## 準備

以下の構成になるようにデータファイルを配置

```txt
.
├── Makefile
├── README.md
├── data
│   ├── sns
│   │   ├── links.txt
│   │   └── nicknames.txt
│   ├── stations
│   │   ├── edges.txt
│   │   └── stations.txt
│   └── wikipedia
│       ├── links.txt
│       └── pages.txt
├── go.mod
├── go.sum
├── lib
├── main.go
├── search
└── topic
```

## 実行

`go run main.go`でプログラムが起動

その後は標準入力を受け取ってのCLI形式で各探索の関数を呼び出す

### topic: SNS

1. 入力した二つのユーザーを最も少ないリンク数で辿るルートを幅優先探索で探索
2. 最も離れたユーザーの組み合わせを幅優先探索のアルゴリズムを使って探索

## topic: Stations

1. 入力した二つの駅を最も少ないリンク数で辿るルートを幅優先探索で探索
2. 入力した二つの駅を最も少ない所要時間で辿るルートをダイクストラ法で探索
3. 入力した駅から入力した時間ちょうどでたどり着ける駅をダイクストラ法で探索

### topic: Wikipedia

1. 入力した二つの単語の記事を最も少ないリンク数で辿るルートを幅優先探索で探索
2. 入力した単語の記事からリンク数が最も離れた記事を幅優先探索のアルゴリズムを使って探索
