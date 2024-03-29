package fp

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"testing"

	A "github.com/IBM/fp-go/array"
	F "github.com/IBM/fp-go/function"
	N "github.com/IBM/fp-go/number"
	I "github.com/IBM/fp-go/number/integer"
	S "github.com/IBM/fp-go/string"
)

var (
	Match   = F.Curry2((*regexp.Regexp).FindStringSubmatch)
	Matches = F.Curry2((*regexp.Regexp).MatchString)
	Split   = F.Curry2(F.Bind3of3((*regexp.Regexp).Split)(-1))

	Add      = N.Add[int]
	ToString = I.ToString
	ToLower  = strings.ToLower
	ToUpper  = strings.ToUpper
	Concat   = F.Curry2(S.Monoid.Concat)
)

// Replace cannot be generated via [F.Curry3] because the order of parameters does not match our desired curried order
func Replace(search *regexp.Regexp) func(replace string) func(s string) string {
	return func(replace string) func(s string) string {
		return func(s string) string {
			return search.ReplaceAllString(s, replace)
		}
	}
}

func TestFp4(t *testing.T) {
	// words :: String -> [String]
	words := Split(regexp.MustCompile(` `))

	fmt.Println(words("Jingle bells Batman smells"))

	// Output:
	// [Jingle bells Batman smells]
}

func TestFp5(t *testing.T) {
	// filterQs :: [String] -> [String]
	filterQs := A.Filter(Matches(regexp.MustCompile(`q`)))

	fmt.Println(filterQs(A.From("quick", "camels", "quarry", "over", "quails")))

	// Output:
	// [quick quarry quails]
}

func TestFp6(t *testing.T) {

	keepHighest := N.Max[int]

	// max :: [Number] -> Number
	max := A.Reduce(keepHighest, math.MinInt)

	fmt.Println(max(A.From(323, 523, 554, 123, 5234)))

	// Output:
	// 5234
}
