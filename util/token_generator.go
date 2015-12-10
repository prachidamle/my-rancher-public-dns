package util

import(
    "crypto/rand"
    "encoding/base64"
)

// generateRandomBytes returns securely generated random bytes. 
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    // Note that err == nil only if we read len(b) bytes.
    if err != nil {
        return nil, err
    }

    return b, nil
}

// generateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomString(s int) (string, error) {
    b, err := generateRandomBytes(s)
    return base64.URLEncoding.EncodeToString(b), err
}

func GenerateNewToken() string {
	// Example: this will give us a 44 byte, base64 encoded output
	token, err := generateRandomString(32)
	if err != nil {
	    // Serve an appropriately vague error to the
	    // user, but log the details internally.
	}
	return token
}