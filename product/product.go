package product

type Product struct {
	Id             int    `json:"id"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerunit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	Name           string `json:"name"`
}
