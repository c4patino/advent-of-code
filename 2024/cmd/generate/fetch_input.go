package generate

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func retrieveInput(day int) error {
	cookieValue := os.Getenv("COOKIE")
	if cookieValue == "" {
		return errors.New("COOKIE environment variable not set")
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}

	client := &http.Client{Jar: jar}

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "github.com/C4theBomb/advent-of-code by c4patino@gmail.com")

	cookie := &http.Cookie{Name: "session", Value: cookieValue}
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Unexpected status code: %d\n%s", resp.StatusCode, string(body)))
	}

	filename := fmt.Sprintf("./day%02d/input.txt", day)
	if err := os.WriteFile(filename, []byte(body), 0755); err != nil {
		return err
	}

	return nil
}
