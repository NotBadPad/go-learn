package main

func test() {
	defer func() {
		if x := recover(); x != nil {
			res.Head.Code = HANDLE_ERR
			res.Head.Desc = "服务器错误"
			beego.Error("error: ", session, x)
		}
	}()
}
