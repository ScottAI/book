package adv

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

//Get user's auth code
func GetAuthCode() Middleware {
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		//Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			code := 0
			//auth code is available only when access root
			if r.URL.Path != "/"{
				code = -1
			}
			//create a new request context containing the auth code, context available >= go 1.7
			ctxWithUser := context.WithValue(r.Context(), code, "User")
			//create a new request using that new context
			rWithUser := r.WithContext(ctxWithUser)
			//call the real handler, passing the new request
			f(w,rWithUser)
		}
	}
}

// Ensure user's auth
func EnsureAuth() Middleware {
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		//Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value(0)
			if user != nil {
				log.Println("auth available!")
			}else {
				http.Error(w,"Please sign in!",http.StatusUnauthorized)
				return
			}
			// Call the next middleware/handler in chain
			f(w,r)
		}
	}
}
// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()
			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}else {
				log.Println("request is :",m)
			}
			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
func Auth(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w,"You r authorized!")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"),GetAuthCode(), Logging()))
	http.HandleFunc("/auth/", Chain(Auth, Method("GET"),GetAuthCode(),EnsureAuth(), Logging()))
	http.ListenAndServe(":7775", nil)
}