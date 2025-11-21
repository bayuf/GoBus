package dto

// struct request
type Request struct {
	Name        string
	Destination string
}

func NewRequest(name, destination string) Request {
	return Request{Name: name, Destination: destination}
}
