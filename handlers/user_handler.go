package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"med-api/models"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	user := models.User{}
	db.DB.First(&user, id)
	json.NewEncoder(w).Encode(user)
}
