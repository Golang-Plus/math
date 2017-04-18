package big

import (
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-plus/errors"
)

// Decimal represents a decimal which can handing fixed precision.
type Decimal struct {
	integer  *big.Int
	decimals int
}

func (d *Decimal) ensureInitialized() {
	if d.integer == nil {
		d.integer = new(big.Int)
	}
}

// IsZero reports whether the value of d is equal to zero.
func (d *Decimal) IsZero() bool {
	d.ensureInitialized()
	return d.integer.Int64() == 0
}

// Sign returns:
// -1: if d <  0
//  0: if d == 0
// +1: if d >  0
func (d *Decimal) Sign() int {
	d.ensureInitialized()
	return d.integer.Sign()
}

// Float32 returns the float32 value nearest to d and a boolean indicating whether is exact.
func (d *Decimal) Float32() (float32, bool) {
	d.ensureInitialized()
	if d.decimals == 0 {
		return float32(d.integer.Int64()), true
	}
	return big.NewRat(d.integer.Int64(), new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(d.decimals)), nil).Int64()).Float32()
}

// Float64 returns the float64 value nearest to d and a boolean indicating whether is exact.
func (d *Decimal) Float64() (float64, bool) {
	d.ensureInitialized()
	if d.decimals == 0 {
		return float64(d.integer.Int64()), true
	}
	return big.NewRat(d.integer.Int64(), new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(d.decimals)), nil).Int64()).Float64()
}

// Int64 returns the int64 value nearest to d and a boolean indicating whether is exact.
func (d *Decimal) Int64() (int64, bool) {
	d.ensureInitialized()
	if d.decimals == 0 {
		return d.integer.Int64(), true
	}
	z, r := new(big.Int).QuoRem(d.integer, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(d.decimals)), nil), new(big.Int))
	return z.Int64(), r.Int64() == 0
}

// String converts the floating-point number d to a string.
func (d *Decimal) String() string {
	d.ensureInitialized()
	str := d.integer.String()
	if d.decimals == 0 {
		return str
	}

	var sign, integer string
	if strings.HasPrefix(str, "-") {
		sign = "-"
		integer = str[1 : len(str)-d.decimals]
	} else {
		integer = str[:len(str)-d.decimals]
	}
	if len(integer) == 0 {
		integer = "0"
	}
	decimals := str[len(str)-d.decimals:]
	return sign + integer + "." + decimals
}

// SetInt64 sets x to y and returns x.
func (x *Decimal) SetInt64(y int64) *Decimal {
	x.ensureInitialized()
	x.integer.SetInt64(y)
	x.decimals = 0
	return x
}

var (
	_DecimalPattern = regexp.MustCompile(`^([-+]?\d+)(\.(\d+))?([eE]([-+]?\d+))?$`)
)

// SetString sets x to the value of y and returns x and a boolean indicating success.
// If the operation failed, the value of d is undefined but the returned value is nil.
func (x *Decimal) SetString(y string) (*Decimal, bool) {
	x.ensureInitialized()

	matches := _DecimalPattern.FindStringSubmatch(y)
	if len(matches) != 6 {
		return nil, false
	}
	integer, _ := strconv.ParseInt(matches[1]+strings.TrimRight(matches[3], "0"), 10, 64)
	decimals := len(matches[3])
	if len(matches[5]) > 0 {
		exponents, _ := strconv.ParseInt(matches[5], 10, 64)
		decimals -= int(exponents)
	}

	x.integer.SetInt64(integer)
	x.decimals = decimals

	return x, true
}

// SetFloat64 sets x to y and returns x.
func (x *Decimal) SetFloat64(y float64) *Decimal {
	x.ensureInitialized()
	str := strconv.FormatFloat(y, 'f', -1, 64)
	x.SetString(str)
	return x
}

// Copy sets x to y and returns x. y is not changed.
func (x *Decimal) Copy(y *Decimal) *Decimal {
	x.ensureInitialized()
	y.ensureInitialized()
	x.integer.SetInt64(y.integer.Int64())
	x.decimals = y.decimals
	return x
}

// Abs sets d to the value |d| (the absolute value of d) and returns d.
func (d *Decimal) Abs() *Decimal {
	d.ensureInitialized()
	d.integer.Abs(d.integer)
	return d
}

// Neg sets d to the value of d with its sign negated, and returns d.
func (d *Decimal) Neg() *Decimal {
	d.ensureInitialized()
	d.integer.Neg(d.integer)
	return d
}

func (x *Decimal) align(y *Decimal) {
	switch {
	case x.decimals < y.decimals:
		x.integer.Mul(x.integer, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(y.decimals-x.decimals)), nil))
		x.decimals = y.decimals
	case x.decimals > y.decimals:
		y.integer.Mul(y.integer, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(x.decimals-y.decimals)), nil))
		y.decimals = x.decimals
	}
}

func (d *Decimal) shrink() {
	for d.decimals > 0 && d.integer.Int64()%10 == 0 {
		d.integer.SetInt64(d.integer.Int64() / 10)
		d.decimals -= 1
	}
}

// Cmp compares x and y and returns:
// -1 if d < y
//  0 if d == y (includes: -0 == 0, -Inf == -Inf, and +Inf == +Inf)
// +1 if d > y
func (x *Decimal) Cmp(y *Decimal) int {
	x.ensureInitialized()
	y.ensureInitialized()
	x.align(y)
	sign := x.integer.Cmp(y.integer)
	x.shrink()
	y.shrink()
	return sign
}

// Add sets d to the sum of d and y and returns x.
func (x *Decimal) Add(y *Decimal) *Decimal {
	x.ensureInitialized()
	y.ensureInitialized()
	x.align(y)
	x.integer.Add(x.integer, y.integer)
	x.shrink()
	y.shrink()
	return x
}

// Sub sets d to the difference x-y and returns x.
func (x *Decimal) Sub(y *Decimal) *Decimal {
	x.ensureInitialized()
	y.ensureInitialized()
	x.align(y)
	x.integer.Sub(x.integer, y.integer)
	x.shrink()
	y.shrink()
	return x
}

// Mul sets x to the product x*y and returns x.
func (x *Decimal) Mul(y *Decimal) *Decimal {
	x.ensureInitialized()
	y.ensureInitialized()
	x.align(y)
	x.integer.Mul(x.integer, y.integer)
	x.decimals *= 2
	x.shrink()
	y.shrink()
	return x
}

// Quo sets x to the quotient x/y and return x.
func (x *Decimal) Quo(y *Decimal) *Decimal {
	x.ensureInitialized()
	y.ensureInitialized()
	x.align(y)
	if q, r := new(big.Int).QuoRem(x.integer, y.integer, new(big.Int)); r.Int64() == 0 { // modulus x%y == 0
		x.integer.SetInt64(q.Int64())
		x.decimals = 0
	} else { // modulus x%y > 0
		f, _ := new(big.Float).Quo(new(big.Float).SetInt(x.integer), new(big.Float).SetInt(y.integer)).Float64()
		x.SetFloat64(f)
	}

	y.shrink()
	return x
}

// Div is same to Quo.
func (x *Decimal) Div(y *Decimal) *Decimal {
	return x.Quo(y)
}

// RoundToNearestEven rounds (IEEE 754-2008 Round to nearest, ties to even) the floating-point number x with given precision (the number of digits after the decimal point).
func (x *Decimal) RoundToNearestEven(precision uint) *Decimal {
	x.ensureInitialized()
	prec := int(precision)
	if x.IsZero() || x.decimals == 0 || x.decimals <= prec { // rounding needless
		return x
	}

	str := strconv.FormatInt(x.integer.Int64(), 10)
	part1 := str[:len(str)-x.decimals]
	part2 := str[len(str)-x.decimals:]
	isRoundUp := false
	switch part2[prec : prec+1] {
	case "6", "7", "8", "9":
		isRoundUp = true
	case "5":
		if len(part2) > prec+1 { // found decimals back of "5"
			isRoundUp = true
		} else {
			var neighbor string
			if prec == 0 { // get neighbor from integer part
				neighbor = part1[len(part1)-1:]
			} else {
				neighbor = part2[prec-1 : prec]
			}
			switch neighbor {
			case "1", "3", "5", "7", "9":
				isRoundUp = true
			}
		}
	}

	z, _ := strconv.ParseInt(part1+part2[:prec], 10, 64)
	if isRoundUp {
		z += int64(x.integer.Sign() * 1)
	}
	x.integer.SetInt64(z)
	x.decimals = prec

	return x
}

// Round is short to RoundToNearestEven.
func (x *Decimal) Round(precision uint) *Decimal {
	return x.RoundToNearestEven(precision)
}

// RoundToNearestAway rounds (IEEE 754-2008, Round to nearest, ties away from zero) the floating-point number x with given precision (the number of digits after the decimal point).
func (x *Decimal) RoundToNearestAway(precision uint) *Decimal {
	x.ensureInitialized()
	prec := int(precision)
	if x.IsZero() || x.decimals == 0 || x.decimals <= prec { // rounding needless
		return x
	}

	x.integer.Quo(x.integer, new(big.Int).Exp(big.NewInt(int64(10)), big.NewInt(int64(x.decimals-int(precision)-1)), nil))
	factor := big.NewInt(int64(5))
	if x.integer.Sign() == -1 {
		factor.Neg(factor)
	}
	x.integer.Add(x.integer, factor)
	x.integer.Quo(x.integer, big.NewInt(int64(10)))
	x.decimals = prec
	return x
}

// RoundToZero rounds (IEEE 754-2008 Round Towards Zero) the floating-point number x with given precision (the number of digits after the decimal point).
func (x *Decimal) RoundToZero(precision uint) *Decimal {
	x.ensureInitialized()
	prec := int(precision)
	if x.IsZero() || x.decimals == 0 || x.decimals <= prec { // rounding needless
		return x
	}

	x.integer.Quo(x.integer, new(big.Int).Exp(big.NewInt(int64(10)), big.NewInt(int64(x.decimals-int(precision))), nil))
	x.decimals = prec
	return x
}

// Truncate is same as RoundToZero.
func (x *Decimal) Truncate(precision uint) *Decimal {
	return x.RoundToZero(precision)
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
