package transform_test

import (
  "nplsolver/transform"
  "testing"
)

const (
   testWord = "hello"
   testWordSorted = "ehllo"
   testWordUniqueSorted = "ehlo"
   testWordPattern = "abccd"
)

func TestSortingCharacters(t *testing.T) {
   sortedWord := transform.SortAllCharacters(testWord)
   if sortedWord != testWordSorted {
      t.Errorf("Sorted word %v does not equal expectation %v",sortedWord,testWordSorted)
   }
}

func TestUniqueSortedCharacters(t *testing.T) {
   sortedWord := transform.UniqueSortedCharacters(testWord)
   if sortedWord != testWordUniqueSorted {
      t.Errorf("Unique sorted word %v does not equal expectation %v",sortedWord,testWordUniqueSorted)
   }
}

func TestToPattern(t *testing.T) {
   patternWord := transform.ToPattern(testWord)
   if patternWord != testWordPattern {
     t.Errorf("Pattern %v does not equal expectation %v", patternWord, testWordPattern)
   }
}

