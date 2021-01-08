package main

import (
    "time"
)

type ServiceFunc func(*Item, *SnipcartData) *Service

type ServiceID int

const (
    PYMAIL ServiceID = iota
    CANVASREGISTRATION
    EXTERNALDONATION
)

type Service struct {
    Name        string
    Success     bool
    Message     string
}

type ItemStatus struct {
    ID          string
    Services    []Service
}

type CustomField struct {
    DisplayValue  string    `json:"displayValue"`
    Value         string    `json:"value"`
    Name          string    `json:"name"`
}

type Metadata struct {
    Sendto               []string   `json:"sendto"`
    FirstName            string     `json:"first_name"`
    LastName             string     `json:"last_name"`
    EmailAddress         string     `json:"email_address"`
    CanvasCourseID       int        `json:"canvas_course_id"`
    ServiceIDs            []int     `json:"serviceIDs"`
}

type Item struct {
    PaymentSchedule struct {
        Interval          int         `json:"interval"`
        IntervalCount     int         `json:"intervalCount"`
        TrialPeriodInDays interface{} `json:"trialPeriodInDays"`
        StartsOn          time.Time   `json:"startsOn"`
    } `json:"paymentSchedule"`
    PausingAction          string        `json:"pausingAction"`
    CancellationAction     string        `json:"cancellationAction"`
    Token                  string        `json:"token"`
    Name                   string        `json:"name"`
    Price                  float64       `json:"price"`
    Quantity               int           `json:"quantity"`
    FileGUID               interface{}   `json:"fileGuid"`
    URL                    string        `json:"url"`
    ID                     string        `json:"id"`
    InitialData            string        `json:"initialData"`
    Description            string        `json:"description"`
    Categories             []interface{} `json:"categories"`
    TotalPriceWithoutTaxes float64       `json:"totalPriceWithoutTaxes"`
    Weight                 interface{}   `json:"weight"`
    Image                  string        `json:"image"`
    OriginalPrice          interface{}   `json:"originalPrice"`
    UniqueID               string        `json:"uniqueId"`
    Stackable              bool          `json:"stackable"`
    MinQuantity            int           `json:"minQuantity"`
    MaxQuantity            int           `json:"maxQuantity"`
    AddedOn                time.Time     `json:"addedOn"`
    ModificationDate       time.Time     `json:"modificationDate"`
    Shippable              bool          `json:"shippable"`
    Taxable                bool          `json:"taxable"`
    Duplicatable           bool          `json:"duplicatable"`
    Width                  interface{}   `json:"width"`
    Height                 interface{}   `json:"height"`
    Length                 interface{}   `json:"length"`
    Metadata               Metadata      `json:"metadata"`
    TotalPrice      float64              `json:"totalPrice"`
    TotalWeight     float64              `json:"totalWeight"`
    Taxes           []interface{} `json:"taxes"`
    AlternatePrices struct {
    } `json:"alternatePrices"`
    CustomFields                       []CustomField `json:"customFields"`
    UnitPrice                          int           `json:"unitPrice"`
    HasDimensions                      bool          `json:"hasDimensions"`
    HasTaxesIncluded                   bool          `json:"hasTaxesIncluded"`
    TotalPriceWithoutDiscountsAndTaxes int           `json:"totalPriceWithoutDiscountsAndTaxes"`
}

type SnipcartData struct {
    EventName string    `json:"eventName"`
    Mode      string    `json:"mode"`
    CreatedOn time.Time `json:"createdOn"`
    Content   struct {
        Discounts []interface{} `json:"discounts"`
        Items   []Item `json:"items"`
        Plans   []interface{} `json:"plans"`
        Refunds []interface{} `json:"refunds"`
        Taxes   []interface{} `json:"taxes"`
        User    struct {
            ID         string `json:"id"`
            Email      string `json:"email"`
            Mode       string `json:"mode"`
            Statistics struct {
                OrdersCount        int         `json:"ordersCount"`
                OrdersAmount       interface{} `json:"ordersAmount"`
                SubscriptionsCount int         `json:"subscriptionsCount"`
            } `json:"statistics"`
            CreationDate                 time.Time   `json:"creationDate"`
            BillingAddressFirstName      string      `json:"billingAddressFirstName"`
            BillingAddressName           string      `json:"billingAddressName"`
            BillingAddressCompanyName    string      `json:"billingAddressCompanyName"`
            BillingAddressAddress1       string      `json:"billingAddressAddress1"`
            BillingAddressAddress2       string      `json:"billingAddressAddress2"`
            BillingAddressCity           string      `json:"billingAddressCity"`
            BillingAddressCountry        string      `json:"billingAddressCountry"`
            BillingAddressProvince       string      `json:"billingAddressProvince"`
            BillingAddressPostalCode     string      `json:"billingAddressPostalCode"`
            BillingAddressPhone          string      `json:"billingAddressPhone"`
            ShippingAddressFirstName     string      `json:"shippingAddressFirstName"`
            ShippingAddressName          string      `json:"shippingAddressName"`
            ShippingAddressCompanyName   string      `json:"shippingAddressCompanyName"`
            ShippingAddressAddress1      string      `json:"shippingAddressAddress1"`
            ShippingAddressAddress2      string      `json:"shippingAddressAddress2"`
            ShippingAddressCity          string      `json:"shippingAddressCity"`
            ShippingAddressCountry       string      `json:"shippingAddressCountry"`
            ShippingAddressProvince      string      `json:"shippingAddressProvince"`
            ShippingAddressPostalCode    string      `json:"shippingAddressPostalCode"`
            ShippingAddressPhone         string      `json:"shippingAddressPhone"`
            ShippingAddressSameAsBilling bool        `json:"shippingAddressSameAsBilling"`
            Status                       string      `json:"status"`
            SessionToken                 interface{} `json:"sessionToken"`
            GravatarURL                  string      `json:"gravatarUrl"`
            BillingAddress               struct {
                FullName    string      `json:"fullName"`
                FirstName   interface{} `json:"firstName"`
                Name        string      `json:"name"`
                Company     string      `json:"company"`
                Address1    string      `json:"address1"`
                Address2    string      `json:"address2"`
                FullAddress string      `json:"fullAddress"`
                City        string      `json:"city"`
                Country     string      `json:"country"`
                PostalCode  string      `json:"postalCode"`
                Province    string      `json:"province"`
                Phone       string      `json:"phone"`
                VatNumber   interface{} `json:"vatNumber"`
            } `json:"billingAddress"`
            ShippingAddress struct {
                FullName    string      `json:"fullName"`
                FirstName   string      `json:"firstName"`
                Name        string      `json:"name"`
                Company     string      `json:"company"`
                Address1    string      `json:"address1"`
                Address2    string      `json:"address2"`
                FullAddress string           `json:"fullAddress"`
                City        string      `json:"city"`
                Country     string      `json:"country"`
                PostalCode  string      `json:"postalCode"`
                Province    string      `json:"province"`
                Phone       string      `json:"phone"`
                VatNumber   string      `json:"vatNumber"`
            } `json:"shippingAddress"`
        } `json:"user"`
        Token                     string      `json:"token"`
        IsRecurringOrder          bool        `json:"isRecurringOrder"`
        ParentToken               interface{} `json:"parentToken"`
        ParentInvoiceNumber       interface{} `json:"parentInvoiceNumber"`
        SubscriptionID            interface{} `json:"subscriptionId"`
        Currency                  string      `json:"currency"`
        CreationDate              time.Time   `json:"creationDate"`
        ModificationDate          time.Time   `json:"modificationDate"`
        RecoveredFromCampaignID   interface{} `json:"recoveredFromCampaignId"`
        Status                    string      `json:"status"`
        PaymentStatus             string      `json:"paymentStatus"`
        Email                     string      `json:"email"`
        WillBePaidLater           bool        `json:"willBePaidLater"`
        BillingAddressFirstName   interface{} `json:"billingAddressFirstName"`
        BillingAddressName        string      `json:"billingAddressName"`
        BillingAddressCompanyName string      `json:"billingAddressCompanyName"`
        BillingAddressAddress1    string      `json:"billingAddressAddress1"`
        BillingAddressAddress2    string      `json:"billingAddressAddress2"`
        BillingAddressCity        string      `json:"billingAddressCity"`
        BillingAddressCountry     string      `json:"billingAddressCountry"`
        BillingAddressProvince    string      `json:"billingAddressProvince"`
        BillingAddressPostalCode  string      `json:"billingAddressPostalCode"`
        BillingAddressPhone       string      `json:"billingAddressPhone"`
        BillingAddress            struct {
            FullName    string      `json:"fullName"`
            FirstName   interface{} `json:"firstName"`
            Name        string      `json:"name"`
            Company     string      `json:"company"`
            Address1    string      `json:"address1"`
            Address2    string      `json:"address2"`
            FullAddress string      `json:"fullAddress"`
            City        string      `json:"city"`
            Country     string      `json:"country"`
            PostalCode  string      `json:"postalCode"`
            Province    string      `json:"province"`
            Phone       string      `json:"phone"`
            VatNumber   interface{} `json:"vatNumber"`
        } `json:"billingAddress"`
        ShippingAddressFirstName     string `json:"shippingAddressFirstName"`
        ShippingAddressName          string `json:"shippingAddressName"`
        ShippingAddressCompanyName   string `json:"shippingAddressCompanyName"`
        ShippingAddressAddress1      string `json:"shippingAddressAddress1"`
        ShippingAddressAddress2      string `json:"shippingAddressAddress2"`
        ShippingAddressCity          string `json:"shippingAddressCity"`
        ShippingAddressCountry       string `json:"shippingAddressCountry"`
        ShippingAddressProvince      string `json:"shippingAddressProvince"`
        ShippingAddressPostalCode    string `json:"shippingAddressPostalCode"`
        ShippingAddressPhone         string `json:"shippingAddressPhone"`
        ShippingAddress              string `json:"shippingAddress"`
        ShippingAddressSameAsBilling bool        `json:"shippingAddressSameAsBilling"`
        CreditCardLast4Digits        string      `json:"creditCardLast4Digits"`
        TrackingNumber               interface{} `json:"trackingNumber"`
        TrackingURL                  interface{} `json:"trackingUrl"`
        ShippingFees                 float64     `json:"shippingFees"`
        ShippingProvider             interface{} `json:"shippingProvider"`
        ShippingMethod               interface{} `json:"shippingMethod"`
        CardHolderName               string      `json:"cardHolderName"`
        PaymentMethod                string      `json:"paymentMethod"`
        Notes                        interface{} `json:"notes"`
        CustomFieldsJSON             string      `json:"customFieldsJson"`
        UserID                       string      `json:"userId"`
        CompletionDate               time.Time   `json:"completionDate"`
        CardType                     string      `json:"cardType"`
        PaymentGatewayUsed           string      `json:"paymentGatewayUsed"`
        PaymentDetails               struct {
            IconURL      interface{} `json:"iconUrl"`
            Display      interface{} `json:"display"`
            Instructions interface{} `json:"instructions"`
        } `json:"paymentDetails"`
        TaxProvider                        string        `json:"taxProvider"`
        Lang                               string        `json:"lang"`
        RefundsAmount                      float64       `json:"refundsAmount"`
        AdjustedAmount                     float64       `json:"adjustedAmount"`
        FinalGrandTotal                    float64       `json:"finalGrandTotal"`
        TotalNumberOfItems                 int           `json:"totalNumberOfItems"`
        InvoiceNumber                      string        `json:"invoiceNumber"`
        BillingAddressComplete             bool          `json:"billingAddressComplete"`
        ShippingAddressComplete            bool          `json:"shippingAddressComplete"`
        ShippingMethodComplete             bool          `json:"shippingMethodComplete"`
        SavedAmount                        float64       `json:"savedAmount"`
        Subtotal                           float64       `json:"subtotal"`
        BaseTotal                          float64       `json:"baseTotal"`
        ItemsTotal                         float64       `json:"itemsTotal"`
        TotalPriceWithoutDiscountsAndTaxes float64       `json:"totalPriceWithoutDiscountsAndTaxes"`
        TaxableTotal                       float64       `json:"taxableTotal"`
        GrandTotal                         float64       `json:"grandTotal"`
        Total                              float64       `json:"total"`
        TotalWeight                        float64       `json:"totalWeight"`
        TotalRebateRate                    float64       `json:"totalRebateRate"`
        CustomFields                       []interface{} `json:"customFields"`
        ShippingEnabled                    bool          `json:"shippingEnabled"`
        NumberOfItemsInOrder               int           `json:"numberOfItemsInOrder"`
        PaymentTransactionID               string        `json:"paymentTransactionId"`
        Metadata                           struct {
        } `json:"metadata"`
        TaxesTotal int `json:"taxesTotal"`
        ItemsCount int `json:"itemsCount"`
        Summary    struct {
            Subtotal                      float64       `json:"subtotal"`
            TaxableTotal                  float64       `json:"taxableTotal"`
            Total                         float64       `json:"total"`
            PayableNow                    float64       `json:"payableNow"`
            PaymentMethod                 string        `json:"paymentMethod"`
            Taxes                         []interface{} `json:"taxes"`
            DiscountInducedTaxesVariation float64       `json:"discountInducedTaxesVariation"`
            AdjustedTotal                 float64       `json:"adjustedTotal"`
            Shipping                      interface{}   `json:"shipping"`
        } `json:"summary"`
        IPAddress        string `json:"ipAddress"`
        UserAgent        string `json:"userAgent"`
        HasSubscriptions bool   `json:"hasSubscriptions"`
    } `json:"content"`
}
