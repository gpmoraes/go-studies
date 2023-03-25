package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {

	// 1) Create a new array (!) that contains three hobbies you have. Output (print) that array in the command line.
	hobbies := [3]string{"Sports", "Reading", "Cooking"}
	fmt.Println(hobbies)

	// 2) Also Output more data aout that array:
	// 	- The first element (standalone)
	// 	- The second and third element combined a new list
	fmt.Printf("First element: %v \n", hobbies[0])
	fmt.Printf("Second and Third elements: %v \n", hobbies[1:3])

	// 3) Create a sliece based on the first element that contains the first and second elements.
	// 	Create that slice in two different ways (i.e. create two slices in the end)
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4) Re-slice the slice from (3) and change it to contain the second and last element of the original array.
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"Learn all the basics", "Learn advanced Go"}
	fmt.Println(courseGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Lear all the details!"
	courseGoals = append(courseGoals, "Learn advanced Go")
	fmt.Println(courseGoals)

	// 7) Bonus: Create a "Product" struct with tittle, id, price and create a dynamic list of products (at least 2 products).
	// 	Then add a third product to the existing list of products.
	products := []Product{
		{
			"first-product",
			"A first product",
			12.99,
		},
		{
			"second-product",
			"A second product",
			13.99,
		},
	}
	fmt.Println(products)

	newProduct := Product{
		"third-product",
		"A third product",
		14.99,
	}

	products = append(products, newProduct)
	fmt.Println(products)
}
