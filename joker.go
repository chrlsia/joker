package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"sort"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	nums:=rand.Perm(45)
	
	for i:=0;i<len(nums);i++{
		nums[i] +=1
	}
	
	firstSixNumbers:=nums[0:5]
	sort.Ints(firstSixNumbers)
	joker:=0
	
	for i:=6;i<len(nums)-1;i++{
		if nums[i]<=20{
			joker=nums[i]
			break
		}
	}
	str:="\nΟι πρώτοι σου ταξινομημένοι 6 τυχεροί αριθμοί είναι: "
	
	for i:=0;i<5;i++{
		str +=strconv.Itoa(firstSixNumbers[i])
		if i!=4{
			str +="-"
		}
	}
	
	fmt.Printf("%s και το τζόκερ είναι: %d\n",str,joker)
}