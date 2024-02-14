package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchProducts(t *testing.T) {
	// Create an instance of our test object
	mockRepository := new(repository.MockRepositoryProducts)

	// Define a query
	query := internal.ProductQuery{
		Id: 1,
	}

	// Prepare the expected products returned by SearchProducts
	products := map[int]internal.Product{
		1: {
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "Product1",
				Price:       100.0,
				SellerId:    1,
			},
		},
	}

	// Setup expectations
	mockRepository.On("SearchProducts", query).Return(products, nil)

	// Call the code we are testing
	actual, err := mockRepository.SearchProducts(query)

	// Asserts
	assert.NoError(t, err)
	assert.Equal(t, products, actual) // Check that the return value from the mocked method is as expected.

	// Confirm that the method was called with the expected parameters
	mockRepository.AssertCalled(t, "SearchProducts", query)
	mockRepository.AssertExpectations(t)
}
