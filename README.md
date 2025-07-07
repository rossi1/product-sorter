# Product Sorting Service
This project implements a flexible and extensible sorting system for products using the **Strategy Pattern** in Go.

It allows A/B testing of different product sorting strategies — e.g. by price or by sales-per-view ratio — to determine which performs best on different parts of the website.

## Features

- Sort products by **price** or **sales/view ratio**
- Easily extendable by registering new sort strategies
- Fully compliant with the **Open/Closed Principle**
- Strategy registry for dynamic resolution
- Clean, idiomatic Go

## Example Strategies

- `price`: Sorts products by ascending price
- `sales_per_view`: Sorts products by sales_count ÷ views_count in descending order

## Sample Output
Sorted by Sales/View Ratio:
ID: 2, Name: Zebra Table, Ratio: 0.0918, Price: $44.49
ID: 3, Name: Coffee Table, Ratio: 0.0520, Price: $10.00
ID: 1, Name: Alabaster Table, Ratio: 0.0438, Price: $12.99

Sorted by Price:
ID: 3, Name: Coffee Table, Price: $10.00
ID: 1, Name: Alabaster Table, Price: $12.99
ID: 2, Name: Zebra Table, Price: $44.49

## Usage

```bash
go run main.go
```

## Testing

```bash
go test -v
```

## Add New Strategy
1. Implement SortStrategy interface
2. Register it in the registry:

```go
registry.RegisterStrategy("created_at", SortByCreatedAt{})
```

3. Use it:

```go
strategy, err := registry.GetStrategy("created_at")
if err != nil {
    fmt.Println(err)
    return
}
sortedProducts := strategy.Sort(products)
```

## Usage (Makefile)

### Build

```bash
make build
```

### Run

```bash
make run
```

### Test

```bash
make test
```

### Docker Build

```bash
make docker-build
```

### Docker Run

```bash
make docker-run
```

### Clean

```bash
make clean
```