FROM golang:1.18 as builder

COPY . /lambda
WORKDIR /lambda/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzlambda.a && \
    clang -fsanitize=fuzzer fuzzlambda.a -o fuzzlambda.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /lambda/mayhem/fuzzlambda.libfuzzer /