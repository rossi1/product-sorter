package main

import (
	"testing"
	"time"
)

func testProducts() []*Product {
	return []*Product{
		{ID: 1, Name: "A", Price: 15, Created: time.Now(), SalesCount: 10, ViewsCount: 100},
		{ID: 2, Name: "B", Price: 5, Created: time.Now(), SalesCount: 5, ViewsCount: 10},
		{ID: 3, Name: "C", Price: 10, Created: time.Now(), SalesCount: 50, ViewsCount: 1000},
	}
}

func TestPriceSortStrategy(t *testing.T) {
	strategy := PriceSortStrategy{}
	products := testProducts()
	sorted := strategy.Sort(products)

	if sorted[0].Price != 5 {
		t.Errorf("Expected product with lowest price first, got %.2f", sorted[0].Price)
	}
	if sorted[2].Price != 15 {
		t.Errorf("Expected product with highest price last, got %.2f", sorted[2].Price)
	}
}

func TestSalesPerViewSortStrategy(t *testing.T) {
	strategy := SalesPerViewRatioSortStrategy{}
	products := testProducts()
	sorted := strategy.Sort(products)

	r1 := float64(sorted[0].SalesCount) / float64(sorted[0].ViewsCount)
	r2 := float64(sorted[1].SalesCount) / float64(sorted[1].ViewsCount)

	if r1 < r2 {
		t.Errorf("Expected descending sales/view ratio but got %.4f < %.4f", r1, r2)
	}
}

func TestRegistry(t *testing.T) {
	reg := NewProductSortingStrategyRegistry()
	reg.SetStrategy(StrategyPrice, PriceSortStrategy{})

	strategy, err := reg.GetStrategy(StrategyPrice)
	if err != nil {
		t.Fatalf("Expected strategy to be found, got error: %v", err)
	}

	products := testProducts()
	sorted := strategy.Sort(products)
	if sorted[0].Price != 5 {
		t.Errorf("Registry did not return correct strategy result")
	}
}

func TestDateSortStrategy(t *testing.T) {
	strategy := DateSortStrategy{}
	products := testProducts()
	sorted := strategy.Sort(products)

	if sorted[0].Created.After(sorted[1].Created) {
		t.Errorf("Expected products to be sorted by creation date")
	}
}

func TestUnknownStrategy(t *testing.T) {
	reg := NewProductSortingStrategyRegistry()
	_, err := reg.GetStrategy("nonexistent")

	if err == nil {
		t.Error("Expected error for unknown strategy")
	}
}
