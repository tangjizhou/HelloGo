package base

import (
	"fmt"
	"os"
)

func Slice() {
	var slice = []int{0, 1}
	slice = append(slice, 0)
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(cap(slice), len(slice))

	for key, arg := range os.Args {
		fmt.Println(key, arg)
	}
	for i, v := range slice {
		fmt.Println(&i, &v, &slice[i])
	}

}
