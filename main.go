package main

import (
	"golang-web/handler"
	"log"
	"net/http"
)

func main()  {
	mux := http.NewServeMux();

	// cara 2
	// aboutHandler := func (w http.ResponseWriter, r *http.Request) {
	// 	aboutPage := `
	// 		<html>
	// 		<head><title>About Page</title></head>
	// 		<body>
	// 			<h1>Welcome to the About Page!</h1>
	// 		</body>
	// 		</html>`

	// 	w.Write([]byte(aboutPage))
	// }

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.MarioHandler)
	// mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// cara 3
	// mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	// 	profilePage := `
    //         <html>
    //         <head><title>Profile Page</title></head>
    //         <body>
    //             <h1>Welcome to the Profile Page!</h1>
    //         </body>
    //         </html>`

	// 	w.Write([]byte(profilePage))
	// })

	log.Println("Starting web on port 8080")
	
	err := http.ListenAndServe(":8080", mux)
	if err!= nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
