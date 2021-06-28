package amh_go_utils

import "math/rand"

var (
	lowercase = make([]rune, 0)
	uppercase = make([]rune, 0)
	numbers   = make([]rune, 0)

	allChars = make([]rune, 0)
)

func init() {
	for i := 'a'; i < 'a'+26; i++ {
		lowercase = append(lowercase, i)
		allChars = append(allChars, i)
	}
	for i := 'A'; i < 'A'+26; i++ {
		uppercase = append(uppercase, i)
		allChars = append(allChars, i)
	}

	for i := '0'; i < '0'+10; i++ {
		numbers = append(numbers, i)
		allChars = append(allChars, i)
	}

}

//CreateRandomString creates a random string of alphanumeric characters of the specified length
func CreateRandomString(length int) string {
	stringToReturn := make([]rune, length)

	for i := 0; i < length; i++ {
		var randomToSelect = rand.Intn(len(allChars))
		stringToReturn[i] = allChars[randomToSelect]
	}

	return string(stringToReturn)
}
