package main
import "fmt"

func main() {
  //var a int
  a := 5
  var b int = 2
  //var c int
  //fmt.Printf("%d\n", a / b)

if a%b == 0 {
        fmt.Println("a is even")
    } else {
        fmt.Println("a is odd")
    }

if ( a + b > 100 ) || ( a - b > 100 ) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }

fmt.Println("Can I have a loan of £200?")

if ( 50 > 100 ) && ( 1000 > 100 ) {
        fmt.Println("Yes, you have enough dosh")
    } else {
        fmt.Println("Bog off")
    }
fmt.Println("Can i have a small loan of one grand?")

if  ( 5 * 52 * 5 >= 1000 ) {
        fmt.Println("Yes i will grant you a small loan of 1000")
    } else {
        fmt.Println("Get a job")
    }
}
