package handler

import (
	"golang-web/entity"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

//  cara 1
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err!= nil {
		log.Println(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	// data := entity.Product{Id: 1, Name: "Arda", Price: 2000000, Stock: 5}
	data := []entity.Product{
		{Id: 1, Name: "Arda", Price: 2000000, Stock: 2},
		{Id: 2, Name: "Bahrul", Price: 3000000, Stock: 11},
		{Id: 3, Name: "Ulum", Price: 4000000, Stock: 15},
	}

	// data := map[string]interface{}{
	// 	"title":   "Home Page",
	// 	"content": "This is a message from the data.",
	// }
	
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
	}

		// homePage := `
	// 	<html>
	// 	<head><title>Home Page</title></head>
	// 	<body>
	// 		<h1>Welcome to the Home Page!</h1>
	// 	</body>
	// 	</html>`
	
	// w.Write([]byte(homePage))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	helloPage := `
		<html>
		<head><title>Hello Page</title></head>
		<body>
			<h1>Welcome to the Hello Page!</h1>
		</body>
		</html>`

	w.Write([]byte(helloPage))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	marioPage := `
		<html>
		<head><title>Mario Page</title></head>
		<body>
			<h1>Welcome to the Mario Page!</h1>
		</body>
		</html>`
	w.Write([]byte(marioPage))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id) // Convert string id to int64
	
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err!= nil {
		log.Println(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	data := map[string]interface{}{
		"title":   "Product Page",
		"content": idNumb,
	}
	
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// fmt.Fprintf(w, "Product page : %d\n", idNumb)

	// productPage  := `
	// 	<html>
	// 	<head><title>Product Page</title></head>
	// 	<body>
	// 		<h1>Welcome to the Product Page!</h1>
	// 	</body>
	// 	</html>`
	// w.Write([]byte(productPage ))
}
