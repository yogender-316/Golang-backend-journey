package main

import (
	"fmt"
)

func main() {
	// Go does not have enums directly.
		// We usually create enum-like constants with iota.
		type LogLevel int
		const (
			Info LogLevel = iota
			Warning
			Error
		)
	fmt.Println("Info level:", Info)
	fmt.Println("Warning level:", Warning)
	fmt.Println("Error level:", Error)	


	userAccess := map[string]bool{
		"admin": true,
		"user":  false,

	}
	fmt.Println("User access:", userAccess)	


	if userAccess["admin"] {
		fmt.Println("Admin access granted")
	}	else {
		fmt.Println("Admin access denied")
	}

	if hasaccess, ok := userAccess["ds"]; ok && !hasaccess {
		fmt.Println("User access granted")
	} else
	{
		fmt.Println("User access denied",9)
	}
}