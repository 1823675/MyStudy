package main

var x, y int
var ( //shegnming quanju bianliang
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

/*不带声明格式的，只能在函数体中出现
g,h :=123,"hello"*/

func main() {
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
}
