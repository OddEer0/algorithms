package main

import (
	"github.com/OddEer0/algorithms/golang/my"
)

func main() {
	// t1 := time.Now()
	// res, err := my.RequestsSum([]string{"yandex.ru", "ozon.ru", "sber.ru"})
	// if err != nil {
	// 	fmt.Println("Ошибка:", err.Error())
	// }
	// fmt.Println("Размер:", len(res))
	// fmt.Println("Запрос выполнен за", time.Now().Sub(t1).Seconds(), "секунд")
	my.BobrWorkerPool()
	// fmt.Println("res:", res)
}
