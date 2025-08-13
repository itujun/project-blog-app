package util

import (
	"regexp"
	"strings"
	"unicode"
)

// Regex slug: huruf/angka lowercase dipisah dash, tanpa dash ganda/awal/akhir.
var slugRegexp = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

// NormalizeSlug mengubah judul/slug bebas menjadi slug-lowercase yang rapi.
// - Lowercase
// - Non-alfanumerik jadi '-'
// - Kompres dash ganda & trim dash pinggir
func NormalizeSlug(s string) string {
	// to lower & ganti spasi/karakter non-alfanumerik jadi '-'
	var b strings.Builder
	prevDash := false
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			prevDash = false
			continue
		}
		// selain alnum: jadikan '-'
		if !prevDash {
			b.WriteByte('-')
			prevDash = true
		}
	}
	out := b.String()
	out = strings.Trim(out, "-")         // hapus dash awal/akhir
	out = strings.ReplaceAll(out, "--", "-") // safety (kalau ada)
	return out
}

// ValidateSlug mengecek format & panjang maksimal slug.
func ValidateSlug(slug string, maxLen int) bool {
	if len(slug) == 0 || len(slug) > maxLen {
		return false
	}
	return slugRegexp.MatchString(slug)
}
