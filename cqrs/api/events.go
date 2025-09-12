package main

type PlaceOrderCommand struct {
	RequestID string   `json:"request_id"`
	Customer  Customer `json:"customer"`
	Products  []Product
	Address   Address
}

type OrderPlaced struct {
	RequestID string    `json:"request_id"`
	OrderID   string    `json:"order_id"`
	Customer  Customer  `json:"customer"`
	Address   Address   `json:"address"`
	Products  []Product `json:"products"`
}

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
