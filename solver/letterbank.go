package solver

import (
   "nplsolver/dict"
   "nplsolver/transform"
)
type LetterBankSolver struct{}
func (solver LetterBankSolver) Solve(pattern string, dictionary dict.Dictionary, results chan<- string, done_channel chan<- interface{}) {
    //transform the pattern into its ordered set of letters
    transformedString := transform.UniqueSortedCharacters(pattern)
    
    if matcher,err := newSameUniqueCharacterMatcher(); err != nil {
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
