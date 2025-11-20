package application

import (
"github.com/dong-tran/docs/ddd-example/domain/model"
"github.com/dong-tran/docs/ddd-example/domain/repository"
"github.com/dong-tran/docs/ddd-example/domain/service"
)

type ProductService struct {
	repo           repository.ProductRepository
	pricingService *service.PricingService
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		repo:           repo,
		pricingService: service.NewPricingService(),
	}
}

type CreateProductDTO struct {
	Name        string
	Description string
	Price       float64
	Currency    string
	Category    string
}

func (s *ProductService) CreateProduct(dto CreateProductDTO) (*model.Product, error) {
	price, err := model.NewMoney(dto.Price, dto.Currency)
	if err != nil {
		return nil, err
	}

	category, err := model.NewCategory(dto.Category)
	if err != nil {
		return nil, err
	}

	product, err := model.NewProduct(dto.Name, dto.Description, price, category)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) ApplyDiscountToProduct(productID model.ProductID, discount float64) error {
	product, err := s.repo.FindByID(productID)
	if err != nil {
		return err
	}

	if err := s.pricingService.ApplyDiscount(product, discount); err != nil {
		return err
	}

	return s.repo.Save(product)
}

func (s *ProductService) GetProduct(id model.ProductID) (*model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) GetAllProducts() ([]*model.Product, error) {
	return s.repo.FindAll()
}
