package middleware

// Middleware represent the data-struct for middleware
type Middleware struct {
	// another stuff , may be needed by middleware
}

// InitMiddleware intialize the middleware
func InitMiddleware() *Middleware {
	return &Middleware{}
}
