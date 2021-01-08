package main

import (
    "io"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

var router *gin.Engine

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
    req, _ := http.NewRequest(method, path, body)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    return w
}

func TestTestRoute(t *testing.T) {
    body := "Success"
    //router := setupRouter()
    w := performRequest(router, "GET", "/test", nil)

    assert.Equal(t, http.StatusOK, w.Code)

    var response string = w.Body.String()

    assert.Equal(t, response, body)
}

func TestTestEmailRoute(t *testing.T) {
    body := "Success"
    //router := setupRouter()
    w := performRequest(router, "GET", "/test-email-logging", nil)
    assert.Equal(t, http.StatusOK, w.Code)

    var response string = w.Body.String()

    assert.Equal(t, response, body)
}

func TestIndexRoute(t *testing.T) {
    body := "No Snipcart token in request."

    w := performRequest(router, "POST", "/", nil)
    assert.Equal(t, http.StatusBadRequest, w.Code)

    var response string = w.Body.String()

    assert.Equal(t, response, body)
}

// func TestMain(m *testing.M) {
//     setTestEmailLogger()
//     s := createServer(client, emailLogger, router, limiter)

//     var client *http.Client = createHTTPClient()
//     var emailLogger *EmailLogger = createEmailLogger(logEmail)
//     var router *gin.Engine = gin.Default()
//     var limiter *IPRateLimiter = NewIPRateLimiter(1, 5)

//     var server *Server = createServer(client, emailLogger, router, limiter)

//     server.serveRoutes()

//     os.Exit(m.Run())
// }
