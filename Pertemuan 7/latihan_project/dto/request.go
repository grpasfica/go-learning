package dto

type Request struct {
	Name        string
	Destination string
}

// Constructor untuk membuat request baru
func NewRequest(name string, destination string) Request {
	return Request{
		Name:        name,
		Destination: destination,
	}
}
