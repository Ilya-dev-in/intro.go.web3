package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	titleJson, err := json.Marshal(struct{ Name string }{Name: "Stay connected"})
	if err != nil {
		fmt.Println(err)
	}

	w.Write(titleJson)
}
