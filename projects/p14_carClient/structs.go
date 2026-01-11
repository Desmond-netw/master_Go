package main

// Manufacture struct

type Manufacturer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Country      string `json:"country"`
	FoundingYear int    `json:"fundingYear"`
}

// Category struct
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"Name"`
}

// Specification ( nested struct)
type Specifications struct {
	Engine       string `json:"engine"`
	Horsepower   string `json:"horsepower"`
	Transmission string `json:"transmission"`
	Drivetrain   string `json:"drivetrain"`
}

// Car Model Struct represent one model from the API
type CarModel struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	ManufacturerID int            `json:"manufacturerId"`
	CategoryId     int            `json:"categoryId"`
	Year           int            `json:"year"`
	Specifications Specifications `json:"specification"`
	Image          string         `json:"image"`
}
