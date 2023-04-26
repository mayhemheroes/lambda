package fuzz_lambda_sort

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/koss-null/lambda/internal/algo/parallel/mergesort"
    "github.com/koss-null/lambda/internal/algo/parallel/qsort"
)

func _less(a, b int) bool { return a < b }

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {

            case 0:
                arrSize, _ := fuzzConsumer.GetInt()
                var intArr []int
                for i := 0; i < arrSize; i++ {
                    temp, _ := fuzzConsumer.GetInt()
                    intArr = append(intArr, temp)
                }
                testThreads, _ := fuzzConsumer.GetInt()

                mergesort.Sort(intArr, _less, testThreads)
                return 0

            default:
                arrSize, _ := fuzzConsumer.GetInt()
                var intArr []int
                for i := 0; i < arrSize; i++ {
                    temp, _ := fuzzConsumer.GetInt()
                    intArr = append(intArr, temp)
                }
                testThreads, _ := fuzzConsumer.GetInt()

                qsort.Sort(intArr, _less, testThreads)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}