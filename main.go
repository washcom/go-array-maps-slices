package main

import "fmt"

func main() {
	userNames := []string{}
	userNames = append(userNames, "Alice")
	userNames = append(userNames, "Bob")
	userNames = append(userNames, "Charlie")

	fmt.Println(userNames)

	for index, value := range userNames {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}
