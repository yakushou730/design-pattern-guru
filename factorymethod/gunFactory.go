package main

import "fmt"

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
