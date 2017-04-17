<a name=""></a>
#  (2017-04-17)


### Bug Fixes

* Float32 Float64 Int64 returns wrong result ([db89e1e](https://github.com/Golang-Plus/math/commit/db89e1e))
* produces unpredictable results sometimes ([f912932](https://github.com/Golang-Plus/math/commit/f912932))


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



<a name=""></a>
#  (2017-04-17)


### Bug Fixes

* produces unpredictable results sometimes ([f912932](https://github.com/Golang-Plus/math/commit/f912932))


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



