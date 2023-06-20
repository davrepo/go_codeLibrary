// Go slices
package main

import (
	"fmt"
)

func main() {
	loons := []string{"bugs", "daffy", "taz"}

	for i := range loons {
		fmt.Println(i)
	} // 0 1 2

	for i, name := range loons {
		fmt.Printf("idx: %d, val: %s; ", i, name)
	} // idx: 0, val: bugs; idx: 1, val: daffy; idx: 2, val: taz;

	for _, name := range loons {
		fmt.Println(name)
	} // bugs daffy taz

	loons = append(loons, "elmer")
	fmt.Println(loons) // [bugs daffy taz elmer]
}
