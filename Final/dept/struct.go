package main
 
// DeptStructure defines the structure for a department
type Dept struct {
    DeptID   int    `json:"dept_id"`
    DeptName string `json:"dept_name"`
    Location string `json:"location"`
}