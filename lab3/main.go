package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// greetHandler handles the root ("/") route.
// It expects query parameters 'name' and 'age' and responds with a greeting message.
func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	if name == "" || age == "" {
		http.Error(w, "Missing 'name' or 'age' query parameter", http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("Меня зовут %s, мне %s лет", name, age)
	fmt.Fprintln(w, response)
}

// addHandler handles the "/add" route.
// It expects two query parameters 'a' and 'b' and responds with their sum.
func addHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseQueryParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := a + b
	fmt.Fprintf(w, "Result of addition: %d", result)
}

// subHandler handles the "/sub" route.
// It expects two query parameters 'a' and 'b' and responds with their difference.
func subHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseQueryParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := a - b
	fmt.Fprintf(w, "Result of subtraction: %d", result)
}

// mulHandler handles the "/mul" route.
// It expects two query parameters 'a' and 'b' and responds with their product.
func mulHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseQueryParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := a * b
	fmt.Fprintf(w, "Result of multiplication: %d", result)
}

// divHandler handles the "/div" route.
// It expects two query parameters 'a' and 'b' and responds with their quotient.
// Includes error handling for division by zero.
func divHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseQueryParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if b == 0 {
		http.Error(w, "Division by zero is not allowed", http.StatusBadRequest)
		return
	}
	result := a / b
	fmt.Fprintf(w, "Result of division: %d", result)
}

// modHandler handles the "/mod" route.
// It expects two query parameters 'a' and 'b' and responds with the modulus (remainder) of a divided by b.
func modHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseQueryParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if b == 0 {
		http.Error(w, "Modulus by zero is not allowed", http.StatusBadRequest)
		return
	}
	result := a % b
	fmt.Fprintf(w, "Result of modulus: %d", result)
}

// parseQueryParams parses 'a' and 'b' query parameters from the request.
// Converts them to integers and returns an error if any parameter is invalid.
func parseQueryParams(r *http.Request) (int, int, error) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid parameter 'a': %v", err)
	}
	b, err := strconv.Atoi(bStr)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid parameter 'b': %v", err)
	}
	return a, b, nil
}

// countCharsHandler handles the "/count" route.
// Expects a JSON body with a 'text' field and returns the count of each character in the text.
func countCharsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var requestData map[string]string
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	text, ok := requestData["text"]
	if !ok {
		http.Error(w, "Missing 'text' field", http.StatusBadRequest)
		return
	}

	// Count each character in the input text
	charCount := make(map[string]int)
	for _, char := range text {
		charCount[string(char)]++
	}

	// Set response content type to JSON and encode the character count map as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(charCount)
}

// main initializes the server and routes.
func main() {
	// Register route handlers for various endpoints
	http.HandleFunc("/", greetHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/sub", subHandler)
	http.HandleFunc("/mul", mulHandler)
	http.HandleFunc("/div", divHandler)
	http.HandleFunc("/mod", modHandler) // New modulus route
	http.HandleFunc("/count", countCharsHandler)

	// Start the server and log startup or fatal errors
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
