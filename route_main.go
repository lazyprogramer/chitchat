package main

import (
	"net/http"
	"chitchat/data"
)

func err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "index")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "index")
	}

}
func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	} else {
		err_message(w, r, "Cannot get threads")
	}
}


