package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key") // Secret key for signing JWT tokens

// In-memory user store (username -> hashed password)
var users = map[string]string{}

// User represents the structure of a user with username and password
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims structure to be included in the JWT token
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Signup handler for registering new users
func Signup(w http.ResponseWriter, r *http.Request) {
	log.Println("Signup handler called")

	var creds User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		log.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request body"))
		return
	}

	// Check if the username already exists
	if _, exists := users[creds.Username]; exists {
		log.Printf("Username %s already exists", creds.Username)
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Username already exists"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error hashing password"))
		return
	}

	users[creds.Username] = string(hashedPassword)
	log.Printf("User %s successfully registered", creds.Username)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User successfully registered"))
}

// Login handler for authenticating users
func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login handler called")

	var creds User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		log.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the username exists
	storedPassword, ok := users[creds.Username]
	if !ok {
		log.Printf("User %s not found", creds.Username)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Username does not exist"))
		return
	}

	// Verify the password
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(creds.Password))
	if err != nil {
		log.Printf("Incorrect password for user %s", creds.Username)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Wrong credentials"))
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error signing the JWT token:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	log.Printf("User %s successfully logged in", creds.Username)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}

// Middleware to protect routes by checking for JWT token
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				log.Println("No JWT token found in the request")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			log.Println("Error retrieving JWT token from cookies")
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				log.Println("Invalid JWT signature")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			log.Println("Error parsing JWT token")
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			log.Println("Invalid JWT token")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		r.Header.Set("username", claims.Username)
		next.ServeHTTP(w, r)
	})
}
