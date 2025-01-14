package file

import (
	"encoding/csv"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/xuri/excelize/v2"
)

func createTempCSV(content [][]string) (string, error) {
	file, err := os.CreateTemp("", "test_*.csv")
	if err != nil {
		return "", err
	}

	writer := csv.NewWriter(file)
	for _, record := range content {
		if err := writer.Write(record); err != nil {
			file.Close()
			os.Remove(file.Name())
			return "", err
		}
	}
	writer.Flush()
	file.Close()
	return file.Name(), nil
}

func createTempExcel(content [][]string) (string, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	for i, row := range content {
		for j, cell := range row {
			cellName, _ := excelize.CoordinatesToCellName(j+1, i+1)
			f.SetCellValue(sheet, cellName, cell)
		}
	}

	file, err := os.CreateTemp("", "test_*.xlsx")
	if err != nil {
		return "", err
	}

	if err := f.SaveAs(file.Name()); err != nil {
		file.Close()
		os.Remove(file.Name())
		return "", err
	}
	file.Close()
	return file.Name(), nil
}

func TestReader_ReadCSV(t *testing.T) {
	content := [][]string{{"Name", "Age"}, {"Alice", "30"}, {"Bob", "25"}}
	filePath, err := createTempCSV(content)
	if err != nil {
		t.Fatalf("Failed to create temp CSV: %v", err)
	}
	defer os.Remove(filePath)

	r, err := NewReader[string](filePath)
	if err != nil {
		t.Fatalf("Failed to create reader: %v", err)
	}

	parseFunc := func(data [][]string) (string, error) {
		return data[1][0], nil
	}

	result, err := r.Read(parseFunc)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != "Alice" {
		t.Errorf("Expected 'Alice', got %s", result)
	}
}

func TestReader_ReadExcel(t *testing.T) {
	content := [][]string{{"Name", "Age"}, {"Alice", "30"}}
	filePath, err := createTempExcel(content)
	if err != nil {
		t.Fatalf("Failed to create temp Excel: %v", err)
	}
	defer os.Remove(filePath)

	r, err := NewReader[string](filePath)
	if err != nil {
		t.Fatalf("Failed to create reader: %v", err)
	}

	parseFunc := func(data [][]string) (string, error) {
		return data[1][0], nil
	}

	result, err := r.Read(parseFunc)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != "Alice" {
		t.Errorf("Expected 'Alice', got %s", result)
	}
}

func TestReader_InvalidFileType(t *testing.T) {
	filePath := "invalid.txt"
	_, err := NewReader[string](filePath)
	if err == nil {
		t.Fatal("Expected error for invalid file type, got nil")
	}
}

func TestReader_ReadErrorHandling(t *testing.T) {
	filePath := filepath.Join(os.TempDir(), "nonexistent.csv")
	r, err := NewReader[string](filePath)
	if err != nil {
		t.Fatalf("Failed to create reader: %v", err)
	}

	parseFunc := func(data [][]string) (string, error) {
		return "", errors.New("parse error")
	}

	_, err = r.Read(parseFunc)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if err.Error() != "open "+filePath+": no such file or directory" {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestWriter_WriteCSV(t *testing.T) {
	data := [][]string{{"Name", "Age"}, {"Alice", "30"}, {"Bob", "25"}}
	file, err := os.CreateTemp("", "test_*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp CSV: %v", err)
	}
	defer os.Remove(file.Name())

	w, err := NewWriter(file.Name())
	if err != nil {
		t.Fatalf("Failed to create writer: %v", err)
	}

	if err := w.Write(data); err != nil {
		t.Fatalf("Failed to write CSV: %v", err)
	}

	file.Close()
	file, err = os.Open(file.Name())
	if err != nil {
		t.Fatalf("Failed to open written CSV: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	readData, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read CSV: %v", err)
	}

	for i, row := range data {
		for j, value := range row {
			if readData[i][j] != value {
				t.Errorf("Expected %s, got %s", value, readData[i][j])
			}
		}
	}
}

func TestWriter_WriteExcel(t *testing.T) {
	data := [][]string{{"Name", "Age"}, {"Alice", "30"}, {"Bob", "25"}}
	file, err := os.CreateTemp("", "test_*.xlsx")
	if err != nil {
		t.Fatalf("Failed to create temp Excel: %v", err)
	}
	defer os.Remove(file.Name())

	w, err := NewWriter(file.Name())
	if err != nil {
		t.Fatalf("Failed to create writer: %v", err)
	}

	if err := w.Write(data); err != nil {
		t.Fatalf("Failed to write Excel: %v", err)
	}

	f, err := excelize.OpenFile(file.Name())
	if err != nil {
		t.Fatalf("Failed to open written Excel: %v", err)
	}
	defer f.Close()

	if f.GetSheetName(0) != "Sheet1" {
		t.Fatalf("Expected sheet name 'Sheet1', got '%s'", f.GetSheetName(0))
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		t.Fatalf("Failed to get rows from Excel: %v", err)
	}

	for i, row := range data {
		for j, value := range row {
			if rows[i][j] != value {
				t.Errorf("Expected %s, got %s", value, rows[i][j])
			}
		}
	}
}

func TestWriter_InvalidFileType(t *testing.T) {
	filePath := "invalid.txt"
	_, err := NewWriter(filePath)
	if err == nil {
		t.Fatal("Expected error for invalid file type, got nil")
	}
}
