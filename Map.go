package main
import(
	"fmt"
)

func main() {
	month := map[string]int{}

	month["januari"] = 31
	month["februari"] = 29
	month["maret"] = 31

	fmt.Println(month["januari"])
	fmt.Println(month["februari"])
	fmt.Println(month["maret"])
	fmt.Println(month["april"])

	fmt.Println()

	// Map, sama seperti array asosiasi
	zodiacs := map[string]string{
		"libra": "september - october",
		"scorpio": "october - november",
		"sagittarius": "november - december",
	}

	for zodiac, month := range zodiacs {
	    fmt.Println(zodiac, ":", month)
	}

	fmt.Println()

	// Multidimension Map
	divisions := []map[string]string{
		{"name":"backend programmer", "skills":"php, golang, nodejs", "salary":"4jt"},
		{"name":"frontend programmer", "skills":"javascript, css, jquery", "salary":"4jt"},
		{"name":"grphic designer", "skills":"photoshop, ilustrator, coreldraw", "salary":"4.5jt"},
	}

	for _, division := range divisions {
	    for key, val := range division {
		    fmt.Println(key, ":", val)
		}

		fmt.Println()
	}
}