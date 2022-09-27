package main

import "fmt"

func main() {

	//В Golang вместо наследования можно использовать композицию или агрегацию

	//метод родителя
	h := Human{}
	h.run(5)

	//композиция
	A := ActionComposition{}
	A.run(10)

	//агрегация
	A2 := ActionAggregation{Human: newHuman("Jim", "male", 18, true)}
	A2.Human.run(15)

}

//родительский тип
type Human struct {
	name   string
	age    int
	alive  bool
	gender string
}
//ооздание структуры Human с заданными полями
func newHuman(name, gender string, age int, alive bool) *Human {
	return &Human{name:name, gender:gender, age:age, alive:alive}
}

//метод, который наследуем
func (h *Human) run(time int) {
	fmt.Printf("running for %d mins\n", time)
}

//композиция
type ActionComposition struct {
	Human
}

//агрегация
type ActionAggregation struct {
	Human *Human
}
