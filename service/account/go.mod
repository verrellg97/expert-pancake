module github.com/expert-pancake/service/account

go 1.18

require (
	github.com/calvinkmts/expert-pancake/engine v0.0.0-20221028183511-de01ce05c4ff // indirect
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/go-chi/cors v1.2.1 // indirect
)

replace (
	github.com/calvinkmts/expert-pancake/engine v0.0.0-20221028183511-de01ce05c4ff => "../../engine"
)