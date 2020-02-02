# Cryptocurrency exchange rates

This service gets relevant information about actual exchange rates of given cryptocurrency pairs from different exchanges and answers **_"/get_rates"_** http queries which return **_.json_** with saved rates.

Supported exchanges: **_Binance_**, **_Exmo_**.

# Stack

**_Postgresql_** has been used to store information about given cryptocurrency pairs (pair name, exchange name, actual rate, update time).


