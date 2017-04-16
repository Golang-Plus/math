package big

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/golang-plus/errors"
)

// Decimal represents a decimal which can handing fixed precision.
type Decimal struct {
	float     *big.Float
	precision int
}

func (x *Decimal) ensureInitialized() {
	if x.float == nil {
		x.float = new(big.Float)
		x.precision = -1
	}
}

// SetInt64 sets x to v and returns x.
func (x *Decimal) SetInt64(v int64) *Decimal {
	x.ensureInitialized()
	x.float.SetInt64(v)
	return x
}

// SetFloat64 sets x to v and returns x.
func (x *Decimal) SetFloat64(v float64) *Decimal {
	x.ensureInitialized()
	x.float.SetFloat64(v)
	return x
}

// SetString sets d to the value of s and returns d and a boolean indicating success.
// If the operation failed, the value of d is undefined but the returned value is nil.
func (x *Decimal) SetString(s string) (*Decimal, bool) {
	x.ensureInitialized()
	if _, ok := x.float.SetString(s); ok {
		return x, true
	}
	return nil, false
}

// Copy sets x to y and returns x. y is not changed.
func (x *Decimal) Copy(y *Decimal) *Decimal {
	x.ensureInitialized()
	x.float.Copy(y.float)
	return x
}

// Sign returns:
// -1: if d <  0
//  0: if d == 0
// +1: if d >  0
func (x *Decimal) Sign() int {
	x.ensureInitialized()
	return x.float.Sign()
}

// Cmp compares d and y and returns:
// -1 if d < y
//  0 if d == y (includes: -0 == 0, -Inf == -Inf, and +Inf == +Inf)
// +1 if d > y
func (x *Decimal) Cmp(y *Decimal) int {
	x.ensureInitialized()
	y.ensureInitialized()
	return x.float.Cmp(y.float)
}

// IsZero reports whether the x is equal to zero.
func (x *Decimal) IsZero() bool {
	x.ensureInitialized()
	return x.float.Cmp(big.NewFloat(0)) == 0
}

// Abs sets x to the value |x| (the absolute value of x) and returns x.
func (x *Decimal) Abs() *Decimal {
	x.ensureInitialized()
	x.float.Abs(x.float)
	return x
}

// Neg sets x to the value of x with its sign negated, and returns x.
func (x *Decimal) Neg() *Decimal {
	x.ensureInitialized()
	x.float.Neg(x.float)
	return x
}

// Add sets d to the sum of d and y and returns x.
func (x *Decimal) Add(y *Decimal) *Decimal {
	x.ensureInitialized()
	y.ensureInitialized()
	x.float.Add(x.float, y.float)
	return x
}

// Sub sets d to the difference x-y and returns x.
func (x *Decimal) Sub(y *Decimal) *Decimal {
	x.ensureInitialized()
	y.ensureInitialized()
	x.float.Sub(x.float, y.float)
	return x
}

// Mul sets x to the product x*y and returns x.
func (x *Decimal) Mul(y *Decimal) *Decimal {
	x.ensureInitialized()
	y.ensureInitialized()
	x.float.Mul(x.float, y.float)
	return x
}

// Quo sets x to the quotient x/y and return x.
func (x *Decimal) Quo(y *Decimal) *Decimal {
	x.ensureInitialized()
	y.ensureInitialized()
	x.float.Quo(x.float, y.float)
	return x
}

// Div is same to Quo.
func (x *Decimal) Div(y *Decimal) *Decimal {
	return x.Quo(y)
}

// RoundToNearestEven rounds (IEEE 754-2008 Round to nearest, ties to even) the floating-point number x with given precision (the number of digits after the decimal point).
func (x *Decimal) RoundToNearestEven(precision uint) *Decimal {
	x.ensureInitialized()
	x.precision = int(precision)

	str := strings.TrimRight(x.float.Text('f', int(precision+1)), "0")
	parts := strings.Split(str, ".")
	if len(parts) == 1 { // x is integer, do nothing
		return x
	}
	if len(parts[1]) <= int(precision) { // round not needed
		return x
	}

	roundUp := false
	var rounded bytes.Buffer
	rounded.WriteString(parts[0])
	if precision > 0 {
		rounded.WriteString(".")
		rounded.WriteString(parts[1][:precision])
	}
	switch parts[1][precision : precision+1] {
	case "6", "7", "8", "9":
		roundUp = true
	case "5":
		if f, _ := new(big.Float).SetString(str); new(big.Float).Abs(x.float).Cmp(new(big.Float).Abs(f)) > 0 { // got decimals back of "5"
			roundUp = true
		} else {
			var neighbor string
			if precision > 1 {
				neighbor = parts[1][precision-1 : precision]
			} else {
				neighbor = parts[0][len(parts[0])-1 : len(parts[0])]
			}
			switch neighbor {
			case "1", "3", "5", "7", "9":
				roundUp = true
			}
		}
	}

	z, _ := new(big.Float).SetString(rounded.String())
	if roundUp {
		factor := new(big.Float).Quo(big.NewFloat(0.1), big.NewFloat(math.Pow10(int(precision)-1)))
		if x.float.Sign() == -1 {
			factor.Neg(factor)
		}
		z.Add(z, factor)
	}
	x.float.Copy(z)

	return x
}

// Round is short to RoundToNearestEven.
func (x *Decimal) Round(precision uint) *Decimal {
	return x.RoundToNearestEven(precision)
}

// RoundToNearestAway rounds (IEEE 754-2008, Round to nearest, ties away from zero) the floating-point number x with given precision (the number of digits after the decimal point).
func (x *Decimal) RoundToNearestAway(precision uint) *Decimal {
	x.ensureInitialized()
	x.precision = int(precision)

	pow := big.NewFloat(math.Pow10(int(precision)))
	factor := big.NewFloat(0.5)
	if x.float.Sign() == -1 {
		factor.Neg(factor)
	}
	y := new(big.Float).Copy(x.float)
	y.Mul(y, pow)
	y.Add(y, factor)
	integer, _ := y.Int64()
	z := new(big.Float).SetInt64(integer)
	z.Quo(z, pow)
	x.float.Copy(z)

	return x
}

// RoundToZero rounds (IEEE 754-2008 Round Towards Zero) the floating-point number x with given precision (the number of digits after the decimal point).
func (x *Decimal) RoundToZero(precision uint) *Decimal {
	x.ensureInitialized()
	x.precision = int(precision)

	pow := big.NewFloat(math.Pow10(int(precision)))
	y := new(big.Float).Copy(x.float)
	y.Mul(y, pow)
	integer, _ := y.Int64()
	z := new(big.Float).SetInt64(integer)
	z.Quo(z, pow)
	x.float.Copy(z)

	return x
}

// Truncate is same as RoundToZero.
func (x *Decimal) Truncate(precision uint) *Decimal {
	return x.RoundToZero(precision)
}

// Float32 returns the float32 value nearest to x and a boolean indicating whether is exact.
func (x *Decimal) Float32() (float32, bool) {
	x.ensureInitialized()
	f, acc := x.float.Float32()
	isExact := acc == big.Exact
	return f, isExact
}

// Float64 returns the float64 value nearest to x and a boolean indicating whether is exact.
func (x *Decimal) Float64() (float64, bool) {
	x.ensureInitialized()
	f, acc := x.float.Float64()
	isExact := acc == big.Exact
	return f, isExact
}

// Int64 returns the int64 value nearest to x and a boolean indicating whether is exact.
func (x *Decimal) Int64() (int64, bool) {
	x.ensureInitialized()
	f, acc := x.float.Int64()
	isExact := acc == big.Exact
	return f, isExact
}

// String converts the floating-point number x to a string.
func (x *Decimal) String() string {
	x.ensureInitialized()
	f, _ := x.float.Float64()
	if x.precision != -1 {
		return fmt.Sprintf("%."+strconv.Itoa(int(x.precision))+"f", f)
	}
	return fmt.Sprintf("%f", f)
}

// NewDecimal returns a new decimal.
func NewDecimal(number float64) *Decimal {
	return new(Decimal).SetFloat64(number)
}

// ParseDecimal returns a new decimal by parsing decimal string.
func ParseDecimal(str string) (*Decimal, error) {
	if d, ok := new(Decimal).SetString(str); ok {
		return d, nil
	}
	return nil, errors.Newf("decimal string %q is invalid", str)
}
