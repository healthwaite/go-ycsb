package measurement

import (
	"fmt"
	"io"
	"os"
	"time"
)

type csvsFile struct {
	file *os.File
}

func (c *csvsFile) GenerateExtendedOutputs() {
}

func InitCSVFile(outputFile string) *csvsFile {
	f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0755)
	if err != nil {
		panic(fmt.Sprintf("failed to open file %s err=%s\n", outputFile, err))
	}

	_, err = f.WriteString(fmt.Sprintf("operation,timestamp_us,latency_us\n"))
	if err != nil {
		panic(fmt.Sprintf("failed to write to file %s err=%v\n", outputFile, err))
	}

	return &csvsFile{
		file: f,
	}
}

func (c *csvsFile) Measure(op string, start time.Time, lan time.Duration) {
	if op == "TOTAL" {
		return
	}

	_, err := c.file.WriteString(fmt.Sprintf("%s,%d,%d\n", op, start.UnixMicro(), lan.Microseconds()))
	if err != nil {
		panic(fmt.Sprintf("failed to write to file /results.csv err=%v\n"))
	}
}

func (c *csvsFile) Output(w io.Writer) error {
	c.file.Sync()
	// Nothing to do
	return nil
}

func (c *csvsFile) Summary() {
	// do nothing as csvsFile don't keep a summary
}
