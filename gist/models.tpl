/**
 * [注释]
 */
func FuncName(/*{参数}*/) (/*{返回值}*/, err error) {
	sql_ := ``
	stmt, err := utils.DB.Prepare(sql_)
	if err != nil {
		fmt.Println("Error prepare sql:", err.Error())
		return
	}
	defer stmt.Close()

	if err = stmt.QueryRow(/*{参数}*/).Scan(/*{赋值}*/); err != nil {
		fmt.Println("GetVmInfo Error:", vmCode, err.Error())
		return
	}
	return
}

 
