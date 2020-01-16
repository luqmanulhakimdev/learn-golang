package main
import(
	"fmt"
)

func main() {
	i := 0
	for i < 5 {
		fmt.Println("Angka:", i)
		i++
	}

	fmt.Println()
	for n := 10; n >= 0; n-- {
		fmt.Println("Angka:", n)
	}

	fmt.Println()
	for x := 0; x < 5; x++ {
	    for j := x; j < 5; j++ {
	        fmt.Print(j, " ")
	    }

	    fmt.Println()
	}

	fmt.Println()
	for y := 5; y >= 0; y-- {
	    for m := y; m < 5; m++ {
	        fmt.Print(m, " ")
	    }

	    fmt.Println()
	}
}