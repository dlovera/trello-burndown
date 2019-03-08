package server

import (
	"net/http"

	"github.com/dlovera/trello-burndown/pkg/trello"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := trello.GetDatabase()
	defer db.Close()
	db.Delete(&trello.Board{
		ID: vars["board"],
	})
	http.Redirect(w, r, viper.GetString("http.baseURL")+"index", 302)
}
