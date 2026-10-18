package main

import (
	"fmt"
	"math"
	"time"
)

const (
	pricePerKm     = 10.0
	pricePerMinute = 2.0
)

type WeatherCondiiton int

const (
	Clear WeatherCondiiton = iota
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondiiton
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	multiplier := 1.0

	if weather.Condition == HeavyRain {
		multiplier += 0.2
	} else if weather.Condition == Snow {
		multiplier += 0.15
	} else if weather.Condition == Rain {
		multiplier += 0.125
	}

	if weather.WindSpeed > 15 {
		multiplier += 0.1
	}

	return multiplier
}

type TripParameters struct {
	Distance float64
	Duration float64
}

func CalculateBasePrice(t TripParameters) float64 {
	return t.Distance*pricePerKm + t.Duration*pricePerMinute
}

func GetTimeMultiplier(t time.Time) float64 {
	hour := t.Hour()
	isWeekend := t.Weekday() == time.Saturday || t.Weekday() == time.Sunday

	switch {
	case hour >= 0 && hour < 5:
		return 1.5
	case hour >= 7 && hour < 10 && !isWeekend:
		return 1.3
	case isWeekend:
		return 1.2
	default:
		return 1.0
	}
}

type TrafficClient interface {
	GetTrafficLevel(lat, lng float64) int
}

func GetTrafficMultiplier(trafficLevel int) float64 {
	return 1.0 + float64(trafficLevel-1)*0.1
}

type PriceCalculator struct {
	TrafficClient TrafficClient
}

type RealTrafficClient struct{}

func (c *RealTrafficClient) GetTrafficLevel(lat, lng float64) int {
	return 3
}

const (
	minPrice = 99.0
	maxPrice = 20000.0
)

func ApplyPriceLimits(price float64) float64 {
	switch {
	case price < minPrice:
		return minPrice
	case price > maxPrice:
		return maxPrice
	default:
		return price
	}
}

func (c *PriceCalculator) CalculatePrice(trip TripParameters, now time.Time, weather WeatherData, lat, lng float64) float64 {
	base := CalculateBasePrice(trip)
	timeMul := GetTimeMultiplier(now)
	weatherMul := GetWeatherMultiplier(weather)
	trafficMul := GetTrafficMultiplier(c.TrafficClient.GetTrafficLevel(lat, lng))

	finalPrice := base * timeMul * weatherMul * trafficMul

	return ApplyPriceLimits(math.Round(finalPrice))
}

func main() {
	calculator := PriceCalculator{
		TrafficClient: &RealTrafficClient{}, // В продакшене используется настоящий клиент, а мы подключим структуру-заглушку для имитации его работы
	}

	price := calculator.CalculatePrice(
		TripParameters{Distance: 108.5, Duration: 120},
		time.Now(),
		WeatherData{Snow, 10},
		55.751244, 37.618423,
	)

	fmt.Printf("Ваша цена: %.0f руб.\n", price)
}
