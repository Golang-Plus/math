package big

import (
	"bytes"
	"io"
	"math/big"
	"strconv"
	"strings"

	"github.com/golang-plus/errors"
)

// alignDecimalScale aligns the scale of two decimals.
func alignDecimalScale(a, b *Decimal) {
	switch {
	case a.scale < b.scale:
		a.integer.Mul(a.integer, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(b.scale-a.scale)), nil))
		a.scale = b.scale
	case a.scale > b.scale:
		b.integer.Mul(b.integer, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(a.scale-b.scale)), nil))
		b.scale = a.scale
	}
}

// Decimal represents a decimal which can handing fixed precision.
type Decimal struct {
	integer *big.Int
	scale   int // scale represents the number of deciaml digits
}

func (d *Decimal) ensureValid() {
	if d.integer == nil {
		d.integer = new(big.Int)
	}
}

// SetInt64 sets d to v and returns d.
func (d *Decimal) SetInt64(v int64) *Decimal {
	d.integer = big.NewInt(v)
	return d
}

// SetString sets d to the valud of v and returns d.
func (d *Decimal) SetString(v string) (*Decimal, error) {
	numberString := v
	var unscaledBuffer bytes.Buffer
	var scale int
	reader := strings.NewReader(numberString)
	index := 1
	for {
		ch, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, errors.Wrap(err, "could not read number string by rune")
		}

		switch ch {
		case '+', '-':
			if index > 1 { // sign must be first character
				return nil, errors.Newf("invalid number string %q", numberString)
			}
			unscaledBuffer.WriteRune(ch)
		case '.':
			if scale != 0 {
				return nil, errors.Newf("invalid number string %q", numberString)
			}
			scale = len(numberString) - index
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			unscaledBuffer.WriteRune(ch)
		default:
			return nil, errors.Newf("invalid number string %q", numberString)
		}

		index++
	}

	integer, ok := new(big.Int).SetString(unscaledBuffer.String(), 10)
	if !ok {
		return nil, errors.Newf("invalid number string %q", numberString)
	}

	d.integer = integer
	d.scale = scale

	return d, nil
}

// SetFloat64 sets d to v and returns d.
func (d *Decimal) SetFloat64(v float64) *Decimal {
	numberString := strconv.FormatFloat(v, 'f', -1, 64)
	d.SetString(numberString)
	return d
}

// Cmp compares d and another and returns:
// -1 if d <  another
//  0 if d == another
// +1 if d >  another
func (d *Decimal) Cmp(another *Decimal) int {
	d.ensureValid()
	another.ensureValid()

	alignDecimalScale(d, another)
	return d.integer.Cmp(another.integer)
}

// Add sets d to the sum of d and another and returns d.
func (d *Decimal) Add(another *Decimal) *Decimal {
	d.ensureValid()
	another.ensureValid()

	alignDecimalScale(d, another)
	d.integer.Add(d.integer, another.integer)
	return d
}

// Sub sets d to the difference d-another and returns d.
func (d *Decimal) Sub(another *Decimal) *Decimal {
	d.ensureValid()
	another.ensureValid()

	alignDecimalScale(d, another)
	d.integer.Sub(d.integer, another.integer)
	return d
}

// Mul sets d to the product d*another and returns d.
func (d *Decimal) Mul(another *Decimal) *Decimal {
	d.ensureValid()
	another.ensureValid()

	d.integer.Mul(d.integer, another.integer)
	d.scale += another.scale
	return d
}

// Div sets d to the quotient d/another and return d.
func (d *Decimal) Div(another *Decimal) *Decimal {
	d.ensureValid()
	another.ensureValid()

	numerator := new(big.Int).Exp(big.NewInt(int64(10)), big.NewInt(int64(another.scale)), nil).Int64()
	denominator := another.integer.Int64()
	b, _ := big.NewRat(numerator, denominator).Float64()

	return d.Mul(new(Decimal).SetFloat64(b))
}

// Sign returns:
// -1: if d <  0
//  0: if d == 0
// +1: if d >  0
func (d *Decimal) Sign() int {
	d.ensureValid()
	return d.integer.Sign()
}

// Float64 returns the nearest float64 value of decimal.
func (d *Decimal) Float64() float64 {
	d.ensureValid()
	resultString := d.String()
	result, _ := strconv.ParseFloat(resultString, 64)
	return result
}

// FloatString returns a string representation of decimal form with precision digits of precision after the decimal point and the last digit rounded.
func (d *Decimal) FloatString(precision uint) string {
	d.ensureValid()

	x := new(big.Rat).SetInt(d.integer)
	y := new(big.Rat).Inv(new(big.Rat).SetInt(new(big.Int).Exp(big.NewInt(int64(10)), big.NewInt(int64(d.scale)), nil)))
	z := new(big.Rat).Mul(x, y)
	return z.FloatString(int(precision))
}

// String returns the string of Decimal.
func (d *Decimal) String() string {
	d.ensureValid()

	unscaledString := strings.TrimLeft(d.integer.String(), "-")
	if d.scale == 0 {
		return unscaledString
	}

	pointIndex := len(unscaledString) - d.scale
	switch {
	case pointIndex < 0:
		if d.integer.Sign() == -1 {
			return "-0." + strings.Repeat("0", -1*pointIndex) + unscaledString
		}

		return "0." + strings.Repeat("0", -1*pointIndex) + unscaledString
	case pointIndex > 0:
		if d.integer.Sign() == -1 {
			return "-" + unscaledString[0:pointIndex] + "." + unscaledString[pointIndex:]
		}

		return unscaledString[0:pointIndex] + "." + unscaledString[pointIndex:]
	default: // pointIndex == 0
		if d.integer.Sign() == -1 {
			return "-0." + unscaledString
		}

		return "0." + unscaledString
	}
}

// NewDecimal returns a new decimal.
func NewDecimal(number float64) *Decimal {
	return new(Decimal).SetFloat64(number)
}

// ParseDecimal returns a new decimal by parse decimal string.
func ParseDecimal(numberString string) (*Decimal, error) {
	return new(Decimal).SetString(numberString)
}
