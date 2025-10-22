package utils

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateEmailFormat(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func ValidateUserFields(name, email, password string, divisionID, departmentID uint) error {
	// Cek field kosong
	if strings.TrimSpace(name) == "" {
		return errors.New("nama wajib diisi")
	}
	if strings.TrimSpace(email) == "" {
		return errors.New("email wajib diisi")
	}
	if strings.TrimSpace(password) == "" {
		return errors.New("password wajib diisi")
	}
	if divisionID == 0 {
		return errors.New("division wajib dipilih")
	}
	if departmentID == 0 {
		return errors.New("department wajib dipilih")
	}

	// Validasi format email
	if !ValidateEmailFormat(email) {
		return errors.New("format email tidak valid")
	}

	// Password minimal 6 karakter
	if len(password) < 6 {
		return errors.New("password minimal 6 karakter")
	}

	return nil
}
