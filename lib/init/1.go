package init

import "fmt"

var ss = make([]string, 0)

func init() {
	fmt.Println("1")
	for _, s := range ss {
		fmt.Printf(">>>%s<<<\n", s)
	}
}
