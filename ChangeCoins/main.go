package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input []string
	input = strings.Split("15.94;16.00", ";")

	coins := make(map[float32]string)
	coins[.01] = "PENNY"
	coins[.05] = "NICKEL"
	coins[.10] = "DIME"
	coins[.25] = "QUARTER"
	coins[.50] = "HALF DOLLAR"
	coins[1.00] = "ONE"
	coins[2.00] = "TWO"
	coins[5.00] = "FIVE"
	coins[10.00] = "TEN"
	coins[20.00] = "TWENTY"
	coins[50.00] = "FIFTY"
	coins[100.00] = "ONE HUNDRED"

	coinsArr := make([]float32, 0)
	coinsArr = append(coinsArr, .01, .05, .10, .25, .50, 1.00, 2.00, 5.00, 10.00, 20.00, 50.00, 100.00)

	PP, _ := strconv.ParseFloat(input[0], 32)
	CH, _ := strconv.ParseFloat(input[1], 32)
	fmt.Println(returnChange(float32(CH), float32(PP), coins, coinsArr))
}

func returnChange(CH float32, PP float32, coins map[float32]string, coinsArr []float32) string {
	if CH < PP {
		return "ERROR"
	}

	if CH == PP {
		return "ZERO"
	}

	balance := CH - PP
	if value, isExist := coins[balance]; isExist {
		return value
	}

	return getChange(balance, coins, coinsArr)
}

func getChange(balance float32, coins map[float32]string, coinsArr []float32) string {
	change := make([]string, 0)
	for balance != 0.0 {
		var max float32 = 0.0
		for _, coin := range coinsArr {
			if coin <= balance {
				max = coin
				continue
			}
			break
		}

		change = append(change, coins[max])
		balance = balance - max
		balance = float32(math.Floor(float64(balance*100)) / 100)
	}

	sort.Strings(change)
	return strings.Join(change, ",")
}
