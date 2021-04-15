// main.go

package main

func main() {
	a := App{}
	// You need to set your Username and Password here
	a.Initialize("admos", "J4n1c#2020", "learn_golang")

	a.Run(":3000")
}
