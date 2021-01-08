package main

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)

type PymailData struct {
    FormName               string        `json:"form_name"`
    Sendto                 []string      `json:"sendto"`
    InvoiceNumber          string        `json:"invoice_number"`
    Product                string        `json:"product"`
    Description            string        `json:"description"`
    Price                  float64       `json:"price"`
    Quantity               int           `json:"quantity"`
    Last4                  string        `json:"last4"`
    BillingName            string        `json:"billing_name"`
    Email                  string        `json:"email"`
    BillingCompanyName     string        `json:"billing_company_name"`
    BillingAddress1        string        `json:"billing_address1"`
    BillingAddress2        string        `json:"billing_address2"`
    BillingCity            string        `json:"billing_city"`
    BillingState           string        `json:"billing_state"`
    BillingCountry         string        `json:"billing_country"`
    BillingZip             string        `json:"billing_zip"`
    BillingPhone           string        `json:"billing_phone"`
    ShippingName           string        `json:"shipping_name"`
    ShippingCompanyName    string        `json:"shipping_company_name"`
    ShippingAddress1       string        `json:"shipping_address1"`
    ShippingAddress2       string        `json:"shipping_address2"`
    ShippingCity           string        `json:"shipping_city"`
    ShippingState          string        `json:"shipping_state"`
    ShippingCountry        string        `json:"shipping_country"`
    ShippingZip            string        `json:"shipping_zip"`
    ShippingPhone          string        `json:"shipping_phone"`
    Metadata               Metadata      `json:"metadata"`
    CustomFields           []CustomField `json:"customFields"`
}

func setPymailData(item *Item, sJSON *SnipcartData) *PymailData {
    data := PymailData{
        FormName: "BSCS Website Payment Form",
        Sendto: item.Metadata.Sendto,
        InvoiceNumber: sJSON.Content.InvoiceNumber,
        Product: item.Name,
        Description: item.Description,
        Price: float64(item.TotalPrice),
        Quantity: item.Quantity,
        Last4: sJSON.Content.CreditCardLast4Digits,
        BillingName: sJSON.Content.BillingAddressName,
        Email: sJSON.Content.User.Email,
        BillingCompanyName: sJSON.Content.BillingAddressCompanyName,
        BillingAddress1: sJSON.Content.BillingAddressAddress1,
        BillingAddress2: sJSON.Content.BillingAddressAddress2,
        BillingCity: sJSON.Content.BillingAddressCity,
        BillingState: sJSON.Content.BillingAddressProvince,
        BillingCountry: sJSON.Content.BillingAddressCountry,
        BillingZip: sJSON.Content.BillingAddressPostalCode,
        BillingPhone: sJSON.Content.BillingAddressPhone,
        ShippingName: sJSON.Content.ShippingAddressName,
        ShippingCompanyName: sJSON.Content.ShippingAddressCompanyName,
        ShippingAddress1: sJSON.Content.ShippingAddressAddress1,
        ShippingAddress2: sJSON.Content.ShippingAddressAddress2,
        ShippingCity: sJSON.Content.ShippingAddressCity,
        ShippingState: sJSON.Content.ShippingAddressProvince,
        ShippingCountry: sJSON.Content.ShippingAddressCountry,
        ShippingZip: sJSON.Content.ShippingAddressPostalCode,
        ShippingPhone: sJSON.Content.ShippingAddressPhone,
        Metadata: item.Metadata,
        CustomFields: item.CustomFields,
    }

    return &data
}

func removeSendToFromPymailData(d *PymailData) {
    (*d).Metadata.Sendto = []string{}
}

func handleNon200Response(responseBody io.ReadCloser) {
    body, err := ioutil.ReadAll(responseBody)
    if err != nil {
        logEmail(fmt.Sprintf("%s\n", err.Error()))
    }
    defer responseBody.Close()

    log.Printf("%s\n", string(body))
}

func checkResponse(res *http.Response, message string) error {
    if res.StatusCode != 200 {
        handleNon200Response(res.Body)

        return errors.New(message)
    }

    return nil
}

func (s *Server) sendRequest(req *http.Request) (*http.Response, error) {
    res, err := s.client.Do(req)
    if err != nil {
        return nil, err
    }

    return res, nil
}

func setAppJSONRequestHeader(req *http.Request) {
    (*req).Header.Set("Content-Type", "application/json")
}

func setupRequest(method, url string, body *bytes.Buffer) (*http.Request, error) {
    req, err := http.NewRequest(method, url, body)
    if err != nil {
        logEmail(err.Error())
        return nil, err
    }

    return req, nil
}

func (s *Server) handlePymail(item *Item, sJSON *SnipcartData) *Service {
    serviceName := "Pymail"
    log.Println("Pymail handler.")

    data := setPymailData(item, sJSON)
    removeSendToFromPymailData(data)

    dataBytes, err := json.Marshal(data)
    if err != nil {
        return failService(serviceName, err)
    }

    req, err := setupRequest(
        "POST",
        "https://pymail.bscs.org/form",
        bytes.NewBuffer(dataBytes),
    )
    if err != nil {
        return failService(serviceName, err)
    }

    setAppJSONRequestHeader(req)

    res, err := s.sendRequest(req)
    if err != nil {
        return failService(serviceName, err)
    }

    err = checkResponse(
        res,
        fmt.Sprintf("Pymail returned %v status code.", res.StatusCode),
    )
    if err != nil {
        return failService(serviceName, err)
    }

    return successService(serviceName)
}

func validateCanvasFields(metadata *Metadata) error {
    missingFields := []string{}

    if (metadata.FirstName) == "" {
        appendToStringArray("first_name", &missingFields)
    }
    if (metadata.LastName) == "" {
        appendToStringArray("last_name", &missingFields)
    }
    if (metadata.EmailAddress) == "" {
        appendToStringArray("email_address", &missingFields)
    }
    if (metadata.CanvasCourseID) == 0 {
        appendToStringArray("canvas_course_id", &missingFields)
    }

    if len(missingFields) != 0 {
        return errors.New(
            fmt.Sprintf("Missing required fields: %v.", missingFields),
        )
    }

    return nil
}

func (s *Server) handleCanvasRegistration(item *Item, sJSON *SnipcartData) *Service {
    log.Println("Canvas Registration handler.")
    serviceName := "Bojack Canvas Registration"

    if err := validateCanvasFields(&item.Metadata); err != nil {
        return failService(serviceName, err)
    }

    dataBytes, err := json.Marshal(item.Metadata)
    if err != nil {
        return failService(serviceName, err)
    }

    req, err := setupRequest(
        "POST",
        "https://bojack.bscs.org/enroll-in-course",
        bytes.NewBuffer(dataBytes),
    )
    if err != nil {
        return failService(serviceName, err)
    }

    setAppJSONRequestHeader(req)

    res, err := s.sendRequest(req)
    if err != nil {
        return failService(serviceName, err)
    }

    err = checkResponse(
        res,
        fmt.Sprintf("Bojack returned %v status code.", res.StatusCode),
    )
    if err != nil {
        return failService(serviceName, err)
    }

    return successService(serviceName)
}

type externalDonationData struct {
    Sendto        []string   `json:"sendto"`
    FormName      string     `json:"form_name"`
    Functions     []string   `json:"functions"`
    DonationType  string     `json:"donation_type"`
    Name          string     `json:"name"`
    TotalPrice    float64    `json:"total_price"`
    Date          time.Time  `json:"date"`
    Email         string     `json:"email"`

}

func (s *Server) handleExternalDonationEmail(item *Item, sJSON *SnipcartData) *Service {
    serviceName := "Pymail External Donation"
    log.Println("External Donation handler.")

    data := externalDonationData{
        Sendto: item.Metadata.Sendto,
        FormName: "External Donation",
        Functions: []string{"sendExternalDonationEmail"},
        DonationType: item.Name,
        Name: sJSON.Content.User.BillingAddressName,
        TotalPrice: item.TotalPrice,
        Date: sJSON.CreatedOn,
        Email: sJSON.Content.User.Email,
    }

    dataBytes, err := json.Marshal(data)
    if err != nil {
        return failService(serviceName, err)
    }

    req, err := setupRequest(
        "POST",
        "https://pymail.bscs.org/form",
        bytes.NewBuffer(dataBytes),
    )
    if err != nil {
        return failService(serviceName, err)
    }

    setAppJSONRequestHeader(req)

    res, err := s.sendRequest(req)
    if err != nil {
        return failService(serviceName, err)
    }

    err = checkResponse(
        res,
        fmt.Sprintf("Pymail returned %v status code.", res.StatusCode),
    )
    if err != nil {
        return failService(serviceName, err)
    }

    return successService(serviceName)
}

func failService(name string, err error) *Service {
    logEmail(err.Error())

    return &Service{
        Name: name,
        Success: false,
        Message: err.Error(),
    }
}

func successService(name string) *Service {
    return &Service{
        Name: name,
        Success: true,
        Message: "",
    }
}
