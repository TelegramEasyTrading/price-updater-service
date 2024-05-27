package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DataProvider interface {
	FetchPrice(symbols string) (CoinGeckoResponse, error)
}
type CoinGeckoProvider struct {
	key string
}

var _ DataProvider = &CoinGeckoProvider{}

func NewCoinGeckoProvider(key string) *CoinGeckoProvider {
	return &CoinGeckoProvider{
		key: key,
	}
}

// symbols is a comma separated list of symbols
func (c *CoinGeckoProvider) FetchPrice(symbols string) (CoinGeckoResponse, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%s&price_change_percentage=1h%%2C24h%%2C7d", symbols)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("x-cg-demo-api-key", c.key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	parsedResponse := CoinGeckoResponse{}
	err = json.Unmarshal(body, &parsedResponse)
	if err != nil {
		fmt.Println("Error parsing response body:", err)
		return nil, err
	}
	return parsedResponse, nil
}

// "id": "bitcoin",
// "symbol": "btc",
// "name": "Bitcoin",
// "image": "https://assets.coingecko.com/coins/images/1/large/bitcoin.png?1696501400",
// "current_price": 68628,
// "market_cap": 1353224556363,
// "market_cap_rank": 1,
// "fully_diluted_valuation": 1442216161205,
// "total_volume": 13727087386,
// "high_24h": 69441,
// "low_24h": 68309,
// "price_change_24h": -657.6060240157676,
// "price_change_percentage_24h": -0.94912,
// "market_cap_change_24h": -11658552583.64795,
// "market_cap_change_percentage_24h": -0.85418,
// "circulating_supply": 19704200,
// "total_supply": 21000000,
// "max_supply": 21000000,
// "ath": 73738,
// "ath_change_percentage": -7.12183,
// "ath_date": "2024-03-14T07:10:36.635Z",
// "atl": 67.81,
// "atl_change_percentage": 100899.04566,
// "atl_date": "2013-07-06T00:00:00.000Z",
// "roi": null,
// "last_updated": "2024-05-27T07:54:28.977Z",
// "price_change_percentage_1h_in_currency": 0.33471027081093724,
// "price_change_percentage_24h_in_currency": -0.9491175807860327,
// "price_change_percentage_7d_in_currency": 3.622635641041768
type CurrencyData struct {
	Symbol       string  `json:"symbol"`
	CurrentPrice float64 `json:"current_price"`
	Change1h     float64 `json:"price_change_percentage_1h_in_currency"`
	Change24h    float64 `json:"price_change_percentage_24h_in_currency"`
	Change7d     float64 `json:"price_change_percentage_7d_in_currency"`
	High24h      float64 `json:"high_24h"`
	Low24h       float64 `json:"low_24h"`
}

type CoinGeckoResponse []CurrencyData
