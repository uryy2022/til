# integers

## 最初にテストを書く

- 整数を利用するため、%dを利用
-

adder_test.go

```
package main

import "testing"

func TestAdder(t *testing.T) {
 sum := Add(2, 2)
 expected := 4

 if sum != expected {
  t.Errorf("expected '%d' but got '%d'", expected, sum)
 }
}

```

adder.go

```
package main

import "fmt"

func Add(x, y int) int {
 return x + y
}

func main() {
 fmt.Println(Add(3, 3))
}

```

```
❯ go run ./adder.go
4
❯ go test -v
=== RUN   TestAdder
--- PASS: TestAdder (0.00s)
PASS
ok   golang-test/integers 0.463s
```

adder_test.goにExampleAddを追加

```
package main

import (
 "fmt"
 "testing"
)

func TestAdder(t *testing.T) {
 sum := Add(2, 2)
 expected := 4

 if sum != expected {
  t.Errorf("expected '%d' but got '%d'", expected, sum)
 }
}

func ExampleAdd() {
 sum := Add(1, 5)
 fmt.Println(sum)
 // Output: 6
}

```

```
❯ go test -v
=== RUN   TestAdder
--- PASS: TestAdder (0.00s)
=== RUN   ExampleAdd
--- PASS: ExampleAdd (0.00s)
PASS
ok   golang-test/integers 0.536s
```

' // Output: 6' を削除してtestしてみるとサンプル関数は実行されない

```
❯ go test -v
=== RUN   TestAdder
--- PASS: TestAdder (0.00s)
PASS
ok   golang-test/integers 0.788s
```
