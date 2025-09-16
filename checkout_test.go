package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckout(t *testing.T) {
	actualPrice := make(map[string]float64)
	actualPrice["ipd"] = 549.99
	actualPrice["mbp"] = 1399.99
	actualPrice["atv"] = 109.50
	actualPrice["vga"] = 30.00

	// Price rules can be modified
	// OfferAcceptableQuantity - quantity for which offer should be enabled
	// Offer price depending on condition
	// conditions > , < , =
	PriceRule := make(map[string]OfferPrice)
	PriceRule["atv"] = OfferPrice{
		OfferAcceptableQuantity: 3,
		OfferPrice:              73,
		Condition:               "=",
	}
	PriceRule["ipd"] = OfferPrice{
		OfferAcceptableQuantity: 4,
		OfferPrice:              499.99,
		Condition:               "<",
	}

	co := Rules{
		pricingRules: PriceRule,
		ActualPrices: actualPrice,
	}
	co.Scan([]string{"atv", "atv", "atv", "vga"})
	// expected 249.00
	assert.Equal(t, 249.00, co.Total())
	co.Scan([]string{"atv", "ipd", "ipd", "atv", "ipd", "ipd", "ipd"})
	// expected 2718.95
	assert.Equal(t, 2718.95, co.Total())
}
