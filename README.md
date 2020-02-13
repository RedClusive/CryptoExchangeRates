# Cryptocurrency rates spectator

This service gets relevant information about actual exchange rates of given cryptocurrency pairs from different exchanges and answers `"/get_rates"` http queries which return **_.json_** with saved rates.

Exchange rates is updated every **m** seconds (where **m** is a given by user **_integer_** value from `input.txt`)

**_Supported exchanges for now:_** _Binance_, _Exmo_.

## Example:
```
[
{"pair":"btc_usdt", "exchange":"Binance", "rate":"9362.69000000", "updated":"2020-02-02 12:08:21.099"}, 
{"pair":"btc_usdt", "exchange":"Exmo", "rate":"9346.36554174", "updated":"2020-02-02 12:08:21.044"}
]
```

# Stack

**_Postgresql_** has been used to store information about given cryptocurrency pairs.

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

3. Initialize `input.txt` 

      ![Sample2](https://sun9-32.userapi.com/c205828/v205828442/51021/MroGCQwTVXo.jpg)

      **_!!!_** Format of `input.txt`:

      You must give integer number **m** in the **first** row **before** cryptocurrency pairs.
4. `go run main.go` or you can use `Dockerfile` or `docker-compose.yml` (If you use `docker-compose.yml` database will be created in container automatically; for connection should be used default `database/config.go`
  
# How to try it

Link: https://ccspectator.herokuapp.com/get_rates

