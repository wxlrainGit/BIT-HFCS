/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"encoding/hex"
	"fmt"
	"strings"
	"math/rand"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Human struct {
	
	ID            string `json:"id"`
	Sex           string `json:"sex"`
	Name          string `json:"name"`
	FatherID      string `json:"fatehrid"`
	MotherID      string `json:"motherid"`
	SpouseID      string `json:"spouseid"`
	ChildID  [10] string `json:"childid"`
	NewChild [10] string `json:"newchild"`
}

type Birth struct {

	BirthID      string `json:"birthid"`
	Date         string `json:"date"`
	Sex          string `json:"sex"`
	FatherID     string `json:"fatherid"`
	MotherID     string `json:"motherid"`
	HosptialID   string `json:"hosptialid"`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryID" {
		return s.queryID(APIstub,args)
	}else if function == "createBirth" {
		return s.createBirth(APIstub, args)
	}else if function == "createHuman" {
		return s.createHuman(APIstub, args)
	}else if function == "marry" {
		return s.marry(APIstub, args)
	}else if function == "divorce" {
		return s.divorce(APIstub, args)
	}else if function == "initLedger" {
		return s.initLedger(APIstub)
	}
	return shim.Error("Invalid Smart Contract function name.")
	
}	

func (s *SmartContract) queryID(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	humanAsBytes, err := APIstub.GetState(args[0])
	var human Human;
	err = json.Unmarshal(humanAsBytes,&human)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(humanAsBytes)+ "\" to Human}")
	}
   return shim.Success(humanAsBytes)
	
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	var humanA Human
	humanA.ID       = "1111111"
	humanA.Sex      = "male"
	humanA.FatherID = "0000000"
	humanA.MotherID = "0000001"
	humanA.ChildID[0] = "0"


	var humanB Human
	humanB.ID       = "1111112"
	humanB.Sex      = "female"
	humanB.FatherID = "0000002"
	humanB.MotherID = "0000003"
	humanB.ChildID[0] = "0"

	
	humanAsBytes, _ := json.Marshal(humanA)
	APIstub.PutState(humanA.ID, humanAsBytes)

	humanBAsBytes, _ := json.Marshal(humanB)
	APIstub.PutState(humanB.ID, humanBAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) createBirth(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//5 paramtes father,mother,childsex,birhdate,hospitalID
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	//whether father is sxisted
	FatherAsBytes, err := APIstub.GetState(args[0])
	var father Human;
	err = json.Unmarshal(FatherAsBytes,&father)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(FatherAsBytes)+ "\" to Human}")
	}
	//whether mother is sxisted
	MotherAsBytes, err := APIstub.GetState(args[1])
	var mother Human;
	err = json.Unmarshal(MotherAsBytes,&mother)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(MotherAsBytes)+ "\" to Human}")
	}

	//whether they are couples
	if 0  != (strings.Compare(father.SpouseID,mother.ID)){
		return shim.Error("{\"Error\":\"They are not couples ")
	}
	if 0  != (strings.Compare(mother.SpouseID,father.ID)){
		return shim.Error("{\"Error\":\"They are not couples ")
	}

	//whether more children
	fnum,err := strconv.Atoi(father.ChildID[0])
	if fnum > 2{
		return shim.Error("{\"Error\":\"They are have enough children")
	}
	mnum,err := strconv.Atoi(mother.ChildID[0])
	if mnum > 2{
		return shim.Error("{\"Error\":\"They are have enough children")
	}
	//create birth certs
	var birth Birth;
	//timestamp := time.Now().Unix()
	//tm := time.Unix(timestamp, 0)
	rd := strconv.Itoa(rand.Intn(100))
	str := strings.Join([]string{args[3],rd},"")
	hashstr := hex.EncodeToString([]byte(str))

	birth.BirthID  = hashstr[0:18]
	birth.Sex      = args[2]
	birth.Date     = args[3]
	birth.FatherID = father.ID
	birth.MotherID = mother.ID
	birth.HosptialID = args[4]
	birthAsBytes, _ := json.Marshal(birth)
	APIstub.PutState(birth.BirthID, birthAsBytes)

	//connected to the parents
	father.NewChild[0] = hashstr[0:18]
	fatherAsBytes, _ := json.Marshal(father)
	APIstub.PutState(father.ID, fatherAsBytes)
	mother.NewChild[0] = hashstr[0:18]
	motherAsBytes, _ := json.Marshal(mother)
	APIstub.PutState(mother.ID, motherAsBytes)
	return shim.Success(nil)
}

func (s *SmartContract) createHuman(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//3 paramtes father or motherID ,1flow father 2 flow mother,name
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	//whether father is sxisted
	FatherAsBytes, err := APIstub.GetState(args[0])
	var father Human;
	err = json.Unmarshal(FatherAsBytes,&father)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(FatherAsBytes)+ "\" to Human}")
	}
	//whether mother is sxisted
	MotherAsBytes, err := APIstub.GetState(father.SpouseID)
	var mother Human;
	err = json.Unmarshal(MotherAsBytes,&mother)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(MotherAsBytes)+ "\" to Human}")
	}
	//get the child birth cert
	ChildAsBytes, err := APIstub.GetState(father.NewChild[0])
	var child Birth
	err = json.Unmarshal(ChildAsBytes,&child)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(ChildAsBytes)+ "\" to Human}")
	}

	//whether they are couples
	if 0  != (strings.Compare(father.SpouseID,mother.ID)){
		return shim.Error("{\"Error\":\"They are not couples ")
	}
	if 0  != (strings.Compare(mother.SpouseID,father.ID)){
		return shim.Error("{\"Error\":\"They are not couples ")
	}

	//whether more children
	fnum,err := strconv.Atoi(father.ChildID[0])
	if fnum > 2{
		return shim.Error("{\"Error\":\"They are have enough children")
	}
	mnum,err := strconv.Atoi(mother.ChildID[0])
	if mnum > 2{
		return shim.Error("{\"Error\":\"They are have enough children")
	}
	
	//whether the same child
	if 0  != (strings.Compare(father.ChildID[0],mother.ChildID[0])){
		return shim.Error("{\"Error\":\"They are not the same children")
	}

	//create new human
	var newhuman Human
	newhuman.Sex      = child.Sex
	newhuman.Name     = args[2]
	newhuman.FatherID = father.ID
	newhuman.MotherID = mother.ID
	newhuman.ChildID[0] = "0"
	
	if 0 == (strings.Compare("1",args[1])){
		address := father.ID[0:6]
		date := child.Date
		if 0 == (strings.Compare("male",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,"123","6"},"")
		}
		if 0 != (strings.Compare("male",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,"122","5"},"")
		}
	}

	if 0 == (strings.Compare("2",args[1])){
		address := mother.ID[0:6]
		date := child.Date
		if 0 == (strings.Compare("male",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,"123","6"},"")
		}
		if 0 != (strings.Compare("male",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,"122","5"},"")
		}
	}
	newhumanAsBytes, _ := json.Marshal(newhuman)
	APIstub.PutState(newhuman.ID, newhumanAsBytes)


	//become father
	fchild := strconv.Itoa(fnum+1)
	father.ChildID[0] = fchild
	father.ChildID[fnum+1] = newhuman.ID
	fatherAsBytes, _ := json.Marshal(father)
	APIstub.PutState(father.ID, fatherAsBytes)
	//become mother
	mchild := strconv.Itoa(mnum+1)
	mother.ChildID[0] = mchild
	mother.ChildID[mnum+1] = newhuman.ID
	motherAsBytes, _ := json.Marshal(mother)
	APIstub.PutState(mother.ID, motherAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) marry(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//2paramtes husbandID wifeID
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	//whether husband is exitd
	husbandAsBytes, err := APIstub.GetState(args[0])
	var husband Human;
	err = json.Unmarshal(husbandAsBytes,&husband)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(husbandAsBytes)+ "\" to Human}")
	}
	//whether wife is exitd
	wifeAsBytes, err := APIstub.GetState(args[1])
	var wife Human;
	err = json.Unmarshal(wifeAsBytes,&wife)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(wifeAsBytes)+ "\" to Human}")
	}
	//whether married
	if 0 != len(husband.SpouseID) {
		return shim.Error("{\"Error\":\"Failed to married")
	}
	if 0 != len(wife.SpouseID) {
		return shim.Error("{\"Error\":\"Failed to married")
	}
	//become husband
	husband.SpouseID = wife.ID
	husbandAsBytes, _ = json.Marshal(husband)
	APIstub.PutState(args[0], husbandAsBytes)
	//become wife
	wife.SpouseID = husband.ID
	wifeAsBytes, _ = json.Marshal(wife)
	APIstub.PutState(args[1], wifeAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) divorce(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	//whether husband is exitd
	husbandAsBytes, err := APIstub.GetState(args[0])
	var husband Human;
	err = json.Unmarshal(husbandAsBytes,&husband)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(husbandAsBytes)+ "\" to Human}")
	}
	//whether wife is exitd
	wifeAsBytes, err := APIstub.GetState(args[1])
	var wife Human;
	err = json.Unmarshal(wifeAsBytes,&wife)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(wifeAsBytes)+ "\" to Human}")
	}
	//whether they are couples
	if 0  != (strings.Compare(husband.SpouseID,wife.ID)){
		return shim.Error("{\"Error\":\"They are not couples ")
	}
	if 0  != (strings.Compare(wife.SpouseID,husband.ID)){
		return shim.Error("{\"Error\":\"They are not couples ")
	}
	//change husband spouse
	husband.SpouseID = ""
	husbandAsBytes, _ = json.Marshal(husband)
	APIstub.PutState(args[0], husbandAsBytes)
	//change wife spouse
	wife.SpouseID = ""
	wifeAsBytes, _ = json.Marshal(wife)
	APIstub.PutState(args[1], wifeAsBytes)

	return shim.Success(nil)
}
// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}