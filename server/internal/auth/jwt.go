package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5" // JWT v5.3.0
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// Claims menyimpan informasi user dalam token.
// - RegisteredClaims: berisi exp, iat, iss, sub, dll.
// - Roles: daftar role user (Reader/Admin/Author/SuperAdmin).
type Claims struct {
	UserID uuid.UUID `json:"uid"`
	Roles  []string  `json:"roles"`
	jwt.RegisteredClaims
}

// TokenPair berisi access & refresh token untuk klien.
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// issuer sederhana untuk identitas penerbit token.
const issuer = "blogapp"

// signAndBuild membangun token dengan secret HMAC (HS256).
func signAndBuild(secret []byte, claims Claims) (string, error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tok.SignedString(secret)
}

// GenerateTokens membuat access & refresh token untuk user tertentu.
// - access: TTL pendek (menit), dipakai ke API.
// - refresh: TTL panjang (jam/hari), untuk minta access baru.
func GenerateTokens(v *viper.Viper, userID uuid.UUID, roles []string) (TokenPair, error) {
	secret := []byte(v.GetString("security.jwt_secret"))

	// Access token (mis: 15 menit)
	accessTTL := time.Duration(v.GetInt("security.access_ttl_minutes")) * time.Minute
	accessClaims := Claims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   userID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	access, err := signAndBuild(secret, accessClaims)
	if err != nil {
		return TokenPair{}, err
	}

	// Refresh token (mis: 30 hari), beri "aud" khusus agar mudah dibedakan
	refreshTTL := time.Duration(v.GetInt("security.refresh_ttl_hours")) * time.Hour
	refreshClaims := accessClaims
	refreshClaims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(refreshTTL))
	refreshClaims.RegisteredClaims.Audience = []string{"refresh"} // penanda tipe refresh

	refresh, err := signAndBuild(secret, refreshClaims)
	if err != nil {
		return TokenPair{}, err
	}
	return TokenPair{AccessToken: access, RefreshToken: refresh}, nil
}

// ParseAndValidate memverifikasi token & mengembalikan Claims bila valid.
// - expectRefresh: true jika token yang diharapkan adalah refresh token (aud=refresh).
func ParseAndValidate(v *viper.Viper, tokenStr string, expectRefresh bool) (*Claims, error) {
	secret := []byte(v.GetString("security.jwt_secret"))

	// Parse token & validasi signature + registered claims.
	tok, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
	if err != nil {
		return nil, err
	}
	claims, ok := tok.Claims.(*Claims)
	if !ok || !tok.Valid {
		return nil, errors.New("invalid token")
	}
	// Jika refresh diharapkan, pastikan audience mengandung "refresh".
	if expectRefresh {
		if !claims.RegisteredClaims.Audience.Contains("refresh") {
			return nil, errors.New("not a refresh token")
		}
	}
	return claims, nil
}
