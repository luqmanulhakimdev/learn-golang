package main
import "fmt"

func main() {
	fruits := [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruits {
	    fmt.Println("nama buah", fruit)
	}

	fmt.Println()
	numbers := [...]int{2, 3, 2, 4, 3}

	for i, num := range numbers {
		fmt.Println("Nomor:", i, num)
	}

	fmt.Println()
	
	// Slice (Referensi dari array)
	members := []string{"joko", "jono", "jamal", "juna"}
	for _, member := range members {
		fmt.Println("Member:", member)
	}
}
