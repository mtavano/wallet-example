# Wallet example

This repo is a simple wallet example that works with USDC, BTC and ETH.


### Run in your local

This service is a simple poc of a transactional system so no deps are needed.
This project uses a simple implementation of an in-memory data store so you can simple run this
as follows

```
PORT=9000 go run cmd/main.go
```

and then you can use the insomnia collection to try it out.

Also you can simple run one of the available builds into the `bin` folder according your OS.
Only MacOSX and Linux supported.
