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
  
  // get a basic solver and run the following tests with the dictionary
  solver,_ := GetSolver(Basic)
  _,resultsMap := solveForBasicSearchPattern("stor*",solver,t)  
  expectedResults := []string{"stork","store","story"}
  if !verifyWordsInMap(expectedResults,resultsMap) {
     t.Errorf("Not all words from %v were in result map",expectedResults)
  }

  _,resultsMap = solveForBasicSearchPattern("stor?",solver,t)
  expectedResults = []string{"stork","store","story"}
  if !verifyWordsInMap(expectedResults,resultsMap) {
     t.Errorf("Not all words from %v were in result map",expectedResults)
  }
  
  _,resultsMap = solveForBasicSearchPattern("st*re",solver,t)
  expectedResults = []string{"stare","store"}  
  if !verifyWordsInMap(expectedResults,resultsMap) {
    t.Errorf("Not all words from %v were in results map", expectedResults)
  }
 
}

// refactored code for getting results from a basic solver with a particular pattern
func solveForBasicSearchPattern(pattern string, solver Solver,t *testing.T) (results []string, resultsMap map[string]bool) {
    results = make([]string,10)
    resultsMap = make(map[string]bool)
    resultsChan := make(chan string)
    doneChan := make(chan interface{})
   
    doneReceived := false
    dictionary,_ := getDictionaryForBasicSolver()
    go solver.Solve(pattern,dictionary,resultsChan,doneChan)

    var curResult string
    for doneReceived != true {
      select {
         case curResult = <-resultsChan:
            results = append(results,curResult)
            resultsMap[curResult] = true
         case item := <- doneChan:
            t.Logf("TODO: test error handling for item: %v\n",item)
            doneReceived = true
      }
    }
    return results,resultsMap
}



// look up each item in the slice and see if it's in the map. return false if any fails
func verifyWordsInMap(words []string, wordsMap map[string]bool) bool {
  if len(words) != len(wordsMap) {
    return false
  }
  
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
