package UID

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

//////////////////////////////////////////////////////////

// Контрольная сумма
func BenchmarkChecksum(b *testing.B) {
	str := strconv.Itoa(time.Now().Nanosecond())

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SYS.CRC(str)
	}
	b.StopTimer()
}

/**/

// потокое созадание указателей
func BenchmarkAdd(b *testing.B) {
	crt := SYS.InitCreator(999)
	defer crt.Close()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, tag := range SYS.TypeMAP {
			crt.New(tag)
		}
	}
	b.StopTimer()
}

// Асинхронное создание указателей в рамках потока
func BenchmarkAddP(b *testing.B) {
	crt := SYS.InitCreator(999)
	defer crt.Close()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(len(SYS.TypeMAP))

		for _, tag := range SYS.TypeMAP {
			bufTag := tag

			go func() {
				crt.New(bufTag)
				wg.Done()
			}()

		}

		wg.Wait()
	}
	b.StopTimer()
}

// Асинхронное создание указателей за пределами потока
func BenchmarkAddPG(b *testing.B) {
	crt := SYS.InitCreator(999)
	defer crt.Close()

	var wg sync.WaitGroup
	wg.Add(len(SYS.TypeMAP) * b.N)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, tag := range SYS.TypeMAP {
			bufTag := tag

			go func() {
				crt.New(bufTag)
				wg.Done()
			}()

		}
	}

	wg.Wait()
	b.StopTimer()
}
