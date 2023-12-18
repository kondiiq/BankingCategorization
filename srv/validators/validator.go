package validators

import (
	"errors"
	"log"
	"regexp"
)

func matchDebitCardProvider(dCardNumber string) (string, error) {
	if !RgxDCard(dCardNumber) {
		err := errors.New("Cannot match Debit Card provider")
		log.SetFlags(1)
		log.Fatal(err)
		return "", err
	}
	if dCardNumber[0] == '3' && dCardNumber[1] == '4' || dCardNumber[1] == '7' {
		return "American Express", nil
	}
	if dCardNumber[0] == '4' {
		return "Visa", nil
	}
	if dCardNumber[0] == '5' {
		return "MasterCard", nil
	}
	if dCardNumber[0] == '6' {
		return "Discover", nil
	}
	err := errors.New("Cannot match Debit Card provider")
	log.SetFlags(1)
	log.Fatal(err)
	return "", err
}

func RgxDCard(dCardNumber string) bool {
	matcher := regexp.MustCompile(`^(5[1-5][0-9]{14}|2(22[1-9][0-9]{12}|2[3-9][0-9]{13}|[3-6][0-9]{14}|7[0-1][0-9]{13}|720[0-9]{12}))$`)
	return matcher.MatchString(dCardNumber)
}
