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
//human ID
type Human struct {
	
	ID            string `json:"身份证号"`
	Sex           string `json:"性别"`
	Name          string `json:"姓名"`
	FatherID      string `json:"父亲"`
	MotherID      string `json:"母亲"`
	SpouseID      string `json:"配偶"`
	Marry_Cert    string `json:"结婚证书"`
	ChildID  [10] string `json:"子女"`
	NewChild [10] string `json:"子女出生证明"`
}

//birth cert
type Birth struct {

	BirthID      string `json:"出生证书编号"`
	Date         string `json:"出生日期"`
	Sex          string `json:"性别"`
	FatherID     string `json:"父亲"`
	MotherID     string `json:"母亲"`
	HosptialID   string `json:"接生机构"`
}

//marry card
type Marry_Card struct {

	Marry_Cert     string `json:"结婚证书编号"`
	Husband_ID     string `json:"丈夫"`
	Wife_ID        string `json:"妻子"`
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
	humanA.ID       = "110105199409026676"
	humanA.Sex      = "male"
	humanA.Name     = "李雷雷"
	humanA.FatherID = "110105197003025376"
	humanA.MotherID = "110105197302055386"
	humanA.ChildID[0] = "0"
 	humanA.NewChild[0] = "0"


	var humanB Human
	humanB.ID       = "110105199409026686"
	humanB.Sex      = "female"
	humanB.Name     = "韩梅梅"
	humanB.FatherID = "110105197107025376"
	humanB.MotherID = "110105197303055386"
	humanB.ChildID[0] = "0"
	humanB.NewChild[0] = "0"

	var humanC Human
	humanC.ID       = "110105199409026656"
	humanC.Sex      = "male"
	humanC.Name     = "王雷雷"
	humanC.FatherID = "110105197003025376"
	humanC.MotherID = "110105197302055386"
	humanC.ChildID[0] = "0"
	humanC.NewChild[0] = "0"


	var humanD Human
	humanD.ID       = "110105199409026646"
	humanD.Sex      = "female"
	humanD.Name     = "张梅梅"
	humanD.FatherID = "110105197107025376"
	humanD.MotherID = "110105197303055386"
	humanD.ChildID[0] = "0"
	humanD.NewChild[0] = "0"

	
	humanAsBytes, _ := json.Marshal(humanA)
	APIstub.PutState(humanA.ID, humanAsBytes)

	humanBAsBytes, _ := json.Marshal(humanB)
	APIstub.PutState(humanB.ID, humanBAsBytes)

	humanCsBytes, _ := json.Marshal(humanC)
	APIstub.PutState(humanC.ID, humanCsBytes)

	humanDAsBytes, _ := json.Marshal(humanD)
	APIstub.PutState(humanD.ID, humanDAsBytes)

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
		return shim.Error("{\"Error\":\"They are not couples }")
	}
	if 0  != (strings.Compare(mother.SpouseID,father.ID)){
		return shim.Error("{\"Error\":\"They are not couples }")
	}

	// //whether more children
	// fnum,err := strconv.Atoi(father.ChildID[0])
	// if fnum > 2{
	// 	return shim.Error("{\"Error\":\"They are have enough children}")
	// }
	// mnum,err := strconv.Atoi(mother.ChildID[0])
	// if mnum > 2{
	// 	return shim.Error("{\"Error\":\"They are have enough children}")
	// }

	//whether more Birthcerts
	fnum,err := strconv.Atoi(father.NewChild[0])
	if fnum > 1{
		return shim.Error("{\"Error\":\"They are have enough children}")
	}
	mnum,err := strconv.Atoi(mother.ChildID[0])
	if mnum > 1{
		return shim.Error("{\"Error\":\"They are have enough children}")
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
	father.NewChild[0] = strconv.Itoa(fnum+1)
	father.NewChild[fnum+1] = birth.BirthID
	fatherAsBytes, _ := json.Marshal(father)
	APIstub.PutState(father.ID, fatherAsBytes)

	mother.NewChild[0] = strconv.Itoa(mnum+1)
	mother.NewChild[mnum+1] = birth.BirthID
	motherAsBytes, _ := json.Marshal(mother)
	APIstub.PutState(mother.ID, motherAsBytes)
	return shim.Success(birthAsBytes)
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

	//whether they are couples
	if 0  != (strings.Compare(father.SpouseID,mother.ID)){
		return shim.Error("{\"Error\":\"They are not couples }")
	}
	if 0  != (strings.Compare(mother.SpouseID,father.ID)){
		return shim.Error("{\"Error\":\"They are not couples }")
	}

	//whether more children
	fnum,err := strconv.Atoi(father.ChildID[0])
	if fnum > 2{
		return shim.Error("{\"Error\":\"They are have enough children}")
	}
	mnum,err := strconv.Atoi(mother.ChildID[0])
	if mnum > 2{
		return shim.Error("{\"Error\":\"They are have enough children}")
	}
	
	//get the child birth cert
	cnum,err := strconv.Atoi(father.NewChild[0])
	if fnum > cnum{
		return shim.Error("{\"Error\":\"The birth cert over time}")
	}

	ChildAsBytes, err := APIstub.GetState(father.NewChild[cnum])
	var child Birth
	err = json.Unmarshal(ChildAsBytes,&child)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of birth cert}")
	}

	//create new human
	var newhuman Human
	newhuman.Sex      = child.Sex
	newhuman.Name     = args[2]
	newhuman.FatherID = father.ID
	newhuman.MotherID = mother.ID
	newhuman.ChildID[0] = "0"
	newhuman.NewChild[0] = "0"
	
	if 0 == (strings.Compare("1",args[1])){
		address := father.ID[0:6]
		date := child.Date
		if 0 == (strings.Compare("男",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,"123",strconv.Itoa(rand.Intn(9))},"")
		}
		if 0 != (strings.Compare("男",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,"122",strconv.Itoa(rand.Intn(9))},"")
		}
	}

	if 0 == (strings.Compare("2",args[1])){
		address := mother.ID[0:6]
		date := child.Date
		if 0 == (strings.Compare("男",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,"123",strconv.Itoa(rand.Intn(9))},"")
		}
		if 0 != (strings.Compare("男",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,"122",strconv.Itoa(rand.Intn(9))},"")
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

	return shim.Success(newhumanAsBytes)
}

func (s *SmartContract) marry(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//3paramtes husbandID ,wifeID,date
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	//whether husband is exitd
	husbandAsBytes, err := APIstub.GetState(args[0])
	var husband Human;
	err = json.Unmarshal(husbandAsBytes,&husband)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of husband}")
	}
	//whether wife is exitd
	wifeAsBytes, err := APIstub.GetState(args[1])
	var wife Human;
	err = json.Unmarshal(wifeAsBytes,&wife)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of wife}")
	}
	//whether married
	if 0 != len(husband.SpouseID) {
		return shim.Error("{\"Error\":\"Failed to married")
	}
	if 0 != len(wife.SpouseID) {
		return shim.Error("{\"Error\":\"Failed to married")
	}

	//generate marry id
	rd := strconv.Itoa(rand.Intn(100))
	str := strings.Join([]string{args[2],rd},"")
	hashstr := hex.EncodeToString([]byte(str))
	marry_cert_id  := strings.Join([]string{"J110101",args[2],hashstr[0:6]},"-")
	
	//become husband
	husband.SpouseID = wife.ID
	husband.Marry_Cert = marry_cert_id
	husbandAsBytes, _ = json.Marshal(husband)
	APIstub.PutState(args[0], husbandAsBytes)
	//become wife
	wife.SpouseID = husband.ID
	wife.Marry_Cert = marry_cert_id
	wifeAsBytes, _ = json.Marshal(wife)
	APIstub.PutState(args[1], wifeAsBytes)

	var card Marry_Card
	card.Marry_Cert = marry_cert_id
	card.Husband_ID = husband.ID
	card.Wife_ID = wife.ID
	MarryAsBytes, _ := json.Marshal(card)
	APIstub.PutState(card.Marry_Cert, MarryAsBytes)

	return shim.Success(MarryAsBytes)
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