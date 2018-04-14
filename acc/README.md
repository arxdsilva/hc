# acc
this cli tool reads 2 csv files, assuming that the first contains unique accounts and It's current states and transactions that will build the account's balance as first and second arguments. Ex:

`$ acc contas.csv transacoes.csv`

# requirements
- Go 1.10.1

# Usage
### Clone
`$ go get github.com/arxdsilva/hc`

### Enter folder
`$ cd $GOPATH/github.com/arxdsilva/hc/acc`

### build
`$ go install`

### Use
`$ acc contas.csv transacoes.csv`

## assumptions
- If a transaction has a account that is not described in `contas.csv` file, It'll be ignored.

- you're ordering the inputs like this:
`$ acc contas.csv transacoes.csv`
