# Price Updater

## Description
This project, **Price Updater**, is a Go application designed to fetch and update cryptocurrency prices using the CoinGecko API. It stores the price data in Redis and provides endpoints to retrieve current prices and historical price changes.

## Getting Started

### Prerequisites
- Go (developed with version 1.21.6)
- Redis server running locally or accessible remotely
- `.env` file with necessary environment variables (e.g., `COINGECKO_API_KEY`)

### Installation
```bash
git clone https://github.com/TropicalDog17/price-updater.git
`cd price-updater
go mod tidy
```

```bash
# Start the server using Makefile
make run
# Alternatively, run directly with Go
go run main.go
```


## Usage
The server starts on port 8081 and provides the following endpoints:
- `GET /price?symbol=<symbol>`: Fetch the current price for a given cryptocurrency symbol.
- `GET /change?symbol=<symbol>`: Fetch the 24-hour price change for a given cryptocurrency symbol.

Replace `<symbol>` with the cryptocurrency symbol you are interested in (e.g., `bitcoin`, `ethereum`).