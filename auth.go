package playstation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ErrNPSSOEmpty is an error indicating that the NPSSO token is empty.
var ErrNPSSOEmpty = errors.New("npsso is empty")

// ErrNPSSOLength is an error indicating that the NPSSO token must be exactly 64 characters long.
var ErrNPSSOLength = errors.New("npsso must be exactly 64 characters")

// validateNPSSO validates the provided NPSSO token.
// It checks if the NPSSO token is empty or if its length is not exactly 64 characters.
//
// Parameters:
//
//	npsso (string): The NPSSO token to be validated.
//
// Returns:
//
//	error: An error indicating whether the NPSSO token is empty or not the correct length.
func validateNPSSO(npsso string) error {
	if npsso == "" {
		return ErrNPSSOEmpty
	}
	if len(npsso) != 64 {
		return ErrNPSSOLength
	}
	return nil
}

// Authenticate authenticates the client using the provided NPSSO token.
// It validates the NPSSO token and performs an authentication request to obtain tokens.
//
// Parameters:
//
//	ctx (context.Context): The context for controlling the request lifetime.
//	npsso (string): The NPSSO token used for authentication.
//
// Returns:
//
//	*ClientAPI: A pointer to the ClientAPI containing the authenticated client and tokens.
//	error: An error indicating whether the authentication was successful or not.
func (c *Client) Authenticate(ctx context.Context, npsso string) (*ClientAPI, error) {
	err := validateNPSSO(npsso)
	if err != nil {
		return nil, fmt.Errorf("invalid npsso: %v", err)
	}
	tokens, err := c.authRequest(ctx, npsso)
	if err != nil {
		return nil, fmt.Errorf("can't do auth request: %w", err)
	}

	var clientAPI = ClientAPI{
		Client: c,
		Tokens: tokens,
		NPSSO:  npsso,
	}

	return &clientAPI, nil
}

// authRequest sends an authentication request using the provided NPSSO token.
// It prepares the authorization URL, sends the request, and handles the response to obtain tokens.
//
// Parameters:
//
//	ctx (context.Context): The context for controlling the request lifetime.
//	npsso (string): The NPSSO token used for authentication.
//
// Returns:
//
//	*Tokens: A pointer to the Tokens containing the authentication tokens.
//	error: An error indicating whether the authentication request was successful or not.
func (c *Client) authRequest(ctx context.Context, npsso string) (*Tokens, error) {
	if npsso == "" {
		return nil, errors.New("npsso parameter is required")
	}

	// Prepare authorization URL parameters
	params := url.Values{}
	params.Add("access_type", "offline")
	params.Add("client_id", "09515159-7237-4370-9b40-3806e67c0891")
	params.Add("response_type", "code")
	params.Add("scope", "psn:mobile.v2.core psn:clientapp")
	params.Add("redirect_uri", "com.scee.psxandroid.scecompcall://redirect")

	authURL := "https://ca.account.sony.com/api/authz/v3/oauth/authorize?" + params.Encode()

	// Create request with context
	req, err := http.NewRequestWithContext(ctx, "GET", authURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add cookie
	req.Header.Add("Cookie", fmt.Sprintf("npsso=%s", npsso))

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("request cancelled: %w", err)
	}

	// Check for redirect and extract code
	if resp.StatusCode != http.StatusFound {
		return nil, errors.New("error: check npsso")
	}

	location := resp.Header.Get("Location")
	locationURL, err := url.Parse(location)
	if err != nil {
		return nil, fmt.Errorf("error parsing location URL: %w", err)
	}

	code := locationURL.Query().Get("code")
	if !strings.HasPrefix(code, "v3") {
		return nil, errors.New("error: check npsso")
	}

	// Prepare token request
	tokenURL := "https://ca.account.sony.com/api/authz/v3/oauth/token"
	tokenData := url.Values{}
	tokenData.Set("code", code)
	tokenData.Set("redirect_uri", "com.scee.psxandroid.scecompcall://redirect")
	tokenData.Set("grant_type", "authorization_code")
	tokenData.Set("token_format", "jwt")

	// Create token request with context
	tokenReq, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, strings.NewReader(tokenData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating token request: %w", err)
	}

	tokenReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	tokenReq.Header.Set("Authorization", "Basic MDk1MTUxNTktNzIzNy00MzcwLTliNDAtMzgwNmU2N2MwODkxOnVjUGprYTV0bnRCMktxc1A=")

	// Send token request
	tokenResp, err := c.httpClient.Do(tokenReq)
	if err != nil {
		return nil, fmt.Errorf("error sending token request: %w", err)
	}
	defer tokenResp.Body.Close()

	// Check for context cancellation again
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("token request cancelled: %w", err)
	}

	if tokenResp.StatusCode != http.StatusOK {
		return nil, errors.New("error: unable to obtain Authentication Token")
	}

	// Read and parse response
	body, err := io.ReadAll(tokenResp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var tokenResponse Tokens
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, fmt.Errorf("error parsing token response: %w", err)
	}

	// Set expiration times
	tokenResponse.AccessExpiresTime = time.Now().Add(time.Second * time.Duration(tokenResponse.AccessExpires))
	tokenResponse.RefreshExpiresTime = time.Now().Add(time.Second * time.Duration(tokenResponse.RefreshExpires))

	return &tokenResponse, nil
}
