package hack_browser_data

import (
	"fmt"
	"testing"
)

type CookieParam struct {
	Name         string `json:"name"`                   // Cookie name.
	Value        string `json:"value"`                  // Cookie value.
	URL          string `json:"url,omitempty"`          // The request-URI to associate with the setting of the cookie. This value can affect the default domain, path, source port, and source scheme values of the created cookie.
	Domain       string `json:"domain,omitempty"`       // Cookie domain.
	Path         string `json:"path,omitempty"`         // Cookie path.
	Secure       bool   `json:"secure,omitempty"`       // True if cookie is secure.
	HTTPOnly     bool   `json:"httpOnly,omitempty"`     // True if cookie is http-only.
	SameSite     string `json:"sameSite,omitempty"`     // Cookie SameSite type.
	Priority     string `json:"priority,omitempty"`     // Cookie Priority.
	SameParty    bool   `json:"sameParty,omitempty"`    // True if cookie is SameParty.
	SourceScheme string `json:"sourceScheme,omitempty"` // Cookie source scheme type.
	SourcePort   int64  `json:"sourcePort,omitempty"`   // Cookie source port. Valid values are {-1, [1, 65535]}, -1 indicates an unspecified port. An unspecified port value allows protocol clients to emulate legacy cookie scope for the port. This is a temporary ability and it will be removed in the future.
	PartitionKey string `json:"partitionKey,omitempty"` // Cookie partition key. The site of the top-level URL the browser was visiting at the start of the request to the endpoint that set the cookie. If not set, the cookie will be set as not partitioned.
}

// test run example
func TestExample(t *testing.T) {
	var cookiesItem []*CookieParam
	fetchChromeCookies, err := GetBrowserAllCookies("chrome")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, cookies := range *fetchChromeCookies {
		for _, cookie := range cookies {
			var cookiesNode CookieParam
			cookiesNode.Name = cookie.KeyName
			cookiesNode.Value = cookie.Value
			cookiesNode.Domain = cookie.Host
			cookiesNode.Path = cookie.Path
			cookiesNode.HTTPOnly = cookie.IsHTTPOnly
			cookiesNode.Secure = cookie.IsSecure
			cookiesNode.SameSite = "None"
			cookiesNode.Priority = "Medium"
			cookiesNode.SameParty = false
			cookiesNode.SourceScheme = "Secure"
			if cookie.IsSecure {
				cookiesNode.SourcePort = 443
			} else {
				cookiesNode.SourcePort = 80
			}
			cookiesItem = append(cookiesItem, &cookiesNode)
		}
	}
	// debug see `cookiesItem`
	fmt.Println(len(cookiesItem))
}
