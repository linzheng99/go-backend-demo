package main

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := &application{
		config: cfg,
	}

	// 运行应用
	if err := app.run(); err != nil {
		panic(err)
	}

}
