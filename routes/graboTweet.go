package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/raulpressel/twittor/bd"
	"github.com/raulpressel/twittor/models"
)

/* GraboTweet premite grabar el twwet en la base de datos*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "mensaje error"+err.Error(), 400)
		return
	}
	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "error al insertar el tweet "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "no se inserto el tweet ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
