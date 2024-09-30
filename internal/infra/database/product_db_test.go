package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/edsonjuniordev/full-cycle-api/internal/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.Product{})

	product, err := entities.NewProduct("Product 1", 10.00)
	assert.Nil(t, err)

	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.Product{})

	for i := 1; i < 24; i++ {
		product, err := entities.NewProduct(fmt.Sprintf("Product %d", i), rand.Int())
		assert.Nil(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.Product{})

	product, err := entities.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	db.Create(product)

	productDB := NewProduct(db)

	product, err = productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.Name, "Product 1")
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.Product{})

	product, err := entities.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	db.Create(product)

	productDB := NewProduct(db)

	product.Name = "Product 2"

	err = productDB.Update(product)
	assert.Nil(t, err)

	product, err = productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, "Product 2", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.Product{})

	product, err := entities.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	db.Create(product)

	productDB := NewProduct(db)

	err = productDB.Delete(product)
	assert.Nil(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}
