package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB     *sql.DB
	Port   string
	Router *mux.Router
}

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello World")
// }

func (a *App) Initialize() {
	DB, err := sql.Open("sqlite3", "../../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	a.DB = DB
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	// Product routes
	a.Router.HandleFunc("/products", a.allProducts).Methods("GET")
	a.Router.HandleFunc("/product/{id}", a.fetchProduct).Methods("GET")
	a.Router.HandleFunc("/products", a.newProduct).Methods("POST")

	// Order routes
	a.Router.HandleFunc("/orders", a.allOrders).Methods("GET")
	a.Router.HandleFunc("/order/{id}", a.fetchOrder).Methods("GET")
	a.Router.HandleFunc("/orders", a.newOrder).Methods("POST")
	a.Router.HandleFunc("/orderitems", a.newOrderItems).Methods("POST")
}

func (a *App) allProducts(w http.ResponseWriter, r *http.Request) {
	// get all products in JSON format
	// curl localhost:8080/products
	products, err := getProducts(a.DB)
	if err != nil {
		fmt.Printf("getProducts error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) fetchProduct(w http.ResponseWriter, r *http.Request) {
	// get a single product in JSON format
	// curl localhost:8080/product/9
	vars := mux.Vars(r)
	id := vars["id"]

	var p product
	p.ID, _ = strconv.Atoi(id)
	err := p.getProduct(a.DB)
	if err != nil {
		fmt.Printf("getProduct error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) newProduct(w http.ResponseWriter, r *http.Request) {
	// create a new product (POST)
	// curl -X POST -H "Content-Type: application/json" -d '{"name":"New Product","price":99.99}' localhost:8080/products
	reqBody, _ := ioutil.ReadAll(r.Body)
	var p product
	json.Unmarshal(reqBody, &p)

	err := p.createProduct(a.DB)
	if err != nil {
		fmt.Printf("createProduct error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) allOrders(w http.ResponseWriter, r *http.Request) {
	// get all orders in JSON format
	// curl localhost:8080/orders
	orders, err := getOrders(a.DB)
	if err != nil {
		fmt.Printf("allOrders error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, orders)
}

func (a *App) fetchOrder(w http.ResponseWriter, r *http.Request) {
	// get a single order in JSON format
	// curl localhost:8080/order/1
	vars := mux.Vars(r)
	id := vars["id"]

	var o order
	o.ID, _ = strconv.Atoi(id)
	err := o.getOrder(a.DB)
	if err != nil {
		fmt.Printf("fetchOrder error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, o)
}

func (a *App) newOrder(w http.ResponseWriter, r *http.Request) {
	// create a new order (POST)
	// curl -X POST -H "Content-Type: application/json" -d '{"customer":"New Customer","items":[{"product_id":1,"quantity":1},{"product_id":2,"quantity":2}]}' localhost:8080/orders
	reqBody, _ := ioutil.ReadAll(r.Body)

	var o order
	json.Unmarshal(reqBody, &o)

	err := o.createOrder(a.DB)
	if err != nil {
		fmt.Printf("newOrder error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, item := range o.Items {
		var oi orderItem = item
		oi.ID = o.ID
		err := oi.createOrderItem(a.DB)
		if err != nil {
			fmt.Printf("newOrder, newOrderItem error: %s\n", err.Error())
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	respondWithJSON(w, http.StatusOK, o)
}

func (a *App) newOrderItems(w http.ResponseWriter, r *http.Request) {
	// for if send in OrderItems separately
	// i.e. that create the Order and later adds the OrderItems
	// first create the Order with empty OrderItems
	// curl -X POST -H "Content-Type: application/json" -d '{"customer":"New Customer", "items": []}' localhost:8080/orders
	// then add the OrderItems separately
	// curl -X POST -H "Content-Type: application/json" -d '[{"order_id":3,"product_id":1,"quantity":1},{"order_id":3,"product_id":2,"quantity":2}]' localhost:8080/orderitems
	reqBody, _ := ioutil.ReadAll(r.Body)

	var ois []orderItem
	json.Unmarshal(reqBody, &ois)

	for _, item := range ois {
		var oi orderItem = item
		err := oi.createOrderItem(a.DB)
		if err != nil {
			fmt.Printf("newOrderItems error: %s\n", err.Error())
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	respondWithJSON(w, http.StatusOK, ois)
}

func (a *App) Run() {
	fmt.Println("Server started and listening on port ", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}

// Helper functions
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
