package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	titleJson, err := json.Marshal(struct{ Name string }{Name: "To my go + angular site"})
	if err != nil {
		fmt.Println(err)
	}

	w.Write(titleJson)
}
