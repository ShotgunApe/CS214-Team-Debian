package main

import (
	"log"
	"net/http"
	"os"
	"student-tracker/student"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Testing string `yaml:"testing"`
}

func main() {
	var settings Settings
	settings.getSettings()

	student.Init()
	student.GetStudents()

	//Handler points to available directories
	http.Handle("/web/html", http.StripPrefix("/web/html", http.FileServer(http.Dir("web/html"))))
	http.Handle("/web/scripts/", http.StripPrefix("/web/scripts/", http.FileServer(http.Dir("web/scripts"))))
	http.Handle("/web/css/", http.StripPrefix("/web/css/", http.FileServer(http.Dir("web/css"))))
	http.Handle("/web/assets/", http.StripPrefix("/web/assets/", http.FileServer(http.Dir("web/assets"))))
	http.Handle("/web/data/", http.StripPrefix("/web/data/", http.FileServer(http.Dir("web/data"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path[1:] == "" {
			http.ServeFile(w, r, "web/html/index.html")
		} else if len(r.URL.Path) == 10 || (len(r.URL.Path) == 11 && r.URL.Path[10:11] == "/") {
			if r.URL.Path[1:10] == "add-entry" {
				http.ServeFile(w, r, "web/html/add-entry/index.html")
			}
		} else if len(r.URL.Path) == 14 || (len(r.URL.Path) == 15 && r.URL.Path[14:15] == "/") {
			if r.URL.Path[1:14] == "view-database" {
				http.ServeFile(w, r, "web/html/view-database/index.html")
			}
		} else {
			http.ServeFile(w, r, "web/html/404/index.html")
		}
	})

	//Serves local webpage for testing
	if settings.Testing == "true" {
		errhttp := http.ListenAndServe("localhost:3000", nil)
		if errhttp != nil {
			log.Fatal("Web server (HTTP): ", errhttp)
		}
	} else {
		//Serves the webpage to the internet
		errhttps := http.ListenAndServeTLS(":443", "certs/cert.pem", "certs/key.pem", nil)
		if errhttps != nil {
			log.Fatal("Web server (HTTPS): ", errhttps)
		}
	}

}

func (c *Settings) getSettings() *Settings {

	yamlFile, err := os.ReadFile("settings.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
