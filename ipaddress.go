package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// "strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	ipfile, err := os.Open("./ip.txt")

	check(err)

	defer ipfile.Close()

	scanner := bufio.NewScanner(ipfile)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		ipRange := scanner.Text()
		// fmt.Println(ipRange)
		spindex := strings.Index(ipRange, "/")
		s := fmt.Sprintf("%s.1", ipRange[0:spindex-2])
		// fmt.Println(s)
		address, err := getAddressInfo(s)
		if err != nil || address.isUSA() {
			continue
		}

		fmt.Printf("%s is in %s \n", ipRange, address.CountryName)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func getAddressInfo(ipAddress string) (*Address, error) {

	url := fmt.Sprintf("https://freeipapi.com/api/json/%s", ipAddress)

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	address := Address{}

	json.Unmarshal(body, &address)

	// fmt.Println(address.CountryName)

	// fmt.Printf("ip in %s %s", address.CountryName, strconv.FormatBool(address.isUSA()))

	return &address, nil
}

type Address struct {
	IpAddress   string `json:"ipAddress"`
	CountryName string `json:"countryName"`
}

func (address *Address) isUSA() bool {
	return strings.Contains(address.CountryName, "America")
}
