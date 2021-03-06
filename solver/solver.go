// The solver package encapsulates all the solvers and support code used by the solver server
package solver

import (
  "fmt"
  "regexp"
  "nplsolver/dict"
  "nplsolver/transform"
  "strings"
  "errors"
)

type SolverType string

const (
  Basic = "basic"
  Transposal = "transposal"
  LetterBank = "letterbank"
  TransAddition = "transaddition"
  Crypt = "crypt"
  Regex = "regex"
)

// get an appropriate solver object
func GetSolver(solverType SolverType) (Solver,error) {
   var solver Solver
   switch strings.ToLower(string(solverType)) {
     case Basic:
        solver = BasicSolver{}
     case Transposal:
        solver = TransposalSolver{}
     case LetterBank:
        solver = LetterBankSolver{}
     case TransAddition:
        solver = TransAdditionSolver{}
     case Crypt:
        solver = CryptSolver{}
   }

   if solver != nil {
     return solver,nil
   }

   return nil,errors.New(fmt.Sprintf("No solver found for type %v",solverType))
}


// The basic behavior for a Solver, a thing that works with a puzzle and a Dictionary to find matches.
// matches are written to the receiving channel. when finished, the solver writes to done_channel (either itself or an error)
type Solver interface {
   Solve(pattern string, dict dict.Dictionary, sols_channel chan<- string, done_channel chan<- interface{})
}

type BasicSolver struct {}
func (solver BasicSolver) Solve(pattern string, dictionary dict.Dictionary, results chan<- string, done_channel chan<- interface{}) {

   regex := convertBasicSearchWildcardsToRegex(pattern)
   if matcher,err := newRegexMatcher(regex); err != nil {
      done_channel <- err
   } else {
        dictionary.Iterate(func(entry dict.Entry) {
           if matcher.Match(regex,entry) {
              results <- entry.Word()
           }
        })
       done_channel <- solver
   }
}

// convert the pattern (which takes ? for single-character matches and * for multi-character matches) to a regex
func convertBasicSearchWildcardsToRegex(basicSearchPattern string) (regexPattern string) {
   regexPattern = strings.Replace(basicSearchPattern,"?",".",-1)
   regexPattern = strings.Replace(regexPattern,"*",".*",-1)
   return "^" + regexPattern
}

type TransposalSolver struct{}
func (solver TransposalSolver) Solve(pattern string, dictionary dict.Dictionary, results chan<- string, done_channel chan<- interface{}) {
    //transform the pattern into its ordered set of letters
    transformedString := transform.SortAllCharacters(pattern)

    if matcher,err := newSameCharacterMatcher(); err != nil {
       done_channel <- err
    } else {
       // search for any item in the dictionary that has the same unordered set (using MatchTransformed)
       dictionary.Iterate(func(entry dict.Entry) {
          if matcher.MatchTransformed(transformedString,entry) {
             results <- entry.Word()
          }
       })
       done_channel <- solver
    }
}

func newIdentityMatcher() (Matcher,error) {
  return identityMatcher{},nil
}
func newSameCharacterMatcher() (Matcher,error) {
  return sameCharactersMatcher{},nil
}
func newRegexMatcher(regex string) (matcher Matcher, err error) {
  var compiled *regexp.Regexp
  if compiled,err = regexp.Compile(regex); err != nil {
     return nil,err
  }
  return regexMatcher{compiled},nil
}

func newSameUniqueCharacterMatcher() (Matcher, error) {
   return sameUniqueCharacterMatcher{},nil
}

func newPatternMatcher() (Matcher, error) {
   return samePatternMatcher{}, nil
}

// The Matcher interface defines the ability to look at a word and see if it lines up with a dictionary entry.
type Matcher interface {
    // Returns true if the given word (which is assumed to be the "raw" word) matches the dictionary entry.
    Match(word string, dictEntry dict.Entry)  bool

    // Return true if the transformed word lines up with the dictionary entry. clients can use this to pass words that are already
    // modified for optimal lookups (e.g., a word already set up with characters in order to match an anagram)
    // and to signal the matcher to use that lookup instead of doing work itself
    MatchTransformed(word string, dictEntry dict.Entry) bool
}


// Verifies that the passed word precisely matches the word in the dictionary. This doesn't have much external use, but provides a testing hook.
type identityMatcher struct{}

func (matcher identityMatcher) Match(word string, dictEntry dict.Entry) bool {
  return word == dictEntry.Word()
}
func (matcher identityMatcher) MatchTransformed(word string, dictEntry dict.Entry) bool {
  // in this case, it's the same
  return matcher.Match(word,dictEntry)
}

// checks that two words have the exact same letters
type sameCharactersMatcher struct{}

func (matcher sameCharactersMatcher) Match(word string, dictEntry dict.Entry) bool {
  return matcher.MatchTransformed(transform.SortAllCharacters(word),dictEntry)
}
func (matcher sameCharactersMatcher) MatchTransformed(word string, dictEntry dict.Entry) bool {
  return word == dictEntry.LettersOrdered()
}

// a matcher that checks to see if an entry's raw word matches the regex
type regexMatcher struct{
   *regexp.Regexp
}
func (matcher regexMatcher) Match(word string, dictEntry dict.Entry) bool {
   return matcher.MatchString(dictEntry.Word())
}
func (matcher regexMatcher) MatchTransformed(word string, dictEntry dict.Entry) bool {
   return matcher.Match(word,dictEntry)
}

// a matcher that ensures two words have the same unique letters
type sameUniqueCharacterMatcher struct{}
func (matcher sameUniqueCharacterMatcher) Match(word string, dictEntry dict.Entry) bool {
  return matcher.MatchTransformed(transform.UniqueSortedCharacters(word),dictEntry)
}
func (matcher sameUniqueCharacterMatcher) MatchTransformed(word string, dictEntry dict.Entry) bool {
  return word == dictEntry.UniqueLettersOrdered()
}

// a matcher that ensures two words have the same pattern
type samePatternMatcher struct{}
func (matcher samePatternMatcher) Match(word string, dictEntry dict.Entry) bool {
  return matcher.MatchTransformed(transform.ToPattern(word), dictEntry)
}

func (matcher samePatternMatcher) MatchTransformed(word string, dictEntry dict.Entry) bool {
  return word == dictEntry.Pattern()
}
