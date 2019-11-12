package example

import (
	"fmt"
	"net/http"
	"strings"
)

func ServerFile() {
	fmt.Println(strings.TrimRight("/", "/"))
	http.Handle("/", http.FileServer(http.Dir("/data/logs")))
	http.ListenAndServe(":4000", nil)
}
func he(resp http.ResponseWriter, req *http.Request) {
	fmt.Println(resp.Header())
	fmt.Println(req.Form)
	fmt.Println(req.RequestURI)
	fmt.Println(req.RemoteAddr)
	fmt.Println(req.Method)
	fmt.Println(req.URL)
	fmt.Println(req.RequestURI)
	fmt.Println(req.Header)
	resp.Write([]byte("<h1>haha</h1>"))

}
func gb(resp http.ResponseWriter, req *http.Request) {
	fmt.Println(resp.Header())

}
func HandlerHttpTest() {
	http.HandleFunc("/hello", he)
	http.HandleFunc("/gb", gb)
	http.ListenAndServe(":4000", nil)
}
