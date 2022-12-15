module myapp

go 1.18

replace github.com/andihoerudin24/goand => ../goand

require github.com/andihoerudin24/goand v0.0.0

require (
	github.com/go-chi/chi/v5 v5.0.8 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
)
