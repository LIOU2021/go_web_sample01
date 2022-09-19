package person

import (
	"fmt"
	"math"
	"net/http"
)

// 宣告名為幾何的interface
type geometry interface {
	//定義了area 面積與perim的function
	area() float64
	perim() float64
}

// 定義名為矩形的結構
type Rect struct {
	width, height float64
}

// 定義名為圓形的結構
type circle struct {
	radius float64
}

// 矩形的求面積 function
func (r Rect) area() float64 {
	return r.width * r.height
}

// 矩形的求周長 function
func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// 圓形的求面積 function
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// 圓形的求周長 function
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 定義一個測量的function來進行計算
func Measure(g geometry, w http.ResponseWriter) {
	fmt.Fprintln(w, g)
	fmt.Fprintln(w, g.area())
	fmt.Fprintln(w, g.perim())
	fmt.Fprintln(w, "==========================")
}
