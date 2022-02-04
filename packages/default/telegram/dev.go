// //go:build dev
// // +build dev

package main

// import (
// 	"fmt"
// 	"net/http"
// )

// // This main function is for testing locally
// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		qp := r.URL.Query()

// 		params := make(map[string]interface{}, len(qp))
// 		for k, v := range qp {
// 			params[k] = v
// 		}
// 		w.Header().Set("Content-Type", "text/html")
// 		payload := Main(params)
// 		if b, ok := payload["body"]; ok {
// 			fmt.Fprint(w, b)
// 		}
// 		if b, ok := payload["error"]; ok {
// 			fmt.Fprint(w, b)
// 		}
// 	})

// 	// push up to web?
// 	// fs := http.FileServer(http.Dir("static/"))
// 	// http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	fmt.Println("Running on port :8080")
// 	http.ListenAndServe(":8080", nil)
// }
