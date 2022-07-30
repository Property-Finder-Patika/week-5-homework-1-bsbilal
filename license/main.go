package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
}

type ILicense interface {
	UseKey(cmd string) (resp string, err error)
}

type Product struct {
	name            string
	validLicenseKey string
	currentCapacity int
	maxCapacity     int
}

func (p *Product) UseKey() {

}

type LicenseKey struct {
	key string
}

type ProductProxy struct {
	product Product
	key     LicenseKey
}

func (p *ProductProxy) UseKey() {

}
