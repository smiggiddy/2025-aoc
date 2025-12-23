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

func GetInput(day, year int) {
	if day == 0 || year == 0 {
		fmt.Println("Please enter a valid day")
		os.Exit(1)
	}

	aocURL := fmt.Sprintf("%s/%d/day/%d/input", url, year, day)

	input := setupRequest(aocURL)

	var dayString string
	if day < 10 {

		dayString = "0" + strconv.Itoa(day)
	} else {
		dayString = strconv.Itoa(day)
	}
	filePath := fmt.Sprintf("day%s/input_test.txt", dayString)
	err := os.WriteFile(filePath, input, 06440)
	if err != nil {
		fmt.Println("Error saving input", err)
		os.Exit(1)
	}
}

func setupRequest(url string) []byte {
	cookieValue, err := os.ReadFile("./cookie.txt")
	if err != nil {
		fmt.Println("Error reading cookie file", err)
		os.Exit(1)
	}
	client := &http.Client{
		Transport: &http.Transport{},
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: string(cookieValue),
		Path:  "/",
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.AddCookie(cookie)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body

}
