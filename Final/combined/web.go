Combined 
	web.go 
	module -> refer to dept module 
	
	get -> return records
	post -> return 201 Created -> no content 
				
			return 409 Conflict -> any error -> 

func main() {
  r := gin.Default()
  r.GET("/dept", func(c *gin.Context) {
    c.JSON(200, getRecords())
  })
  r.POST("/dept", func(c *gin.Context) {
    val, _ := io.ReadAll(c.Request.Body)
    newDept := Dept{}
    _ = json.Unmarshal(val, &newDept)
    insertRecord(newDept)

  })
  r.Run()
}
 