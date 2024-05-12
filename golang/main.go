package main

import (
	"fmt"

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
	res := my.MaxSumEvery2nd([]int{8, 11, 2, 2, 11})
	fmt.Println("res:", res)
}
