package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/random"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type movie struct{
	ID string `json:"id"`
	isbn string `json:"isbn"`
	title string `json:"title"`
	Director *Director `json:"Director"`
}

type Director struct{
	firstName string `json:"firstName"`
	lastName string `json:"lastName"`
}

var movies []movie


func main() {
	
}