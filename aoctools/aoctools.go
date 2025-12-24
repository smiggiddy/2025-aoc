package aoctools

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	url string = "https://adventofcode.com"
)

func LoadFile(file string) string {
	data, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	str := strings.Trim(string(data), "\n")

	return str
}

func AllSame(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, ch := range s {
		if ch != rune(s[0]) {
			return false
		}
	}

	return true
}

func SetupDay(path string) {

	err := os.MkdirAll(path, 0750)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path+"/input.txt", []byte(""), 0640)
	err = os.WriteFile(path+"/sample.txt", []byte(""), 0640)
	err = os.WriteFile(path+"/solve.go", []byte(""), 0640)

}

func GetInput(day, year int) ([]byte, error) {
	if day == 0 || year == 0 {
		return nil, fmt.Errorf("Please enter a valid day\n")
	}

	aocURL := fmt.Sprintf("%s/%d/day/%d/input", url, year, day)

	cookie, err := loadCookieFile("./aoctools/cookie.txt")
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{},
	}

	req, err := setupRequest(aocURL, cookie)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return body, nil
}

func loadCookieFile(filePath string) (string, error) {
	cookieValue, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	cookie := strings.TrimSpace(string(cookieValue))
	return cookie, nil
}

func FormatDayString(day int) (dayString string) {
	if day < 10 {
		dayString = "0" + strconv.Itoa(day)
	} else {
		dayString = strconv.Itoa(day)
	}
	return
}

func setupRequest(url, cookieValue string) (*http.Request, error) {

	cookie := &http.Cookie{
		Name:  "session",
		Value: cookieValue,
		Path:  "/",
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error Creating Request%v", err)
	}

	req.AddCookie(cookie)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")

	return req, nil

}
