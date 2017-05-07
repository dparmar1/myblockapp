package main

import (
	"errors"
	"fmt"
	"strconv"
	"encoding/json"
//	"time"
//	"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	
)


// SimpleChaincode example simple Chaincode implementation
type SampleChaincode struct {
	
}

type User struct{
		Userid			string 	 `json:"Userid"`
		FirstName   string   `json:"firstName"`
		LastName    string   `json:"lastName"` 
		Dob			string   `json:"dob"`
		PanNo		string   `json:"pan"`
		
}

type Review struct{
		IdUser			string 	 `json:"idUser"`
		IdDep			string   `json:"idDepartment"`
		Type			string   `json:"type"`
		Flag			string   `json:"flag"`
}

type Approval struct{
		IdUser			string 	 `json:"idUser"`
		IdDep			string   `json:"idDepartment"`
		Type			string   `json:"type"`
		ApproveFlag		string   `json:"approveFlag"`
}

type Organisation struct{
		OrgCode			int32 	 `json:"orgCode"`
		IdDep			string   `json:"idDepartment"`
		Type			string   `json:"type"`
		
}

func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
     if len(args) == 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}
	 
	 _, err := stub.GetTable("UserDetail")
	if err == nil {
		// Table already exists; do not recreate
		
		return nil, nil
	}
	
	err = stub.CreateTable("UserDetail", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "Userid", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "FirstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "LastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PanNo", Type: shim.ColumnDefinition_STRING, Key: false},
		
		})
	
	if err != nil {
		return  nil,errors.New("Failed creating UserDetail table.")
	}
	
	 _, err = stub.GetTable("Approval")
	if err == nil {
		// Table already exists; do not recreate
		
		return nil, nil
	}
	
	err = stub.CreateTable("Approval", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "IdUser", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "IdDep", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Type", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ApproveFlag", Type: shim.ColumnDefinition_STRING, Key: false},
		
		
		})
	
	if err != nil {
		return  nil,errors.New("Failed creating Review table.")
	}
	 _, err = stub.GetTable("Review")
	if err == nil {
		// Table already exists; do not recreate
		
		return nil, nil
	}
	
	err = stub.CreateTable("Review", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "IdUser", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "IdDep", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Type", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Flag", Type: shim.ColumnDefinition_STRING, Key: false},
		
		
		})
	
	if err != nil {
		return  nil,errors.New("Failed creating Review table.")
	}
	
	 _, err = stub.GetTable("Organisation")
	if err == nil {
		// Table already exists; do not recreate
		
		return nil, nil
	}
	
	err = stub.CreateTable("Organisation", []*shim.ColumnDefinition{
	&shim.ColumnDefinition{Name: "Type", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Oid", Type: shim.ColumnDefinition_INT32, Key: false},
		&shim.ColumnDefinition{Name: "Name", Type: shim.ColumnDefinition_STRING, Key: false},		
		})
	
	if err != nil {
		return  nil,errors.New("Failed creating Organisation table.")
	}
	fmt.Println("table created") 
	
	
	
	 fmt.Println("Deployment end ") 
    return nil, nil
}
 
func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    

					if function == "user_detail" {
					return t.getCustomerInfo(stub,args)
					}else if function == "approve"{
					return t.getRecordForApproval(stub,args)
					}else if function == "review"{
					return t.getRecordForReview(stub,args)
					}else if function == "org_detail"{
					return t.getOrgList(stub,args)
					}
					return nil,nil
}
 
 
 
func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("function called " + function) 
    
				    if function == "updateApprovaStatus" {													
						return t.update_approval_status(stub, args )
					}else if function == "updateReviewStatus" {	
						return t.update_review_status(stub, args )
					}else if function == "addUser" {	
						_,err := t.add_customer(stub, args)
						return nil ,err
					}else if function == "approvereq" {	
						_,err := t.insert_For_Approval(stub, args)
						return nil ,err
					}else if function == "verify" {	
						_,err := t.insert_review(stub, args)
						return nil ,err
					}else if function == "addOrg" {	
						_,err := t.insert_Organisation(stub, args)
						return nil ,err
					}
    return nil, errors.New("invalid method call")
}


func (t *SampleChaincode) insert_row(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	if len(args) != 5 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}
		ffId := args [0]
		firstname := args [1]
		lastname := args [2]
		dob := args [3]
		pan := args [4]
		
		
		ok, err := stub.InsertRow("UserDetail", shim.Row{
		Columns: []*shim.Column{
		    &shim.Column{Value: &shim.Column_String_{String_: ffId}},
			&shim.Column{Value: &shim.Column_String_{String_: firstname}},
			&shim.Column{Value: &shim.Column_String_{String_: lastname}},
			&shim.Column{Value: &shim.Column_String_{String_: dob}},
			&shim.Column{Value: &shim.Column_String_{String_: pan}},
			},
		})
		
		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
		
			return nil, errors.New("Row already exists.")
		}
	
		return nil, nil
}

func (t *SampleChaincode) insert_For_Approval(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}
	
	
		userIdval 		:= args [0]
		orgId 			:= args [1]
		orgType 		:= args [2]
		ApprovedStatus 	:= "F"
		
		ok, err := stub.InsertRow("Approval", shim.Row{
		Columns: []*shim.Column{
		    &shim.Column{Value: &shim.Column_String_{String_: userIdval}},
			&shim.Column{Value: &shim.Column_String_{String_: orgId}},
			&shim.Column{Value: &shim.Column_String_{String_: orgType}},
			&shim.Column{Value: &shim.Column_String_{String_: ApprovedStatus}},
			},
	})
		
		if err != nil {
		
			return nil, err 
		}
			if !ok && err == nil {
		
			return nil, errors.New("Row already exists.")
		}
		return nil, nil
}

func (t *SampleChaincode) insert_review(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}
	
	
		userIdval 		:= args [0]
		orgId 			:= args [1]
		orgType 		:= args [2]
		ApprovedStatus 	:= "F"
		
		ok, err := stub.InsertRow("Review", shim.Row{
		Columns: []*shim.Column{
		    &shim.Column{Value: &shim.Column_String_{String_: userIdval}},
			&shim.Column{Value: &shim.Column_String_{String_: orgId}},
			&shim.Column{Value: &shim.Column_String_{String_: orgType}},
			&shim.Column{Value: &shim.Column_String_{String_: ApprovedStatus}},
			},
	})
		
		if err != nil {
		
			return nil, err 
		}
			if !ok && err == nil {
		
			return nil, errors.New("Row already exists.")
		}
		return nil, nil
}



func (t *SampleChaincode) insert_Organisation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}
	
	ffIdval, err := strconv.ParseInt(args[1], 10, 32)

	if err != nil {
			return nil, errors.New("insertRowTable failed. arg[0] must be convertable to int32")
	}
		
		orgcode := int32(ffIdval) 
		orgname := args [1]
		orgType := args [2]
		
		
		ok, err := stub.InsertRow("Organisation", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: orgType}},
			&shim.Column{Value: &shim.Column_Int32{Int32: orgcode}},
			&shim.Column{Value: &shim.Column_String_{String_: orgname}},
			},
	})
		
		if err != nil {
		
			return nil, err 
		}
			if !ok && err == nil {
		
			return nil, errors.New("Row already exists.")
		}
		return nil, nil
}

func (t *SampleChaincode) update_review_status(stub shim.ChaincodeStubInterface, args []string ) ([]byte, error) {
		fmt.Println("inside update_review_status") 
	if len(args) != 3 {
	fmt.Println("less than three")
			return nil, errors.New("replaceRow failed. Must include 2 column values")
		}
	

		
		col2Val := args[0]
		col3Val := args[1]
		
		
		var columns []shim.Column
	//col1 := shim.Column{Value: &shim.ColumnDefinition_INT64{INT64: s}}    Int64
		col1 := shim.Column { Value: &shim.Column_String_{String_: col2Val}}
		col2 := shim.Column { Value: &shim.Column_String_{String_: col3Val}}
		columns = append(columns, col1)
		columns = append(columns, col2)

		row, err := stub.GetRow("Review", columns)
		if err != nil {
	
			return nil, fmt.Errorf("Error: Failed retrieving application with applicationId %s. Error" , err.Error())
		}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		
		return nil, nil
	}
	
	//fmt.Printf("%T, %v", row, row)
	

	

	//fmt.Printf("%T, %v", prvOwner, prvOwner)
	//stub.ReplaceRow(tableName, row)
		col1_Val := row.Columns[0].GetString_()
		col2_Val := row.Columns[1].GetString_()
		col3_Val := row.Columns[2].GetString_()
		col4_Val := args[2]
	
	
	
	var columns_rep []*shim.Column
		col1 = shim.Column{Value: &shim.Column_String_{String_: col1_Val}}
		col2 = shim.Column{Value: &shim.Column_String_{String_: col2_Val}}
		col3 := shim.Column{Value: &shim.Column_String_{String_: col3_Val}}
		col4 := shim.Column{Value: &shim.Column_String_{String_: col4_Val}}
		columns_rep = append(columns_rep, &col1)
		columns_rep = append(columns_rep, &col2)
		columns_rep = append(columns_rep, &col3)
		columns_rep = append(columns_rep, &col4)
		row = shim.Row{Columns: columns_rep}
		//fmt.Println("prvOwner 1",row.Columns[10])
		ok, err := stub.ReplaceRow("Review", row)
		if err != nil {
			return nil, fmt.Errorf("replaceRowTableOne operation failed. %s", err)
		}
		if !ok {
			return nil, errors.New("replaceRowTableOne operation failed. Row with given key does not exist")
		}
	
	fmt.Println("reward_point updated")
	return nil ,nil
}

func (t *SampleChaincode) update_approval_status(stub shim.ChaincodeStubInterface, args []string ) ([]byte, error) {
		fmt.Println("inside approval") 
	if len(args) != 3 {
	fmt.Println("less than three")
			return nil, errors.New("replaceRow failed. Must include 2 column values")
		}
	

		
	col2Val := args[0]
	col3Val := args[1]
		
		
	var columns []shim.Column
	//col1 := shim.Column{Value: &shim.ColumnDefinition_INT64{INT64: s}}    Int64
	col1 := shim.Column { Value: &shim.Column_String_{String_: col2Val}}
	col2 := shim.Column { Value: &shim.Column_String_{String_: col3Val}}
	columns = append(columns, col1)
	columns = append(columns, col2)

	row, err := stub.GetRow("Approval", columns)
	if err != nil {
	
		return nil, fmt.Errorf("Error: Failed retrieving application with applicationId %s. Error" , err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		
		return nil, nil
	}
	
	//fmt.Printf("%T, %v", row, row)
	

	

	//fmt.Printf("%T, %v", prvOwner, prvOwner)
	//stub.ReplaceRow(tableName, row)
	col1_Val := row.Columns[0].GetString_()
	col2_Val := row.Columns[1].GetString_()
	col3_Val := row.Columns[2].GetString_()
	col4_Val := args[2]
	
	
	
	var columns_rep []*shim.Column
		col1 = shim.Column{Value: &shim.Column_String_{String_: col1_Val}}
		col2 = shim.Column{Value: &shim.Column_String_{String_: col2_Val}}
		col3 := shim.Column{Value: &shim.Column_String_{String_: col3_Val}}
		col4 := shim.Column{Value: &shim.Column_String_{String_: col4_Val}}
		columns_rep = append(columns_rep, &col1)
		columns_rep = append(columns_rep, &col2)
		columns_rep = append(columns_rep, &col3)
		columns_rep = append(columns_rep, &col4)
		row = shim.Row{Columns: columns_rep}
		ok, err := stub.ReplaceRow("Approval", row)
		if err != nil {
			return nil, fmt.Errorf("replaceRowTableOne operation failed. %s", err)
		}
		if !ok {
			return nil, errors.New("replaceRowTableOne operation failed. Row with given key does not exist")
		}
	
	fmt.Println("reward_point updated")
	return nil ,nil
}


func (t *SampleChaincode) add_customer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	return t.insert_row(stub,args)
	
}





func (t *SampleChaincode) getCustomerInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
			if len(args) != 1 {
				return nil, errors.New("Incorrect number of arguments. Expecting 1.")
			}
			
			col2Val := args[0]
	

	
			var columns []shim.Column
				col1 := shim.Column { Value: &shim.Column_String_{String_: col2Val}}
				columns = append(columns, col1)
						row, err := stub.GetRow("UserDetail", columns)
			
			if err != nil {
				return nil, fmt.Errorf("Failed to retrieve row")
			}
			if len(row.Columns) == 0 {
				return nil, nil
			}
	
			res := User{}
			res.Userid = row.Columns[0].GetString_()
			res.FirstName = row.Columns[1].GetString_()
			res.LastName = row.Columns[2].GetString_()
			res.Dob = row.Columns[3].GetString_()
			res.PanNo = row.Columns[4].GetString_()
			
	
		
			resJson, _ := json.Marshal(res)
    
	
	return resJson, nil
}

func (t *SampleChaincode) getRecordForReview(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
			if len(args) != 2 {
				return nil, errors.New("Incorrect number of arguments. Expecting 2.")
			}
					
			col2Val := args[0]
			col3Val := args[1]
			var columns []shim.Column
				col1 := shim.Column { Value: &shim.Column_String_{String_: col2Val}}
				col2 := shim.Column { Value: &shim.Column_String_{String_: col3Val}}
				columns = append(columns, col1)
				columns = append(columns, col2)
				row, err := stub.GetRow("Review", columns)
			
			if err != nil {
				return nil, fmt.Errorf("Failed to retrieve row")
			}
			if len(row.Columns) == 0 {
				return nil, nil
			}
	
			res := Review{}
			res.IdUser 	= row.Columns[0].GetString_()
			res.IdDep 	= row.Columns[1].GetString_()
			res.Type 	= row.Columns[2].GetString_()
			res.Flag  	= row.Columns[3].GetString_()
			
			resJson, _ := json.Marshal(res)
    
	
	return resJson, nil
}

func (t *SampleChaincode) getRecordForApproval(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
			if len(args) != 2 {
				return nil, errors.New("Incorrect number of arguments. Expecting 2.")
			}
					
			col2Val := args[0]
			col3Val := args[1]
			var columns []shim.Column
				col1 := shim.Column { Value: &shim.Column_String_{String_: col2Val}}
				col2 := shim.Column { Value: &shim.Column_String_{String_: col3Val}}
				columns = append(columns, col1)
				columns = append(columns, col2)
				row, err := stub.GetRow("Approval", columns)
			
			if err != nil {
				return nil, fmt.Errorf("Failed to retrieve row")
			}
			if len(row.Columns) == 0 {
				return nil, nil
			}
	
			res := Approval{}
			res.IdUser = row.Columns[0].GetString_()
			res.IdDep = row.Columns[1].GetString_()
			res.Type = row.Columns[2].GetString_()
			res.ApproveFlag = row.Columns[3].GetString_()
			
			resJson, _ := json.Marshal(res)
    
	
	return resJson, nil
}


func (t *SampleChaincode) getOrgList(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
			if len(args) != 1 {
				return nil, errors.New("Incorrect number of arguments. Expecting 1.")
			}
			fmt.Println("inside count = %v")
			fmt.Println("total count = %s", args[0])			
			col2Val := args[0]
			var columns []shim.Column
				col1 := shim.Column { Value: &shim.Column_String_{String_: col2Val}}
				columns = append(columns, col1)
							
			rowChannel, err := stub.GetRows("Organisation", columns)
		if err != nil {
			return nil, fmt.Errorf("Customer operation failed. %s", err)
		}
		
		fmt.Println("total count = %v", len(rowChannel))
		
			var rows []shim.Row
		for {
			select {
			case row, ok := <-rowChannel:
				if !ok {
				fmt.Println("no rowChannel")
					rowChannel = nil
				} else {
					rows = append(rows, row)
					fmt.Println("total rows in else = %v", len(rows))
				}
			}
			fmt.Println("for loop = %v", rowChannel)
			if rowChannel == nil {
				break
			}
		}

		jsonRows, err := json.Marshal(rows)
		if err != nil {
			return nil, fmt.Errorf("getRowsTableFour operation failed. Error marshaling JSON: %s", err)
		}

		return jsonRows, nil
}


func (t *SampleChaincode) getCustomerInfo1(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
			
			if len(args) != 1 {
				return nil, errors.New("Incorrect number of arguments. Expecting 1.")
			}
			col1Int, err := strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				return nil, errors.New("fetch record failed. arg[0] must be convertable to int32")
			}
			col2Val := int32(col1Int)
		
		var columns []shim.Column
				col1 := shim.Column { Value: &shim.Column_Int32{Int32: col2Val}}
				columns = append(columns, col1)

		rowChannel, err := stub.GetRows("Customer", columns)
		if err != nil {
			return nil, fmt.Errorf("Customer operation failed. %s", err)
		}

		var rows []shim.Row
		for {
			select {
			case row, ok := <-rowChannel:
				if !ok {
					rowChannel = nil
				} else {
					rows = append(rows, row)
				}
			}
			if rowChannel == nil {
				break
			}
		}

		jsonRows, err := json.Marshal(rows)
		if err != nil {
			return nil, fmt.Errorf("getRowsTableFour operation failed. Error marshaling JSON: %s", err)
		}

		return jsonRows, nil
		
		

}
func main() {
    err := shim.Start(new(SampleChaincode))
    if err != nil {
        fmt.Println("Could not start SampleChaincode")
    } else {
        fmt.Println("SampleChaincode successfully started")
    }
 
}