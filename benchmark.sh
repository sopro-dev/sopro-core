#!/bin/bash

# Benchmark my_executable
result1=$(perf stat -r 1000 --format csv ./executable | tee /dev/tty)

# Benchmark my_other_executable
result2=$(perf stat -r 1000 --format csv sox -r 8000 -c 1 -t ul internal/samples/recording.ulaw internal/samples/recording.wav | tee /dev/tty)

# Compare the results
diff <(echo "$result1") <(echo "$result2")
