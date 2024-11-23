[![Go Reference](https://pkg.go.dev/badge/github.com/FrostBreker/go-playstation-api.svg)](https://pkg.go.dev/github.com/FrostBreker/go-playstation-api)
# go-playstation-api

This is a simple API that allows you to interact with the PlayStation API.

## Read First
Corresponding to my research how PSN works you need npsso to interact with Sony servers.
Instructions how to get it below.  

### How to get NPSSO

1. Open your browser and go to https://my.playstation.com/
2. Login to your account
3. Open https://ca.account.sony.com/api/v1/ssocookie in new tab

### Functionality
- You can get user profile info
- You can get user games info


## Installation

```bash
go get github.com/FrostBreker/go-playstation-api
```

## Usage

```go
package main

import (
	"context"
	"log"
	"net/http"
	"github.com/FrostBreker/go-playstation-api"
	"time"
)

func main() {
	// Handle errors from option functions
	regionOpt, err := playstation.WithRegion(playstation.RegionFR)
	if err != nil {
		log.Fatalf("Error setting region: %v", err)
	}

	langOpt, err := playstation.WithLanguage(playstation.LangFR)
	if err != nil {
		log.Fatalf("Error setting language: %v", err)
	}

	clientOpt, err := playstation.WithClient(&http.Client{})
	if err != nil {
		log.Fatalf("Error setting HTTP client: %v", err)
	}

	// Create the client using the package and use the options to set the region and lang to fr and use an HTTP client
	client := playstation.NewClient(regionOpt, langOpt, clientOpt)

	// Use the client variable to avoid unused variable error
	log.Printf("Client created: %+v", client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use the client to authenticate
	clientAPI, err := client.Authenticate(ctx, "your-npsso")
	if err != nil {
		log.Fatalf("Error authenticating: %v", err)
	}

	// Use the clientAPI to get user account, profile, and games

	// Use the clientAPI to get the user account id
	userAccount, err := clientAPI.GetUserAccountId(ctx, "james")
	if err != nil {
		log.Fatalf("Error getting user account id: %v", err)
	}

	log.Printf("User account id: %+v", userAccount)

	// Use the user account to get the user profile
	userProfile, err := clientAPI.GetUserProfile(ctx, userAccount.Profile.AccountID)
	if err != nil {
		log.Fatalf("Error getting user profile: %v", err)
	}

	log.Printf("User profile: %+v", userProfile)

	// Use the user account to get the user games
	userGames, err := clientAPI.GetUserGames(ctx, userAccount.Profile.AccountID)
	if err != nil {
		log.Fatalf("Error getting user games: %v", err)
	}

	log.Printf("User games: %+v", userGames)
}
```

This project highly inspired by https://github.com/Tustin/psn-php and https://github.com/sizovilya/go-psn-api.