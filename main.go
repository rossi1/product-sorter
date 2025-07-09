package main

import (
	"fmt"
	"sort"
	"time"
)

const (
	StrategyPrice             = "price"
	StrategySalesPerViewRatio = "sales_per_view_ratio"
	StrategyDate              = "date"
)

type Product struct {
	ID         int
	Name       string
	Price      float64
	Created    time.Time
	SalesCount int
	ViewsCount int
}

type SortStrategy interface {
	Sort([]*Product) []*Product
}

type PriceSortStrategy struct{}

func (p PriceSortStrategy) Sort(products []*Product) []*Product {
	if len(products) <= 1 {
		return products
	}
	sorted := make([]*Product, len(products))
	copy(sorted, products)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Price < sorted[j].Price
	})
	return sorted
}

type SalesPerViewRatioSortStrategy struct{}

func (s SalesPerViewRatioSortStrategy) Sort(products []*Product) []*Product {
	if len(products) <= 1 {
		return products
	}
	sorted := make([]*Product, len(products))
	copy(sorted, products)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].ViewsCount == 0 || sorted[j].ViewsCount == 0 {
			return sorted[i].ViewsCount < sorted[j].ViewsCount
		}
		r1 := float64(sorted[i].SalesCount) / float64(sorted[i].ViewsCount)
		r2 := float64(sorted[j].SalesCount) / float64(sorted[j].ViewsCount)
		return r1 > r2
	})
	return sorted
}

type DateSortStrategy struct{}

func (d DateSortStrategy) Sort(products []*Product) []*Product {
	if len(products) <= 1 {
		return products
	}
	sorted := make([]*Product, len(products))
	copy(sorted, products)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Created.Before(sorted[j].Created)
	})
	return sorted
}

type ProductSortingStrategyRegistry struct {
	strategies map[string]SortStrategy
}

func NewProductSortingStrategyRegistry() *ProductSortingStrategyRegistry {
	return &ProductSortingStrategyRegistry{
		strategies: make(map[string]SortStrategy),
	}
}

func (r *ProductSortingStrategyRegistry) SetStrategy(key string, strategy SortStrategy) {
	r.strategies[key] = strategy
}

func (r *ProductSortingStrategyRegistry) GetStrategy(key string) (SortStrategy, error) {
	strategy, ok := r.strategies[key]
	if !ok {
		return nil, fmt.Errorf("strategy not found for key: %s", key)
	}
	return strategy, nil
}

func parseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

func products() ([]*Product, error) {
	date1, err := parseDate("2019-01-04")
	if err != nil {
		return nil, fmt.Errorf("invalid date for product 1: %w", err)
	}

	date2, err := parseDate("2012-01-04")
	if err != nil {
		return nil, fmt.Errorf("invalid date for product 2: %w", err)
	}

	date3, err := parseDate("2014-05-28")
	if err != nil {
		return nil, fmt.Errorf("invalid date for product 3: %w", err)
	}

	products := []*Product{
		{
			ID:         1,
			Name:       "Alabaster Table",
			Price:      12.99,
			Created:    date1,
			SalesCount: 32,
			ViewsCount: 730,
		},
		{
			ID:         2,
			Name:       "Zebra Table",
			Price:      44.49,
			Created:    date2,
			SalesCount: 301,
			ViewsCount: 3279,
		},
		{
			ID:         3,
			Name:       "Coffee Table",
			Price:      10.00,
			Created:    date3,
			SalesCount: 1048,
			ViewsCount: 20123,
		},
	}

	return products, nil
}

func main() {
	products, err := products()
	if err != nil {
		fmt.Println(err)
		return
	}
	registry := NewProductSortingStrategyRegistry()
	registry.SetStrategy(StrategyPrice, PriceSortStrategy{})
	registry.SetStrategy(StrategySalesPerViewRatio, SalesPerViewRatioSortStrategy{})
	registry.SetStrategy(StrategyDate, DateSortStrategy{})

	strategy, err := registry.GetStrategy(StrategySalesPerViewRatio)
	if err != nil {
		fmt.Println(err)
		return
	}
	sortedProducts := strategy.Sort(products)

	fmt.Println("Sorted by Sales/View Ratio:")

	for _, product := range sortedProducts {
		fmt.Printf("ID: %d, Name: %s, Price: $%.2f, SalesCount: %d, ViewsCount: %d\n",
			product.ID, product.Name, product.Price, product.SalesCount, product.ViewsCount)
	}

	strategy, err = registry.GetStrategy(StrategyPrice)
	if err != nil {
		fmt.Println(err)
		return
	}
	sortedProducts = strategy.Sort(products)

	fmt.Println("Sorted by Price:")

	for _, product := range sortedProducts {
		fmt.Printf("ID: %d, Name: %s, Price: $%.2f, SalesCount: %d, ViewsCount: %d\n",
			product.ID, product.Name, product.Price, product.SalesCount, product.ViewsCount)
	}

	strategy, err = registry.GetStrategy(StrategyDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	sortedProducts = strategy.Sort(products)

	fmt.Println("Sorted by Date:")

	for _, product := range sortedProducts {
		fmt.Printf("ID: %d, Name: %s, Price: $%.2f, SalesCount: %d, ViewsCount: %d\n",
			product.ID, product.Name, product.Price, product.SalesCount, product.ViewsCount)
	}
}
