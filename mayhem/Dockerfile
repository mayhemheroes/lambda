FROM golang:1.18 as builder

COPY . /lambda
WORKDIR /lambda/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN cd fuzz_lambda_batch && go-fuzz-build -libfuzzer -o fuzz_lambda_batch.a && \
    clang -fsanitize=fuzzer fuzz_lambda_batch.a -o fuzz_lambda_batch.libfuzzer

RUN cd fuzz_lambda_bitmap && go-fuzz-build -libfuzzer -o fuzz_lambda_bitmap.a && \
    clang -fsanitize=fuzzer fuzz_lambda_bitmap.a -o fuzz_lambda_bitmap.libfuzzer

RUN cd fuzz_lambda_sort && go-fuzz-build -libfuzzer -o fuzz_lambda_sort.a && \
    clang -fsanitize=fuzzer fuzz_lambda_sort.a -o fuzz_lambda_sort.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /lambda/mayhem/fuzz_lambda_batch/fuzz_lambda_batch.libfuzzer /
COPY --from=builder /lambda/mayhem/fuzz_lambda_bitmap/fuzz_lambda_bitmap.libfuzzer /
COPY --from=builder /lambda/mayhem/fuzz_lambda_sort/fuzz_lambda_sort.libfuzzer /