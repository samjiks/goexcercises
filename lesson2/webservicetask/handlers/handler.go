package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/samjiks/lesson2/webservicetask/responses"
	"github.com/samjiks/lesson2/webservicetask/developers"

)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Build the response struct
	name := r.FormValue("name")
	code := 200
	res := ""
	if len(name) == 0 {
		code = 422
	} else {
		res = fmt.Sprintf("%s %s", "hello", name)
	}

	responses.SendReponse(res, w, code)
}


func DevelopersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
		case "GET": {
			dev := developers.Developer{}

			vars := mux.Vars(r)
			id, _ := vars["id"]

			if id != "" {
				ids, _ := strconv.Atoi(id)
				buf, err := json.Marshal(dev.GetById(ids))
				if err != nil {
					fmt.Println(err)
					return
				}
			responses.SendReponse(string(buf), w, 200)
			} else {
				buf, err := json.Marshal(dev.Get())
				if err != nil {
					fmt.Println(err)
					return
				}
			responses.SendReponse(string(buf), w, 200)
			}


		}
		case "POST": {
			i, _ := strconv.Atoi(r.FormValue("age"))
			flo, _ := strconv.Atoi(r.FormValue("floor"))

			d := developers.Developer{
				Name : r.FormValue("name"),
				Age: i,
				Language: r.FormValue("language"),
				Floor: flo,
			}

			buffer := d.Create()
			buf, err := json.Marshal(buffer)
			if err != nil {
				fmt.Println(err)
				return
			}
			responses.SendReponse(string(buf), w, 200)
		}
	}
}