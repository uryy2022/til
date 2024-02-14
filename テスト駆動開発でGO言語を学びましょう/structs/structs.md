#### 最初にテストを書く

- <https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#nitesutowoku>

<https://pkg.go.dev/fmt>
> %f decimal point but no exponent, e.g. 123.456

```
    t.Errorf("got %.2f want %.2f", got, want)
```

#### 成功させるのに十分なコードを書く

- <https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#saserunoninakdowoku>

```
package main

import "testing"

func TestPerimeter(t *testing.T) {
 got := Perimeter(10.0, 10.0)
 want := 40.0

 if got != want {
  t.Errorf("got %.2f want %.2f", got, want)
 }
}

func TestArea(t *testing.T) {
 got := Area(12.0, 6.0)
 want := 72.0

 if got != want {
  t.Errorf("got %.2f want %.2f", got, want)
 }
}

```

```
package main

import "fmt"

func Perimeter(width float64, height float64) float64 {
 return 2 * (width + height)
}

func Area(width float64, height float64) float64 {
 return width * height
}

func main() {
 fmt.Println("--result--")
 fmt.Println("Perimeter:", Perimeter(10.0, 10.0))
 fmt.Println("Area:", Area(12.0, 6.0))
}

```

#### リファクタリング1

- <https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#rifakutaringu>

```
/Users/abe-hiroaki/sandbox_uryy/til/テスト駆動開発でGO言語を学びましょう/structs/structs_test.go:22:14: not enough arguments in call to Area
 have (Rectangle)
 want (float64, float64)
FAIL golang-test/structs [build failed]
FAIL
```

```
package main

import "testing"

func TestPerimeter(t *testing.T) {
 rectangle := Rectangle{10.0, 10.0}
 got := Perimeter(rectangle)
 want := 40.0

 if got != want {
  t.Errorf("got %.2f want %.2f", got, want)
 }
}

func TestArea(t *testing.T) {
 rectangle := Rectangle{12.0, 6.0}
 got := Area(rectangle)
 want := 72.0

 if got != want {
  t.Errorf("got %.2f want %.2f", got, want)
 }
}

```

```
package main

import "fmt"

type Rectangle struct {
 Width  float64
 Height float64
}

func Perimeter(rectangle Rectangle) float64 {
 return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
 return rectangle.Width * rectangle.Height
}

func main() {
 fmt.Println("--result--")
 fmt.Println("Perimeter:", Perimeter(Rectangle{10.0, 10.0}))
 fmt.Println("Area:", Area(Rectangle{12.0, 6.0}))
}

```

#### リファクタリング

- <https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#nitesutowoku-1>

```
package main

import "testing"

func TestArea(t *testing.T) {

 t.Run("rectangles", func(t *testing.T) {
  rectangle := Rectangle{12.0, 6.0}
  got := Area(rectangle)
  want := 72.0

  if got != want {
   t.Errorf("got %g want %g", got, want)
  }

 })
 t.Run("circles",func(t *testing.T)) {
  circle:= Circle{10}
  got:= Area(circle)
  want:= 314.1592653589793

  if got != want {
   t.Errorf("got %g want %g", got, want)
  }
 }
}

```

- f は gに置き換えられています。fを使用すると、正確な10進数を知るのが難しい場合があります。gを使用すると、エラーメッセージで完全な10進数が表示されます

ここから

- <https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#tesutowosurutamenonokdowoshishitatesutowoshimasu-1>

```
 t.Run("circles", func(t *testing.T) {
  circle := Circle{10.0}
  got := CircleCalculator(circle)
  //want := 0.0
  want := 314.1592653589793

  if got != want {
   t.Errorf("got %g want %g", got, want)
  }
 })
```

以下で、Piを使えば良さそう

```
func CircleCalculator(circle Circle) float64 {
 return math.Pi * circle.Radius * circle.Radius
}
```

## メソッドとは？

- メソッドは、レシーバーを持つ関数

メソッドレシーバーの構文
> func (receiverName ReceiverType) MethodName(args)

```
type Rectangle struct {
 Width  float64
 Height float64
}

func (r Rectangle) Area() float64 {
 return r.Width * r.Height
}

type Circle struct {
 Radius float64
}

func (c Circle) Area() float64 {
 return 0
}
```

ここでは、Areaメソッドが定義されており、レシーバ型として、Rectangleと Circleがバインドされている。r Rectangleここで、レシーバとして宣言されている。rはレシーバ変数。

```

func TestArea(t *testing.T) {

 t.Run("rectangles", func(t *testing.T) {
  rectangle := Rectangle{12, 6}
  got := rectangle.Area()
  want := 72.0

  if got != want {
   t.Errorf("got %g want %g", got, want)
  }
 })

 t.Run("circles", func(t *testing.T) {
  circle := Circle{10}
  got := circle.Area()
  want := 314.1592653589793

  if got != want {
   t.Errorf("got %g want %g", got, want)
  }
 })

}

```

```

type Rectangle struct {
 Width  float64
 Height float64
}

func (r Rectangle) Area() float64 {
 return r.Width * r.Height
}

type Circle struct {
 Radius float64
}

func (c Circle) Area() float64 {
 return math.Pi * c.Radius * c.Radius
}


```

golangのレシーバにはポインタと値の２種類がある
<https://recursionist.io/learn/languages/go/oop/method>

## リファクタリング

<https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#rifakutaringu-1>

#### インターフェイス

- Goのインターフェースは、メソッドシグネチャの集合体です。
<https://recursionist.io/learn/languages/go/oop/interface>
<https://qiita.com/rtok/items/46eadbf7b0b7a1b0eb08>
- Goでは、interfaceの中にある関数名と同一の関数が全て実装されている構造体に自動的に実装される

```
func TestArea(t *testing.T) {

 checkArea := func(t *testing.T, shape Shape, want float64) {
  t.Helper()
  got := shape.Area()
  if got != want {
   t.Errorf("got %g want %g", got, want)
  }
 }

 t.Run("rectangles", func(t *testing.T) {
  rectangle := Rectangle{12, 6}
  checkArea(t, rectangle, 72.0)
 })

 t.Run("circles", func(t *testing.T) {
  circle := Circle{10}
  checkArea(t, circle, 314.1592653589793)
 })
}
```

```
type Shape interface {
 Area() float64
}
```

を追加すればOK

```
type Shape interface {
 Area() float64
}

type Rectangle struct {
 Width  float64
 Height float64
}

func (r Rectangle) Area() float64 {
 return r.Width * r.Height
}

type Circle struct {
 Radius float64
}

func (c Circle) Area() float64 {
 return math.Pi * c.Radius * c.Radius
}

func main() {
 rectangle := Rectangle{12, 6}
 fmt.Println("Area:", rectangle.Area())

 circle := Circle{10}
 fmt.Println("Area:", circle.Area())
}


```

- Area() メソッドはShapeインターフェイスの一部として定義されている
- func (r Rectangle) Area() float64　の中では Area()はfloat型の戻り値となる
- main関数内の処理
  - rectangle := Rectangle{12, 6}
    - `Rectangle`型の変数 `rectangle` がRectangle{12, 6}で初期化される
  - `rectangle.Area()`
    - `rectangle.Area()`　メソッドが呼び出される。`rectangle`は`Width`が12、`Height`が6の`Rectangle`オブジェクト。
    - `Area()`メソッドの実行として、Shape interfaceが実装されている`Rectangle`型に対する`Area`メソッドの実装に従って、`Rectangle`オブジェクトの面積が計算される。12 * 6 = 72 となる。

```
❯ go run ./structs.go
Area: 72
Area: 314.1592653589793
```

## ちょっとまってなぜ

<https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#chottenaze>

> Goでは、インターフェースの解決は暗黙的です。 渡したタイプがインターフェースが要求するものと一致する場合、それはコンパイルされます。

下記の`Rectangle`型に対する`Area`メソッドの場合、r.Width * r.Heightはfloat64型なので、`Area`メソッドの戻り値であるfloat64型と一致されるのでコンパイルされる。

```
func (r Rectangle) Area() float64 {
 return r.Width * r.Height
}
```

### 切り離し

<https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#rishidecoupling>
