package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type sign struct {
	name string
	sign string
}

func formatSign(s sign) string {
	return fmt.Sprint(s.name, " | ", s.sign)
}

type input struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func parseToSign(s string) sign {
	var newSign input
	err := json.Unmarshal([]byte(s), &newSign)
	if err != nil {
		log.Fatal(err)
	}
	return sign{
		name: newSign.Name,
		sign: matchSignByDate(newSign.Date),
	}
}

func matchSignByDate(d string) string {
	t, _ := time.Parse("2006-01-02", d)

	switch true {
	case int(t.Month()) >= int(time.Month(4)) && t.Day() >= 21 &&
		int(t.Month()) <= int(time.Month(5)) && t.Day() <= 21:
		return "Taurus"
	case int(t.Month()) >= int(time.Month(5)):
		return "Gemini"
	default:
		return "Aries"
	}
}
