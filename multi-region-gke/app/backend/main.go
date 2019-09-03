package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "time"
    "encoding/json"
    "os"
    "os/signal"
    "context"
)

type Category struct {
    CategoryId string `json:"CategoryId"`
    CategoryName string `json:"CategoryName"`
    CategoryDescription string `json:"CategoryDescription"`
}

type Categories struct {
    Categories []Category
}

type Product struct {
    CategoryId string `json:"CategoryId"`
    CategoryName string `json:"CategoryName"`
    CategoryDescription string `json:"CategoryDescription"`
    ProductId string `json:"ProductId"`
    ProductName string `json:"ProductName"`
    ProductDescription string `json:"ProductDescription"`
}

type Products struct {
    Products []Product
}

type Option struct {
    CategoryId string `json:"CategoryId"`
    CategoryName string `json:"CategoryName"`
    CategoryDescription string `json:"CategoryDescription"`
    ProductId string `json:"ProductId"`
    ProductName string `json:"ProductName"`
    ProductDescription string `json:"ProductDescription"`
    OptionId string `json:"OptionId"`
    OptionName string `json:"OptionName"`
    OptionDescription string `json:"OptionDescription"`
    OptionValue string `json:"OptionValue"`
}

type Options struct {
    Options []Option
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.RequestURI)
        next.ServeHTTP(w, r)
    })
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
    cat1 := Category{
        CategoryId: "42614149-427b-4857-8957-e8f7ca0b3b8b",
        CategoryName: "Nisi porta lorem",
        CategoryDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Fermentum dui faucibus in ornare. Commodo odio aenean sed adipiscing diam donec adipiscing tristique.",
    }
    cat2 := Category{
        CategoryId: "28dde213-cc32-4fe7-80c5-763c49decbb4",
        CategoryName: "Ipsum consequat nisl vel pretium",
        CategoryDescription: "Nibh mauris cursus mattis molestie a iaculis at. Morbi tincidunt augue interdum velit euismod in pellentesque massa placerat.",
    }
    cats := Categories{[]Category{cat1,cat2}}
    json.NewEncoder(w).Encode(cats)
}

func GetOneCategory(w http.ResponseWriter, r *http.Request) {
    cat1 := Category{
        CategoryId: "42614149-427b-4857-8957-e8f7ca0b3b8b",
        CategoryName: "Nisi porta lorem",
        CategoryDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Fermentum dui faucibus in ornare. Commodo odio aenean sed adipiscing diam donec adipiscing tristique.",
    }
    json.NewEncoder(w).Encode(cat1)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
    prod1 := Product{
        CategoryId: "42614149-427b-4857-8957-e8f7ca0b3b8b",
        CategoryName: "Nisi porta lorem",
        CategoryDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Fermentum dui faucibus in ornare. Commodo odio aenean sed adipiscing diam donec adipiscing tristique.",
        ProductId: "afc94081-5d5d-4bfe-90fb-365f685751f3",
        ProductName: "Phasellus vestibulum lorem",
        ProductDescription: "Nunc congue nisi vitae suscipit tellus. Volutpat diam ut venenatis tellus in metus vulputate eu scelerisque.",
    }
    prod2 := Product{
        CategoryId: "28dde213-cc32-4fe7-80c5-763c49decbb4",
        CategoryName: "Ipsum consequat nisl vel pretium",
        CategoryDescription: "Nibh mauris cursus mattis molestie a iaculis at. Morbi tincidunt augue interdum velit euismod in pellentesque massa placerat.",
        ProductId: "3c1ccd0b-00ee-41b2-a8fa-396a3f47300e",
        ProductName: "Egestas quis ipsum",
        ProductDescription: "Nulla posuere sollicitudin aliquam ultrices sagittis orci a scelerisque.",
    }
    prods := Products{[]Product{prod1,prod2}}
    json.NewEncoder(w).Encode(prods)
}

func GetOneProduct(w http.ResponseWriter, r *http.Request) {
    prod1 := Product{
        CategoryId: "42614149-427b-4857-8957-e8f7ca0b3b8b",
        CategoryName: "Nisi porta lorem",
        CategoryDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Fermentum dui faucibus in ornare. Commodo odio aenean sed adipiscing diam donec adipiscing tristique.",
        ProductId: "3c1ccd0b-00ee-41b2-a8fa-396a3f47300e",
        ProductName: "Erat velit scelerisque",
        ProductDescription: "Lorem dolor sed viverra ipsum nunc aliquet bibendum enim. Dui sapien eget mi proin sed libero enim.",
    }
    json.NewEncoder(w).Encode(prod1)
}

func GetOneProductOptions(w http.ResponseWriter, r *http.Request) {
    opt1 := Option{
        CategoryId: "42614149-427b-4857-8957-e8f7ca0b3b8b",
        CategoryName: "Nisi porta lorem",
        CategoryDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Fermentum dui faucibus in ornare. Commodo odio aenean sed adipiscing diam donec adipiscing tristique.",
        ProductId: "3c1ccd0b-00ee-41b2-a8fa-396a3f47300e",
        ProductName: "Erat velit scelerisque",
        ProductDescription: "Lorem dolor sed viverra ipsum nunc aliquet bibendum enim. Dui sapien eget mi proin sed libero enim.",
        OptionId: "5a3633a6-9fba-446e-8814-2189456d75f8",
        OptionName: "Morbi tristique senectus",
        OptionDescription: "Commodo viverra maecenas accumsan lacus vel facilisis. ",
        OptionValue: "Viverra orci sagittis eu volutpat odio facilisis mauris.",
    }
    opt2 := Option{
        CategoryId: "42614149-427b-4857-8957-e8f7ca0b3b8b",
        CategoryName: "Nisi porta lorem",
        CategoryDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Fermentum dui faucibus in ornare. Commodo odio aenean sed adipiscing diam donec adipiscing tristique.",
        ProductId: "afc94081-5d5d-4bfe-90fb-365f685751f3",
        ProductName: "Phasellus vestibulum lorem",
        ProductDescription: "Nunc congue nisi vitae suscipit tellus. Volutpat diam ut venenatis tellus in metus vulputate eu scelerisque.",
        OptionId: "030a9531-c2ba-4cf8-aca0-e2cffe374f5e",
        OptionName: "Arcu felis bibendum",
        OptionDescription: "Non pulvinar neque laoreet suspendisse interdum consectetur libero id faucibus.",
        OptionValue: "Nisl nunc mi ipsum faucibus vitae aliquet. Lorem mollis aliquam ut porttitor leo.",
    }
    opts := Options{[]Option{opt1,opt2}}
    json.NewEncoder(w).Encode(opts)
}

func main() {
    r := mux.NewRouter()
    r.Use(LoggingMiddleware)
    r.HandleFunc("/categories", GetAllCategories)
    r.HandleFunc("/categories/{category}", GetOneCategory)
    r.HandleFunc("/categories/{category}/products", GetAllProducts)
    r.HandleFunc("/categories/{category}/products/{product}", GetOneProduct)
    r.HandleFunc("/categories/{category}/products/{product}/options", GetOneProductOptions)
    srv := &http.Server{
        Handler: r,
        Addr: "0.0.0.0:8000",
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
    ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
    defer cancel()
    srv.Shutdown(ctx)
    log.Println("Server shutting down")
    os.Exit(0)
}
