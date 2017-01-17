package main

import (
	"net/http"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"log"
	"encoding/json"
)

var (
	httpPort int
	portStr bytes.Buffer
)

func init() {
	flag.IntVar(&httpPort, "http", -1, "HTTP Listen Address")
}

func main() {

	flag.Parse()

	//Default Port as -1 to define custom port - Couldn't find a better way to define a custom port
	if httpPort == -1 {
		fmt.Fprintf(os.Stderr, "Usage --http [:PORT](string) %v \n", (flag.Lookup("http").Name))
            	os.Exit(0)
	}

	//Convert Integer to String
	listenPort := strconv.Itoa(httpPort)

	//Combing strings
	portStr.WriteString(":")
	portStr.WriteString(listenPort)

	http.HandleFunc("/hello", HelloHandler)

	err := http.ListenAndServe(portStr.String(), nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}

type Response struct {
	Result string `json:"result"`
	Code int `json:"code"`
}

func (res *Response) createResponse(code int, result string) Response {

	return Response{
		Result: result,
		Code: code,
	}
}

func (res *Response) sendResponse(w http.ResponseWriter, response Response){
	resp, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(resp)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	res := Response{}

	if len(name) != 0 {
		res = res.createResponse(200, fmt.Sprintf("%s %s", "Hello", name))
	} else {
		res = res.createResponse(401, "")
	}

	res.sendResponse(w, res)
}

