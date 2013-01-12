package transform_test

import (
  "nplsolver/transform"
  "testing"
)

const (
   testWord = "hello"
   testWordSorted = "ehllo"
)

func TestSortingCharacters(t *testing.T) {
   sortedWord := transform.SortAllCharacters(testWord)
   if sortedWord != testWordSorted {
      t.Errorf("Sorted word %v does not equal expectation %v",sortedWord,testWordSorted)
   }
}

