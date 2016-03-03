package api

import (
	"github.com/artshpakov/drawclash/app"
	"github.com/artshpakov/drawclash/app/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func UsersShowAction(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	user, err := model.FindUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	app.Render(w, http.StatusOK, user)
}
