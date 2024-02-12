package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"asciiart/asciiart"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		RenderTemplate(w, "main")
		// Write custom response for status code 200
		w.Write([]byte("Status code 200: Ok"))
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		Result(w, r)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func Result(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "main")
	switch r.Method {
	case "POST":
		// Call your function and get the output
		str := r.FormValue("str")
		radio := r.FormValue("radio")
		html := `<div id="containerresult">`
		output := [][8]string{}
		str = strings.ReplaceAll(str, "\r\n", "\\n")
		data := strings.Split(str, "\\n")
		for _, v := range data {
			if len(v) > 0 {
				if !isValidInput(v) { // Check if the input contains unsupported characters
					http.Error(w, "error 400: character not supported", http.StatusBadRequest)
					return
				}
				output = append(output, asciiart.Asciiart(v, "./asciiart/"+radio+".txt"))
			} else {
				output = append(output, [8]string{})
			}
		}
		// Generate the HTML content dynamically
		for _, v := range output {
			if len(v[0]) > 0 {
				for _, subV := range v {
					html += `<div class="result" style="font-family: monospace">` + subV + `</div>`
				}
			} else {
				html += `<div>&nbsp</div>`
			}
		}
		html += `</div>Status code 200: Ok`
		w.Header().Set("Content-Type", "text/html") // Set the appropriate headers
		w.Write([]byte(html))                       // Send the HTML content as the response
	}
}

func isValidInput(str string) bool {
	for _, char := range str {
		if char < 32 || char > 127 {
			return false
		}
	}
	return true
}

// appeler ascii art et afficher le résultat sur la page html
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".tmpl")
	if err != nil {
		http.Error(w, "error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/result", Result)
	fs := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css", fs))
	http.ListenAndServe(":8800", nil)
}
