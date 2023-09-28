package fp

import (
	"fmt"
	"regexp"
	"testing"

	A "github.com/IBM/fp-go/array"
	F "github.com/IBM/fp-go/function"
	I "github.com/IBM/fp-go/number/integer"
	O "github.com/IBM/fp-go/option"
	"github.com/IBM/fp-go/ord"
	S "github.com/IBM/fp-go/string"
)

var (
	Exclaim   = S.Format[string]("%s!")
	Shout     = F.Flow2(ToUpper, Exclaim)
	Dasherize = F.Flow4(
		Replace(regexp.MustCompile(`\s{2,}`))(" "),
		Split(regexp.MustCompile(` `)),
		A.Map(ToLower),
		A.Intercalate(S.Monoid)("-"),
	)
)

func TestExample_shout(t *testing.T) {
	fmt.Println(Shout("send in the clowns"))

	// Output: SEND IN THE CLOWNS!
}

func TestExample_dasherize(t *testing.T) {
	fmt.Println(Dasherize("The world is a vampire"))

	// Output: the-world-is-a-vampire
}

func TestExample_pipe(t *testing.T) {
	output := F.Pipe2(
		"send in the clowns",
		ToUpper,
		Exclaim,
	)

	fmt.Println(output)

	// Output: SEND IN THE CLOWNS!
}

func TestExample_solution05A(t *testing.T) {
	IsLastInStock := F.Flow2(
		A.Last[Car],
		O.Map(Car.getInStock),
	)

	fmt.Println(IsLastInStock(Cars[0:3]))
	fmt.Println(IsLastInStock(Cars[3:]))

	// Output:
	// Some[bool](true)
	// Some[bool](false)
}

func TestExample_solution05B(t *testing.T) {
	// averageDollarValue :: [Car] -> Int
	averageDollarValue := F.Flow2(
		A.Map(Car.getDollarValue),
		average,
	)

	fmt.Println(averageDollarValue(Cars))

	// Output:
	// 790700
}

func TestExample_solution05C(t *testing.T) {
	// order by horsepower
	ordByHorsepower := ord.Contramap(Car.getHorsepower)(I.Ord)

	// fastestCar :: [Car] -> Option[String]
	fastestCar := F.Flow3(
		A.Sort(ordByHorsepower),
		A.Last[Car],
		O.Map(F.Flow2(
			Car.getName,
			S.Format[string]("%s is the fastest"),
		)),
	)

	fmt.Println(fastestCar(Cars))

	// Output:
	// Some[string](Aston Martin One-77 is the fastest)
}
