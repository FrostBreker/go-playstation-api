package playstation

import (
	"fmt"
	"net/http"
)

// defaultConfig initializes a Client with default configuration.
// It sets the default language, region, and HTTP client.
//
// Returns:
//
//	(Client): A Client struct with default settings.
func defaultConfig() Client {
	return Client{
		lang:       Languages[0],
		region:     Regions[0],
		httpClient: http.DefaultClient,
	}
}

// NewClient creates a new Client with the provided options.
// It initializes the Client with default configuration and applies each option function to the Client.
//
// Parameters:
//
//	opts (...Options): A variadic list of option functions to customize the Client.
//
// Returns:
//
//	(*Client): A pointer to the newly created Client with the applied options.
func NewClient(opts ...Options) *Client {
	c := defaultConfig()
	for _, opt := range opts {
		opt(&c)
	}

	httpClient := *c.httpClient // Create a copy of the existing client
	httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return &Client{
		httpClient: &httpClient,
		lang:       c.lang,
		region:     c.region,
	}
}

// WithLanguage sets a custom language for the Client.
// It returns an Options function that sets the lang field of the Client struct.
// If the provided language is not supported, it returns an error.
//
// Parameters:
//
//	lang (playstation.Language): The custom language to be used.
//
// Returns:
//
//	(Options, error): A function that sets the lang field of the Client struct, or an error if the language is unsupported.
func WithLanguage(lang Language) (Options, error) {
	if !IsContain(Languages, lang) {
		return nil, fmt.Errorf("unsupported lang %s", lang)
	}
	return func(c *Client) {
		c.lang = lang
	}, nil
}

// WithRegion sets a custom region for the Client.
// It returns an Options function that sets the region field of the Client struct.
// If the provided region is not supported, it returns an error.
//
// Parameters:
//
//	region (playstation.Region): The custom region to be used.
//
// Returns:
//
//	(Options, error): A function that sets the region field of the Client struct, or an error if the region is unsupported.
func WithRegion(region Region) (Options, error) {
	if !IsContain(Regions, region) {
		return nil, fmt.Errorf("unsupported region %s", region)
	}
	return func(c *Client) {
		c.region = region
	}, nil
}

// WithClient sets a custom HTTP client for the Client.
// It returns an Options function that sets the httpClient field of the Client struct.
// If the provided client is nil, it returns an error.
//
// Parameters:
//
//	client (*http.Client): The custom HTTP client to be used.
//
// Returns:
//
//	(Options, error): A function that sets the httpClient field of the Client struct, or an error if the client is nil.
func WithClient(client *http.Client) (Options, error) {
	if client == nil {
		return nil, fmt.Errorf("cannot use nil httpClient")
	}
	return func(c *Client) {
		c.httpClient = client
	}, nil
}
