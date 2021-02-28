package repository

import (
	"encoding/json"
	"log"
	"restful/domain"
	"restful/infrastructure/kvs"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	"golang.org/x/xerrors"
)

type CustomerRepository struct {
	kvs *kvs.Redis
}

func NewCustomerRepository() (*CustomerRepository, error) {
	db, err := kvs.NewRedis()
	if err != nil {
		return nil, xerrors.Errorf("fail to create customerRepository : %w", err)
	}
	return &CustomerRepository{kvs: db}, nil
}

func (c *CustomerRepository) GetByID(id string) (*domain.Customer, error) {
	str, err := c.kvs.Get(id)

	if err == redis.Nil {
		return nil, xerrors.Errorf("customerID (%s) doesn't exist", id)
	}
	if err != nil {
		return nil, err
	}

	var customer domain.Customer
	if err := json.Unmarshal([]byte(str), &customer); err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *CustomerRepository) Create(customer *domain.Customer) error {
	expire := 5 * time.Hour
	bytes, err := json.Marshal(*customer)
	if err != nil {
		return err
	}
	return c.kvs.Set(customer.ID, bytes, expire)
}

func (c *CustomerRepository) Update(
	id string,
	updateData *domain.Customer,
) (*domain.Customer, error) {
	str, err := c.kvs.Get(id)
	if err != nil {
		return nil, xerrors.Errorf("customer(id: %s) not found : %w", id, err)
	}
	var customer domain.Customer
	if err := json.Unmarshal([]byte(str), &customer); err != nil {
		return nil, xerrors.Errorf("failed to Unmarshal id(%s) : %w", id, err)
	}
	if err := copier.Copy(&customer, updateData); err != nil {
		log.Printf("%+v\n %+v", customer, updateData)
		return nil, xerrors.Errorf("failed to update : %w", err)
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		return nil, err
	}
	expire := 5 * time.Hour

	if err := c.kvs.Set(customer.ID, bytes, expire); err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *CustomerRepository) Delete(id string) error {
	return c.kvs.Del(id)
}
