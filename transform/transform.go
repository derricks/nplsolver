// a variety of refactored tools for transforming strings into other, more-optimized-for-matching strings
package transform

import (
  "strings"
  "sort"
)

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

func splitStringIntoCharacters(word string) []string {
  return strings.Split(word,"")
}

func joinCharsIntoString(chars []string) string {
  return strings.Join(chars,"")
}