// Реализовать паттерн «адаптер» на любом примере.
package main

import "fmt"

// Есть наша структура, у неё есть метод
type OurStruct struct {
}

// Метод нашей структуры
func (o *OurStruct) OurRequest() string {
	return "our req"
}

// Конструктор нашей структуры
func NewOur() OurInterface {
	return &OurStruct{}
}

// Есть наш интерфейс с нашим методом, менять нельзя
type OurInterface interface {
	OurRequest() string
}

// Есть внешняя структура которая содержит нужный нам метод, её тоже нельзя изменять
type NeedToAdapt struct {
}

// Вот нужный метод
func (n *NeedToAdapt) AnotherRequest() string {
	return "another req"
}

// Создаём адаптер содержащий нужную нам структуру, у которой есть нужный нам метод
type Adapter struct {
	*NeedToAdapt
}

// Реализуем наш интерфейс у адаптера с использованием метода внешней структуры
func (a *Adapter) OurRequest() string {
	return a.AnotherRequest()
}

// Конструктор адаптера
func NewAdapter(need *NeedToAdapt) OurInterface {
	return &Adapter{
		NeedToAdapt: need,
	}
}

func DoRequest(reqs ...OurInterface) {
	for _, req := range reqs {
		fmt.Println(req.OurRequest())
	}
}

func main() {
	our := NewOur()
	need := new(NeedToAdapt)
	adapter := NewAdapter(need)

	DoRequest(our, adapter)
}
