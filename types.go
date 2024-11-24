package playstation

import (
	"net/http"
	"time"
)

// Options is a type alias for a function that configures a Client.
// It takes a pointer to a Client as its parameter and modifies it.
type Options func(c *Client)

// Client represents a client for interacting with the PlayStation API.
// It holds the configuration for the HTTP client, language and region.
//
// Fields:
//
//	httpClient (*http.Client): The HTTP client used for making requests.
//	lang (Language): The language used for the client.
//	region (Region): The region used for the client.
type Client struct {
	httpClient *http.Client
	lang       Language
	region     Region
}

// Tokens represents the authentication tokens used for accessing the PlayStation API.
// It includes both access and refresh tokens along with their expiration times.
//
// Fields:
//
//	AccessToken (string): The token used for accessing the API.
//	RefreshToken (string): The token used for refreshing the access token.
//	AccessExpires (int64): The expiration time of the access token in seconds.
//	RefreshExpires (int64): The expiration time of the refresh token in seconds.
type Tokens struct {
	AccessToken        string `json:"access_token"`
	RefreshToken       string `json:"refresh_token"`
	AccessExpires      int64  `json:"expires_in"`
	RefreshExpires     int64  `json:"refresh_token_expires_in"`
	AccessExpiresTime  time.Time
	RefreshExpiresTime time.Time
}

// ClientAPI represents a client for interacting with the PlayStation API that includes authentication tokens and NPSSO.
//
// Fields:
//
//	Client (*Client): The embedded client for interacting with the PlayStation API.
//	Tokens (*Tokens): The authentication tokens used for accessing the API.
//	NPSSO (string): The NPSSO token used for authentication.
type ClientAPI struct {
	Client *Client
	Tokens *Tokens
	NPSSO  string
}

type UserAccountResponse struct {
	Profile struct {
		OnlineID        string `json:"onlineId"`
		AccountID       string `json:"accountId"`
		CurrentOnlineID string `json:"currentOnlineId"`
	} `json:"profile"`
}

type UserProfileResponse struct {
	OnlineID       string `json:"onlineId"`
	PersonalDetail struct {
		FirstName       string `json:"firstName"`
		LastName        string `json:"lastName"`
		DisplayName     string `json:"displayName"`
		ProfilePictures []struct {
			Size string `json:"size"`
			URL  string `json:"url"`
		} `json:"profilePictures"`
	} `json:"personalDetail"`
	AboutMe string `json:"aboutMe"`
	Avatars []struct {
		Size string `json:"size"`
		URL  string `json:"url"`
	} `json:"avatars"`
	Languages            []string `json:"languages"`
	IsPlus               bool     `json:"isPlus"`
	IsOfficiallyVerified bool     `json:"isOfficiallyVerified"`
	IsMe                 bool     `json:"isMe"`
}

type UserGamesResponse struct {
	Titles []struct {
		TitleID           string `json:"titleId"`
		Name              string `json:"name"`
		LocalizedName     string `json:"localizedName"`
		ImageURL          string `json:"imageUrl"`
		LocalizedImageURL string `json:"localizedImageUrl"`
		Category          string `json:"category"`
		Service           string `json:"service"`
		PlayCount         int    `json:"playCount"`
		Concept           struct {
			ID       int      `json:"id"`
			TitleIds []string `json:"titleIds"`
			Name     string   `json:"name"`
			Media    struct {
				Audios []interface{} `json:"audios"`
				Videos []interface{} `json:"videos"`
				Images []struct {
					URL    string `json:"url"`
					Format string `json:"format"`
					Type   string `json:"type"`
				} `json:"images"`
			} `json:"media"`
			Genres        []string `json:"genres"`
			LocalizedName struct {
				DefaultLanguage string `json:"defaultLanguage"`
				Metadata        struct {
					FiFI   string `json:"fi-FI"`
					UkUA   string `json:"uk-UA"`
					DeDE   string `json:"de-DE"`
					EnUS   string `json:"en-US"`
					KoKR   string `json:"ko-KR"`
					PtBR   string `json:"pt-BR"`
					EsES   string `json:"es-ES"`
					ArAE   string `json:"ar-AE"`
					NoNO   string `json:"no-NO"`
					FrCA   string `json:"fr-CA"`
					ItIT   string `json:"it-IT"`
					PlPL   string `json:"pl-PL"`
					RuRU   string `json:"ru-RU"`
					ZhHans string `json:"zh-Hans"`
					NlNL   string `json:"nl-NL"`
					PtPT   string `json:"pt-PT"`
					ZhHant string `json:"zh-Hant"`
					SvSE   string `json:"sv-SE"`
					DaDK   string `json:"da-DK"`
					TrTR   string `json:"tr-TR"`
					FrFR   string `json:"fr-FR"`
					EnGB   string `json:"en-GB"`
					Es419  string `json:"es-419"`
					JaJP   string `json:"ja-JP"`
				} `json:"metadata"`
			} `json:"localizedName"`
			Country  string `json:"country"`
			Language string `json:"language"`
		} `json:"concept"`
		Media struct {
			Audios []interface{} `json:"audios"`
			Videos []interface{} `json:"videos"`
			Images []struct {
				URL    string `json:"url"`
				Format string `json:"format"`
				Type   string `json:"type"`
			} `json:"images"`
		} `json:"media"`
		FirstPlayedDateTime time.Time `json:"firstPlayedDateTime"`
		LastPlayedDateTime  time.Time `json:"lastPlayedDateTime"`
		PlayDuration        string    `json:"playDuration"`
	} `json:"titles"`
	NextOffset     int `json:"nextOffset"`
	PreviousOffset int `json:"previousOffset"`
	TotalItemCount int `json:"totalItemCount"`
}

type RequestError struct {
	Error struct {
		Reason      string `json:"reason"`
		Source      string `json:"source"`
		Code        int    `json:"code"`
		Message     string `json:"message"`
		ReferenceID string `json:"referenceId"`
	} `json:"error"`
}
