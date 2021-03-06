# Cryptocurrency rates spectator

This service gets relevant information about actual exchange rates of given cryptocurrency pairs from different exchanges and answers `"/get_rates"` http queries which return **_.json_** with saved rates.

Exchange rates is updated every **SleepSec** seconds (where **SleepSec** is a given by user **_integer_** value from `config.toml`)

**_Supported exchanges for now:_** _Binance_, _Exmo_.

## Example:
```
[
 {
  "Pair": "btc_usdt",
  "Exchange": "Binance",
  "Rate": "10249.83000000",
  "Updated": "2020-02-15 14:35:24.315"
 },
 {
  "Pair": "btc_usdt",
  "Exchange": "Exmo",
  "Rate": "10278.33958778",
  "Updated": "2020-02-15 14:35:24.139"
 }
]
```

# Stack

Application uses **_Postgresql_** to store information about given cryptocurrency pairs.

# How to use it

1. Create a postgresql database

2. Initialize environment variables `PORT`, `DBHOST`, `DBPORT`, `DBUSER`, `DBPASSWORD`, `DBNAME`

      Name | Description
      -----|------------
      PORT | Port that app should listen
      DBHOST | database host
      DBPORT | database port (default Postgresql port: 5432)
      DBUSER  | database username
      DBPASSWORD | database password
      DBNAME | database name

3. Initialize `config.toml` 

4. `go run main.go` or you can use `Dockerfile` or `docker-compose.yml` (If you use `docker-compose.yml` database will be created in container automatically; for connection should be used default `database/config.go`)
  
# How to try it

Link: https://ccspectator.herokuapp.com/get_rates

