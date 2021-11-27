package main

import (
	"fmt"
	"strconv"
	"strings"
)
func RegisterAdmin(name string,mobileNo string,pincode string){
	//var user User
	if CheckIfAdminUserExitsWithSameMobile(mobileNo)==1{
		fmt.Println("User with same mobileNo already exists")
		return
	}
	var user User
	user.name=strings.ToLower(name)
	user.mobileNo=mobileNo
	user.pincode=pincode
	user.userid=strconv.Itoa(len(adminUsers)+1)
	adminUsers=append(adminUsers,user)
	fmt.Println("Admin User with id "+ user.userid +"has been added successfully")
}

func CheckIfAdminUserExitsWithSameMobile( mobileNo string) int {
	for _,user:=range adminUsers{
		if strings.EqualFold(mobileNo,user.mobileNo) {
			return 1
		}
	}
	return 0
}
func checkIfAdminExists(userid string) int {
	for _,user:=range adminUsers{
		if strings.EqualFold(user.userid,userid){
			return 1
		}
	}
	return 0
}


func UpdateCovidResult(userid string,adminid string,result string){
	if checkIfAdminExists(adminid)==0{
		fmt.Println("Sorry No admin exists with such id")
	}
	if checkIfUserExists(userid)==0{
		fmt.Println("Sorry No user exists with such userid")
	}
	user:=GetUserDetails(userid)
	user.result=strings.ToLower(result)
	//Update num of cases in zones
	if strings.EqualFold(user.result,"positive"){
		zones[user.pincode]=zones[user.pincode]+1
	}
	fmt.Println("Updated covid results successfully")
}