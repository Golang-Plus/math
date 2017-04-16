<a name=""></a>
#  (2017-04-16)


### Features

* add functions for Decimal type ([5f6844f](https://github.com/Golang-Plus/math/commit/5f6844f))


### BREAKING CHANGES

* some fucntions has changed or removed.

Changed:

    Before: *Decimal.SetString(string) (*Decimal, error)
    After : *Decimal.SetString(string) (*Decimal, bool)

    Before: *Decimal.Float64() float64
    After : *Decimal.Float64() (float64, bool)

Removed:

    *Decimal.FloatString(uint) string



