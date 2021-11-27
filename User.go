package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	name string
	mobileNo string
	pincode string
	userid string
	riskPercentage int
	role string
	result string
}


func RegisterUser(name string,mobileNo string,pincode string){
	//var user User
	if checkIfUserExitsWithSameMobile(mobileNo)==1{
		fmt.Println("User with same mobileNo already exists")
		return
	}
	var user User
	user.name=strings.ToLower(name)
	user.mobileNo=mobileNo
	user.pincode=pincode
	user.userid=strconv.Itoa(len(users)+1)
	users=append(users,user)
	fmt.Println("User with id "+ user.userid +"has been added successfully")
}

func checkIfUserExitsWithSameMobile( mobileNo string) int {
	for _,user:=range users{
		if strings.EqualFold(mobileNo,user.mobileNo) {
			return 1
		}
	}
	return 0
}

func checkIfUserExists(userid string) int {
	for _,user:=range users{
		if strings.EqualFold(user.userid,userid){
			return 1
		}
	}
	return 0
}

func checkSymptoms(symptoms []string) int {
	numofSymptoms:=0
	for _,s:=range symptoms{
		if strings.EqualFold(strings.ToLower(s),"fever") || strings.Contains(strings.ToLower(s),"cough")||
			strings.EqualFold(strings.ToLower(s),"cold"){
			numofSymptoms++
		}
	}
	return numofSymptoms
}

func GetUserDetails(userid string) User {
	for _,user:=range users{
		if strings.EqualFold(strings.ToLower(userid),user.userid){
			return user
		}
	}
	return User{}
}

func SelfAssessment(userid string,symptoms []string,travelHistory bool,contactwithCovid bool){
	if checkIfUserExists(userid)==0{
		fmt.Println("Sorry No user exists with such userid")
	}
	noofSymptoms:=checkSymptoms(symptoms)
	user:= GetUserDetails(userid)
	if noofSymptoms==0 && travelHistory==false && contactwithCovid==false{
		user.riskPercentage=5
	}
	if noofSymptoms==1 && travelHistory==true && contactwithCovid==true{
		user.riskPercentage=50
	}
	if noofSymptoms==2 && travelHistory==true && contactwithCovid==true{
		user.riskPercentage=75
	}
	if noofSymptoms>2 && travelHistory==true && contactwithCovid==true{
		user.riskPercentage=95
	}
	if user.riskPercentage>=50 {
		//Update result in the respective zones
		zones[user.pincode] = zones[user.pincode] + 1
	}
	fmt.Println("Risk Percentage= ",user.riskPercentage)
}