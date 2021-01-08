package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleIndexRoute(client Doer) gin.HandlerFunc {
    return func(c *gin.Context) {
        sJSON := SnipcartData{}

        snipcartRequestHeaders := c.Request.Header["X-Snipcart-Requesttoken"]
        if len(snipcartRequestHeaders) != 1 {
            log.Println("No Snipcart token in request.")
            c.JSON(http.StatusBadRequest, gin.H{
                "message": "No Snipcart token in request.",
            })
            return
        }

        snipcartRequestToken := snipcartRequestHeaders[0]

        err := validateSnipcartToken(client, snipcartRequestToken)
        if err != nil {
            log.Printf("%v\n", err)
            c.JSON(http.StatusBadRequest, gin.H{
                "message": "Unable to validate snipcart token.",
            })
            return
        }

        if err := c.ShouldBindJSON(&sJSON); err != nil {
            log.Printf("%v\n", err)
        }

        if sJSON.EventName == "order.completed" && len(sJSON.Content.Items) != 0 {
            allItemStatus := s.handleItems(&sJSON.Content.Items, &sJSON)
            message, statusCode := checkServices(allItemStatus)

            c.JSON(statusCode, gin.H{
                "message": message,
                "data": allItemStatus,
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "BSCS Snipcart no action taken.",
        })
    }
}

func (s *Server) handleTestRoute() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Success",
        })
    }
}

func (s *Server) handleTestEmailLogging() gin.HandlerFunc {
    return func(c *gin.Context) {
        content, err := ioutil.ReadFile("myEmail.html")
        if err != nil {
            log.Printf("%v\n", err)
            c.JSON(http.StatusInternalServerError, gin.H{
                "message": "Internal Server Error",
            })
            return
        }

        text := string(content)
        s.emailLogger.LogEmail(text)

        c.JSON(http.StatusOK, gin.H{
            "message": "Success",
        })
    }
}

func (s *Server) serveRoutes() {
    s.router.Use(s.limitRequestsFromIPAddress())

    s.router.POST("/", s.handleIndexRoute(s.client))
    s.router.GET("/test", s.handleTestRoute())
    s.router.GET("/test-email-logging", s.handleTestEmailLogging())
}
