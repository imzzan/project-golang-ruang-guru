package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Pay(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	product, err := s.database.GetProductByName(productName)

	if err != nil {
		return err
	}

	products, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	if quantity <= 0 {
		return errors.New("invalid quantity")
	}

	products = append(products, entity.CartItem{
		ProductName: productName,
		Price:       product.Price,
		Quantity:    quantity,
	})

	err = s.database.SaveCartItems(products)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveCart(productName string) error {
	_, err := s.database.GetProductByName(productName)

	if err != nil {
		return err
	}

	products, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	isProductAvailable := false
	newProducts := []entity.CartItem{}
	for i, item := range products {
		if item.ProductName != productName {
			newProducts = append(newProducts, products[i])
			isProductAvailable = true
		}
	}

	if !isProductAvailable {
		return errors.New("product not found")
	}

	err = s.database.SaveCartItems(newProducts)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.GetCartItems()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s *Service) ResetCart() error {
	itemEmpty := []entity.CartItem{}
	err := s.database.SaveCartItems(itemEmpty)
	return err
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	product := s.database.GetProductData()
	return product, nil
}

func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	products, err := s.ShowCart()
	if err != nil {
		return entity.PaymentInformation{}, err
	}

	totalPrice := 0
	for _, item := range products {
		totalPrice += item.Price * item.Quantity
	}

	kembalian := money - totalPrice

	if kembalian < 0 {
		return entity.PaymentInformation{}, errors.New("money is not enough")
	}

	s.ResetCart()
	return entity.PaymentInformation{
		TotalPrice:  totalPrice,
		Change:      kembalian,
		ProductList: products,
		MoneyPaid:   money,
	}, nil
}
