# iteration

## リファクタリング

iteration_test.go

```
package main

import "testing"

func TestRepeat(t *testing.T) {
 repeated := Repeat("a")
 expected := "aaaaa"

 if repeated != expected {
  t.Errorf("expected %q but got %q", expected, repeated)
 }
}
```

iteration.go

```
package main

import "fmt"

const repeatCount = 5

func Repeat(character string) string {
 var repeated string
 for i := 0; i < repeatCount; i++ {
  repeated = repeated + character
  fmt.Println(repeated)
 }
 return repeated
}

func main() {
 fmt.Printf("result=%q", Repeat("a"))
}

```

## ベンチマーク

- ベンチマーク追加

```
package main

import "testing"

func TestRepeat(t *testing.T) {
 repeated := Repeat("a")
 expected := "aaaaa"

 if repeated != expected {
  t.Errorf("expected %q but got %q", expected, repeated)
 }
}
func BenchmarkRepeat(b *testing.B) {
 for i := 0; i < b.N; i++ {
  Repeat("a")
 }
}
```

ベンチマークを実行してみる

- ベンチマーク用の関数はBenchmarkで始める必要あり
- `*testing.B` はtestingパッケージから提供される型
- testing.BはNというフィールドを持っており、このフィールドはベンチマークを実行する回数を示す。
- <https://sagantaf.hatenablog.com/entry/2023/07/21/235340>

```
func BenchmarkRepeat(b *testing.B) {
 for i := 0; i < b.N; i++ {
  Repeat("a")
 }
 //Output: aaaaaa
}

```

```
❯ go test -bench=.
goos: darwin
goarch: amd64
pkg: golang-test/iteration
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkRepeat-8    10324899        111.9 ns/op
PASS
ok   golang-test/iteration 1.760s
```

## 練習問題

- テストを変更して、発信者が文字が繰り返される回数を指定し、コードを修正できる
- const repeatCount を追加

```
package main

import "fmt"

const repeatCount = 5

func Repeat(character string) string {
 var repeated string
 for i := 0; i < repeatCount; i++ {
  repeated = repeated + character
  //  fmt.Println(repeated)
 }
 return repeated
}

func main() {
 Repeats := Repeat("a")
 fmt.Printf("result=%q", Repeats)
}

```

- 関数をドキュメント化するために ExampleRepeat を記述します

- go.modが必要

```
❯ pwd
/Users/abe-hiroaki/sandbox_uryy/til/テスト駆動開発でGO言語を学びましょう/iteration
❯ go mod init iteration
❯ cat go.mod
module iteration

go 1.21.4
```

godocコマンドを実施

```
❯ godoc -http=:6060
using module mode; GOMOD=/Users/abe-hiroaki/sandbox_uryy/til/テスト駆動開発でGO言語を学びましょう/iteration/go.mod
go: no module dependencies to download

```

<http://localhost:6060/pkg/>
にアクセスすると、`iteration` がある
