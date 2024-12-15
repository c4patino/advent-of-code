package generate

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

const (
	baseURL       = "https://adventofcode.com/2024/day/%d/input"
	host          = "adventofcode.com"
	filePerm      = 0644
	userAgent     = "github.com/C4theBomb/advent-of-code by c4patino@gmail.com"
	cookieEnvName = "COOKIE"
)

func getCookie() (string, error) {
	cookieValue := os.Getenv(cookieEnvName)
	if cookieValue == "" {
		return "", fmt.Errorf("%s environment variable not set", cookieEnvName)
	}

	return cookieValue, nil

}

func createHTTPClient(cookieValue string) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create cookie jar: %w", err)
	}

	domainURL := &url.URL{Scheme: "https", Host: host}

	client := &http.Client{Jar: jar}
	jar.SetCookies(domainURL, []*http.Cookie{{Name: "session", Value: cookieValue}})

	return client, nil
}

func fetchInput(client *http.Client, day int) ([]byte, error) {
	url := fmt.Sprintf(baseURL, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func retrieveInput(day int) error {
	cookieValue, err := getCookie()
	if err != nil {
		return err
	}

	client, err := createHTTPClient(cookieValue)
	if err != nil {
		return err
	}

	body, err := fetchInput(client, day)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("./day%02d/input.txt", day)
	if err := os.WriteFile(filename, []byte(body), 0755); err != nil {
		return err
	}

	return nil
}
