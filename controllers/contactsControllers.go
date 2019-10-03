package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
