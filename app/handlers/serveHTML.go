package handlers

import (
	"io/ioutil"
	"net/http"
	"os"
)

func ServeHTML(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("app/templates/index.html")
	defer f.Close()
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	o, err := ioutil.ReadAll(f)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(o)
}
