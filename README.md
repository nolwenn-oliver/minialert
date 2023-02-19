# minialert
Simple alerting system in Golang


## Build

```make build```

## Run
**Launch in server mode**

```./minialert --server```

You can optionally specify the port with the flag
```--port=8080``` (default is 8080).


**Launch in client mode**

Send metric

```./minialert --client send cpu 90```

List alerts

```./minialert --client alerts-list```

## Test

````make test````

## Implementation choices
For simplicity purposes, I considered alerts trigger only when value is greater than threshold.

## Next steps

Non-exhaustive list of enhancements to make:
- CLI: add exclusion rules between flags
- Improve test coverage
- Improve HTTP status code & headers returned by server
- Pass config file as parameter

