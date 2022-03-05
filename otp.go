package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/xlzd/gotp"
)

var (
	secret string
)

func init() {
	flag.StringVar(&secret, "secret", "default secret", "otp secret")
}

func main() {
	flag.Parse()
	defaultTOTPUsage()
}

func defaultTOTPUsage() {
	tt := int(time.Now().Unix() / 60)
	hotp := gotp.NewDefaultHOTP(secret)
	fmt.Print(hotp.At(tt))
}

func googleTOTPUsage() {
	tt := int(time.Now().Unix() / 60)
	// secret := "KAVPEZYTZ7VR77DUI"
	hotp := gotp.NewDefaultHOTP(secret)
	fmt.Print(hotp.At(tt))
}
