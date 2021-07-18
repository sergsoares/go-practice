package main

import (
	"fmt"

	"github.com/alexellis/hmac"
)

func main() {
	input := []byte(`input message from api`)
	secret := []byte(`so secret`)

	digest := hmac.Sign(input, secret)
	fmt.Printf("Digest: %x\n ", digest)
}
