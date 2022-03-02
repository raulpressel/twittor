package routes

import (
	"encoding/json"
	"net/http"

	"github.com/raulpressel/twittor/bd"
	"github.com/raulpressel/twittor/models"
)

/* ModificarPefil modifica el perfil de usuario*/

func ModificarPefil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar modif el registro "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se logro modificar el registro del usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
