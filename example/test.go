package main

import (
	"fmt"
	"runtime"

	"github.com/ip2location/ip2location-go"
	"github.com/zut/x/xx"
)

func main() {
	dbPath := fmt.Sprintf("%v/%v", xx.IfStr(runtime.GOOS == "darwin", "/Users/d/z", "/d/z"), "IP2LOCATION-LITE-DB3.BIN")
	db, err := ip2location.OpenDB(dbPath)
	if err != nil {
		fmt.Print(err)
		return
	}
	ip := "113.104.251.44"
	results, err := db.Get_all(ip)

	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("country_short: %s\n", results.Country_short)
	fmt.Printf("country_long: %s\n", results.Country_long)
	fmt.Printf("region: %s\n", results.Region)
	fmt.Printf("city: %s\n", results.City)
	db.Close()
}
