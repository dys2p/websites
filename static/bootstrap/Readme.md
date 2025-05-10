# Readme

* download Bootstrap source: `https://github.com/twbs/bootstrap/archive/v5.3.6.zip`
* extract `bootstrap-5.3.6/scss`
* modify `scss/_tables.scss`
  * `.table {` => `table, .table {`
* run `sassc custom.scss | python -m rcssmin -b > ../bootstrap.min.css`
