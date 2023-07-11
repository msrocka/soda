# soda

`soda` is a client library for the soda4LCA API written in Go. Here is a small
example how it can be used:

```go
import (
  "fmt"
  "github.com/msrocka/soda"
)

func main() {
  client := soda.NewClient("https://oekobaudat.de/OEKOBAU.DAT/resource")
  stocks, err := client.GetDataStocks()
  for _, stock := range stocks.DataStocks {
    fmt.Println(stock.ShortName)
  }
}
```
