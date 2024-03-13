package search

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"regexp"
	"strings"
	"unicode"
)

type KagomeAnalyzer struct{}

var pattern = `[\s　]+`

func AnalyzerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Analyzer, error) {
	return &KagomeAnalyzer{}, nil
}

func (a *KagomeAnalyzer) Analyze(input []byte) analysis.TokenStream {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}

	regex := regexp.MustCompile(pattern)

	str := string(input)
	str = strings.ToLower(str)
	str = removeLowercase2(str)
	str = RemoveNumbers(str)
	str = RemoveFullWidthCharacters(str)
	// Replace all whitespace characters with an empty string
	str = regex.ReplaceAllString(str, "")
	str = removePunctuation(str)
	//sWithoutSpaces := strings.ReplaceAll(string(input), " ", "")
	segments := t.Wakati(str)
	segments = removeSingleHiragana(segments)
	//for i, segment := range segments {
	//	if segment == "" {
	//		fmt.Println(segments)
	//	}
	//	if segment == " " {
	//		segments = append(segments[:i], segments[i+1:]...)
	//		fmt.Println(segments)
	//
	//	}
	//	if segment == "  " {
	//		fmt.Println(segments)
	//
	//	}
	//}
	tokens := make(analysis.TokenStream, len(segments))
	for i, segment := range segments {
		tokens[i] = &analysis.Token{
			Term:     []byte(segment),
			Position: i + 1,
			Start:    0,
			End:      len(segment),
		}
	}
	return tokens
}

func removePunctuation(input string) string {
	var result strings.Builder
	for _, char := range input {
		if !unicode.IsPunct(char) && !unicode.IsSymbol(char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func removeLowercase(input string) string {
	// Define a regular expression pattern to match lowercase letters
	pattern := "[a-z]"

	// Compile the regular expression pattern
	regex := regexp.MustCompile(pattern)

	// Replace all lowercase letters with an empty string
	cleanedString := regex.ReplaceAllString(input, "")

	return cleanedString
}

func removeLowercase2(input string) string {
	var cleanedString string
	for _, char := range input {
		if char < 'a' || char > 'z' {
			cleanedString += string(char)
		}
	}
	return cleanedString
}

func RemoveNumbers(input string) string {
	var cleanedString string
	for _, char := range input {
		if char < '0' || char > '9' {
			cleanedString += string(char)
		}
	}
	return cleanedString
}

// Function to remove single-letter strings containing only Hiragana characters
func removeSingleHiragana(slice []string) []string {
	var cleanedSlice []string

	for _, str := range slice {
		// Check if the string is a single character and contains only Hiragana characters
		if len([]rune(str)) > 1 || !isHiragana(str) {
			cleanedSlice = append(cleanedSlice, str)
		}
	}

	return cleanedSlice
}

// Function to check if a string contains only Hiragana characters
func isHiragana(str string) bool {
	for _, char := range str {
		if !unicode.Is(unicode.Hiragana, char) {
			return false
		}
	}
	return true
}

func RemoveFullWidthCharacters(input string) string {
	var cleanedString string
	for _, char := range input {
		// Exclude characters in the range of full-width digits and full-width English characters
		if (char < '０' || char > '９') && (char < 'ａ' || char > 'ｚ') {
			cleanedString += string(char)
		}
	}
	return cleanedString
}
