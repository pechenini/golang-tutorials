package main
import "fmt"

func main() {
	var personPhones = map[string]string{
		"John":  "+33-8273658526",
		"Steve": "+1-8579822345",
		"David": "+44-9462834443",
	}

	var mobileNo = personPhones["Steve"]
	fmt.Println("Steve's Mobile No : ", mobileNo)

	// If a key doesn't exist in the map, we get the zero value of the value type
	// there is also second param returned, which will be true if value exists
	mobileNo, exists := personPhones["Jack"]
	if exists {
		fmt.Println("Jack's PhoneNo : ", mobileNo)
	}

	//delete key
	delete(personPhones, "John")
	fmt.Println(personPhones)

	for key, value := range personPhones {
		fmt.Println(key, value)
	}
}