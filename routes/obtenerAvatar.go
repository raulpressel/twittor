package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/raulpressel/twittor/bd"
)

/*ObtenerAvatar envia el avatar al http*/

func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "usuario no encontrado", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars" + perfil.Avatar)
	if err != nil {
		http.Error(w, "imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)

	if err != nil {
		http.Error(w, "error al copiar la imagen", http.StatusBadRequest)
	}

}
