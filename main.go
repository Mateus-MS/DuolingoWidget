package main

import (
	"fmt"
	"net/http"

	"github.com/Mateus-MS/DuolingoWidget/routes"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", routes.StreakRoute)

	startServer(router)
}

func startServer(router *http.ServeMux) {
	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("certs"),
		HostPolicy: autocert.HostWhitelist("cubonauta.com"),
	}

	go func() {
		httpServer := &http.Server{
			Addr:    ":80",
			Handler: certManager.HTTPHandler(nil),
		}
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %s", err)
		}
	}()

	httpsServer := &http.Server{
		Addr:      ":443",
		Handler:   router,
		TLSConfig: certManager.TLSConfig(),
	}

	if err := httpsServer.ListenAndServeTLS("", ""); err != nil {
		fmt.Println("HTTPS server error:", err)
	}

	fmt.Println("Starting SERVER")
}
