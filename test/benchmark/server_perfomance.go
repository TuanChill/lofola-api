package main

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"strings"
	"sync"
)

const (
	domain        = "https://www.luongtuan.xyz"
	requestNumber = 10000
	emailDomain   = "@gmail.com"
	alphabet      = "abcdefghiklmnopqrstuvwxyz"
	usernameLen   = 10
	passwordLen   = 10
	emailLen      = 10
)

// RandomString generates a random string of the specified length.
func RandomString(n int) string {
	var sb strings.Builder
	alphabetLen := len(alphabet)
	for i := 0; i < n; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(alphabetLen)))
		if err != nil {
			log.Fatalf("Failed to generate random index: %v", err)
		}
		sb.WriteByte(alphabet[index.Int64()])
	}
	return sb.String()
}

// sendRegister sends a registration request to the given domain.
func sendRegister(domain string) {
	url := domain + "/api/v1/auth/register"

	// Create random email, username, and password
	email := RandomString(emailLen) + emailDomain
	username := RandomString(usernameLen)
	password := RandomString(passwordLen)

	// Prepare the request body
	data := map[string]string{
		"email":    email,
		"username": username,
		"password": password,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return
	}

	// Send the POST request
	res, err := http.Post(url, "application/json", strings.NewReader(string(jsonData)))
	if err != nil {
		log.Printf("Failed to send request: %v", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Received non-OK response: %v", res.Status)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(requestNumber)

	for i := 0; i < requestNumber; i++ {
		go func() {
			defer wg.Done()
			sendRegister(domain)
		}()
	}

	wg.Wait()
	log.Printf("Completed sending %d registration requests.", requestNumber)
}
