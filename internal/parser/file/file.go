package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dougdoenges/flexion-coding-challenge/internal/util"
	"github.com/xuri/excelize/v2"
)

type FileType string

const (
	CSV   FileType = ".csv"
	EXCEL FileType = ".xslx"
)

var FILE_TYPES = map[FileType]struct{}{
	CSV:   {},
	EXCEL: {},
}

func IsValidFileType(ft string) bool {
	_, ok := FILE_TYPES[FileType(ft)]
	return ok
}

func validFileTypes() []FileType {
	vft := make([]FileType, 0, len(FILE_TYPES))
	for key := range FILE_TYPES {
		vft = append(vft, key)
	}
	return vft
}

type Reader[T any] struct {
	path     string
	fileType FileType
}

func NewReader[T any](path string) (Reader[T], error) {
	typ := strings.ToLower(filepath.Ext(path))
	if !IsValidFileType(typ) {
		return Reader[T]{}, fmt.Errorf("invalid file given '%s'. allowed types: %s",
			path, validFileTypes())
	}
	return Reader[T]{path, FileType(typ)}, nil
}

func (r Reader[T]) Read(parseFunc func([][]string) (T, error)) (T, error) {
	var gridValues [][]string
	var err error

	switch r.fileType {
	case CSV:
		gridValues, err = readCSV(r.path)
	case EXCEL:
		gridValues, err = readExcel(r.path)
	default:
		return util.ZeroValue[T](), fmt.Errorf("unsupported file type: %s", r.fileType)
	}
	if err != nil {
		return util.ZeroValue[T](), err
	}

	result, err := parseFunc(gridValues)
	if err != nil {
		return util.ZeroValue[T](), err
	}

	return result, nil
}

func readCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func readExcel(filePath string) ([][]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	return f.GetRows(sheetName)
}

type Writer struct {
	path     string
	fileType FileType
}

// NewWriter creates a new Writer instance based on the file type
func NewWriter(path string) (Writer, error) {
	typ := strings.ToLower(filepath.Ext(path))
	if !IsValidFileType(typ) {
		return Writer{}, fmt.Errorf("invalid file type '%s'. Supported types: .csv, .xlsx", typ)
	}
	return Writer{path, FileType(typ)}, nil
}

func (w Writer) Write(data [][]string) error {
	switch w.fileType {
	case CSV:
		return writeCSVData(w.path, data)
	case EXCEL:
		return writeExcelData(w.path, data)
	default:
		return fmt.Errorf("unsupported file type '%s'", w.fileType)
	}
}

func writeCSVData(path string, data [][]string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write record to CSV: %v", err)
		}
	}

	return nil
}

func writeExcelData(path string, data [][]string) error {
	f := excelize.NewFile()
	sheet := "Output"

	for i, row := range data {
		for j, value := range row {
			cell, err := excelize.CoordinatesToCellName(j+1, i+1)
			if err != nil {
				return fmt.Errorf("failed to convert coordinates to cell: %v", err)
			}
			if err := f.SetCellValue(sheet, cell, value); err != nil {
				return fmt.Errorf("failed to set cell value: %v", err)
			}
		}
	}

	if err := f.SaveAs(path); err != nil {
		return fmt.Errorf("failed to save Excel file: %v", err)
	}
	return nil
}
