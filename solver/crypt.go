package solver

import (
   "nplsolver/dict"
   "nplsolver/transform"
)

type CryptSolver struct{}
func (solver CryptSolver) Solve(pattern string, dictionary dict.Dictionary, results chan<- string, done_channel chan<- interface{}) {
    transformedString := transform.ToPattern(pattern)
    
    if matcher,err := newPatternMatcher(); err != nil {
       done_channel <- err
    } else {
       // search for any item in the dictionary that has the same unordered set (using MatchTransformed)
       dictionary.Iterate(func(entry dict.Entry) {
          if matcher.MatchTransformed(transformedString, entry) {
             results <- entry.Word()
          }
       })
       done_channel <- solver
    }
}
