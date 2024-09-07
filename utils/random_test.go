package utils

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRandomInt(t *testing.T) {
	min := int64(1)
	max := int64(10)

	// Test that RandomInt(min, max) returns a value within the range [min, max]
	for i := 0; i < 100; i++ {
		result := RandomInt(min, max)
		assert.True(t, result >= min && result <= max, "RandomInt(%d, %d) = %d; want between %d and %d", min, max, result, min, max)
	}
}

func TestRandomString(t *testing.T) {
	length := 10
	randomStr := RandomString(length)

	// Test that RandomString(n) returns a string of length n
	assert.Equal(t, length, len(randomStr), "RandomString(%d) = %s; want length %d", length, randomStr, length)

	// Test that RandomString(n) contains only allowed characters
	for _, c := range randomStr {
		assert.True(t, strings.Contains(alphabet, string(c)), "RandomString(%d) contains disallowed character '%c'", length, c)
	}
}

func TestRandomOwner(t *testing.T) {
	owner := RandomUser()

	// Test that RandomOwner() returns a string of length 16
	assert.Equal(t, 16, len(owner), "RandomOwner() = %s; want length 16", owner)

	// Test that RandomOwner() contains only allowed characters
	for _, c := range owner {
		assert.True(t, strings.Contains(alphabet, string(c)), "RandomOwner() contains disallowed character '%c'", c)
	}
}

func TestRandomMoney(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := RandomMoney()
		assert.True(t, result >= 0 && result <= 1000, "RandomMoney() = %d; want between 0 and 1000", result)
	}
}

func TestRandomCurrency(t *testing.T) {
	currencies := []string{"EUR", "GBP", "JPY", "NZD", "USA", "AUD", "VND"}
	for i := 0; i < 100; i++ {
		result := RandomCurrency()
		assert.Contains(t, currencies, result, "RandomCurrency() = %s; want one of %v", result, currencies)
	}
}
