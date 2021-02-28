package usecase

import (
	"restful/domain"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type CustomerInputBoundary interface {
	Create(CustomerInputData) (*domain.Customer, error)
	Read(id string) (*domain.Customer, error)
	Update(id string, customer *domain.Customer) (*domain.Customer, error)
	Delete(id string) error
}

type CustomerInputData struct {
	FirstName string
	LastName  string
}

type CustomerInteractor struct {
	crepo CustomerRepository
}

func NewCustomerInteractor(crepo CustomerRepository) *CustomerInteractor {
	return &CustomerInteractor{crepo}
}

func (c CustomerInteractor) Create(input CustomerInputData) (*domain.Customer, error) {
	customer := domain.Customer{
		ID: uuid.New().String(),
		Name: domain.Name{
			First: input.FirstName,
			Last:  input.LastName,
		},
	}
	if err := c.crepo.Create(&customer); err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c CustomerInteractor) Read(id string) (*domain.Customer, error) {
	customer, err := c.crepo.GetByID(id)
	if err != nil {
		return nil, xerrors.Errorf("failed to read : %w", err)
	}
	return customer, nil
}

func (c CustomerInteractor) Update(
	id string,
	updateData *domain.Customer,
) (*domain.Customer, error) {
	customer, err := c.crepo.Update(id, updateData)
	if err != nil {
		return nil, xerrors.Errorf("failed to update customer id: %s : %w", err)
	}
	return customer, nil
}

func (c CustomerInteractor) Delete(id string) error {
	return c.crepo.Delete(id)
}
