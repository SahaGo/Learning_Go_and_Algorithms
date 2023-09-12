package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Group struct {
	ID       string
	Number   string
	Year     int
	Students []Student
}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []float32
}

type Rating struct {
	Average float32
}

func main() {

	//var GroupJsonType []byte

	//fmt.Scan(&GroupJsonType)
	//	fmt.Fscan(os.Stdin, &GroupJsonType)
	//GroupJsonType, err := io.ReadAll(os.Stdin)
	//if err != nil {
	//	log.Fatalf("failed to read stdin: %s", err)
	//}
	var GroupGoStruct Group
	//var GroupJsonType = new(bytes.Buffer)
	json.NewDecoder(os.Stdin).Decode(&GroupGoStruct)
	//err := json.Unmarshal(GroupJsonType, &GroupGoStruct)
	//if err != nil {
	//	fmt.Println(err)
	//}

	var countStudents float32
	var countMarks float32
	for _, v := range GroupGoStruct.Students {
		countStudents++
		for _, _ = range v.Rating {
			countMarks++
		}
	}

	//fmt.Println(countStudents, countMarks)
	var answer Rating
	answer.Average = countMarks / countStudents
	data, err := json.MarshalIndent(answer, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
	//	io.WriteString(os.Stdout, strconv.Itoa(data))
}

/*
{
    "ID":134,
    "Number":"ИЛМ-1274",
    "Year":2,
    "Students":[
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3]
            },
    ]
}

1) Создаем 3 структуры: Group, Student, Rating (при этом слайс типа Student является типом для поля Group.Students);

2) Считываем данные из ввода через ioutil.ReadAll;

3) Десерелизуем полученные данные в структуру Group;

4) Создаем 2 цикла - один вложен в другой: в первом проходим по полю Students и считаем количество студентов в счетчик, во втором проходим по полю Rating каждого студента и считаем количество оценок в другой счетчик;

5) Делим количество оценок на количество студентов - среднее значение и записываем в объект типа Rating (третья структура);

6) Серелизуем этот объект с помощью json.MarshalIndent(a, "", "    ");

7) Выводим с помощью w.WriteString(string(data2)):*/
