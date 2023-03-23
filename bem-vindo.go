package main

import (
    "fmt"
    "time"
)

//By: Erick Saraiva

func main() {
    // mensagem de boas-vindas
    fmt.Println("Bem-vindo(a) ao programa de data e hora em Go!")

    // Obter a data e hora atual
    now := time.Now()

    // Apresentar o dia da semana e a data
    fmt.Println("Hoje é", now.Weekday().String(), ", ", now.Format("02/01/2006"))

    // Calcular quantos dias restam até o final do ano
    daysLeft := time.Date(now.Year(), 12, 31, 0, 0, 0, 0, time.UTC).YearDay() - now.YearDay()

    // Apresentar a quantidade de dias restantes e já percorridos no ano
    fmt.Println("Já passaram", now.YearDay(), "dias neste ano e restam", daysLeft, "dias até o final do ano.")
}