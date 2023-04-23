package fuzzlambda

import (
    "strconv"
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/koss-null/lambda/internal/bitmap"
    "github.com/koss-null/lambda/internal/algo/batch"
    "github.com/koss-null/lambda/internal/algo/parallel/mergesort"
    "github.com/koss-null/lambda/internal/algo/parallel/qsort"
    "github.com/koss-null/lambda/pkg/pipe"
)

func _less(a, b int) bool { return a < b }

func mayhemit(data []byte) int {

    var num int
    if len(data) > 2 {
        num, _ = strconv.Atoi(string(data[0]))
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            case 0:
                testInt, _ := fuzzConsumer.GetInt()

                bitmap.NewNaive(testInt)
                return 0

            case 1:
                int1, _ := fuzzConsumer.GetInt()
                int2, _ := fuzzConsumer.GetInt()

                var test = bitmap.NewNaive(int1)
                fuzzConsumer.GenerateStruct(&test)

                test.Get(int2)
                return 0

            case 2:
                int1, _ := fuzzConsumer.GetInt()
                int2, _ := fuzzConsumer.GetInt()
                int3, _ := fuzzConsumer.GetInt()
                bool1, _ := fuzzConsumer.GetBool()

                var test = bitmap.NewNaive(int1)
                fuzzConsumer.GenerateStruct(&test)

                test.Set(int2, int3, bool1)
                return 0

            case 3:
                arrSize, _ := fuzzConsumer.GetInt()
                var intArr []int
                for i := 0; i < arrSize; i++ {
                    temp, _ := fuzzConsumer.GetInt()
                    intArr = append(intArr, temp)
                }

                pipe.Slice(intArr)
                return 0

            case 4:
                arrSize, _ := fuzzConsumer.GetInt()
                var intArr []int
                for i := 0; i < arrSize; i++ {
                    temp, _ := fuzzConsumer.GetInt()
                    intArr = append(intArr, temp)
                }
                testBatch, _ := fuzzConsumer.GetInt()

                batch.Do(intArr, testBatch)
                return 0

            case 5:
                arrSize, _ := fuzzConsumer.GetInt()
                var intArr []int
                for i := 0; i < arrSize; i++ {
                    temp, _ := fuzzConsumer.GetInt()
                    intArr = append(intArr, temp)
                }
                testThreads, _ := fuzzConsumer.GetInt()

                mergesort.Sort(intArr, _less, testThreads)
                return 0

            case 6:
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