package routes

import (
	"encoding/json"
	"net/http"

	"github.com/raulpressel/twittor/bd"
	"github.com/raulpressel/twittor/models"
)

/*Registro es la funcion para crear en la bd el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) //body es un obj string de solo lectura, una vez q se utiliza body se destruye
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido "+err.Error(), 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Contraseña no valida < 6 "+err.Error(), 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Email ya registrado "+err.Error(), 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "no se ha logrado insertar el registro status = false ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
