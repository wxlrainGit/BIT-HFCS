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

// Written by Xu Chen Hao
// Building on windows:
// 1. Install cygwin64 with gcc/g++
// 2. Set system Path env to C:\cygwin64\bin
// 3. Install golang v1.9.2 & set GOPATH env to D:\go
// 4. mkdir D:\go\src\sacc\ & put this chaincoe in there
// 5. go get -u --tags nopkcs11 github.com/hyperledger/fabric/core/chaincode/shim
// 6. go build --tags nopkcs11
package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"time"
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
	Date          string `json:"出生日期"`
	FatherName    string `json:"父亲姓名"`
	FatherID      string `json:"父亲身份证号"`
	MotherName    string `json:"母亲姓名"`
	MotherID      string `json:"母亲身份证号"`
	MarryState    string `json:"婚姻状态"`
	SpouseName    string `json:"配偶姓名"`
	SpouseID      string `json:"配偶身份证号"`
	Marry_Cert    string `json:"结婚证书"`
	ChildID  [10] string `json:"子女身份证号"`
	ChildName[10] string `json:"子女姓名"`
	NewChild [10] string `json:"子女出生证明"`
}

//birth cert
type Birth struct {
    
    BirthName    string `json:"新生儿姓名"`
	BirthID      string `json:"出生证书编号"`
	Date         string `json:"出生日期"`
	Sex          string `json:"性别"`
	Weight       string `json:"体重"`
	Health       string `json:"健康情况"`
	Place        string `json:"出生地"`
	FatherName   string `json:"父亲姓名"`
	FatherID     string `json:"父亲身份证号"`
	MotherName   string `json:"母亲姓名"`
	MotherID     string `json:"母亲身份证号"`
	HosptialID   string `json:"接生机构"`
}

//marry card
type Marry_Card struct {

	Marry_Cert     string `json:"证书编号"`
	State          string `json:"状态"`
	Husband_Name   string `json:"丈夫姓名"`
	Husband_ID     string `json:"丈夫身份证号"`
	Wife_Name      string `json:"妻子姓名"`
	Wife_ID        string `json:"妻子身份证号"`
	Date           string `json:"登记日期"`
}

//marry check 
type Marry_Check struct{
	CheckID  		 string `json:"审查编号"`
	Husband_Name     string `json:"丈夫姓名"`
	Husband_ID       string `json:"丈夫身份证号"`
	HusbandState     string `json:"丈夫婚姻状态"`
	Wife_Name        string `json:"妻子姓名"`
	Wife_ID          string `json:"妻子身份证号"`
	WifeState        string `json:"妻子婚姻状态"`
	Check [6]        string `json:"判断结果"`
	CheckStae        string `json:"审查表状态"`
	Marry_Cert       string `json:"结婚证书"`
}

type Creat_Check struct{
	CheckID            string `json:"审查编号"`
	Name               string `json:"姓名"`
	FatherName         string `json:"父亲姓名"`
	FatherID           string `json:"父亲身份证号"`
	MotherName         string `json:"母亲姓名"`
	MotherID           string `json:"母亲身份证号"`
	Marry_Cert         string `json:"父母结婚证书编号"`
	BirthID            string `json:"出生证书编号"`
	BirthDate          string `json:"出生日期"`
	Sex                string `json:"性别"`
	HosptialID         string `json:"接生机构"`
	Check [9]          string `json:"判断结果"`
	CheckStae          string `json:"审查表状态"`
	ID                 string `json:"身份证号"`
}

type Divorce_Check struct{
	CheckID  		 string `json:"审查编号"`
	Husband_Name     string `json:"丈夫姓名"`
	Husband_ID       string `json:"丈夫身份证号"`
	Wife_Name        string `json:"妻子姓名"`
	Wife_ID          string `json:"妻子身份证号"`
	Marry_Cert       string `json:"结婚证书编号"`
	Check [5]        string `json:"判断结果"`
	CheckStae        string `json:"审查表状态"`
}

type R_Err struct{
	Reason  		 string `json:"原因"`
}
/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	var humanA Human
	humanA.ID          = "110105199409026676"
	humanA.Sex         = "男"
	humanA.Name        = "李雷雷"
	humanA.FatherName  = "李父"
	humanA.FatherID    = "110105197003025376"
	humanA.MotherName  = "李母"
	humanA.MotherID    = "110105197302055386"
	humanA.ChildID[0]  = "0"
 	humanA.NewChild[0] = "0"


	var humanB Human
	humanB.ID          = "110105199409026686"
	humanB.Sex         = "女"
	humanB.Name        = "韩梅梅"
	humanB.FatherName  = "韩父"
	humanB.FatherID    = "110105197003025376"
	humanB.MotherName  = "韩母"
	humanB.MotherID    = "110105197302055386"
	humanB.ChildID[0]  = "0"
	humanB.NewChild[0] = "0"

	var humanC Human
	humanC.ID          = "110105199409026616"
	humanC.Sex         = "男"
	humanC.Name        = "王雷雷"
	humanC.FatherName  = "王父"
	humanC.FatherID    = "110105197003025376"
	humanC.MotherName  = "王母"
	humanC.MotherID    = "110105197302055386"
	humanC.ChildID[0]  = "0"
	humanC.NewChild[0] = "0"


	var humanD Human
	humanD.ID          = "110105199409026626"
	humanD.Sex         = "女"
	humanD.Name        = "张梅梅"
	humanD.FatherName  = "张父"
	humanD.FatherID    = "110105197003025376"
	humanD.MotherName  = "张母"
	humanD.MotherID    = "110105197302055386"
	humanD.ChildID[0]  = "0"
	humanD.NewChild[0] = "0"

	var humanE Human
	humanE.ID          = "110105199409026636"
	humanE.Sex         = "男"
	humanE.Name        = "张雷雷"
	humanE.FatherName  = "张父"
	humanE.FatherID    = "110105197003025376"
	humanE.MotherName  = "张母"
	humanE.MotherID    = "110105197302055386"
	humanE.ChildID[0]  = "0"
	humanE.NewChild[0] = "0"

	var humanF Human
	humanF.ID          = "110105199409026646"
	humanF.Sex         = "女"
	humanF.Name        = "宋梅梅"
	humanF.FatherName  = "宋父"
	humanF.FatherID    = "110105197003025376"
	humanF.MotherName  = "宋母"
	humanF.MotherID    = "110105197302055386"
	humanF.ChildID[0]  = "0"
	humanF.NewChild[0] = "0"

	var humanG Human
	humanG.ID          = "110105199409026656"
	humanG.Sex         = "男"
	humanG.Name        = "赵雷雷"
	humanG.FatherName  = "赵父"
	humanG.FatherID    = "110105197003025376"
	humanG.MotherName  = "赵母"
	humanG.MotherID    = "110105197302055386"
	humanG.ChildID[0]  = "0"
	humanG.NewChild[0] = "0"
	
	var humanH Human
	humanH.ID          = "110105199409026666"
	humanH.Sex         = "女"
	humanH.Name        = "孙梅梅"
	humanH.FatherName  = "孙父"
	humanH.FatherID    = "110105197003025376"
	humanH.MotherName  = "孙母"
	humanH.MotherID    = "110105197302055386"
	humanH.ChildID[0]  = "0"
	humanH.NewChild[0] = "0"

	var humanI Human
	humanI.ID          = "110105199409026696"
	humanI.Sex         = "男"
	humanI.Name        = "钱雷雷"
	humanI.FatherName  = "钱父"
	humanI.FatherID    = "110105197003025376"
	humanI.MotherName  = "钱母"
	humanI.MotherID    = "110105197302055386"
	humanI.ChildID[0]  = "0"
	humanI.NewChild[0] = "0"

	var humanK Human
	humanK.ID          = "110105199409026606"
	humanK.Sex         = "女"
	humanK.Name        = "刘梅梅"
	humanK.FatherName  = "刘父"
	humanK.FatherID    = "110105197003025376"
	humanK.MotherName  = "刘母"
	humanK.MotherID    = "110105197302055386"
	humanK.ChildID[0]  = "0"
 	humanK.NewChild[0] = "0"

 	var human1 Human
	human1.ID          = "110105199409026675"
	human1.Sex         = "男"
	human1.Name        = "李二"
	human1.FatherName  = "李父"
	human1.FatherID    = "110105197003025376"
	human1.MotherName  = "李母"
	human1.MotherID    = "110105197302055386"
	human1.ChildID[0]  = "0"
 	human1.NewChild[0] = "0"


	var human2 Human
	human2.ID          = "110105199409026685"
	human2.Sex         = "女"
	human2.Name        = "韩娟"
	human2.FatherName  = "韩父"
	human2.FatherID    = "110105197003025376"
	human2.MotherName  = "韩母"
	human2.MotherID    = "110105197302055386"
	human2.ChildID[0]  = "0"
	human2.NewChild[0] = "0"

	var human3 Human
	human3.ID          = "110105199409026615"
	human3.Sex         = "男"
	human3.Name        = "王雷"
	human3.FatherName  = "王父"
	human3.FatherID    = "110105197003025376"
	human3.MotherName  = "王母"
	human3.MotherID    = "110105197302055386"
	human3.ChildID[0]  = "0"
	human3.NewChild[0] = "0"


	var human4 Human
	human4.ID          = "110105199409026625"
	human4.Sex         = "女"
	human4.Name        = "张梅"
	human4.FatherName  = "张父"
	human4.FatherID    = "110105197003025376"
	human4.MotherName  = "张母"
	human4.MotherID    = "110105197302055386"
	human4.ChildID[0]  = "0"
	human4.NewChild[0] = "0"

	var human5 Human
	human5.ID          = "110105199409026635"
	human5.Sex         = "男"
	human5.Name        = "张雷"
	human5.FatherName  = "张父"
	human5.FatherID    = "110105197003025376"
	human5.MotherName  = "张母"
	human5.MotherID    = "110105197302055386"
	human5.ChildID[0]  = "0"
	human5.NewChild[0] = "0"

	var human6 Human
	human6.ID          = "110105199409026645"
	human6.Sex         = "女"
	human6.Name        = "宋梅"
	humanF.FatherName  = "宋父"
	humanF.FatherID    = "110105197003025376"
	humanF.MotherName  = "宋母"
	humanF.MotherID    = "110105197302055386"
	human6.ChildID[0]  = "0"
	human6.NewChild[0] = "0"

	var human7 Human
	human7.ID          = "110105199409026655"
	human7.Sex         = "男"
	human7.Name        = "赵雷雷"
	humanG.FatherName  = "赵父"
	humanG.FatherID    = "110105197003025376"
	humanG.MotherName  = "赵母"
	humanG.MotherID    = "110105197302055386"
	human7.ChildID[0]  = "0"
	human7.NewChild[0] = "0"
	
	var human8 Human
	human8.ID          = "110105199409026665"
	human8.Sex         = "女"
	human8.Name        = "孙梅"
	human8.FatherName  = "孙父"
	human8.FatherID    = "110105197003025376"
	human8.MotherName  = "孙母"
	human8.MotherID    = "110105197302055386"
	human8.ChildID[0]  = "0"
	human8.NewChild[0] = "0"

	var human9 Human
	human9.ID          = "110105199409026695"
	human9.Sex         = "男"
	human9.Name        = "钱雷"
	human9.FatherName  = "钱父"
	human9.FatherID    = "110105197003025376"
	human9.MotherName  = "钱母"
	human9.MotherID    = "110105197302055386"
	human9.ChildID[0]  = "0"
	human9.NewChild[0] = "0"

	var human10 Human
	human10.ID          = "110105199409026605"
	human10.Sex         = "女"
	human10.Name        = "刘梅"
	human10.FatherName  = "刘父"
	human10.FatherID    = "110105197003025376"
	human10.MotherName  = "刘母"
	human10.MotherID    = "110105197302055386"
	human10.ChildID[0]  = "0"
 	human10.NewChild[0] = "0"

	
	humanAsBytes, _ := json.Marshal(humanA)
	APIstub.PutState(humanA.ID, humanAsBytes)

	humanBAsBytes, _ := json.Marshal(humanB)
	APIstub.PutState(humanB.ID, humanBAsBytes)

	humanCsBytes, _ := json.Marshal(humanC)
	APIstub.PutState(humanC.ID, humanCsBytes)

	humanDAsBytes, _ := json.Marshal(humanD)
	APIstub.PutState(humanD.ID, humanDAsBytes)

	humanEAsBytes, _ := json.Marshal(humanE)
	APIstub.PutState(humanE.ID, humanEAsBytes)

	humanFAsBytes, _ := json.Marshal(humanF)
	APIstub.PutState(humanF.ID, humanFAsBytes)

	humanGAsBytes, _ := json.Marshal(humanG)
	APIstub.PutState(humanG.ID, humanGAsBytes)

	humanHAsBytes, _ := json.Marshal(humanH)
	APIstub.PutState(humanH.ID, humanHAsBytes)

	humanIAsBytes, _ := json.Marshal(humanI)
	APIstub.PutState(humanI.ID, humanIAsBytes)

	humanKAsBytes, _ := json.Marshal(humanK)
	APIstub.PutState(humanK.ID, humanKAsBytes)

	human1sBytes, _ := json.Marshal(human1)
	APIstub.PutState(human1.ID, human1sBytes)

	human2AsBytes, _ := json.Marshal(human2)
	APIstub.PutState(human2.ID, human2AsBytes)

	human3sBytes, _ := json.Marshal(human3)
	APIstub.PutState(human3.ID, human3sBytes)

	human4AsBytes, _ := json.Marshal(human4)
	APIstub.PutState(human4.ID, human4AsBytes)

	human5AsBytes, _ := json.Marshal(human5)
	APIstub.PutState(human5.ID, human5AsBytes)

	human6AsBytes, _ := json.Marshal(human6)
	APIstub.PutState(human6.ID, human6AsBytes)

	human7AsBytes, _ := json.Marshal(human7)
	APIstub.PutState(human7.ID, human7AsBytes)

	human8AsBytes, _ := json.Marshal(human8)
	APIstub.PutState(human8.ID, human8AsBytes)

	human9AsBytes, _ := json.Marshal(human9)
	APIstub.PutState(human9.ID, human9AsBytes)

	human10AsBytes, _ := json.Marshal(human10)
	APIstub.PutState(human10.ID, human10AsBytes)

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
	}else if function == "marryCheck" {
		return s.marryCheck(APIstub, args)
	}else if function == "divorceCheck" {
		return s.divorceCheck(APIstub, args)
	}else if function == "divorce" {
		return s.divorce(APIstub, args)
	}else if function == "createCheck" {
		return s.createCheck(APIstub, args)
	}else if function == "queryMarryCheck" {
		return s.queryMarryCheck(APIstub,args)
	}else if function == "queryCreatCheck" {
		return s.queryCreatCheck(APIstub,args)
	}
	return shim.Error("Invalid Smart Contract function name.")
	
}	

func (s *SmartContract) queryID(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var re R_Err

	if len(args) != 1 {
		re.Reason = "参数数量不正确"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}

	humanAsBytes, err := APIstub.GetState(args[0])
	var human Human;
	err = json.Unmarshal(humanAsBytes,&human)//反序列化
	if err != nil {
		re.Reason = "此人不存在"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}
   return shim.Success(humanAsBytes)
	
}

func (s *SmartContract) queryMarryCheck(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var re R_Err

	if len(args) != 1 {
		re.Reason = "参数数量不正确"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}

	
	checkAsBytes, err := APIstub.GetState(args[0])
	var check Marry_Check;
	err = json.Unmarshal(checkAsBytes,&check)//反序列化
	if err != nil {
		re.Reason = "申请不存在"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}

	if 0 != strings.Compare(check.CheckStae,"0"){
		marryAsBytes, _:= APIstub.GetState(check.Marry_Cert)
		return shim.Success(marryAsBytes)
	}else{
		re.Reason = "申请未被处理"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}	
}

func (s *SmartContract) queryCreatCheck(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var re R_Err

	if len(args) != 1 {
		re.Reason = "参数数量不正确"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}

	
	checkAsBytes, err := APIstub.GetState(args[0])
	var check Creat_Check;
	err = json.Unmarshal(checkAsBytes,&check)//反序列化
	if err != nil {
		re.Reason = "申请不存在"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}

	if 0 != strings.Compare(check.CheckStae,"0"){
		humanAsBytes ,_:= APIstub.GetState(check.ID)
		return shim.Success(humanAsBytes)
	}else{
		re.Reason = "申请未被处理"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}	
}


func (s *SmartContract) createBirth(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//9 paramtes father,mother,childsex,birhdate 20171223,hospitalID，Place,weight,health,name
	var re R_Err

	if len(args) !=  9{
		re.Reason = "参数数量不正确"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}
	//whether father is sxisted
	FatherAsBytes, err := APIstub.GetState(args[0])
	var father Human;
	err = json.Unmarshal(FatherAsBytes,&father)//反序列化
	if err != nil {
		re.Reason = "父亲不存在"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}
	//whether mother is sxisted
	MotherAsBytes, err := APIstub.GetState(args[1])
	var mother Human;
	err = json.Unmarshal(MotherAsBytes,&mother)//反序列化
	if err != nil {
		re.Reason = "母亲不存在"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}

	//whether they are couples
	if 0  != (strings.Compare(father.SpouseID,mother.ID)){
		re.Reason = "不是夫妻"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}
	if 0  != (strings.Compare(mother.SpouseID,father.ID)){
		re.Reason = "不是夫妻"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}

	fnum,err := strconv.Atoi(father.NewChild[0])
	if fnum > 1{
		re.Reason = "超生"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}
	mnum,err := strconv.Atoi(mother.ChildID[0])
	if mnum > 1{
		re.Reason = "超生"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}

	//create birth certs
	var birth Birth;
	//timestamp := time.Now().Unix()
	//tm := time.Unix(timestamp, 0)
	birth.BirthID  = strconv.FormatInt(time.Now().Unix(),10)
	birth.Sex      = args[2]
	birth.Date     = args[3]
	birth.FatherID = father.ID
	birth.FatherName = father.Name
	birth.MotherID = mother.ID
	birth.MotherName = mother.Name
	birth.HosptialID = args[4]
	birth.Place = args[5]
	birth.Weight = args[6]
	birth.Health = args[7]
	birth.BirthName = args[8]
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

func (s *SmartContract) createCheck(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	//3 paramtes father or motherID ,1flow father 2 flow mother,name
	var re R_Err

	if len(args) != 3 {
		re.Reason = "参数数量不正确"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}
	var check Creat_Check

	//whether father is sxisted
	FatherAsBytes, err := APIstub.GetState(args[0])
	var father Human;
	err = json.Unmarshal(FatherAsBytes,&father)//反序列化
	if err != nil {
		check.Check[0] = "0"
		check.Check[1] = "0"
		check.FatherName   = "无"      
		check.FatherID     = "无"
	}else{
		check.Check[0] = "1"
		check.Check[1] = "1"
		check.FatherName   = father.Name      
		check.FatherID     = father.ID 
	}
	//whether mother is sxisted
	MotherAsBytes, err := APIstub.GetState(father.SpouseID)
	var mother Human;
	err = json.Unmarshal(MotherAsBytes,&mother)//反序列化
	if err != nil {
		check.Check[2] = "0"
		check.Check[3] = "0"
		check.MotherName   = "无"      
		check.MotherID     = "无" 
	}else{
		check.Check[2] = "1"
		check.Check[3] = "1"
		check.MotherName   = mother.Name      
		check.MotherID     = mother.ID 
	}

	//whether they are couples
	//whether married
	if 0 == strings.Compare(father.SpouseID,"") {
		check.Check[4] = "0"
		check.Marry_Cert = "不是夫妻"
	}else if 0 == strings.Compare(mother.SpouseID,"") {
		check.Check[4] = "0"
		check.Marry_Cert = "不是夫妻"
	}else if 0  != (strings.Compare(father.Marry_Cert,mother.Marry_Cert)){
		check.Check[4] = "0"
		check.Marry_Cert = "不是夫妻"
	}else {
	check.Check[4] = "1"
	check.Marry_Cert = father.Marry_Cert
	}
	// //get the child birth cert
	 cnum,err := strconv.Atoi(father.NewChild[0])
	// if fnum > cnum{
	// 	check.Check[5] = "0"
	// 	check.Check[6] = "0"
	// 	check.Check[7] = "0"
	// 	check.Check[8] = "0"
	// }else{

	ChildAsBytes, err := APIstub.GetState(father.NewChild[cnum])
	var child Birth
	err = json.Unmarshal(ChildAsBytes,&child)//反序列化
	if err != nil {
		check.Check[5] = "0"
		check.Check[6] = "0"
		check.Check[7] = "0"
		check.Check[8] = "0"
		check.BirthID      = "无"    
		check.BirthDate    = "无"      
		check.Sex          = "无"      
		check.HosptialID   = "无"
	}else{
		check.Check[5] = "1"
		check.Check[6] = "1"
		check.Check[7] = "1"
		check.Check[8] = "1"
		check.BirthID      = child.BirthID     
		check.BirthDate    = child.Date      
		check.Sex          = child.Sex      
		check.HosptialID   = child.HosptialID
	}
	    check.CheckStae = "0"
	    check.Name = args[2]
		check.CheckID  =  strconv.FormatInt(time.Now().Unix(),10)
		checkAsBytes, _ := json.Marshal(check)
	    APIstub.PutState(check.CheckID, checkAsBytes)    

		return shim.Success(checkAsBytes)
	}


func (s *SmartContract) createHuman(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//3 paramtes checkId , yes or no
	var re R_Err

	if len(args) != 2{
		re.Reason = "参数数量不正确"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}
	//whether check is exitd
	checkAsBytes, err := APIstub.GetState(args[0])
	var check Creat_Check;
	err = json.Unmarshal(checkAsBytes,&check)//反序列化
	if err != nil {
		re.Reason = "申请表不存在"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}
	if 0 == strings.Compare(check.CheckStae,"1"){
			re.Reason = "申请表已处理"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}
	for i := 0; i < 9; i++{
		if 0 != strings.Compare(check.Check[i],"1"){
			re.Reason = "条件不符"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
		}
	}

	if 0 != strings.Compare(args[1],"1"){
			re.Reason = "不被批准"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}

	//whether father is sxisted
	FatherAsBytes, err := APIstub.GetState(check.FatherID)
	var father Human;
	err = json.Unmarshal(FatherAsBytes,&father)//反序列化
	if err != nil {
		    re.Reason = "父亲不存在"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}
	//whether mother is sxisted
	MotherAsBytes, err := APIstub.GetState(check.MotherID)
	var mother Human;
	err = json.Unmarshal(MotherAsBytes,&mother)//反序列化
	if err != nil {
		    re.Reason = "母亲不存在"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}
	
	//get the child birth cert
	fnum,err := strconv.Atoi(father.ChildID[0])
	mnum,err := strconv.Atoi(mother.ChildID[0])
	cnum,err := strconv.Atoi(father.NewChild[0])
	ChildAsBytes, err := APIstub.GetState(father.NewChild[cnum])
	var child Birth
	err = json.Unmarshal(ChildAsBytes,&child)//反序列化
	if err != nil {
		    re.Reason = "出生证不存在"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}

	//create new human
	var newhuman Human
	newhuman.Sex      = child.Sex
	newhuman.Name     = check.Name
	newhuman.Date     = child.Date
	newhuman.MarryState = "未婚"
	newhuman.SpouseName = "无"
	newhuman.SpouseID = "无"
	newhuman.FatherID = father.ID
	newhuman.FatherName = father.Name
	newhuman.MotherID = mother.ID
	newhuman.MotherName = mother.Name
	newhuman.ChildID[0] = "0"
	newhuman.NewChild[0] = "0"
	
	if 0 == (strings.Compare("1",args[1])){
		address := father.ID[0:6]
		date := child.Date
		if 0 == (strings.Compare("男",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,strconv.FormatInt(time.Now().Unix(),10)[7:10],strconv.Itoa(rand.Intn(9))},"")
		}
		if 0 != (strings.Compare("男",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,strconv.FormatInt(time.Now().Unix(),10)[7:10],strconv.Itoa(rand.Intn(9))},"")
		}
	}

	if 0 == (strings.Compare("2",args[1])){
		address := mother.ID[0:6]
		date := child.Date
		if 0 == (strings.Compare("男",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,strconv.FormatInt(time.Now().Unix(),10)[7:10],strconv.Itoa(rand.Intn(9))},"")
		}
		if 0 != (strings.Compare("男",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,strconv.FormatInt(time.Now().Unix(),10)[7:10],strconv.Itoa(rand.Intn(9))},"")
		}
	}
	newhumanAsBytes, _ := json.Marshal(newhuman)
	APIstub.PutState(newhuman.ID, newhumanAsBytes)

	//change check state
	check.CheckStae = "1"
	check.ID = newhuman.ID
	CheckAsBytes , _:= json.Marshal(check)
	APIstub.PutState(check.CheckID, CheckAsBytes)


	//become father
	fchild := strconv.Itoa(fnum+1)
	father.ChildID[0] = fchild
	father.ChildID[fnum+1] = newhuman.ID
	father.ChildName[fnum+1] = newhuman.Name
	fatherAsBytes, _ := json.Marshal(father)
	APIstub.PutState(father.ID, fatherAsBytes)
	//become mother
	mchild := strconv.Itoa(mnum+1)
	mother.ChildID[0] = mchild
	mother.ChildID[mnum+1] = newhuman.ID
	mother.ChildName[mnum+1] = newhuman.Name
	motherAsBytes, _ := json.Marshal(mother)
	APIstub.PutState(mother.ID, motherAsBytes)

	return shim.Success(newhumanAsBytes)
}

func (s *SmartContract) marryCheck(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//2paramtes husbandID ,wifeID
	var re R_Err

	if len(args) != 2 {
		re.Reason = "参数数量不正确"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}

	var check Marry_Check
	//whether husband is exitd
	husbandAsBytes, err := APIstub.GetState(args[0])
	var husband Human;
	err = json.Unmarshal(husbandAsBytes,&husband)//反序列化
	if err != nil {
		check.Check[0] = "0"
		check.Check[1] = "0"
		check.Check[2] = "0"
		check.Husband_Name = "无"    
	    check.Husband_ID   = "无"
	}else{
		check.Check[0] = "1"
		check.Check[1] = "1"
		check.Husband_Name = husband.Name    
		check.Husband_ID   = husband.ID 
	}
	//whether wife is exitd
	wifeAsBytes, err := APIstub.GetState(args[1])
	var wife Human;
	err = json.Unmarshal(wifeAsBytes,&wife)//反序列化
	if err != nil {
		check.Check[3] = "0"
		check.Check[4] = "0"
		check.Check[5] = "0"
		check.Wife_Name    = "无"   
		check.Wife_ID      = "无"   
	}else{
		check.Check[3] = "1"
		check.Check[4] = "1"
		check.Wife_Name    = wife.Name    
		check.Wife_ID      = wife.ID  
	}
	//whether married
	if 0 != strings.Compare(husband.SpouseID,"") {
		check.Check[2] = "0"
		check.HusbandState = "已婚"
	}else{
		check.Check[2] = "1"
		check.HusbandState = "未婚"
	}
	if 0 != strings.Compare(wife.SpouseID,""){
		check.Check[5] = "0"
		check.WifeState = "已婚"
	}else{
		check.Check[5] = "1"
		check.WifeState = "未婚"
	}

	check.CheckStae = "0"
	check.CheckID      =  strconv.FormatInt(time.Now().Unix(),10)
	checkAsBytes, _ := json.Marshal(check)
	APIstub.PutState(check.CheckID, checkAsBytes)

	return shim.Success(checkAsBytes)
}

func (s *SmartContract) divorceCheck(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//2paramtes husbandID ,wifeID
	var re R_Err

	if len(args) != 2 {
		re.Reason = "参数数量不正确"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}

	var check Divorce_Check
	//whether husband is exitd
	husbandAsBytes, err := APIstub.GetState(args[0])
	var husband Human;
	err = json.Unmarshal(husbandAsBytes,&husband)//反序列化
	if err != nil {
		check.Check[0] = "0"
		check.Check[1] = "0"
	}
	check.Check[0] = "1"
	check.Check[1] = "1"
	//whether wife is exitd
	wifeAsBytes, err := APIstub.GetState(args[1])
	var wife Human;
	err = json.Unmarshal(wifeAsBytes,&wife)//反序列化
	if err != nil {
		check.Check[2] = "0"
		check.Check[3] = "0"
	}
	check.Check[2] = "1"
	check.Check[3] = "1"
	//whether married
	if 0 == strings.Compare(husband.SpouseID,"") {
		check.Check[4] = "0"
		check.Marry_Cert = "不是夫妻"
	}else if 0 == strings.Compare(wife.SpouseID,"") {
		check.Check[4] = "0"
		check.Marry_Cert = "不是夫妻"
	}else if 0  != (strings.Compare(husband.SpouseID,wife.SpouseID)){
		check.Check[4] = "0"
		check.Marry_Cert = "不是夫妻"
	}else {
	check.Check[4] = "1"
	check.Marry_Cert = husband.Marry_Cert
	}

	check.CheckID      =  strconv.FormatInt(time.Now().Unix(),10)
	check.Husband_Name = husband.Name    
	check.Husband_ID   = husband.ID       
	check.Wife_Name    = wife.Name    
	check.Wife_ID      = wife.ID    
	 
	check.CheckStae = "0"
	checkAsBytes, _ := json.Marshal(check)
	APIstub.PutState(check.CheckID, checkAsBytes)

	return shim.Success(checkAsBytes)
}

func (s *SmartContract) marry(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//3par  checkId , yes or no , date 20171223
	var re R_Err

	if len(args) != 3 {
		re.Reason = "参数数量不正确"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}
	//whether check is exitd
	checkAsBytes, err := APIstub.GetState(args[0])
	var check Marry_Check;
	err = json.Unmarshal(checkAsBytes,&check)//反序列化
	if err != nil {
		    re.Reason = "申请表不存在"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}
	if 0 == strings.Compare(check.CheckStae,"1"){
			re.Reason = "申请表已处理"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}
	
	for i := 0; i < 6; i++{
		if 0 != strings.Compare(check.Check[i],"1") {
			re.Reason = "条件不符"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
		}
	}

	if 0 != strings.Compare(args[1],"1"){
		    re.Reason = "未被批准"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}

	//whether husband is exitd
	husbandAsBytes, err := APIstub.GetState(check.Husband_ID)
	var husband Human;
	err = json.Unmarshal(husbandAsBytes,&husband)//反序列化
	if err != nil {
		    re.Reason = "丈夫不存在"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}
	//whether wife is exitd
	wifeAsBytes, err := APIstub.GetState(check.Wife_ID)
	var wife Human;
	err = json.Unmarshal(wifeAsBytes,&wife)//反序列化
	if err != nil {
		    re.Reason = "妻子不存在"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}


	//generate marry id
	marry_cert_id  := strings.Join([]string{"J110101",args[2][0:4],strconv.FormatInt(time.Now().Unix(),10)[4:10]},"-")

	//change check state
	check.CheckStae = "1"
	check.Marry_Cert = marry_cert_id
	CheckAsBytes , _:= json.Marshal(check)
	APIstub.PutState(check.CheckID, CheckAsBytes)

	
	//become husband
	husband.SpouseID = wife.ID
	husband.SpouseName = wife.Name
	husband.MarryState = "已婚"
	husband.Marry_Cert = marry_cert_id
	husbandAsBytes, _ = json.Marshal(husband)
	APIstub.PutState(husband.ID, husbandAsBytes)
	//become wife
	wife.SpouseID = husband.ID
	wife.SpouseName = husband.Name
	wife.MarryState = "已婚"
	wife.Marry_Cert = marry_cert_id
	wifeAsBytes, _ = json.Marshal(wife)
	APIstub.PutState(wife.ID, wifeAsBytes)

	var card Marry_Card
	card.Marry_Cert = marry_cert_id
	card.State = "结婚"
	card.Husband_ID = husband.ID
	card.Husband_Name = husband.Name
	card.Wife_ID = wife.ID
	card.Wife_Name = wife.Name
	card.Date = strings.Join([]string{args[2][0:4],args[2][4:6],args[2][5:7]},"-")
	MarryAsBytes, _ := json.Marshal(card)
	APIstub.PutState(card.Marry_Cert, MarryAsBytes)

	return shim.Success(MarryAsBytes)
}

func (s *SmartContract) divorce(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//3par  checkId , yes or no , date 20171223
	var re R_Err

	if len(args) != 3 {
		re.Reason = "参数数量不正确"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}
	//whether check is exitd
	checkAsBytes, err := APIstub.GetState(args[0])
	var check Divorce_Check;
	err = json.Unmarshal(checkAsBytes,&check)//反序列化
	if err != nil {
		    re.Reason = "申请表不存在"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}
	if 0 == strings.Compare(check.CheckStae,"1"){
			re.Reason = "申请表已处理"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}
	
	for i := 0; i < 5; i++{
		if 0 != strings.Compare(check.Check[i],"1") {
			re.Reason = "条件不符"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
		}
	}

	if 0 != strings.Compare(args[1],"1"){
		    re.Reason = "未被批准"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}

	//whether husband is exitd
	husbandAsBytes, err := APIstub.GetState(check.Husband_ID)
	var husband Human;
	err = json.Unmarshal(husbandAsBytes,&husband)//反序列化
	if err != nil {
		    re.Reason = "丈夫不存在"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}
	//whether wife is exitd
	wifeAsBytes, err := APIstub.GetState(check.Wife_ID)
	var wife Human;
	err = json.Unmarshal(wifeAsBytes,&wife)//反序列化
	if err != nil {
		    re.Reason = "妻子不存在"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}

	//change Marry card
	var card Marry_Card
	cardAsBytes, err := APIstub.GetState(check.Marry_Cert)
	err = json.Unmarshal(cardAsBytes,&card)//反序列化
	if err != nil {
		    re.Reason = "结婚证书不存在"
		    reAsBytes, _ := json.Marshal(re)    
		    return shim.Success(reAsBytes)
	}
	card.State = "离婚"
	card.Date = strings.Join([]string{args[2][0:4],args[2][4:6],args[2][5:7]},"-") 
	cardAsBytes, _ = json.Marshal(card)
	APIstub.PutState(husband.Marry_Cert, cardAsBytes)

	//change divorce checkState
	check.CheckStae = "1";
	CheckAsBytes, _ := json.Marshal(check)
	APIstub.PutState(check.CheckID, CheckAsBytes)

	//change husband spouse
	husband.SpouseID = "无"
	husband.SpouseName = "无"
	husband.Marry_Cert = "无"
	husbandAsBytes, _ = json.Marshal(husband)
	APIstub.PutState(check.Husband_ID, husbandAsBytes)
	//change wife spouse
	wife.SpouseID = "无"
	wife.SpouseName = "无"
	wife.Marry_Cert = "无"
	wifeAsBytes, _ = json.Marshal(wife)
	APIstub.PutState(check.Wife_ID, wifeAsBytes)


	return shim.Success(cardAsBytes)
}

func (s *SmartContract) addInter(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var re R_Err

	if len(args) != 3 {
		re.Reason = "参数数量不正确"
		reAsBytes, _ := json.Marshal(re)    
		return shim.Success(reAsBytes)
	}
	var humanA Human
	humanA.ID       = args[0]
	humanA.Sex      = args[1]
	humanA.Name     = args[2]
	humanA.ChildID[0] = "0"
 	humanA.NewChild[0] = "0"

	humanAsBytes, _ := json.Marshal(humanA)
	APIstub.PutState(humanA.ID, humanAsBytes)

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