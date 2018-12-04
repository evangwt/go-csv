### go-csv

Simple CSV lib writen by golang solves the encoding nightmare while exporting records. Both on Windows and MacOS with Microsoft Excel, Apple Numbers.

## Why

The encoding of CSV files is a nightmare. Some old versions of Microsoft Excel drop the Unicode BOM, althrogh you try to add '0xEF, 0xBB, 0xBF' with UTF-8 encoding. For Example, Excel 2003, Excel 2007. However the newer versions have no problems. Maybe you will use "encoding/csv" lib, but it does not solve the encoding problem.

## Solution

#### [Unicode Encoding Form](http://unicode.org/faq/utf_bom.html#BOM)

| Bytes       | Encoding Form         |
| ----------- | --------------------- |
| 00 00 FE FF | UTF-32, big-endian    |
| FF FE 00 00 | UTF-32, little-endian |
| FE FF       | UTF-16, big-endian    |
| FF FE       | UTF-16, little-endian |
| EF BB BF    | UTF-8                 |

####UTF-16

UTF-16 is the key point. But the comma delimiter does not work. We should replace "," with "\t". So this is not a standard CSV format as "Comma Seperated Values". And this is why when you use standard lib "encoding/csv" and encode your records with UTF-16 and the result is still not what you want. Here I choose  UTF-16 little-endian.

## Usage

### Basic
``` go
buf := bytes.NewBuffer(nil)
w := csv.NewWriter(buf)
titles := []string{ "id", "name" }
if err := w.Write(titles); err != nil {
    return err
}

// fetch data records
rows := db.getRecords()
for _, row := range rows {
    row := []string{ row.ID, row.name  }
    if err := w.Write(row); err != nil {
        return err
    }
}
w.Flush()
fmt.Println(buf.String())
```

### Export File
``` go
// gin framework
func ExportCsv(ctx *gin.Context, name string, data string) {
	ctx.Header("Content-Disposition", "attachment;filename="+name+".csv")
	ctx.Header("Content-Transfer-Encoding", "binary")
	// UTF-16LE
	ctx.Data(http.StatusOK, "text/csv;charset=UTF-16LE", []byte(data))
}

ExportCsv(ctx, 'file', buf.String())
```

## License
MIT License

Copyright (c) 2018

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
