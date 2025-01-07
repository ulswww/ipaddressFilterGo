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
	"os/exec"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	// optionally, resize scanner's capacity for lines over 64K, see next example
	// fmt.Println(ipRange)
	// fmt.Println(s)
	addresses := findIpRange("./IpRange.txt")

	writeToIpTxt("./Ip.txt", addresses)

	runSpeedProgran()

}

func runSpeedProgran() {

	app := "./CloudflareST.exe"
	arg0 := "-tll"
	arg1 := "50"
	arg2 := "-tl"
	arg3 := "200"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func writeToIpTxt(filePath string, addresses []string) {

	if fileExists(filePath) {
		os.Remove(filePath)
	}

	ipfile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)

	check(err)

	defer ipfile.Close()

	for _, v := range addresses {
		ipfile.WriteString(v + "\n")
	}
}

func findIpRange(filePath string) []string {
	ipfile, err := os.Open(filePath)

	check(err)

	defer ipfile.Close()

	scanner := bufio.NewScanner(ipfile)

	addresses := []string{}

	for scanner.Scan() {
		ipRange := scanner.Text()

		spindex := strings.Index(ipRange, "/")
		s := fmt.Sprintf("%s.1", ipRange[0:spindex-2])

		address, err := getAddressInfo(s)
		if err != nil || address.isUSA() {
			continue
		}

		fmt.Printf("%s is in %s \n", ipRange, address.CountryName)
		addresses = append(addresses, ipRange)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return addresses
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
