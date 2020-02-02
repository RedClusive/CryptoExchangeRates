# Cryptocurrency exchange rates

This service gets relevant information about actual exchange rates of given cryptocurrency pairs from different exchanges and answers **_"/get_rates"_** http queries which return **_.json_** with saved rates.

**_Supported exchanges for now:_** _Binance_, _Exmo_.

## Example:
```
[{"pair":"btc_usdt", "exchange":"Binance", "rate":"9362.69000000", "updated":"2020-02-02 12:08:21.099"}, {"pair":"btc_usdt", "exchange":"Exmo", "rate":"9346.36554174", "updated":"2020-02-02 12:08:21.044"}]
```

# Stack

**_Postgresql_** has been used to store information about given cryptocurrency pairs (pair name, exchange name, actual rate, update time).


