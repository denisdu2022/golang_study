package main

func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func main() {

	foo, bar := test(10)
	foo() //10
	bar() //20

}
