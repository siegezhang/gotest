package fp

import (
	"fmt"
	E "github.com/IBM/fp-go/either"
	"github.com/IBM/fp-go/errors"
	F "github.com/IBM/fp-go/function"
	O "github.com/IBM/fp-go/option"
	S "github.com/IBM/fp-go/string"
	"testing"
)

func foo(bar func() string) func() string {
	return func() string {
		return "foo" + "  " + bar()
	}
}

func Filter[T any](f func(T) bool, src []T) []T {
	var dst []T
	for _, v := range src {
		if f(v) {
			dst = append(dst, v)
		}
	}
	return dst
}

type Thing struct {
	Name string
}

func (t Thing) GetName() string {
	return t.Name
}

var (
	// func(Thing) Either[error, string]
	getName = F.Flow2(
		Thing.GetName,
		E.FromPredicate(S.IsNonEmpty, errors.OnSome[string]("value [%s] is empty")),
	)

	// func(option.Option[Thing]) Either[error, string]
	GetName = F.Flow2(
		E.FromOption[Thing](errors.OnNone("value is none")),
		E.Chain(getName),
	)
)

func TestFp1(t *testing.T) {
	bar := func() string {
		return "bar"
	}
	foobar := foo(bar)
	fmt.Println(foobar())
}

func TestFp2(t *testing.T) {
	src := []int{-2, -1, -0, 1, 2}
	f := func(v int) bool { return v >= 0 }
	dst := Filter(f, src)
	fmt.Println(dst)
}

func TestFp3(t *testing.T) {

	oThing := O.Of(Thing{"Carsten"})

	res := F.Pipe2(
		oThing,
		GetName,
		E.Fold(S.Format[error]("failed with error %v"), S.Format[string]("get value %s")),
	)
	fmt.Println(res)

	// Output:
	// get value Carsten

}
