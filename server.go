package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fileServer := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fileServer) //Sends HTML response to the main page (GET)
	StyleServer := http.FileServer(http.Dir("./styles"))
	http.Handle("/styles/", http.StripPrefix("/styles", StyleServer))
	http.HandleFunc("/ascii-art", formHandler) // goes to templates to receive and display data from the server (index.html) (GET Tip)

	
	fmt.Printf("Starting server at port 8080 ... \n")

	if err := http.ListenAndServe(":8080", nil); err != nil { // Set up server, listen for localhost:8080, display error if unable to
		log.Fatal("HTTP status 500 - Internal server error: %s", err)
	}

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// http status 200 - OK
	w.WriteHeader(http.StatusOK)

	if r.URL.Path != "/ascii-art" { // error 404 handling
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	Font := r.FormValue("Banner")
	f, _ := os.Open(Font + ".txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	PrintStr := r.FormValue("Text")
	var PrintMulti []string

	if strings.Contains(PrintStr, "\\n") {
		PrintMulti = strings.Split(PrintStr, "\\n")
	} else {
		PrintMulti = append(PrintMulti, PrintStr)
	}

	if strings.Contains(PrintStr, "ä") || strings.Contains(PrintStr, "ü") || strings.Contains(PrintStr, "õ") || strings.Contains(PrintStr, "ö") {
		fmt.Println()
		fmt.Println("http status 400, Bad request")
		fmt.Println()
	}

	charS := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	var Fontstyle []string
	for _, x := range PrintMulti {
		PrintStr = x
		for i := 0; i < len(charS)*9; i++ {
			scanner.Scan()
			Fontstyle = append(Fontstyle, scanner.Text())

		}
		var Onscreen [8]string
		for j := 0; j < len(PrintStr); j++ {
			for i := 0; i < 8; i++ {
				Onscreen[i] += Fontstyle[strings.Index(charS, string(PrintStr[j]))*9+i+1]
			}
		}
		for _, y := range Onscreen {
			fmt.Fprintf(w, y+"\n")
		}
		Onscreen = [8]string{}
	}

}
