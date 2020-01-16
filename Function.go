package main
import(
	"fmt"
	"strconv"
)

func main() {
	divisions := []map[string]string{
		{"name":"backend programmer", "skills":"php, golang, nodejs", "salary":"4jt"},
		{"name":"frontend programmer", "skills":"javascript, css, jquery", "salary":"4jt"},
		{"name":"grphic designer", "skills":"photoshop, ilustrator, coreldraw", "salary":"4.5jt"},
	}

	newDivisions := multiMapToMap(divisions)
	for index, division := range newDivisions {
		fmt.Println(index, "=>", division)
	}
}

func multiMapToMap(divisions []map[string]string) (map[string]string) {
	newDivisions := map[string]string{}
	for index, division := range divisions {
		valString := ""
	    for key, val := range division {
		    valString += " "+key+" : "+val+","
		}

		newDivisions[strconv.Itoa(index)] = valString
	}

	return newDivisions
}