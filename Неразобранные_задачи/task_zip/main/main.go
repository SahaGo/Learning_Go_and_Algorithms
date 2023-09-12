/* Задача 1. Поиск файла в заданном формате и его обработка
Данная задача поможет вам разобраться в пакете encoding/csv и path/filepath,
хотя для решения может быть использован также пакет archive/zip (поскольку файл с заданием предоставляется именно
в этом формате).

В тестовом архиве, который мы можете скачать из нашего репозитория на github.com,
содержится набор папок и файлов. Один из этих файлов является файлом с данными в формате CSV,
прочие же файлы структурированных данных не содержат.

Требуется найти и прочитать этот единственный файл со структурированными данными (это таблица 10х10,
	разделителем является запятая), а в качестве ответа необходимо указать число,
находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).*/

package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func mywalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil // Проигнорируем директории
	}
	
	bites, _ := ioutil.ReadFile(path)
	if len(bites) == 0 {
		return nil
	}

	buffer := bytes.NewBuffer(bites)
	r := csv.NewReader(buffer)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if len(records) != 10 {
		return nil
	}

	//	fmt.Println(records)
	fmt.Println(records[4][2])
	//	fmt.Println(path)
	return nil
}

func main() {
	const root = "/Users/kisats/projects/task_zip/task"
	if err := filepath.Walk(root, mywalkFunc); err != nil {
		fmt.Printf("ошибка: %v ", err)
	}
}
