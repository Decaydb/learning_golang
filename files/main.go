//Получить именя файлов в директории
files, err := ioutil.ReadDir(".")
if err != nil{
	log.Fatal(err)
}

for _, file := range files{
	fmt.Println(file.Name())
}
