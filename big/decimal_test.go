package big

import (
	//	"fmt"
	"testing"

	testing2 "github.com/golang-plus/testing"
)

func TestDecimal(t *testing.T) {
	/*x := 1.23456789*/
	//f32, _ := NewDecimal(x).Float32()
	//f64, _ := NewDecimal(x).Float64()
	//i64, _ := NewDecimal(x).Int64()
	/*fmt.Println(x, f32, f64, i64)*/

	// round to nearest away
	data := map[string]string{
		"3.141592653589793238":  "3.1416",
		"-3.141592653589793238": "-3.1416",
		"11.25555":              "11.2556",
		"-11.25555":             "-11.2556",
		"22.25554":              "22.2555",
		"-22.25554":             "-22.2555",
		"33.25556":              "33.2556",
		"-33.25556":             "-33.2556",
		"0":                     "0",
		"-0":                    "0",
		"+0":                    "0",
	}
	for k, v := range data {
		d, err := ParseDecimal(k)
		if err != nil {
			t.Fatal(err)
		}
		testing2.AssertEqual(t, d.RoundToNearestAway(4).String(), v)
	}

	// round to nearest even 0
	data = map[string]string{
		"3.141592653589793238":  "3.1416",
		"-3.141592653589793238": "-3.1416",
		"11.25555":              "11.2556",
		"-11.25555":             "-11.2556",
		"22.25554":              "22.2555",
		"-22.25554":             "-22.2555",
		"33.25556":              "33.2556",
		"-33.25556":             "-33.2556",
		"2.34564":               "2.3456",
		"-2.34564":              "-2.3456",
		"2.34566":               "2.3457",
		"-2.34566":              "-2.3457",
		"2.34605001":            "2.3461",
		"-2.34605001":           "-2.3461",
		"2.34635":               "2.3464",
		"-2.34635":              "-2.3464",
		"2.34645":               "2.3464",
		"-2.34645":              "-2.3464",
		"2.34555":               "2.3456",
		"2.34525":               "2.3452",
		"2.1234":                "2.1234",
		"2.1":                   "2.1",
		"2":                     "2",
	}
	for k, v := range data {
		d, err := ParseDecimal(k)
		if err != nil {
			t.Fatal(err)
		}
		testing2.AssertEqual(t, d.RoundToNearestEven(4).String(), v)
	}

	data = map[string]string{
		"3.141592653589793238":  "3.1",
		"-3.141592653589793238": "-3.1",
		"11.25555":              "11.3",
		"-11.25555":             "-11.3",
		"22.25554":              "22.3",
		"-22.25554":             "-22.3",
		"33.25556":              "33.3",
		"-33.25556":             "-33.3",
		"2.34564":               "2.3",
		"-2.34564":              "-2.3",
		"2.34566":               "2.3",
		"-2.34566":              "-2.3",
		"2.34605001":            "2.3",
		"-2.34605001":           "-2.3",
		"2.34635":               "2.3",
		"-2.34635":              "-2.3",
		"2.34645":               "2.3",
		"-2.34645":              "-2.3",
		"2.34555":               "2.3",
		"2.34525":               "2.3",
		"2.1234":                "2.1",
		"2.1":                   "2.1",
		"2":                     "2",
	}
	for k, v := range data {
		d, err := ParseDecimal(k)
		if err != nil {
			t.Fatal(err)
		}
		testing2.AssertEqual(t, d.RoundToNearestEven(1).String(), v)
	}

	data = map[string]string{
		"3.141592653589793238":  "3",
		"-3.141592653589793238": "-3",
		"11.25555":              "11",
		"-11.25555":             "-11",
		"22.25554":              "22",
		"-22.25554":             "-22",
		"33.25556":              "33",
		"-33.25556":             "-33",
		"2.34564":               "2",
		"-2.34564":              "-2",
		"2.34566":               "2",
		"-2.34566":              "-2",
		"2.34605001":            "2",
		"-2.34605001":           "-2",
		"2.34635":               "2",
		"-2.34635":              "-2",
		"2.34645":               "2",
		"-2.34645":              "-2",
		"2.34555":               "2",
		"2.34525":               "2",
		"2.1234":                "2",
		"2.1":                   "2",
		"2":                     "2",
		"2.6":                   "3",
		"-2.6":                  "-3",
		"2.51":                  "3",
		"-2.51":                 "-3",
		"2.5":                   "2",
		"1.5":                   "2",
	}
	for k, v := range data {
		d, err := ParseDecimal(k)
		if err != nil {
			t.Fatal(err)
		}
		testing2.AssertEqual(t, d.RoundToNearestEven(0).String(), v)
	}

	// round to zero
	data = map[string]string{
		"3.141592653589793238":  "3.1415",
		"-3.141592653589793238": "-3.1415",
		"11.25555":              "11.2555",
		"-11.25555":             "-11.2555",
		"22.25554":              "22.2555",
		"-22.25554":             "-22.2555",
		"33.25556":              "33.2555",
		"-33.25556":             "-33.2555",
	}
	for k, v := range data {
		d, err := ParseDecimal(k)
		if err != nil {
			t.Fatal(err)
		}
		testing2.AssertEqual(t, d.RoundToZero(4).String(), v)
	}

	data2 := map[string][4]float64{
		"-0.825":     {1.1, 2.2, 3.3, 4.4},
		"2.2":        {4.4, 3.3, 2.2, 1.1},
		"22.22":      {44.44, 33.33, 22.22, 11.11},
		"2222.2222":  {4444.4444, 3333.3333, 2222.2222, 1111.1111},
		"-833.3333":  {1111.1111, 2222.2222, 3333.3333, 4444.4444},
		"-1481.4815": {1111.1111, 2222.2222, 4444.4444, 3333.3333},
	}
	for k, v := range data2 {
		// + - * /
		d := new(Decimal).Add(NewDecimal(v[0])).Sub(NewDecimal(v[1])).Mul(NewDecimal(v[2])).Div(NewDecimal(v[3]))
		testing2.AssertEqual(t, d.Round(4).String(), k)
	}
}
