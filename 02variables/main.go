package main

import "fmt"

const LoginToken string = "uabuvhsja"

func main(){
	var username string = "Aniket";
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggenIn bool = false;
	fmt.Println(isLoggenIn)
	fmt.Printf("Variable is of type: %T \n", isLoggenIn)

	var smallVal uint8 = 255;
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45556262;
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat);

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable);

	// implicit type
	var website = "anikxt.codes"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken);

} 