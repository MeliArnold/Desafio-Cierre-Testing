package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"bytes"
	"errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductsDefault_Get(t *testing.T) {
	assert := assert.New(t)

	t.Run("success", func(t *testing.T) {
		query := internal.ProductQuery{Id: 1}
		products := map[int]internal.Product{
			1: {
				// el contenido del product
			},
		}

		mockRepo := new(repository.MockRepositoryProducts)
		mockRepo.On("SearchProducts", query).Return(products, nil)

		h := handler.NewProductsDefault(mockRepo)

		req, _ := http.NewRequest("GET", "?id="+strconv.Itoa(query.Id), bytes.NewBuffer(nil))
		resp := httptest.NewRecorder()

		h.Get()(resp, req)

		mockRepo.AssertExpectations(t)
		assert.Equal(http.StatusOK, resp.Code)
		// Puedes agregar más aserciones para verificar la respuesta aquí
	})

	t.Run("internal_error", func(t *testing.T) {
		mockRepo := new(repository.MockRepositoryProducts)
		mockRepo.On("SearchProducts", mock.Anything).Return(nil, errors.New("some error"))
		h := handler.NewProductsDefault(mockRepo)

		req, _ := http.NewRequest("GET", "?id=1", bytes.NewBuffer(nil))
		resp := httptest.NewRecorder()

		h.Get()(resp, req)

		mockRepo.AssertExpectations(t)
		assert.Equal(http.StatusInternalServerError, resp.Code)
	})

	// Y así para los demás casos de prueba
}
