package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("")
}

type ILicense interface {
	UseKey() (resp string, err error)
}

type Product struct {
	id              string
	name            string
	validLicenseKey string
	currentCapacity int
	maxCapacity     int
}

func (p *Product) UseKey() (string, error) {
	return "License key used.", nil
}

type LicenseKey struct {
	key string
}

type ProductProxy struct {
	product Product
	key     *LicenseKey
}

func NewKeyProxy(k *LicenseKey, p *Product) *ProductProxy {
	return &ProductProxy{*p, k}
}

func (p *ProductProxy) UseKey() (string, error) {
	if p.product.validLicenseKey != p.key.key {
		return "Error", errors.New("This key does not work for this product.")

	} else if p.product.currentCapacity+1 > p.product.currentCapacity {
		return "Error", errors.New("The capacity is full for this product.")
	} else {
		return "OK", nil
	}
}
