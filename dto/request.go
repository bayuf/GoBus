package dto

// struct request
type Request struct {
	Name        string
	Destination string
	Price       float64
}

func NewRequest(name, destination string) Request {
	return Request{Name: name, Destination: destination}
}

func NewDest(destination string, price float64) Request {
	return Request{Destination: destination, Price: price}
}
