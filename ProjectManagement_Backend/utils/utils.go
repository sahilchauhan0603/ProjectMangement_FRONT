package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
)

// JWKS represents a JSON Web Key Set
type JWKS struct {
	Keys []JWK `json:"keys"`
}

// JWK represents a JSON Web Key
type JWK struct {
	Kty   string   `json:"kty"`
	Use   string   `json:"use"`
	KeyId string   `json:"kid"`
	X5t   string   `json:"x5t"`
	N     string   `json:"n"`
	E     string   `json:"e"`
	X5c   []string `json:"x5c"`
}

// FetchJWKS fetches the JSON Web Key Set (JWKS) from the given URL
//
// This function makes an HTTP GET request to the specified URL to retrieve the JWKS.
//
// Parameters:
//   - url: The URL from which to fetch the JWKS.
//
// Returns:
//   - A pointer to a JWKS object containing the keys.
//   - An error if there is an issue with the HTTP request or JSON decoding.
func FetchJWKS(url string) (*JWKS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jwks JWKS
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return nil, err
	}

	return &jwks, nil
}

// FindKey finds the RSA public key in the JWKS with the given key ID (kid)
//
// This function iterates through the keys in the JWKS and returns the RSA public key
// corresponding to the given key ID.
//
// Parameters:
//   - kid: The key ID to search for in the JWKS.
//
// Returns:
//   - A pointer to an rsa.PublicKey object if a matching key is found.
//   - An error if the key is not found or there is an issue converting the key.
func (jwks *JWKS) FindKey(kid string) (*rsa.PublicKey, error) {
	for _, jwk := range jwks.Keys {
		if jwk.KeyId == kid {
			return jwk.rsaPublicKey()
		}
	}
	return nil, fmt.Errorf("unable to find appropriate key")
}

// rsaPublicKey converts the JWK to an RSA public key
//
// This function decodes the modulus (N) and exponent (E) from the JWK and constructs
// an rsa.PublicKey object.
//
// Returns:
//   - A pointer to an rsa.PublicKey object.
//   - An error if there is an issue decoding the modulus or exponent.
func (jwk *JWK) rsaPublicKey() (*rsa.PublicKey, error) {
	nBytes, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return nil, err
	}
	eBytes, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		return nil, err
	}

	e := 0
	for _, b := range eBytes {
		e = e*256 + int(b)
	}

	pubKey := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: e,
	}

	return pubKey, nil
}
