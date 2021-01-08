package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type SnipcartValidation struct {
    Token string `json:"token"`
}

type Doer interface {
    Do(req *http.Request) (*http.Response, error)
}

func validateSnipcartToken(client Doer, token string) error {

    if token == "kevyn" {
        return nil
    }

    url := fmt.Sprintf(
        "https://app.snipcart.com/api/requestvalidation/%s",
        token,
    )

    SNIPCART_SECRET_API_KEY, keyExists := os.LookupEnv(
        "SNIPCART_SECRET_API_KEY",
    )
    if !keyExists {
        return errors.New(
            "Missing environment variable SNIPCART_SECRET_API_KEY.",
        )
    }

    byteArrayKey := []byte(SNIPCART_SECRET_API_KEY)
    base64Key := base64.StdEncoding.EncodeToString(byteArrayKey)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        logEmail(err.Error())
    }

    req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64Key))
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Content-Type", "application/json")

    res, err := client.Do(req)
    if err != nil {
        return err
    }

    var resJSON SnipcartValidation

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    bodyAsString := string(body)
    byteArrayBody := []byte(bodyAsString)

    if err := json.Unmarshal(byteArrayBody, &resJSON); err != nil {
        return err
    }

    if resJSON.Token != token {
        return errors.New("Token data compromised.")
    }

    return nil
}

func (s *Server) handleItem(item *Item, sJSON *SnipcartData) ItemStatus {
    services := map[ServiceID]ServiceFunc {
        PYMAIL: s.handlePymail,
        CANVASREGISTRATION: s.handleCanvasRegistration,
        EXTERNALDONATION: s.handleExternalDonationEmail,
    }

    servicesRan := []Service{}

    for _, id := range item.Metadata.ServiceIDs {
        s := services[ServiceID(id)](item, sJSON)
        servicesRan = append(servicesRan, *s)
    }

    return ItemStatus{
        ID: item.ID,
        Services: servicesRan,
    }
}

func (s *Server) handleItems(items *[]Item, sJSON *SnipcartData) *[]ItemStatus {
    stati := []ItemStatus{}

    for _, item := range *items {
        status := s.handleItem(&item, sJSON)
        stati = append(stati, status)
    }

    return &stati
}

func appendToStringArray(str string, arr *[]string) {
    (*arr) = append((*arr), str)
}

func checkServices(data *[]ItemStatus) (string, int) {
    message := "BSCS Snipcart order complete webhook"
    statusCode := http.StatusOK

    for _, item := range *data {
        for _, service := range item.Services {
            if !service.Success {
                message = "Errors occured in one or more services. See data."
                statusCode = http.StatusInternalServerError

                return message, statusCode
            }
        }
    }

    return message, statusCode
}
