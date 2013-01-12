package solver_test

import (
  "nplsolver/dict"
  "nplsolver/solver"
  "testing"
)

const (
   testWord = "hello"
)

func TestIdentityMatching(t *testing.T) {
  entry,_ := dict.NewEntryFromWord(testWord)
  
  matcher,_ := solver.NewIdentityMatcher()
  
  if !matcher.Match(testWord,entry) {
     t.Errorf("IdentityMatcher did not show a match for test word.")
  }
}

func TestSameCharactersMatching(t *testing.T) {
  entry,_ := dict.NewEntryFromWord(testWord)
  matcher,_ := solver.NewSameCharacterMatcher()
  if !matcher.Match("loleh",entry) {
     t.Errorf("%v does not have same letters as loleh")
  }
}

func TestRegexMatching(t *testing.T) {
  regex := "he..o"
  entry,_ := dict.NewEntryFromWord(testWord)
  matcher,_ := solver.NewRegexMatcher(regex)
  if !matcher.Match(regex,entry) {
     t.Errorf("%v did not match regex %v",entry.Word(),regex)
  } 
}