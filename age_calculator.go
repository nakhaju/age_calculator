package age_calculator

import (
	"time"
)

func calculateAge(birthDate time.Time) int {
	today := time.Now()
	age := today.Year() - birthDate.Year()

	// Check if birthday hasn't happened yet this year
	if today.Month() < birthDate.Month() ||
		(today.Month() == birthDate.Month() && today.Day() < birthDate.Day()) {
		age--
	}

	return age
}
