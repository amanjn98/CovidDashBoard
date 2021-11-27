package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var zones =make(map[string]int)
var users []User
var adminUsers []User


func main() {
	for {
		fmt.Println("Enter the cmd")
		reader := bufio.NewReader(os.Stdin)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		cmd=strings.ToLower(cmd)
		command:= strings.Split(cmd, " ")
		if strings.Contains(strings.ToLower(cmd),strings.ToLower("Register User")){
			RegisterUser(command[2],command[3],command[4])
		}else if strings.Contains(cmd,"self assessment"){
			SelfAssessment(command[2],[]string{command[3],command[4],command[5]},parseStringstoBool(command[6]),parseStringstoBool(command[7]))
		}else if strings.Contains(cmd,"add admin"){
			RegisterAdmin(command[2],command[3],command[4])
		}else if strings.Contains(cmd,"update covidresult "){
			UpdateCovidResult(command[2],command[3],command[4])
		}else if strings.Contains(cmd,"get zoneinfo"){
			GetNumCasesForZone(command[2])
		} else {
			os.Exit(1)
		}

	}
}
func parseStringstoBool(value string) bool {
	b, err := strconv.ParseBool(value)
	if err == nil {
		return b
	}
	return false
}
