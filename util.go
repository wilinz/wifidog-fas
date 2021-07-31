package main

import (
	"crypto/sha256"
	"fmt"
)

func Sha256Sum(text string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
}

func ConvertFlowUnit(count int64) string {
	floatCount := float64(count)
	if floatCount < 1024 {
		return fmt.Sprint(count, "B")
	} else if floatCount >= 1024 && floatCount < 1024*1024 {
		return fmt.Sprint(TwoDecimalPlaces(floatCount/1024), "kB")
	} else if floatCount >= 1024*1024 && floatCount < 1024*1024*1024 {
		return fmt.Sprint(TwoDecimalPlaces(floatCount/(1024*1024)), "MB")
	} else {
		return fmt.Sprint(TwoDecimalPlaces(floatCount/(1024*1024*1024)), "GB")
	}
}

func TwoDecimalPlaces(value float64) string {
	return fmt.Sprintf("%.2f", value)
}
