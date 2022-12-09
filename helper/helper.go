package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func DownloadInput() {
	fInfo, _ := os.Stat("input.txt")
	if fInfo != nil {
		return
	}

	p, _ := os.Getwd()
	d := filepath.Base(p)
	dNum, _ := strconv.Atoi(strings.TrimPrefix(d, "day"))
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", dNum), nil)
	req.Header.Set("Cookie", getSessionCookie())
	req.Header.Set("User-Agent", "neilo40's golang client")
	c := &http.Client{Timeout: 60 * time.Second}
	r, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	ioutil.WriteFile("input.txt", bodyBytes, 0644)
}

func getSessionCookie() string {
	session, err := ioutil.ReadFile("../helper/session")
	if err != nil {
		log.Fatal("Error opening session cookie")
	}
	return string(session)
}
