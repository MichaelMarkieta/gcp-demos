package main

import (
    "log"
    "net/http"
    "html/template"
    "github.com/gorilla/mux"
    "time"
)

var templates = template.Must(template.ParseGlob("templates/*"))

type PageData struct {
    PageTitle string
    ApiResponse string
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.RequestURI)
        next.ServeHTTP(w, r)
    })
}

func Index(w http.ResponseWriter, r *http.Request) {
    data := PageData{
        PageTitle: "Home | GCP Demos: multi-region-gke",
    }
    err := templates.ExecuteTemplate(w, "index.html", data)
    if err != nil {
        log.Fatal("Issue rendering page. ", err)
    }
}

func ApiResponse(w http.ResponseWriter, r *http.Request) {
    data := PageData{
        PageTitle: "API Response | GCP Demos: multi-region-gke",
        ApiResponse: "{hello:world}",
    }
    err := templates.ExecuteTemplate(w, "apiresponse.html", data)
    if err != nil {
        log.Fatal("Issue rendering page. ", err)
    }
}

func main() {
    r := mux.NewRouter()
    fs := http.FileServer(http.Dir("assets/"))
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
    r.Use(LoggingMiddleware)
    r.HandleFunc("/", Index)
    r.HandleFunc("/categories", ApiResponse)
    r.HandleFunc("/categories/{category}", ApiResponse)
    r.HandleFunc("/categories/{category}/products", ApiResponse)
    r.HandleFunc("/categories/{category}/products/{product}", ApiResponse)
    srv := &http.Server{
        Handler: r,
        Addr: "0.0.0.0:8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second,
        IdleTimeout: 60 * time.Second,
    }
    go func() {
        if err := srv.ListenAndServe(); err != nil {
            log.Println(err)
        }
    }()
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()
    srv.Shutdown(ctx)
    log.Println("Server shutting down")
    os.Exit(0)
}