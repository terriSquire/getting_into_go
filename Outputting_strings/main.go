package main

import "fmt"

func main() {
	// The Println function, and others on which it's based, will return the
	// requested value and an error object.
	// Because the Println function will return these two variables, you need
	// to declare them before calling the function.
	str1 := "Once upon a time"
	str2 := "not long ago"
	str3 := "I worked with dough"
	aNumber := 47
	isTrue := true

	stringLength, err := fmt.Println(str1, str2, str3)
	if err == nil {
		fmt.Println("String length:", stringLength)

		// There will be times when you're getting values back from a function
		// and you know you're not going to get an error object back. However,
		// the function will still return the error object, and because you've
		// already declared that value (err), you must address it in your code.

		// To get around this, you can "throw away" the variable by naming
		// the variable with a single underscore character like so:
		//stringLength, _ := fmt.Println(str1, str2, str3)

		// You can also format stings using the Printf function, which lets
		// you create strings that have placeholders known as verbs:
		fmt.Printf("Value of aNumber: %v\n", aNumber)
		fmt.Printf("Value of isTrue: %v\n", isTrue)
		fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber))

		fmt.Printf("Data types: %T, %T, %T, %T, %T\n",
			str1, str2, str3, aNumber, isTrue)
		// Be sure to add linefeed (\n) when using Printf bc it does not
		// automatically create its own new line

		// Sprintf will return a string
		myString := fmt.Sprintf("Data types as var: %T, %T, %T, %T, %T\n",
			str1, str2, str3, aNumber, isTrue)
		fmt.Print(myString)
	}

}
