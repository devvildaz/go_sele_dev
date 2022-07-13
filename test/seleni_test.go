package seleni_test

import (
	"testing"

	operations "github.com/devvildaz/seleni/internal/foo"
	"github.com/ysmood/got"
)

func TestAssertion(t *testing.T) {
	result := operations.Sum("10", "1024")
	got.T(t).Eq(result, "2")
}
