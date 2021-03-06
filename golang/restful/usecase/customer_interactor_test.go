package usecase

import (
	"reflect"
	"restful/domain"
	"testing"
)

type testRepo struct{}

func (t testRepo) Create(customer *domain.Customer) error {
	return nil
}

func (t testRepo) GetByID(id string) (*domain.Customer, error) {
	switch id {
	case "test":
		return &domain.Customer{
			ID:   "test",
			Name: domain.Name{First: "test1", Last: "test2"},
		}, nil
	}
	return nil, nil
}

func (t testRepo) Update(id string, customer *domain.Customer) (*domain.Customer, error) {
	switch id {
	case "test":
		customer.ID = "test"
		return customer, nil
	}
	return nil, nil
}

func (t testRepo) Delete(id string) error {
	return nil
}

func TestCreate(t *testing.T) {
	testcases := []struct {
		name  string
		input CustomerInputData
		want  domain.Name
	}{
		{
			name:  "InputDataで設定したデータが正しくCustomer構造体にsetされている",
			input: CustomerInputData{FirstName: "firstname", LastName: "lastname"},
			want:  domain.Name{First: "firstname", Last: "lastname"},
		},
	}

	repo := NewCustomerInteractor(testRepo{})
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := repo.Create(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tc.want, got.Name) {
				t.Errorf("\nwant: %v\ngot: %v\n", tc.want, got.Name)
			}
		})
	}
}
