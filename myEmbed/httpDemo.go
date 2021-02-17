package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed static
var locall embed.FS

func main()  {
	http.Handle("/",http.FileServer(http.FS(locall)))
	log.Fatal(http.ListenAndServe(":8999",nil))
}
