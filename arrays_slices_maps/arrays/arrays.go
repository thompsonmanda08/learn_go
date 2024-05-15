package arrays

import "errors"

type Product struct {
	id    string
	title string
	price float64
}

func New(title string, price float64, id string) (Product, error) {
	if title == "" || price == 0 || id == "" {
		return Product{}, errors.New("Product must have a title, Price and ID.")

	}

	return Product{
		id:    id,
		title: title,
		price: price,
	}, nil
}
