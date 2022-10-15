package dashboard

import (
	"embed"
	"log"
	"net/http"
)

//go:embed yacd
var yacd embed.FS

func Serve() {
	http.Handle("/yacd", http.FileServer(http.FS(yacd)))
	http.ListenAndServe(":8080", nil)
	log.Println("Listen on *:8080")
}
