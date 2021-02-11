package service

import (
	"github.com/abhishek70/golang-rest-api/model"
	"github.com/abhishek70/golang-rest-api/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)


type MockRepository struct {
	mock.Mock
}


func (mock *MockRepository) FindAll() ([]model.Product, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Product), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepository := new(MockRepository)

	product := model.Product{Id:1, Sku:"ABC", Name:"Product Title"}

	mockRepository.On("FindAll").Return([]model.Product{product}, nil)

	testProductService := service.NewProductService(mockRepository)

	result, _ := testProductService.FindAll()

	mockRepository.AssertExpectations(t)

	assert.Equal(t, 1, result[0].Id)
	assert.Equal(t, "ABC", result[0].Sku)
	assert.Equal(t, "Product Title", result[0].Name)

}