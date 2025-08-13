package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// slugRegex sama dengan aturan di util/slug.go
var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

// New membuat instance validator & mendaftarkan tag custom "slug".
func New() *validator.Validate {
	v := validator.New()

	// Tag "slug" -> memvalidasi format slug
	_ = v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		s := fl.Field().String()
		if len(s) == 0 || len(s) > 200 {
			return false
		}
		return slugRegex.MatchString(s)
	})

	return v
}
