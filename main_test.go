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
	reg := NewProductSortingStrategyRegistry()
	reg.SetStrategy(StrategyPrice, PriceSortStrategy{})
	sortedProducts, err := reg.ExecuteStrategy(StrategyPrice, testProducts())
	if err != nil {
		t.Fatalf("Expected strategy to be found, got error: %v", err)
	}

	if sortedProducts[0].Price != 5 {
		t.Errorf("Expected product with lowest price first, got %.2f", sortedProducts[0].Price)
	}
	if sortedProducts[2].Price != 15 {
		t.Errorf("Expected product with highest price last, got %.2f", sortedProducts[2].Price)
	}
}

func TestSalesPerViewRatioSortStrategy(t *testing.T) {
	reg := NewProductSortingStrategyRegistry()
	reg.SetStrategy(StrategySalesPerViewRatio, SalesPerViewRatioSortStrategy{})
	sortedProducts, err := reg.ExecuteStrategy(StrategySalesPerViewRatio, testProducts())
	if err != nil {
		t.Fatalf("Expected strategy to be found, got error: %v", err)
	}

	r1 := float64(sortedProducts[0].SalesCount) / float64(sortedProducts[0].ViewsCount)
	r2 := float64(sortedProducts[1].SalesCount) / float64(sortedProducts[1].ViewsCount)

	if r1 < r2 {
		t.Errorf("Expected descending sales/view ratio but got %.4f < %.4f", r1, r2)
	}
}

func TestRegistry(t *testing.T) {
	reg := NewProductSortingStrategyRegistry()
	reg.SetStrategy(StrategyPrice, PriceSortStrategy{})

	sortedProducts, err := reg.ExecuteStrategy(StrategyPrice, testProducts())
	if err != nil {
		t.Fatalf("Expected strategy to be found, got error: %v", err)
	}

	if sortedProducts[0].Price != 5 {
		t.Errorf("Registry did not return correct strategy result")
	}
}

func TestDateSortStrategy(t *testing.T) {
	reg := NewProductSortingStrategyRegistry()
	reg.SetStrategy(StrategyDate, DateSortStrategy{})
	sortedProducts, err := reg.ExecuteStrategy(StrategyDate, testProducts())
	if err != nil {
		t.Fatalf("Expected strategy to be found, got error: %v", err)
	}

	if sortedProducts[0].Created.After(sortedProducts[1].Created) {
		t.Errorf("Expected products to be sorted by creation date")
	}
}

func TestUnknownStrategy(t *testing.T) {
	reg := NewProductSortingStrategyRegistry()
	_, err := reg.ExecuteStrategy("nonexistent", testProducts())

	if err == nil {
		t.Error("Expected error for unknown strategy")
	}
}
