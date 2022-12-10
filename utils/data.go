package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetNFTData() ([]NftDetails, error) {
	client := &http.Client{}

	var nftDetails []NftDetails

	apiData := "https://collections.rarity.tools/collectionsStats"
	req, err := http.NewRequest(http.MethodGet, apiData, nil)
	if err != nil {
		return nil, fmt.Errorf("error making a request %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting a response %w", err)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			return
		}
	}()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(response, &nftDetails); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return nftDetails, nil
}

func GetCryptoData() ([]CoinDetails, error) {
	client := &http.Client{}

	var coinDetails []CoinDetails

	coingecko := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false&price_change_percentage=1h"
	req, err := http.NewRequest(http.MethodGet, coingecko, nil)
	if err != nil {
		return nil, fmt.Errorf("error making a request %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting a response %w", err)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			return
		}
	}()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(response, &coinDetails); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return coinDetails, nil

}

func GetStocksdata() (*StockDetails, error) {
	client := http.Client{}
	var stockDetails StockDetails

	yahooApi := "https://query1.finance.yahoo.com/v8/finance/chart/GME"

	req, err := http.NewRequest(http.MethodGet, yahooApi, nil)
	if err != nil {
		return nil, fmt.Errorf("error making a request %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting a response %w", err)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			return
		}
	}()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(response, &stockDetails); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &stockDetails, nil
}
