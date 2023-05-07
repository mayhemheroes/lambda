package fuzz_lambda_bitmap

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/koss-null/lambda/internal/bitmap"
)

func _less(a, b int) bool { return a < b }

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
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

            default:
                int1, _ := fuzzConsumer.GetInt()
                int2, _ := fuzzConsumer.GetInt()
                int3, _ := fuzzConsumer.GetInt()
                bool1, _ := fuzzConsumer.GetBool()

                var test = bitmap.NewNaive(int1)
                fuzzConsumer.GenerateStruct(&test)

                test.Set(int2, int3, bool1)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}