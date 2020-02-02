# Cryptocurrency exchange rates

This service gets relevant information about actual exchange rates of given cryptocurrency pairs from different exchanges and answers "/get_rates" http queries which return .json with saved rates

Supported exchanges: Binance, Exmo

# Stack

Postgresql has been used to store information about given cryptocurrency pairs (pair name, exchange name, actual rate, update time)


