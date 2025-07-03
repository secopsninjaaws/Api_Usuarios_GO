package config

import "github.com/joho/godotenv"

func CarregandoVariaveis() {
	err := godotenv.Load()
	if err != nil {
		panic("Error ao carregar as variáveis.")
	}
}
