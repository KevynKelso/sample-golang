package main

import (
    "net/http"
    "sync"

    "golang.org/x/time/rate"
    "github.com/gin-gonic/gin"
)

// IPRateLimiter .
type IPRateLimiter struct {
        ips map[string]*rate.Limiter
        mu  *sync.RWMutex
        r   rate.Limit
        b   int
    }

// NewLimiter returns a new Limiter that allows events up to rate r and permits bursts of at most b tokens.
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
    i := &IPRateLimiter{
        ips: make(map[string]*rate.Limiter),
        mu:  &sync.RWMutex{},
        r:   r,
        b:   b,
    }

    return i
}

 //  AddIP creates a new rate limiter and adds it to the ips map,
 // using the IP address as the key
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
    i.mu.Lock()
    defer i.mu.Unlock()

    limiter := rate.NewLimiter(i.r, i.b)

    i.ips[ip] = limiter

    return limiter
}

// GetLimiter returns the rate limiter for the provided IP address if it exists.
// Otherwise calls AddIP to add IP address to the map
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
    i.mu.Lock()
    limiter, exists := i.ips[ip]

    if !exists {
        i.mu.Unlock()
        return i.AddIP(ip)
    }

    i.mu.Unlock()

    return limiter
}


func (s *Server) limitRequestsFromIPAddress() gin.HandlerFunc {
    // either returns context that gets passed to route handler or aborts with 429
    return func(c *gin.Context) {
        limiter := s.limiter.GetLimiter(c.Request.RemoteAddr)

        if !limiter.Allow() {
            c.AbortWithStatus(http.StatusTooManyRequests)
        }
    }
}
