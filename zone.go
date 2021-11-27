package main

import (
	"fmt"
	"strconv"
)

func GetNumCasesForZone(pincode string){
	var numCases int
	var zoneType string
	val,ok:=zones[pincode]
	if ok{
		if val>5{
			zoneType="ORANGE"
		}else{
			zoneType="RED"
		}
	}else{
		fmt.Println("Sorry No Zones exits with that pincode")
	}
	fmt.Println("Numofcases=",strconv.Itoa(numCases)+ " "+"zoneType"+ zoneType)
}