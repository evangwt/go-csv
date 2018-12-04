package csv

import (
	"bufio"
	"io"

	"golang.org/x/text/encoding/unicode"
)

// This is not a standard CSV as "Comma Seperated Values" lib.
// The delimiter this lib used is "\t" instead of "," accroding to the compatible of
// different versions Microsoft Excel.
// Export CSV with UTF-16LE(LE: LittleEndian) does solvethe encoding problems. It works
// well both Windows and MacOS. However the comma delimiter does not work now. After
// replacing "," with "\t", everything is ok now.
const delimiter = "\t"

// line terminater
const terminater = "\n"

// Writer simple csv writer struct
type Writer struct {
	w *bufio.Writer
}

// NewWriter returns a new simple csv writer
func NewWriter(w io.Writer) *Writer {
	// Conver UTF-8 to UTF-16 LittleEndian.
	// Since some old versions of Excel can not decode UTF-8 correctly.
	// For example, Excel 2003, Excel 2007.
	encoder := unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewEncoder()
	return &Writer{
		w: bufio.NewWriter(encoder.Writer(w)),
	}
}

// Write writes a row to writer
func (w *Writer) Write(row []string) error {
	for i, field := range row {
		if i > 0 {
			if _, err := w.w.WriteString(delimiter); err != nil {
				return err
			}
		}

		if field == "" {
			if _, err := w.w.WriteString(""); err != nil {
				return err
			}
		} else {
			if _, err := w.w.WriteString(field); err != nil {
				return err
			}
		}
	}
	if _, err := w.w.WriteString(terminater); err != nil {
		return err
	}
	return nil
}

// Flush flush any bufferd data to writer
func (w *Writer) Flush() error {
	return w.w.Flush()
}

// Error gets the error occours during write or flush
func (w *Writer) Error() error {
	_, err := w.w.Write(nil)
	return err
}

// WriteAll writes all records to writer and flush
func (w *Writer) WriteAll(all [][]string) error {
	for _, row := range all {
		if err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}

