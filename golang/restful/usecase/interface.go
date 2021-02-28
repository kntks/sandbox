package usecase

import "restful/domain"

type CustomerRepository interface {
	GetByID(id string) (*domain.Customer, error)
	Create(*domain.Customer) error
	Update(id string, customer *domain.Customer) (*domain.Customer, error)
	Delete(id string) error
}
