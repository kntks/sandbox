package repository

import (
	"encoding/json"
	"log"
	"reflect"
	"restful/domain"
	"restful/infrastructure/kvs"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
)

var mock = CreateMock()

func CreateMock() *kvs.Redis {
	s, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}
	kvs, err := kvs.NewRedis(s.Addr())
	if err != nil {
		log.Fatal(err)
	}
	return kvs
}

func TestNewCustomerRepository(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	repo, err := NewCustomerRepository(s.Addr())
	if err != nil {
		t.Fatal(err)
	}
	if repo.kvs == nil {
		t.Fatalf("kvs is nil")
	}
}

func TestGetByID(t *testing.T) {
	testcases := []struct {
		name         string
		customerJSON string
		customer     *domain.Customer
	}{
		{
			name:         "正しくcustomerデータを取得できる",
			customerJSON: `{"id": "test", "name": {"first": "firstname", "last": "lastname"}}`,
			customer: &domain.Customer{
				ID: "test",
				Name: domain.Name{
					First: "firstname",
					Last:  "lastname",
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mock.Set(tc.customer.ID, tc.customerJSON, time.Minute)
			got, err := (&CustomerRepository{kvs: mock}).GetByID(tc.customer.ID)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(*tc.customer, *got) {
				t.Errorf("\nwant: %v\ngot: %v\n", *tc.customer, *got)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	testcases := []struct {
		name     string
		customer *domain.Customer
	}{
		{
			name: "正しくcustomerデータを作成できる",
			customer: &domain.Customer{
				ID: "test",
				Name: domain.Name{
					First: "firstname",
					Last:  "lastname",
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := (&CustomerRepository{kvs: mock}).Create(tc.customer)
			if err != nil {
				t.Fatal(err)
			}
			str, err := mock.Get(tc.customer.ID)
			var got domain.Customer
			if err := json.Unmarshal([]byte(str), &got); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(*tc.customer, got) {
				t.Errorf("\nwant: %v\ngot: %v\n", *tc.customer, got)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	testcases := []struct {
		name         string
		customerJSON string
		updateData   *domain.Customer
		want         *domain.Customer
	}{
		{
			name:         "正しくcustomerデータを更新できる",
			customerJSON: `{"id": "test", "name": {"first": "firstname", "last": "lastname"}}`,
			updateData: &domain.Customer{
				ID: "test",
				Name: domain.Name{
					First: "xxxx",
					Last:  "yyyy",
				},
			},
			want: &domain.Customer{
				ID: "test",
				Name: domain.Name{
					First: "xxxx",
					Last:  "yyyy",
				},
			},
		},
		{
			name:         "customerIDはUpdateされないことチェック",
			customerJSON: `{"id": "test1", "name": {"first": "firstname", "last": "lastname"}}`,
			updateData: &domain.Customer{
				ID: "dummy",
				Name: domain.Name{
					First: "firstname2",
					Last:  "lastname2",
				},
			},
			want: &domain.Customer{
				ID: "test1",
				Name: domain.Name{
					First: "firstname2",
					Last:  "lastname2",
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := mock.Set(tc.want.ID, tc.customerJSON, time.Minute)
			if err != nil {
				t.Fatal(err)
			}
			got, err := (&CustomerRepository{kvs: mock}).Update(tc.want.ID, tc.updateData)
			if !reflect.DeepEqual(*tc.want, *got) {
				t.Errorf("\nwant: %v\ngot: %v\n", *tc.want, *got)
			}
		})
	}
}
