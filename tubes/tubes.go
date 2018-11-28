package main

import "fmt"
import "io"

type PriceType struct { date int64; time int64; opening float64; high float64; low float64; closing float64; volume int }

var MarketData[29734] PriceType
var i int64
var date int64
var time int64
var opening float64
var high float64
var low float64
var closing float64
var err error
var ind1, ind2 float64

func main() {
    fmt.Println("PREDICTING/EVALUATING MARKET PRICES WITH HOOK REVERSAL PATTERN")
    i = 0
    for err != io.EOF {
	    fmt.Scanf("%d %d;%v;%v;%v;%v;0", &date, &time, &opening, &high, &low, &closing)
		MarketData[i].date = date
		MarketData[i].time = time
		MarketData[i].opening = opening
		MarketData[i].high = high
		MarketData[i].low = low
		MarketData[i].closing = closing
		MarketData[i].volume = 0
		HookReversal()
		i = i + 1
	}
}

func HookReversal() {
    ind1 = MarketData[i].low + (0.2 * (MarketData[i].high - MarketData[i].low))
	ind2 = MarketData[i].high - (0.2 * (MarketData[i].high - MarketData[i].low))
	if (MarketData[i].opening < ind1) && (MarketData[i].closing > ind2) && (MarketData[i].high < MarketData[i-1].high) && (MarketData[i].low > MarketData[i-1].low) {
	    fmt.Println("Date:", MarketData[i].date)
		fmt.Println("Time:", MarketData[i].time)
		fmt.Println("Close price:", MarketData[i].closing)
		fmt.Println("Indicator 1 value=", ind1)
		fmt.Println("Indicator 2 value=", ind2)
		fmt.Println("***buy")
	}
	if (MarketData[i].opening > ind1) && (MarketData[i].closing < ind2) && (MarketData[i].high < MarketData[i-1].high) && (MarketData[i].low > MarketData[i-1].low) {
	    fmt.Println("Date:", MarketData[i].date)
		fmt.Println("Time:", MarketData[i].time)
		fmt.Println("Close price:", MarketData[i].closing)
		fmt.Println("Indicator 1 value=", ind1)
		fmt.Println("Indicator 2 value=", ind2)
		fmt.Println("***sell")
	}
}

/*Hook Reversals Automatic Analysis by Larry Lovrencic
buy=O<(L+0.2*(H-L)) AND C>(H-0.2*(H-L)) AND H<ref(H,-1) AND L>ref(L,-1);
sell=O>(L+0.8*(H-L)) AND C<(H-0.8*(H-L)) AND H<ref(H,-1) AND L>ref(L,-1);
*/