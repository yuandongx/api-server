package main

func main() {
	app := SetUp()
	err := app.Run()
	if err != nil {
		return
	}
}
