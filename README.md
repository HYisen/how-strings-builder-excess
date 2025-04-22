# How strings.Builder Excess

## Goal

Support my idea that `strings.Builder` has a mild advantage with complicity.

## Usage

Optionally benchmark yourself.

```shell
go test -bench=. -count=10 -timeout 40m | tee new.txt
```

Install the tool if you haven't.

```shell
go install golang.org/x/perf/cmd/benchstat@latest
```

Analyse.

```shell
benchstat -table /partSize -row /count -col /method new.txt
```

Notability, we chose a random part size with ceil `partSize`,
which means the average part size would be (partSize-1)/2.

```shell
benchstat -filter /method:"(Basic OR Build)" -table /partSize -row /count -col /method new.txt
```

## Conclusion

1. Simple is same to Merge, there is no skip return value optimization, or I failed to drop it.
2. For large count and small partSize, Build could be even better than Merge.
3. When does Basic cost double the time of Merge? count@averagePartSize 20@0.5 20@2 5@5 5@10 5@25
4. When does Basic cost double the time of Build? count 10 at all averagePartSizes.
5. On average 5 letters per English word, which is closest to partSize 10,
   Basic cost more than double since count 5, Build excesses Merge since count 20.
6. Can Basic be the better? Never on Merge. partSize 10 or 20 and count 2 on Build.

## Decision

Without more memory footprint, prefer Merge.

If predict more than a hand of parts, switch from Basic to Build.
