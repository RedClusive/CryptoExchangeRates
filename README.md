# Cryptocurrency exchange rates spectator

This service gets relevant information about actual exchange rates of given cryptocurrency pairs from different exchanges and answers **_"/get_rates"_** http queries which return **_.json_** with saved rates.

Exchange rates is updated every **m** seconds (where **m** is a given by user **_integer_** value from **input.txt**)

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

1. Without docker

  * Create a postgresql database

  * Set environment variables **DBHOST**, **DBPORT**, **DBUSER**, **DBPASSWORD**, **DBNAME** to appropriate you

       **_Alt way without changing env. vars_**: Change connection constants in **database/config.go**

       Sample with constants that appropriate only to my own local database:

      ![Sample1](https://sun9-70.userapi.com/c850416/v850416442/1a877f/Fz5cWGZ1KmU.jpg)

  * Change **input.txt** 

      ![Sample2](https://sun9-32.userapi.com/c205828/v205828442/51021/MroGCQwTVXo.jpg)

      **_!!!_** Format of **input.txt**:

      You must give integer number **m** in the **first** row **before** cryptocurrency pairs.

2. Docker (database will be created in container; for connection should be used default **database/config.go**)
    
  * just "docker-compose up"
  
# How to try it

Link: https://ccspectator.herokuapp.com/get_rates

