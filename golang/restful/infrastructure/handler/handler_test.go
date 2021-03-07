package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"restful/domain"
	"restful/usecase"
	"strconv"
	"testing"
)

type mockUseCase struct {
	inmemory map[string]domain.Customer
}

func (m mockUseCase) Create(cid usecase.CustomerInputData) (*domain.Customer, error) {
	id := strconv.Itoa(len(m.inmemory) + 1)
	customer := domain.Customer{
		ID:   id,
		Name: domain.Name{First: cid.FirstName, Last: cid.LastName},
	}
	m.inmemory[id] = customer
	return &customer, nil
}

func (m mockUseCase) Read(id string) (*domain.Customer, error) {
	return nil, nil
}

func (m mockUseCase) Update(id string, customer *domain.Customer) (*domain.Customer, error) {
	return nil, nil
}

func (m mockUseCase) Delete(id string) error {
	return nil
}

func TestCrate(t *testing.T) {
	testcases := []struct {
		name        string
		input       usecase.CustomerInputData
		want        domain.Customer
		statusCode  int
		expectError bool
	}{
		{
			name: "適切なデータがPOSTされたときのテスト",
			input: usecase.CustomerInputData{
				FirstName: "firstName",
				LastName:  "lastName",
			},
			want: domain.Customer{
				ID:   "1",
				Name: domain.Name{First: "firstName", Last: "lastName"},
			},
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name: "バリデーションが適切に動いている:min=1",
			input: usecase.CustomerInputData{
				FirstName: "",
				LastName:  "",
			},
			want:        domain.Customer{},
			statusCode:  http.StatusBadRequest,
			expectError: true,
		},
		{
			name: "バリデーションが適切に動いている:max=15",
			input: usecase.CustomerInputData{
				FirstName: "firstnamefirstname",
				LastName:  "lastnamelastname",
			},
			want:        domain.Customer{},
			statusCode:  http.StatusBadRequest,
			expectError: true,
		},
	}

	handlers := NewHandlers(mockUseCase{inmemory: make(map[string]domain.Customer)})

	for _, tc := range testcases {
		b, _ := json.Marshal(tc.input)
		req := httptest.NewRequest(http.MethodPost, "http://example.com/create", bytes.NewBuffer(b))
		w := httptest.NewRecorder()
		handlers.Create(w, req)

		if tc.expectError {
			t.Run(tc.name, func(t *testing.T) {
				if w.Result().StatusCode == tc.statusCode {
					return
				}
				t.Errorf("\nwant: %d\ngot: %d", tc.statusCode, w.Result().StatusCode)
			})
			continue
		}

		t.Run(tc.name, func(t *testing.T) {
			var got domain.Customer
			if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("\nwant: %+v\ngot: %+v\n", tc.want, got)
			}
		})

	}
}
