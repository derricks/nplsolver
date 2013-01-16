package solver

import (
  "nplsolver/dict"
  "testing"
)

const (
   testWord = "hello"
   resultsChannelSize = 100
)

func TestIdentityMatching(t *testing.T) {
  entry,_ := dict.NewEntryFromWord(testWord)
  
  matcher,_ := newIdentityMatcher()
  
  if !matcher.Match(testWord,entry) {
     t.Errorf("IdentityMatcher did not show a match for test word.")
  }
}

func TestSameCharactersMatching(t *testing.T) {
  entry,_ := dict.NewEntryFromWord(testWord)
  matcher,_ := newSameCharacterMatcher()
  if !matcher.Match("loleh",entry) {
     t.Errorf("%v does not have same letters as loleh")
  }
}

func TestRegexMatching(t *testing.T) {
  regex := "he..o"
  entry,_ := dict.NewEntryFromWord(testWord)
  matcher,_ := newRegexMatcher(regex)
  if !matcher.Match(regex,entry) {
     t.Errorf("%v did not match regex %v",entry.Word(),regex)
  } 
}

func TestBasicSolverBehavior(t *testing.T) {
  // create a dictionary of words we want to test against
  dict,_ := getDictionaryForBasicSolver()
  
  //create the channels for results and signalling done status
  resultsChan := make(chan string)
  doneChan := make(chan interface{})
  
  doneReceived := false
  results := make([]string,10)
  resultsMap := make(map[string]bool) // just to see if various words have been found
   
  // get a basic solver and run the following tests with the dictionary
  solver,_ := GetSolver(Basic)
  
  //  "stor*" should return "stork","store",and "story"
  go solver.Solve("stor*",dict,resultsChan,doneChan)
  
  var curResult string
  for doneReceived != true {
     select {
        case curResult = <-resultsChan:
          results = append(results,curResult)
          resultsMap[curResult] = true
        case item := <-doneChan:
          t.Logf("Set up error handling for item: %v\n",item)
          doneReceived = true
     }
  }
  
  // verify results  
  t.Logf("results: %v\n",results)
  expectedResults := []string{"stork","store","story"}
  if !verifyWordsInMap(expectedResults,resultsMap) {
     t.Error("Not all words were in result map")
  }


//      "stor?" should return the same thing
//      "st?r*" should return the same list plus "stare"
//      "*" should return every word in the dictionary
      

  
}

// look up each item in the slice and see if it's in the map. return false if any fails
func verifyWordsInMap(words []string, wordsMap map[string]bool) bool {
  for _,word := range words {
     if _,found := wordsMap[word]; !found {
        return false
     }
  }
  return true
}

func getDictionaryForBasicSolver() (dict.Dictionary,error) {
   return dict.NewDictionaryFromWords("stork","store","stare","story","whistle")
}
