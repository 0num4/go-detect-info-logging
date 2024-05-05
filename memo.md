# error

## Validate: analyzer "awesomeAnalyzer" is undocumented

必ずドキュメントを書く必要がありそう

analyzer 構造体の Doc フィールドにドキュメントを書く必要がある

```
Doc:              "this is awesome analyzer",
```

error message

```
❯ go test -run ^TestRun$
Running TestRun(t *testing.T)-: cannot find package "a" in any of:
        /Users/oonuma/.anyenv/envs/goenv/versions/1.22.0/src/a (from $GOROOT)
        /Users/oonuma/work/private/go-detect-info-logging/testdata/src/a (from $GOPATH)
Running the analyzer
--- FAIL: TestRun (0.12s)
    --- FAIL: TestRun/TestRun(t_*testing.T) (0.12s)
        analysistest.go:329: Validate: analyzer "awesomeAnalyzer" is undocumented
FAIL
exit status 1
FAIL    github.com/0num4/go-detect-info-logging 0.132s
```

## 動くようになった

```
❯ go test -run ^TestRun$
Running TestRun(t *testing.T)-: cannot find package "a" in any of:
        /Users/oonuma/.anyenv/envs/goenv/versions/1.22.0/src/a (from $GOROOT)
        /Users/oonuma/work/private/go-detect-info-logging/testdata/src/a (from $GOPATH)
Running the analyzer
--- FAIL: TestRun (0.12s)
    --- FAIL: TestRun/TestRun(t_*testing.T) (0.12s)
        analysistest.go:336: error analyzing awesomeAnalyzer@a: analysis skipped due to errors in package
FAIL
exit status 1
FAIL    github.com/0num4/go-detect-info-logging 0.129s
```
このエラーの最初はmain.goの下記の"a"のこと。その下の文でこれはtestdata/src/aやほかのディレクトリが無いよみたいな話をしている。
    
```go
analysistest.Run(t, testdata, mychecker.Analyzer, "sampleSrc")
```

適当にディレクトリを作ることで消えた。
mkdir -p testdata/src/sampleSrc

## 他のparseについて
標準のgo/parserパッケージに色々入ってる。
https://github.com/golang/go/tree/master/src/go
parser.parseFile() とかがある。
tokenはgo/tokenにある。
型情報はgo/typesにある。
configもgo/typesにある。

go/astにもfileが生えててast.Fileでアクセスできる
go/types.configのconfig.checkのast.Fileにはparser.parseFileの引数の返り値を入れられる