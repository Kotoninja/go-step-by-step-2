package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// u := User{Name: "Bob", Age: 18}

	users := []User{{Name: "Bob", Age: 18}, {Name: "Alex", Age: 20}}
	fmt.Println(users)
	jsonData, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Create("config.json")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", " ")

	fmt.Println(string(jsonData))

	if err := encoder.Encode(users); err != nil {
		fmt.Println("Ошибка записи в файл")
	}
}
