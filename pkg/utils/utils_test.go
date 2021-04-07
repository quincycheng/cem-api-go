package utils_test

import (
	"testing"

	"github.com/quincycheng/cem-api-go/pkg/test"
	"github.com/quincycheng/cem-api-go/pkg/utils"
)

type TestQuery struct {
	One   string `query_key:"one"`
	Two   int64  `query_key:"two"`
	Three int    `query_key:"three"`
}

func TestGetUrlQuery(t *testing.T) {
	query := &TestQuery{
		One:   "1",
		Two:   2,
		Three: 3,
	}

	result := utils.GetURLQuery(query)

	test.AssertStringStartsWith(result, "?", t)
	test.AssertStringContains(t, result, "one=1")
	test.AssertStringContains(t, result, "two=2")
	test.AssertStringContains(t, result, "three=3")
}

func TestGetUrlQueryNil(t *testing.T) {
	result := utils.GetURLQuery(nil)
	test.AssertStringIsEmpty(result, t)
}

func TestGetUrlQueryNilValues(t *testing.T) {
	result := utils.GetURLQuery(&TestQuery{})
	test.AssertStringIsEmpty(result, t)
}
