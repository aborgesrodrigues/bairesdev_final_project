package main

import "../internal/handler"

func main() {
	//cmd.CallDao()
	//cmd.CallService()
	//cmd.CreateDatabase()
	//cmd.PopulateUser()

	handler.CreateUserRouters()
}
