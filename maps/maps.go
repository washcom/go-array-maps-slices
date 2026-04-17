package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://www.google.com",
		"amazon": "https://www.amazon.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["amazon"])

	websites["linkedin"] = "https://www.linkedin.com"
	websites["whatsapp"] = "https://www.whatsapp.com"
	fmt.Println(websites)

	delete(websites, "amazon")
	fmt.Println(websites)

}
