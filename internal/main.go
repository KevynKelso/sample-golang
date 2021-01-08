package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "gopkg.in/gomail.v2"
)

type Server struct {
    client                              *http.Client
    emailLogger                         *EmailLogger
    router                              *gin.Engine
    limiter                             *IPRateLimiter
}

func createHTTPClient() *http.Client {
    client := &http.Client{
        Timeout: time.Second * 10,
    }

    return client
}

func logEmail(message string) error {
    SENDER_EMAIL, sender_email_exists := os.LookupEnv("SENDER_EMAIL")
    SENDER_PASSWORD, sender_password_exists := os.LookupEnv("SENDER_PASSWORD")
    MARKHOR_ADMINS, markhor_admins_exists := os.LookupEnv("MARKHOR_ADMINS")

    if !sender_email_exists || !sender_password_exists || !markhor_admins_exists {
        log.Fatal("Missing environment variables.")
    }

    hostname := "smtp.office365.com"
    port := 587
    to := strings.Split(MARKHOR_ADMINS, ",")
    d := gomail.NewDialer(hostname, port, SENDER_EMAIL, SENDER_PASSWORD)

    s, err := d.Dial()
    if err != nil {
        log.Println("Error connecting to email client.")
        return err
    }

    m := gomail.NewMessage()
    for _, recipient := range to {
        m.SetHeader("From", SENDER_EMAIL)
        m.SetHeader("To", recipient)
        m.SetHeader("Subject", "Markhor")
        m.SetBody("text/html", message)

        if err := gomail.Send(s, m); err != nil {
            log.Printf("Could not send email to %q: %v", recipient, err)
            return err
        }
        m.Reset()
    }

    return nil
}

func createEmailLogger(logEmail LogFunction) *EmailLogger {
    emailLogger := NewEmailLogger(logEmail)

    return emailLogger
}

func createServer(
                  client *http.Client,
                  emailLogger *EmailLogger,
                  router *gin.Engine,
                  limiter *IPRateLimiter,
                 ) *Server {
    s := &Server{
        client: client,
        emailLogger: emailLogger,
        router: router,
        limiter: limiter,
    }

    return s
}

func init() {
    if err := godotenv.Load("../.env"); err != nil {
        log.Fatal("No .env file found")
    }
}

func main() {
    PORT, port_exists := os.LookupEnv("PORT")
    ENVIRONMENT, environment_exists := os.LookupEnv("ENVIRONMENT")

    if !port_exists {
        log.Fatal("Missing PORT environment variable.")
    }

    if !environment_exists {
        log.Fatal("Missing ENVIRONMENT environment variable.")
    }

    if ENVIRONMENT == "production" {
        gin.SetMode(gin.ReleaseMode)
    }

    server := createServer(
        createHTTPClient(),
        createEmailLogger(logEmail),
        gin.Default(),
        NewIPRateLimiter(1,5), // 1 event per second w/ 5 tokens
    )

    server.serveRoutes()

    server.router.Run(fmt.Sprintf(":%s", PORT))
}
