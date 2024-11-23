package playstation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// request sends an HTTP GET request to the specified URL using the provided tokens for authentication.
// It handles the creation of the request, adding necessary headers, sending the request, and processing the response.
//
// Parameters:
//
//	ctx (context.Context): The context for controlling the request lifetime.
//	tokens (*Tokens): The tokens used for authentication, must contain a valid access token.
//	url (string): The URL to which the request is sent.
//
// Returns:
//
//	[]byte: The response body as a byte slice if the request is successful.
//	error: An error indicating whether the request was successful or not.
func (c *ClientAPI) request(ctx context.Context, url string) ([]byte, error) {
	if c.Tokens == nil || c.Tokens.AccessToken == "" {
		return nil, errors.New("invalid tokens: access token is required")
	}

	//TODO: Implement token refresh logic without using npsso (with storing refresh token)
	if c.Tokens.AccessExpiresTime.Before(time.Now()) {
		newTokens, err := c.Client.authRequest(ctx, c.NPSSO)
		if err != nil {
			return nil, fmt.Errorf("error refreshing tokens: %w", err)
		}
		c.Tokens = newTokens
	}

	// Create new request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Tokens.AccessToken))
	req.Header.Set("Accept", "application/json")
	if c.Client.lang != "" {
		req.Header.Set("Accept-Language", string(c.Client.lang))
	}

	// Send request
	resp, err := c.Client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing body")
		}
	}(resp.Body)

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Handle different status codes
	switch resp.StatusCode {
	case http.StatusOK:
		return body, nil
		//case http.StatusUnauthorized:
		//	return nil, fmt.Errorf("unauthorized: invalid or expired token")
		//case http.StatusForbidden:
		//	return nil, fmt.Errorf("forbidden: insufficient permissions")
	case http.StatusUnauthorized:
		return body, nil
	case http.StatusForbidden:
		return body, nil
	case http.StatusNotFound:
		return nil, fmt.Errorf("resource not found: %s", url)
	case http.StatusTooManyRequests:
		return nil, fmt.Errorf("rate limit exceeded")
	default:
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}
}

// requestAndUnmarshal sends an HTTP GET request to the specified URL using the provided tokens for authentication,
// and unmarshals the response body into the provided interface.
//
// Parameters:
//
//	ctx (context.Context): The context for controlling the request lifetime.
//	tokens (*Tokens): The tokens used for authentication, must contain a valid access token.
//	url (string): The URL to which the request is sent.
//	v (interface{}): The interface into which the response body is unmarshaled.
//
// Returns:
//
//	error: An error indicating whether the request or unmarshaling was successful or not.
func (c *ClientAPI) requestAndUnmarshal(ctx context.Context, url string, v interface{}) error {
	body, err := c.request(ctx, url)
	if err != nil {
		return err
	}

	var requestError RequestError
	if err := json.Unmarshal(body, &requestError); err != nil {
		return fmt.Errorf("error unmarshaling response: %w, body: %s", err, string(body))
	}

	if requestError.Error.Code != 0 {
		return fmt.Errorf("request error: %s", requestError.Error.Message)
	}

	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("error unmarshaling response: %w, body: %s", err, string(body))
	}

	return nil
}

// GetUserAccountId retrieves the user account ID and other details for the specified online ID.
//
// Parameters:
//
//	ctx (context.Context): The context for controlling the request lifetime.
//	onlineId (string): The online ID of the user whose account details are being retrieved.
//
// Returns:
//
//	*UserAccountResponse: A pointer to the UserAccountResponse containing the user's account details.
//	error: An error indicating whether the request was successful or not.
func (c *ClientAPI) GetUserAccountId(ctx context.Context, onlineId string) (*UserAccountResponse, error) {
	url := fmt.Sprintf("https://us-prof.np.community.playstation.net/userProfile/v1/users/%s/profile2?fields=accountId,onlineId,currentOnlineId", onlineId)

	var response UserAccountResponse
	if err := c.requestAndUnmarshal(ctx, url, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetUserProfile retrieves the user profile details for the specified account ID.
//
// Parameters:
//
//	ctx (context.Context): The context for controlling the request lifetime.
//	accountId (string): The account ID of the user whose profile details are being retrieved.
//
// Returns:
//
//	*UserProfileResponse: A pointer to the UserProfileResponse containing the user's profile details.
//	error: An error indicating whether the request was successful or not.
func (c *ClientAPI) GetUserProfile(ctx context.Context, accountId string) (*UserProfileResponse, error) {
	url := fmt.Sprintf("https://m.np.playstation.com/api/userProfile/v1/internal/users/%s/profiles", accountId)

	var response UserProfileResponse
	if err := c.requestAndUnmarshal(ctx, url, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetUserGames retrieves the list of games for the specified account ID.
//
// Parameters:
//
//	ctx (context.Context): The context for controlling the request lifetime.
//	accountId (string): The account ID of the user whose game list is being retrieved.
//
// Returns:
//
//	*UserGamesResponse: A pointer to the UserGamesResponse containing the user's game list.
//	error: An error indicating whether the request was successful or not.
func (c *ClientAPI) GetUserGames(ctx context.Context, accountId string) (*UserGamesResponse, error) {
	url := fmt.Sprintf("https://m.np.playstation.com/api/gamelist/v2/users/%s/titles?limit=10", accountId)

	var response UserGamesResponse
	if err := c.requestAndUnmarshal(ctx, url, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
