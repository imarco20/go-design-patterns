package main

// Purpose:
//		- Allows for the construction of objects when the types of
//		  those objects is not predetermined at runtime
//
// Scenarios:
//		- Produces code that is more readable when there are multiple
//		  ways of creating a particular object
//		- Situations where object creation needs to be flexible and
//		  cannot be known beforehand

import "fmt"

func newPublication(pubType string, name string, pg int, pub string) (iPublication, error) {
	switch pubType {
	case "newspaper":
		return createNewspaper(name, pg, pub), nil
	case "magazine":
		return createMagazine(name, pg, pub), nil
	default:
		return nil, fmt.Errorf("no such publication type [%s]", pubType)
	}
}
