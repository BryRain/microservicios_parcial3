package controller

import (
	"context"
	"delete-service/service"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LibroController struct {
	Service *service.LibroService
}

func (c *LibroController) EliminarLibro(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	if err := c.Service.EliminarLibro(context.Background(), oid); err != nil {
		http.Error(w, "Error al eliminar libro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Libro eliminado correctamente"))
}
