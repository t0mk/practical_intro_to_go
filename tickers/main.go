package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ExchangeTicker struct {
	Getter       TickerGetter
	ExchangeName string
	Ticker       string
}

var btcUsdtTickers = []ExchangeTicker{
	{CoinbaseGetter, "Coinbase", "BTC-USDT"},
	{BinanceGetter, "Binance", "BTCUSDT"},
	{KrakenGetter, "Kraken", "XBTUSDT"},
	{BitstampGetter, "Bitstamp", "btcusdt"},
	{HuobiGetter, "Huobi", "btcusdt"},
	{FTXGetter, "FTX", "BTC/USDT"},
	{KUCoinGetter, "KUCoin", "BTC-USDT"},
	{GateIOGetter, "GateIO", "BTC_USDT"},
	{BitfinexGetter, "Bitfinex", "btcust"},
}

func getExchangeTicker(et ExchangeTicker) float64 {
	start := time.Now()
	price, err := et.Getter(et.Ticker)
	if err != nil {
		fmt.Println("error getting:", et.ExchangeName, et.Ticker, err)
	}
	elapsed := time.Since(start)
	fmt.Printf("[%10s]\t[%10s]\t[%5s]\t%.2f\n",
		et.ExchangeName, et.Ticker, elapsed, price)
	return price
}

func getExchangeTickerAsync(et ExchangeTicker, channel chan float64) {
	price := getExchangeTicker(et)
	channel <- price
}

func floatSliceMean(slice []float64) float64 {
	sum := 0.0
	i := 0
	for _, v := range slice {
		if v == 0 {
			continue
		}
		sum += v
		i++
	}
	return sum / float64(i)
}

func getBtcUsdt() float64 {
	prices := []float64{}
	for _, et := range btcUsdtTickers {
		price := getExchangeTicker(et)
		prices = append(prices, price)
	}
	return floatSliceMean(prices)
}

func getBtcUsdtAsync() float64 {
	prices := []float64{}
	resultChan := make(chan float64)
	for _, et := range btcUsdtTickers {
		go getExchangeTickerAsync(et, resultChan)
	}
	for i := 0; i < len(btcUsdtTickers); i++ {
		price := <-resultChan
		prices = append(prices, price)
	}
	return floatSliceMean(prices)
}

func main() {
	if len(os.Args) == 0 {
		log.Fatal("usage: tickers <sync|async|api>")
	}
	cmd := os.Args[1]
	if cmd == "sync" {
		start := time.Now()
		price := getBtcUsdt()
		elapsed := time.Since(start)
		fmt.Printf("[%10s]\t[%10s]\t[%5s]\t%.2f\n",
			"Sync", "BTC-USDT Mean", elapsed, price)
	} else if cmd == "async" {
		start := time.Now()
		price := getBtcUsdtAsync()
		elapsed := time.Since(start)
		fmt.Printf("[%10s]\t[%10s]\t[%5s]\t%.2f\n",
			"Async", "BTC-USDT Mean", elapsed, price)
	} else if cmd == "api" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "BTC-USD: %.2f", getBtcUsdtAsync())
		})
		log.Fatal(http.ListenAndServe(":10000", nil))
	}
}
