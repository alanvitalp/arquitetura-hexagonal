package application_test

import (
	"testing"

	"github.com/alanvitalp/go-hexagonal/application"
	mock_application "github.com/alanvitalp/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	productPersistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	productPersistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	productService := application.ProductService{
		ProductPersistence: productPersistence,
	}

	result, err := productService.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}
func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	productPersistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	productPersistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	productService := application.ProductService{
		ProductPersistence: productPersistence,
	}

	result, err := productService.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil).AnyTimes()
	product.EXPECT().Disable().Return(nil).AnyTimes()
	productPersistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	productPersistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	productService := application.ProductService{
		ProductPersistence: productPersistence,
	}

	result, err := productService.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = productService.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
