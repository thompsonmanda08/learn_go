package main

import "fmt"

func main() {

	website := map[string]string{
		"Google":  "http://www.google.com",
		"Youtube": "http://www.youtube.com",
	}

	fmt.Println(website)
}

/* ARRAYS */
// func main() {

// 	hobbies := [3]string{"Coding", "Gaming", "Sleeping"}

// 	fmt.Println("Hobbies: ", hobbies)
// 	fmt.Println("First Element ", hobbies[0])
// 	fmt.Println("Second and Third ", hobbies[1:])

// 	products := []arrays.Product{}
// 	prod1, err := arrays.New("Product 1", 100.0, "prod1")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	prod2, err := arrays.New("Product 2", 200.0, "prod2")

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	products = append(products, prod1, prod2)

// 	fmt.Println(products)
// }
