package deptdb

// DeptStructure defines the structure for a department
type Dept struct {
	Deptno int    `json:"dept_id"`
	Dname  string `json:"dept_name"`
	Loc    string `json:"location"`
}
