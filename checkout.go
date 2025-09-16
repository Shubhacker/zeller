package main

// interface for methods
type Checkout interface {
	Scan([]string)
	Total() float64
}

// Required structs - Later can be moved to independent package specific for model
type Rules struct {
	pricingRules map[string]OfferPrice
	ActualPrices map[string]float64
	order        map[string]float64
}

type OfferPrice struct {
	OfferAcceptableQuantity int
	OfferPrice              float64
	Condition               string
}

func (r *Rules) Scan(ord []string) {
	order := make(map[string]float64)

	for _, dt := range ord {
		num, ok := order[dt]
		if ok {
			order[dt] = num + 1
			continue
		}
		order[dt] = 1
	}

	r.order = order
}

func (r *Rules) Total() float64 {
	var orderValue float64
	for item, quantity := range r.order {
		var price float64
		price = quantity * r.ActualPrices[item]
		data, ok := r.pricingRules[item]
		if ok {
			price = RuleConditions(quantity, r.ActualPrices[item], data)
		}
		orderValue = orderValue + price
	}
	return orderValue
}

// RuleConditions - Helper function
func RuleConditions(orderQuantity float64, actualPrice float64, data OfferPrice) float64 {
	var orderValue float64

	switch data.Condition {
	case "<":
		var price float64
		price = orderQuantity * actualPrice

		if data.OfferAcceptableQuantity <= int(orderQuantity) {
			price = data.OfferPrice * orderQuantity
		}
		orderValue = orderValue + price
	case ">":
		var price float64
		price = orderQuantity * actualPrice

		if data.OfferAcceptableQuantity >= int(orderQuantity) {
			price = data.OfferPrice * orderQuantity
		}
		orderValue = orderValue + price
	case "=":
		var price float64
		price = orderQuantity * actualPrice

		if data.OfferAcceptableQuantity == int(orderQuantity) {
			price = data.OfferPrice * orderQuantity
		}
		orderValue = orderValue + price
	}

	return orderValue
}
