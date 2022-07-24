
package main
import "fmt"

// type Walker interface{
//     printMe()
// }

type Test struct {
				arg string
}

func (t *Test)printMe() {
				fmt.Println(t.arg)
				fmt.Println("printed")
}

func main() {
				var testTest = Test{"testArg"}
				testTest.printMe()
}
