// a variety of refactored tools for transforming strings into other, more-optimized-for-matching strings
package transform

import (
  "strings"
  "sort"
)

// the letters to use when constructing a pattern
var patternLetters = [...]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}

// given "hello" this will return "ehllo"
func SortAllCharacters(word string) string {
  chars := splitStringIntoCharacters(word)
  sort.Strings(chars)
  return joinCharsIntoString(chars)
}

// given "hello", this will return "ehlo"
func UniqueSortedCharacters(word string) string {
   letterMap := make(map[string]bool)
   
   for _,char := range word {
      letterMap[string(char)] = true
   }
   
   chars := make([]string,len(word))
   for char := range letterMap {
      chars = append(chars,char)
   }
   sort.Strings(chars)
   return joinCharsIntoString(chars)   
}

// given hello, this will return abccd
func ToPattern(word string) string {
  // holds the letter -> letter mapping
  patternMap := make(map[string]string)
 
  // the return value
  pattern := make([]string, len(word))

  // the current offset in patternLetters
  patternOffset := 0

  // for each character in word, find its entry in patternMap. If it's not there, add it using the next letter in patternLetters
  for _,char := range word {
     _, found := patternMap[string(char)]
     if !found {
       patternMap[string(char)] = patternLetters[patternOffset]
       patternOffset++
     }
     pattern = append(pattern, patternMap[string(char)])
  }

  return joinCharsIntoString(pattern)
}

func splitStringIntoCharacters(word string) []string {
  return strings.Split(word,"")
}

func joinCharsIntoString(chars []string) string {
  return strings.Join(chars,"")
}