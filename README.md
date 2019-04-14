# Foreign Currency Excercise

## Setup with Docker
1. Change `config.example.json` to `config.json`
2. `config.example.json` has default configuration for deployment in docker, so no need to change
3. Run docker compose command
```
$ docker-compose up --build
```

## Setup without Docker
1. Change `config.example.json` to `config.json`, and change the configuration
2. Make sure you have mysql running on your machine
3. Make sure this repo placed in proper $GOPATH
4. Run command
```
$ make build
$ make start
```

## Testing (without Docker only)
1. Make sure you have mysql running on your machine
2. Run command
```
$ make build
$ make test
```

## Database Design

<div style="text-align:center"><img src ="docs/img/db.png" /></div>
<p style="text-align: center;">Database design</p>

It uses 2 tables, currencies and rates. It has cardinality one to many, since one currency can has many daily rates.
Currencies has special attribute, `"tracked"` and `"tracked_rev"`. These attribute uses for marking currency to be tracked. Currency always store unique pair based on attribute `"from"` and `"to"`. It guarantee by the app that `"from"` is always less than `"to"` lexicographically. If we want to track in reverse order (`"from"` is greater than `"to"`), tracked_rev should be marked as true. This can reduce database storage, since we only store 1 tuple for the same symbols instead of 2.

Rates uses 4 attributes, `"id"` as primary key, `"date"`, `"rate_value"`, and `"currency_id`". Rate always store unique tuple based on attribute `"date"` and `"currency_id"`. So, there's no rates of the same pair currency symbol in the same date.

## API Documentation

Could be found in folder docs, `swagger.json` for API Documentation.