package main

func main() {
	app := SetUp()
	err := app.Run(":9090")
	if err != nil {
		return
	}
}
