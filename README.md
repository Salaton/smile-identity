# Smile Identity Go SDK (unofficial)

> Simple, idiomatic Go client for **Smile ID** server-side products i.e Basic KYC, Enhanced KYC, Document Verification and Phone Number Verification.  

---

## Installation
To start using the SDK in your Go project, run;

```bash
go get github.com/Salaton/smile-identity
```

## Getting Started
### Setting up Environment variables
Optionally, you can supply the base URL, APIKey, PartnerID and CallbackURL via environment variables. The following environment variables are recognized
```bash
SMILE_ID_API_KEY
SMILE_ID_PARTNER_ID
SMILE_ID_BASE_URL
SMILE_ID_CALLBACK_URL
```
This allows you to create a client with zero arguments;
```go
client, err := smile.NewClientFromEnvVars()
if err != nil{
    // handle error
}
```
If you prefer to supply the values manually;
```go
client, err := smile.NewClient(smile.Config{
    APIKey:      os.Getenv("SMILE_API_KEY"),
    PartnerID:   os.Getenv("SMILE_PARTNER_ID"),
    BaseURL:     "https://api.smileidentity.com/v1",
    CallbackURL: "https://example.com/webhook",
})
if err != nil{
    // handle err
}
```

## Usage Example
Below is a simple example showing how you can run a basic kyc request
```go
package main

import (
    "context"
    "fmt"
    "log"

    smile "github.com/Salaton/smile-identity"
)

func main() {
    // 1. Configure the client once at start-up
    client, err := smile.NewClient(smile.Config{
        APIKey:      os.Getenv("SMILE_API_KEY"),
        PartnerID:   os.Getenv("SMILE_PARTNER_ID"),
        BaseURL:     "https://api.smileidentity.com/v1",
        CallbackURL: "https://example.com/webhook",
    })
    if err != nil{
        // handle err
    }

    // 2. Build your request (Basic KYC example)
    req := smile.BasicKYCRequest{
        Country:   "KE",
        IDType:    smile.NATIONAL_ID,
        IDNumber:  "12345678901",
        FirstName: "John",
        LastName:  "Doe",
        DOB:       "1994-07-13",
    }

    // 3. Send the job & handle the response
    resp, err := client.BasicKYCVerification(context.Background(), req)
    if err != nil {
        log.Fatalf("kyc failed: %v", err)
    }
}
```

## How to Release
We use [release-please](https://github.com/googleapis/release-please) to manage our releases. This tool automates the process of creating a pull request with the next semantic version, updating the CHANGELOG, and tagging a release
****