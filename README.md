# zeller

Coding challenge for zeller.

Clone repository locally. 
Run go test ./...
Can manually trigger from checkout_test.go.

actualPrice := make(map[string]float64) 
Map for actual price of product 
key-product name/id
Value - actual value of product

PriceRule := make(map[string]OfferPrice)
type OfferPrice struct {
	OfferAcceptableQuantity int
	OfferPrice              float64
	Condition               string
}
for price rules. more details in test file.
