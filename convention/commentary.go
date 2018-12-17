/*
Package convention contends some examples for "2 Coding Conventions" of the book "Go At Hand"

Includes:
	- Formatting
	- Commentary
	- Names
And other stuff.
*/
package main

// Language represent an ISO 639-1 language
type Language struct {
	code       string //ISO 639-1 language code, such as en, zh
	name       string //name in English
	nativeName string //self-explained?? need extra documentation??
}

// GetLanguageName returns the English name of the code
func GetLanguageName(code string) (name string) {
	return "pseudo"
}
