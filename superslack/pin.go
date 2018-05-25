package superslack

// Pin describes a single pin
type Pin struct {
	ID     string
	Author Author
	Text   string
}

// PinStore can be used to store multiple Pins
type PinStore []Pin
