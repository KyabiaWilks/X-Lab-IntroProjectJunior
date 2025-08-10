package main

func main() {
	//http.HandleFunc("/ping", Ping)
	//http.ListenAndServe(":8080", nil)
	initDB()
	startServer()

}
