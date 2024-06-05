package main

import (
	"fmt"
	pf "github.com/Bookshelf-Writer/PointerFactory"
	"sync"
)

/* Добавляем типы */

const (
	TypeN1 pf.TypeTag = iota + 1
	TypeN2
	TypeN3
)

func init() {
	pf.TypeMAP['a'] = TypeN1
	pf.TypeMAP['b'] = TypeN2
	pf.TypeMAP['c'] = TypeN3
}

//////////////////////////////////////////////////////////////////////////////

func main() {
	crt := pf.InitCreator(12) // Задаем сервер
	defer crt.Close()

	arrCrt := []pf.TypeTag{
		TypeN1,
		TypeN2,
		TypeN3,
	}

	var wg sync.WaitGroup
	wg.Add(len(arrCrt) * len(arrCrt))

	//Создание массива указателей асинхронно
	bufID := []string{}
	for pos, tag := range arrCrt {
		for i := 0; i < len(arrCrt); i++ {
			bufTag := tag
			bufPos := pos
			bufCos := i

			go func() {
				id := crt.New(bufTag)
				bufID = append(bufID, id.String())
				fmt.Println(bufPos, bufCos, "|\t", id, "\t", id.String(), id.StringINT(), id.Uint())
				wg.Done()
			}()
		}
	}

	wg.Wait()

	//Проверка на дубли
	validStatus := "is Valid"
	mapID := map[string]int{}
	for _, id := range bufID {
		_, status := mapID[id]
		if !status {
			mapID[id] = 1
		} else {
			mapID[id] += 1
			validStatus = "not Valid"
		}
	}
	fmt.Println("\n", validStatus)
}
