package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword meng-hash password plaintext dengan bcrypt.
// bcrypt.DefaultCost cukup untuk sebagian besar kasus dev.
func HashPassword(plain string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
    return string(hash), err
}

// CheckPassword membandingkan hash tersimpan dengan password input.
// Nil error artinya cocok; error berarti mismatch/invalid.
func CheckPassword(hash string, plain string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
