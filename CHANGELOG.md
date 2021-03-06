<a name=""></a>
# [](https://github.com/Golang-Plus/math/compare/v1.2.0...v) (2017-04-23)


### Bug Fixes

* *Decimal.String bug in special case ([91559be](https://github.com/Golang-Plus/math/commit/91559be))
* bugs in special cases ([3c4ca23](https://github.com/Golang-Plus/math/commit/3c4ca23))
* RoundAwayFromZero bug ([3e38e04](https://github.com/Golang-Plus/math/commit/3e38e04))
* RoundAwayFromZero bug in special case ([38149d8](https://github.com/Golang-Plus/math/commit/38149d8))


### Features

* add more round functions ([907d20a](https://github.com/Golang-Plus/math/commit/907d20a))
* add MustParseDecimal func ([0dcf854](https://github.com/Golang-Plus/math/commit/0dcf854))



<a name=""></a>
# [](https://github.com/Golang-Plus/math/compare/v1.2.0...v) (2017-04-23)


### Bug Fixes

* *Decimal.String bug in special case ([91559be](https://github.com/Golang-Plus/math/commit/91559be))
* bugs in special cases ([3c4ca23](https://github.com/Golang-Plus/math/commit/3c4ca23))
* RoundAwayFromZero bug in special case ([38149d8](https://github.com/Golang-Plus/math/commit/38149d8))


### Features

* add more round functions ([907d20a](https://github.com/Golang-Plus/math/commit/907d20a))
* add MustParseDecimal func ([0dcf854](https://github.com/Golang-Plus/math/commit/0dcf854))



<a name=""></a>
# [](https://github.com/Golang-Plus/math/compare/v1.2.0...v) (2017-04-23)


### Bug Fixes

* *Decimal.String bug in special case ([91559be](https://github.com/Golang-Plus/math/commit/91559be))
* bugs in special cases ([3c4ca23](https://github.com/Golang-Plus/math/commit/3c4ca23))


### Features

* add more round functions ([907d20a](https://github.com/Golang-Plus/math/commit/907d20a))
* add MustParseDecimal func ([0dcf854](https://github.com/Golang-Plus/math/commit/0dcf854))



<a name=""></a>
# [](https://github.com/Golang-Plus/math/compare/v1.2.0...v) (2017-04-21)


### Bug Fixes

* *Decimal.String bug in special case ([91559be](https://github.com/Golang-Plus/math/commit/91559be))
* bugs in special cases ([3c4ca23](https://github.com/Golang-Plus/math/commit/3c4ca23))


### Features

* add more round functions ([907d20a](https://github.com/Golang-Plus/math/commit/907d20a))



<a name=""></a>
# [](https://github.com/Golang-Plus/math/compare/v1.2.0...v) (2017-04-21)


### Bug Fixes

* bugs in special cases ([3c4ca23](https://github.com/Golang-Plus/math/commit/3c4ca23))


### Features

* add more round functions ([907d20a](https://github.com/Golang-Plus/math/commit/907d20a))



<a name=""></a>
# [](https://github.com/Golang-Plus/math/compare/v1.2.0...v) (2017-04-21)


### Features

* add more round functions ([907d20a](https://github.com/Golang-Plus/math/commit/907d20a))



<a name=""></a>
#  (2017-04-20)


### Bug Fixes

* *Decimal.Quo has problem when divisible ([339c686](https://github.com/Golang-Plus/math/commit/339c686))
* *Decimal.String panics when integer part is 0 ([49c8228](https://github.com/Golang-Plus/math/commit/49c8228))
* Float32 Float64 Int64 returns wrong result ([db89e1e](https://github.com/Golang-Plus/math/commit/db89e1e))
* produces unpredictable results sometimes ([f912932](https://github.com/Golang-Plus/math/commit/f912932))


### Features

* add functions for Decimal type ([5f6844f](https://github.com/Golang-Plus/math/commit/5f6844f))
* support unlimited decimal digits ([c941b7b](https://github.com/Golang-Plus/math/commit/c941b7b))


### BREAKING CHANGES

* some fucntions has changed or removed.

Changed:

    Before: *Decimal.SetString(string) (*Decimal, error)
    After : *Decimal.SetString(string) (*Decimal, bool)

    Before: *Decimal.Float64() float64
    After : *Decimal.Float64() (float64, bool)

Removed:

    *Decimal.FloatString(uint) string



