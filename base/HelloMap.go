package base

import "fmt"

func Map() {
	kv := map[string]string{"1": "2"}
	for key, value := range kv {
		fmt.Printf("key %s,value %s", key, value)
		fmt.Println()
	}
}
