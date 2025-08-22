# **orenngi middleware - A Simple Middleware Router for Go**

orenngi is a lightweight, flexible, and easy-to-use HTTP request router for Go. It allows you to route HTTP requests to specific handlers, while also applying middleware to your routes for added functionality like logging, authentication, and more.

## **Features**

* **Simple routing**: Supports both simple and complex route patterns.
* **Middleware support**: Chain multiple middleware functions to routes for enhanced request processing.
* **Flexible pattern matching**: Supports variables in route patterns (e.g., `/user/{id}`).
* **Context-aware**: Access and modify the context throughout the middleware stack.
* **Customizable error handling**: Handles errors with your own logic and templates.

## **Installation**

```bash
go get github.com/orenngi/mw
```

## **Basic Usage**

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/orenngi/mw"
)

func main() {
    r := mw.NewRouter()

    // Simple route
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to the homepage!")
    })

    // Route with middleware
    r.HandleFunc("/profile", userProfileHandler).Use(loggingMiddleware)

    // Start server
    http.ListenAndServe(":8080", r)
}

func userProfileHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "User Profile Page")
}

func loggingMiddleware(next mw.Handler) mw.Handler {
    return mw.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("Request: %s %s\n", r.Method, r.URL)
        next.ServeHTTP(w, r)
    })
}
```

## **Routing**

### Path Matching

orenngi middleware allows you to define dynamic routes with path variables:

```go
r.HandleFunc("/user/{id}", userHandler)
```

In this case, `{id}` will be a path parameter that can be accessed in the handler.

### HTTP Methods

You can handle specific HTTP methods with `Method`:

```go
r.Method("GET", "/profile", profileHandler)
r.Method("POST", "/user", createUserHandler)
```

### Route Groups

Route groups allow you to apply middleware to a set of routes:

```go
adminRoutes := r.Group("/admin").Use(authMiddleware)
adminRoutes.HandleFunc("/dashboard", dashboardHandler)
```

## **Middleware**

Middleware functions allow you to intercept and modify requests. Middleware can be added globally or per-route.

### Global Middleware

```go
r.Use(loggingMiddleware)
r.Use(authMiddleware)
```

### Per-Route Middleware

```go
r.HandleFunc("/admin", adminHandler).Use(authMiddleware)
```

You can also stack multiple middleware:

```go
r.HandleFunc("/secure", secureHandler).Use(authMiddleware).Use(loggingMiddleware)
```

### Built-in Middleware

* **Logging**: Logs every incoming request to the server.
* **Recovery**: Recovers from panics and returns a 500 Internal Server Error.
* **CORS**: Enables Cross-Origin Resource Sharing headers for API requests.

```go
r.Use(mw.Recovery())
r.Use(mw.CORS())
```

## **Error Handling**

mw allows you to define custom error handling logic, including status code and response body.

```go
r.HandleError(404, func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Page not found", http.StatusNotFound)
})

r.HandleError(500, func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Internal server error", http.StatusInternalServerError)
})
```

## **Context**

mw allows you to pass values into the request context. This is useful for things like storing user authentication details.

```go
r.HandleFunc("/profile", profileHandler).Use(func(next mw.Handler) mw.Handler {
    return mw.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := context.WithValue(r.Context(), "userID", 12345)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
})
```

Inside the handler, you can access the context:

```go
func profileHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID")
    fmt.Fprintf(w, "User Profile: %v", userID)
}
```

## **Development**

To run tests, use the following command:

```bash
go test ./...
```

## **Contributing**

Feel free to submit issues and pull requests! If you have any ideas, improvements, or feature requests, we welcome contributions to the project.

1. Fork the repository
2. Create a new branch (`git checkout -b feature-name`)
3. Commit your changes (`git commit -am 'Add feature'`)
4. Push to the branch (`git push origin feature-name`)
5. Create a new pull request

## **License**

mw is open-source software licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
