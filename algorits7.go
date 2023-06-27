package main

import ("fmt"
	"time")

func main()  {
	t := time.Now()
	horaAtual := t.Hour()
	minutosAtuais := t.Minute()
	mesAtual := t.Month()
	diaAtual := t.Day()
	fmt.Printf("Time now is: %v:%v of: %v %vst", horaAtual, minutosAtuais, mesAtual, diaAtual)
}