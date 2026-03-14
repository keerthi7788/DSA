package strings

import "fmt"


//count the cosecutive characters in a string if the achat appera in between and repared the smae again shoud be new.

//s:="AAABBCCAABB"
//output: A3B2C2A2B2   --make sure output formate string.

func main() {
	s := "AAABBCCAABB"
	value := CountCharInString(s)
	fmt.Println("Hello", value)
}
func CountCharInString(s string) string {
	if len(s) == 0 {
		return ""
	}
	count := 1
	result := ""
	for i := 1; i < len(s); i++ {

		if s[i] == s[i-1] {
			count++
		} else {

			result = result + fmt.Sprintf("%c%d", s[i-1], count)

			count = 1

		}
	}

	result = result + fmt.Sprintf("%c%d", s[len(s)-1], count)

	return result

}
