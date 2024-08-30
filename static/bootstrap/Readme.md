# Readme

* download Bootstrap source: `https://github.com/twbs/bootstrap/archive/v5.2.3.zip`
* extract `bootstrap-5.2.3/scss`
* modify `scss/_tables.scss`
  * `.table {` => `table, .table {`
* run `sassc custom.scss | python -m rcssmin -b > ../bootstrap.min.css`
