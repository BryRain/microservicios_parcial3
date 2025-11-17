package tests

import (
	"context"
	"errors"
	"read-service/model"
	"read-service/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockLibroRepository is a mock implementation of the LibroRepositorio interface
type MockLibroRepository struct {
	mock.Mock
}

// ObtenerTodos is the mock implementation of the ObtenerTodos method
func (m *MockLibroRepository) ObtenerTodos(ctx context.Context) ([]model.Libro, error) {
	args := m.Called(ctx)
	// Handle the case where the first return value is nil but needs to be casted
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Libro), args.Error(1)
}

func TestLibroService_ObtenerLibros_Exitoso(t *testing.T) {
	// 1. Setup
	mockRepo := new(MockLibroRepository)
	libroService := service.LibroService{Repo: mockRepo}

	// Expected result
	expectedLibros := []model.Libro{
		{ID: "1", Titulo: "El Quijote", Autor: "Cervantes", Anio: 1605, Editorial: "Santillana"},
		{ID: "2", Titulo: "Cien Años de Soledad", Autor: "García Márquez", Anio: 1967, Editorial: "Sudamericana"},
	}

	// Configure the mock
	mockRepo.On("ObtenerTodos", context.Background()).Return(expectedLibros, nil)

	// 2. Execution
	libros, err := libroService.ObtenerLibros(context.Background())

	// 3. Assertions
	assert.NoError(t, err, "No se esperaba un error")
	assert.NotNil(t, libros, "Se esperaba una lista de libros, no nula")
	assert.Equal(t, len(expectedLibros), len(libros), "La cantidad de libros no coincide")
	assert.Equal(t, expectedLibros[0].Titulo, libros[0].Titulo, "El título del primer libro no coincide")

	// Verify that the mock was called as expected
	mockRepo.AssertExpectations(t)
}

func TestLibroService_ObtenerLibros_Error(t *testing.T) {
	// 1. Setup
	mockRepo := new(MockLibroRepository)
	libroService := service.LibroService{Repo: mockRepo}

	// Expected error
	expectedError := errors.New("error en la base de datos")

	// Configure the mock to return an error
	mockRepo.On("ObtenerTodos", context.Background()).Return(nil, expectedError)

	// 2. Execution
	libros, err := libroService.ObtenerLibros(context.Background())

	// 3. Assertions
	assert.Error(t, err, "Se esperaba un error")
	assert.Nil(t, libros, "No se esperaba una lista de libros")
	assert.Equal(t, expectedError, err, "El error devuelto no es el esperado")

	// Verify that the mock was called as expected
	mockRepo.AssertExpectations(t)
}
