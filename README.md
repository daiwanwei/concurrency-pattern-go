# concurrency-pattern-go
## Table of Contents

 * [專案描述](#專案描述)
 * [執行專案](#執行專案)

## 專案描述

### concurrency pattern
- [x] barrier
- [x] future
- [x] pipeline
- [x] worker pool
- [x] publish/subscribe

## 執行專案


#### 執行應用程式

```bash
#到專案目錄下
$ cd path_to_dir/concurrency-pattern-go

# 下載第三方套件
$ go mod download

# 編譯專案(輸出到當前目錄下,檔案名為main)
$ go build -o main ./{pattern} 

# 執行應用程式
$ ./main 
```