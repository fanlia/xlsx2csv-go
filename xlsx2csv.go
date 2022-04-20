
package main

import (
    "fmt"
    "encoding/csv"
    "os"
    "github.com/xuri/excelize/v2"
)

func main() {
    f, err := excelize.OpenReader(os.Stdin)
    names := f.GetSheetList()
    rows, err := f.Rows(names[0])
    if  err != nil {
        fmt.Fprintf(os.Stderr, err.Error())
        return
    }
    writer := csv.NewWriter(os.Stdout)
    for rows.Next() {
        row, err := rows.Columns()
        if  err != nil {
            fmt.Fprintf(os.Stderr, err.Error())
            break
        }
        if len(row) == 0 {
            continue
        }
        writer.Write(row)
    }
    rows.Close()
    writer.Flush()

    if err = writer.Error(); err != nil {
        fmt.Fprintf(os.Stderr, err.Error())
    }
}
