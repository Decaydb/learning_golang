//Структуры используются для разбора данных, к примеру json:

//Объявление структуры для разбора строки text, которая содержит json разметку
type info struct{
	Markets []struct{
		Name status
		Price int
	}
}

//Структура вызывается в структуре выше
type status struct{
	Name string
}

func main(){
	text:= `{"markets":[{"name":"1", "price":100}, {"name":"2", "price":101}, {"name":"3", "price":102}]}`
	var infos info 
	json.Unmarshal([]byte(text), &infos)
	fmt.Println(infos)
	fmt.Println(infos.Markets)
	fmt.Println(infos.Markets[0].Name)
	for i:= range infos.Markets{
		fmt.Println(i, infos.Markets[1].Name, infos.Markets[i].Price)
	}
}
