package subpackage_test

import (
	. "github.com/liamsi/testgo17"
	. "github.com/liamsi/testgo17/subpackage"
	"testing"
)

func TestUsingSimpleType(t *testing.T) {
	_ = SimpleType{1, 10}
	_ = SimpleType2{2, 3}
}
