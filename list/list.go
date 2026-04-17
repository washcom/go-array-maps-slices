package list

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	hobbies := [3]string{"coding", "gaming", "travelling"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[1:])

	mainhobbies := hobbies[0:2]
	fmt.Println(mainhobbies)

	fmt.Println(cap(mainhobbies))
	mainhobbies = mainhobbies[1:3]
	fmt.Println(mainhobbies)

	courseGoals := []string{"Learn Go", "Build Projects", "Contribute to Open Source"}
	fmt.Println(courseGoals)

	courseGoals[1] = "learn all the details"
	courseGoals = append(courseGoals, "learn all the basics")
	fmt.Println(courseGoals)

	products := []Product{
		{
			id:    "p1",
			title: "Book",
			price: 9.99,
		},
		{
			id:    "p2",
			title: "Laptop",
			price: 999.99,
		},
		{
			id:    "p3",
			title: "Phone",
			price: 499.99,
		},
	}
	fmt.Println(products)
	fourthProduct := Product{
		id:    "p4",
		title: "Headphones",
		price: 199.99,
	}
	anyotherProduct := Product{
		id:    "p5",
		title: "Monitor",
		price: 299.99,
	}
	products = append(products, fourthProduct, anyotherProduct)
	fmt.Println(products)

}
