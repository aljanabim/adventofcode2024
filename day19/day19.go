package day19

import (
	"fmt"
	"math"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

/*
r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb

r, wr, b, g, bwu, rb, gb, br
1, 21, 3, 4, 325, 13, 43, 31

r=1
w=2
b=3
g=4
u=5

31211 = .*10^5 + .*10^4 + .*10^3 + .*10^2 + .*10^1 + .*10^0

brwrr
31211

..w.. remove 1 len
....r remove 2 len

bggr
gbbr
rrbgbr
ubwu
5325
bwurrg
brgr
bbrgwb


want 123
has 0 12 23 3

should pick 12*10 + 3*1

order -1
=> return 0

order 0, want 12[3]
=> stripe*1+0, try 23

order 1, want 1[23]
=> stripe*10+23, try 0

order 2, want [123]
=> stripe*100+23


== opposite direction ==
want: 312
has: 0 3 31 12 2
should pick: 3*100 + 12*1

order > 2
=> 0

order 2, want 300
=> stripe*100 + 0 : pick 3 return 300

order 1, want 310
=> stripe*10 + 300 : pick none return 300

order 0, want 312
=> stripe*1 + 300 : pick 12 return 312

done

== opposite direction ==
want: 312
has: 0 31 30 12
should pick: 30*10 + 12*1

order > 2
=> 0

order 2, want 300
=> stripe*100 + 0 : pick none return 0

order 1, want 310
=> stripe*10 + 0 : pick 31 return 310

order 0, want 312
=> stripe*1 + 31 : pick none return 312

done

== new approach ==
want: 312
has: 0 22 3 12
should pick: 3*100 + 12*1

top-down analysis

order=0, val=0
=> (stipe*1+val % 10) / 1 == 2 : pick 22
val += stripe(=22)
add = call(order=1,val=22)
if add != -1 {
	return val + add
}

order=1, val=22
=> (stripe*10+val % 100) / 10 == 1 : return -1

order=0, val=0
=> (stipe*1+val % 10) / 1 == 2 : pick 12
val += stripe(=12)*1
add = call(order=1,val=12)
if add != -1 {
	out+=add
}

order=1, val=12
=> (stripe*10+12 % 100) / 10 == 1 : pick 0
val += stripe(=0)*10
add = call(order=1,val=12+0)
if add != -1 {
	out+=add
}

order=2, val=12
=> (stripe*100+(12+1) % 1000) / 100 == 3 : pick 3
val=
add = val + stipe*100 = 312
add = call(order=3,val=312)

order=3

*/

func getDigit(num int, order int) int {
	return num % int(math.Pow10(order+1)) / int(math.Pow10(order))
}

func getOrder(num int) int {
	order := 0
	for num > 0 {
		num /= 10
		order++
	}
	return order
}

func helper(stripes []int, want, order, maxOrder, val int) int {
	if order == maxOrder {
		return 0
	}
	// Only check available stripes if we haven't already got a stripe at this order
	if getOrder(val) == order+1 {
		// fmt.Println("Already has the right order. Skipping order", order)
		if getDigit(val, order) == getDigit(want, order) {
			add := helper(stripes, want, order+1, maxOrder, val)
			if add != -1 {
				// fmt.Println("==== Returning", val+add, "at order", order)
				return add
			}
		} else {
			return -1
		}

	} else {
		for _, stripe := range stripes {
			// fmt.Println("= Trying stripe", stripe, "at order", order)
			stripePow := stripe * int(math.Pow10(order))
			if getDigit(stripePow+val, order) == getDigit(want, order) && getOrder(stripePow) <= maxOrder {
				// fmt.Println("== Stripe", stripe, "works. Going deeper.")
				add := helper(stripes, want, order+1, maxOrder, val+stripePow)
				if add != -1 {
					// fmt.Println("==== Returning", val+stripePow+add, "at order", order)
					return stripePow + add
				}
				// fmt.Println("xx Oopps", stripe, "was bad further down the line")
			}
		}
	}
	return -1
}

func parseStripes(stripesStr string) ([]int, map[rune]int) {
	stripes := strings.Split(stripesStr, ",")
	color2Int := map[rune]int{}
	stripesNum := []int{0}
	for _, stripe := range stripes {
		stripe = strings.TrimSpace(stripe)
		num := 0
		for i, c := range stripe {
			if _, ok := color2Int[c]; !ok {
				color2Int[c] = len(color2Int) + 1
			}
			num += color2Int[c] * int(math.Pow10(len(stripe)-i-1))
		}
		stripesNum = append(stripesNum, num)
	}
	return stripesNum, color2Int
}

func parseFlags(flags []string, color2Int map[rune]int) []int {
	flagNums := []int{}
	for _, flag := range flags {

		flag = strings.TrimSpace(flag)
		num := 0
		for i, c := range flag {
			num += color2Int[c] * int(math.Pow10(len(flag)-i-1))
		}
		flagNums = append(flagNums, num)
	}
	return flagNums
}

func solvePart1(stripes []int, flags []int) int {
	count := 0
	for _, flag := range flags {
		maxOrder := getOrder(flag)
		res := helper(stripes, flag, 0, maxOrder, 0)
		if res != -1 {
			count++
		}
	}
	return count
}
func Solve() {
	lines, err := utils.ReadLines("day19/input.txt")
	if err != nil {
		panic(err)
	}
	stripesStr := lines[0]
	flags := lines[2:]
	stripes, color2Int := parseStripes(stripesStr)
	flagNums := parseFlags(flags, color2Int)
	fmt.Println(color2Int)
	fmt.Println(stripes)
	for _, flag := range flagNums {
		fmt.Println(flag)
	}

	res := solvePart1(stripes, flagNums)
	// too low - 309
	utils.PrintSolution(19, 1, res)
}
