# Hello, World

## テストを書く

テストの作成は、関数の作成と同様であり、いくつかのルールがあります。

- xxx_test.goのような名前のファイルにある必要があります。
- テスト関数はTestという単語で始まる必要があります。
- テスト関数は1つの引数のみをとります。 t *testing.T
- *testing.T 型を使うには、他のファイルの fmt と同じように import "testing" が必要です。

### t.errorf

```
`t.Errorf("got %q want %q", got, want)`
```

- `got`変数と`want`変数を使用して、得られた結果（`got`）と期待する結果（`want`）を比較

## その他

- `go test -v` で -v　オプションをつけて実施すると、テストが成功したにも関わらずログを見れる

```
❯ go test -v
=== RUN   TestHello
=== RUN   TestHello/saying_hello_to_people
    hello_test.go:20: want: "Hello, Chris"
=== RUN   TestHello/say_'Hello,_World'_when_an_empty_string_is_supplied
--- PASS: TestHello (0.00s)
    --- PASS: TestHello/saying_hello_to_people (0.00s)
    --- PASS: TestHello/say_'Hello,_World'_when_an_empty_string_is_supplied (0.00s)
PASS
ok   hello 0.177s
```
