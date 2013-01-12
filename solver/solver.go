// The solver package encapsulates all the solvers and support code used by the solver server
package solver

import (
  "regexp"
  "nplsolver/dict"
  "nplsolver/transform"
)

type MatcherType int

const (
   Identity = iota
)

func NewIdentityMatcher() (Matcher,error) {
  return IdentityMatcher{},nil
}
func NewSameCharacterMatcher() (Matcher,error) {
  return SameCharactersMatcher{},nil
}
func NewRegexMatcher(regex string) (matcher Matcher, err error) {
  var compiled *regexp.Regexp
  if compiled,err = regexp.Compile(regex); err != nil {
     return nil,err
  }
  return RegexMatcher{compiled},nil
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
type IdentityMatcher struct{}

func (matcher IdentityMatcher) Match(word string, dictEntry dict.Entry) bool {
  return word == dictEntry.Word()
}
func (matcher IdentityMatcher) MatchTransformed(word string, dictEntry dict.Entry) bool {
  // in this case, it's the same
  return matcher.Match(word,dictEntry)
}

// checks that two words have the exact same letters
type SameCharactersMatcher struct{}

func (matcher SameCharactersMatcher) Match(word string, dictEntry dict.Entry) bool {
  return matcher.MatchTransformed(transform.SortAllCharacters(word),dictEntry)
}
func (matcher SameCharactersMatcher) MatchTransformed(word string, dictEntry dict.Entry) bool {
  return word == dictEntry.LettersOrdered()
}

// a matcher that checks to see if an entry's raw word matches the regex
type RegexMatcher struct{
   *regexp.Regexp
}
func (matcher RegexMatcher) Match(word string, dictEntry dict.Entry) bool {
   return matcher.MatchString(dictEntry.Word())
}
func (matcher RegexMatcher) MatchTransformed(word string, dictEntry dict.Entry) bool {
   return matcher.Match(word,dictEntry)
}


     




