package utils

import (
	"math/rand"
	"strings"
	"time"
)

// global seed for random generation of numbers
var seed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomChoice returns a random selection from a list of strings
func RandomChoice(choices []string) string {
	n := seed.Int() % len(choices)
	return choices[n]
}

// RandomBool returns a random boolean choice
func RandomBool() bool {
	return rand.Float32() < 0.5
}

func RandomInt(maxval int) int {
	return seed.Intn(maxval)
}

func RandomUint64() uint64 {
	return seed.Uint64()
}

// RandomBool returns a boolean choice at the user's threshold
func RandomBoolWeight(chanceTrue float32) bool {
	return rand.Float32() < chanceTrue
}

// characters to choose from
const charset = "abcdefghijklmnopqrstuvwxyz"
const capitals = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Get a random character
func RandomChar() string {
	b := make([]rune, 1)
	b[0] = letters[seed.Intn(len(letters))]
	return string(b)
}

func RandomIntRange(min int, max int) int {
	return seed.Intn(max-min) + min
}

// RandomFloatRange generates a random float in a rnage
func RandomFloatRange(min uint64, max uint64) uint64 {
	return randomFloatHelper(max-min) + min
}

const maxInt64 uint64 = 1<<63 - 1

func randomFloatHelper(n uint64) uint64 {
	if n < maxInt64 {
		return uint64(seed.Int63n(int64(n + 1)))
	}
	x := seed.Uint64()
	for x > n {
		x = seed.Uint64()
	}
	return x
}

// RandomFloat returns a random floating point
func RandomFloat() uint64 {
	return seed.Uint64()
}

// RandomName generates a random name for a function, variaable, etc.
func RandomName() string {
	length := seed.Intn(20-10) + 10
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seed.Intn(len(charset))]
	}
	return strings.Trim(string(result), " ")
}
