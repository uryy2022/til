# 配列とスライス

## 配列初期化

- numbers := [5]int{1, 2, 3, 4, 5}
- numbers := [...]int{1, 2, 3, 4, 5}

array_test.go

```
package main

import "testing"

func TestSum(t *testing.T) {

 numbers := [5]int{1, 2, 3, 4, 5}

 got := Sum(numbers)
 want := 15

 if got != want {
  t.Errorf("got %d want %d given, %v", got, want, numbers)
 }
}

```

array.goを書く

```
package main

import "fmt"

func Sum(numbers [5]int) int {
    sum := 0
    for i := 0; i < 5; i++ {
        sum += numbers[i]
    }
    return sum
}

func main() {
 numbers := [5]int{1, 2, 3, 4, 5}
 fmt.Println(Sum(numbers))
}

```

- for i:= 0 : カウンター初期化
- i < 5 : 繰り返し条件
- i++  : 1ループ毎実行

リファクタリング

<https://gobyexample.com/>

```
func Sum(numbers [5]int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}
```

<https://go.dev/doc/effective_go#blank>

`for _` の `_` は「ブランク識別子」。宣言したが無視したい変数がある場合に設定。
range は indexとindexにある要素の２つの値を返すが、indexを無視する。

`[5]int` で配列のサイズも型に含まれるので長さが固定される
「配列」ではなく任意のサイズをエンコードできる「slices」がよい

#### テストを書く

```
package main

import "testing"

func TestSum(t *testing.T) {
 t.Run("collection of 5 numbers", func(t *testing.T) {

  numbers := [5]int{1, 2, 3, 4, 5}

  got := Sum(numbers)
  want := 15

  if got != want {
   t.Errorf("got %d want %d given, %v", got, want, numbers)

  }
 })
 t.Run("collection of any size", func(t *testing.T) {
  numbers := []int{1, 2, 3, 4}

  got := Sum(numbers)
  want := 10

  if got != want {
   t.Errorf("got %d want %d given, %v", got, want, numbers)

  }
 })
}

```

コードを修正

```
package main

func Sum(numbers []int) int {
 sum := 0
 for _, number := range numbers {
  sum += number
 }
 return sum
}

```

テストを修正

```
package main

import "testing"

func TestSum(t *testing.T) {
 t.Run("collection of 5 numbers", func(t *testing.T) {

  //  numbers := [5]int{1, 2, 3, 4, 5}
  numbers := []int{1, 2, 3, 4, 5}

  got := Sum(numbers)
  want := 15

  if got != want {
   t.Errorf("got %d want %d given, %v", got, want, numbers)

  }
 })
 t.Run("collection of any size", func(t *testing.T) {
  numbers := []int{1, 2, 3, 4}

  got := Sum(numbers)
  want := 10

  if got != want {
   t.Errorf("got %d want %d given, %v", got, want, numbers)

  }
 })
}

```

```
❯ go test -v
=== RUN   TestSum
=== RUN   TestSum/collection_of_5_numbers
=== RUN   TestSum/collection_of_any_size
--- PASS: TestSum (0.00s)
    --- PASS: TestSum/collection_of_5_numbers (0.00s)
    --- PASS: TestSum/collection_of_any_size (0.00s)
PASS
ok   array 0.241s
```

カバレッジツール
<https://go.dev/blog/cover>

- カバーしていないコードの領域を特定するのに役立つ

```
❯ go test -cover
PASS
coverage: 100.0% of statements
ok   array 0.573s
```

テストケースを１つ削って実施

```
❯ go test -cover
PASS
coverage: 100.0% of statements
ok   array 0.515s```

#### リファクタリング♪

```

func TestSumAll(t *testing.T) {

 got := SumAll([]int{1, 2}, []int{0, 9})
 want := []int{3, 9}

 if got != want {
  t.Errorf("got %v want %v", got, want)
 }
}

```

```

❯ go test -v

# array [array.test]

./array_test.go:37:5: invalid operation: got != want (slice can only be compared to nil)
FAIL array [build failed]

```

- Goでは、スライスで等号演算子を使うことはできないのでDeepEqualを利用する
https://pkg.go.dev/reflect#DeepEqual


```

❯ go test -v
(略)
    array_test.go:41: got [] want [3 9]
--- FAIL: TestSumAll (0.00s)
FAIL
exit status 1
FAIL array 0.217s

```

テスト結果は下記となる。
slice と string を比較しようとしている。
sum_test.go:30: got [] want [3 9]

func SumAllを修正する

make関数
- makeを使用することでsliceを指定することができる
- https://y-hiroyuki.xyz/go/slice/make-func



```

func Sum(numbers []int) int {
 sum := 0
 for _, number := range numbers {
  sum += number
 }
 return sum
}

func SumAll(numbersToSum ...[]int) []int {
 lengthOfNumbers := len(numbersToSum)

 sums := make([]int, lengthOfNumbers)

 for i, numbers := range numbersToSum {
  sums[i] = Sum(numbers) // sums[i] = Sum(numbers)では、`int` のスライス（`[]int{11, 2}` と `[]int{0, 1, 9}`）の全要素の和を計算。
 }

 return sums // 合計値を含むスライスを呼び出し元に返す
}

func main() {
 fmt.Println(SumAll([]int{11, 2}, []int{0, 1, 9}))
}

```

#### リファクタリング2
https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#rifakutaringu-2

```

func SumAllTails(numbersToSum ...[]int) []int {
    var sums []int
    for _, numbers := range numbersToSum {
        // 1つ目の要素を除いたスライスを作成
        tail := numbers[1:]
        sums = append(sums, Sum(tail))
    }

    return sums
}

```

#### リファクタリング3
https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#rifakutaringu-3

```

func TestSumAllTails(t *testing.T) {

 t.Run("make the sums of some slices", func(t *testing.T) {
  got := SumAllTails([]int{1, 2}, []int{0, 9})
  want := []int{2, 9}

  if !reflect.DeepEqual(got, want) {
   t.Errorf("got %v want %v", got, want)
  }
 })

 t.Run("safely sum empty slices", func(t *testing.T) {
  got := SumAllTails([]int{}, []int{3, 4, 5})
  want := []int{0, 9}

  if !reflect.DeepEqual(got, want) {
   t.Errorf("got %v want %v", got, want)
  }
 })

}

```
```

func SumAllTails(numbersToSum ...[]int) []int {
 var sums []int
 for _, numbers := range numbersToSum {
  // numbersの長さが0の場合
  if len(numbers) == 0 {
   // sumsに0を追加
   sums = append(sums, 0)
   fmt.Println("sums", sums)
  } else {
   // tailsに1つ目の要素を除いたスライスを作成
   tail := numbers[1:]
   fmt.Println("tail", tail)

   // sumsにSum(tail)を追加
   sums = append(sums, Sum(tail))
  }
 }

 return sums
}
func main() {
 fmt.Println(SumAllTails([]int{}, []int{0, 1, 9}))
}

```


#### リファクタリング4
https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#rifakutaringu-4

```

func TestSumAllTails(t *testing.T) {

 // checkSumsに関数を代入
 checkSums := func(t *testing.T, got, want []int) {
  t.Helper()
  if !reflect.DeepEqual(got, want) {
   t.Errorf("got %v want %v", got, want)
  }
 }

 t.Run("make the sums of some slices", func(t *testing.T) {
  got := SumAllTails([]int{1, 2}, []int{0, 9})
  want := []int{2, 9}
  checkSums(t, got, want)

 })

 t.Run("safely sum empty slices", func(t *testing.T) {
  got := SumAllTails([]int{}, []int{3, 4, 5})
  want := []int{0, 9}
  checkSums(t, got, want)

 })
}

```
