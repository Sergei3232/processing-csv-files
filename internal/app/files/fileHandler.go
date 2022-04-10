package fileHandler

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type CsvFile struct {
	Sku            int64
	MapiItem       int64
	VerticaVariant int64
	Height         int64
	Width          int64
}

type FileWorker interface {
	ReadCSVFile(file string) ([]CsvFile, error)
}

type FileHandler struct {
}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

func (f *FileHandler) GetListFilesProcess(path string) ([]string, error) {
	lst, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	files := make([]string, 0, len(lst))

	for _, val := range lst {
		if val.IsDir() {
			fmt.Printf("[%s]\n", val.Name())
		} else {
			files = append(files, path+"/"+val.Name())
		}
	}

	if len(files) == 0 {
		return files, errors.New("No files to process!")
	}
	return files, nil
}

func (f *FileHandler) FileTransfer(oldpath, newpath string) error {
	err := os.Rename(oldpath, newpath)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileHandler) CombiningFiles(listCombiningFiles []string) error {
	dataCombining := [][]string{{"sku", "mapi_item", "vertica_variant", "height", "width"}}

	for _, val := range listCombiningFiles {
		combiningFiles, err := f.ReadToCSVFile(val)
		if err != nil {
			return err
		}

		for _, combiningFile := range combiningFiles {
			sku := strconv.Itoa(int(combiningFile.Sku))
			mapiItem := strconv.Itoa(int(combiningFile.MapiItem))
			verticaVariant := strconv.Itoa(int(combiningFile.VerticaVariant))
			width := strconv.Itoa(int(combiningFile.Width))
			height := strconv.Itoa(int(combiningFile.Height))
			dataCombining = append(dataCombining, []string{sku, mapiItem, verticaVariant, width, height})
		}
	}

	return nil
}

func (f *FileHandler) ReadToCSVFile(file string) ([]CsvFile, error) {
	filesCSV, err := os.Open(file)
	defer filesCSV.Close()

	if err != nil {
		return nil, err
	}

	lines, err := csv.NewReader(filesCSV).ReadAll()
	if err != nil {
		return nil, err
	}

	dataCSV := make([]CsvFile, 0, len(lines))

	for nl, line := range lines { //line
		linesCSV := make([]int64, 0, 5)
		if nl > 0 {
			a1, _ := strconv.ParseInt(line[0], 10, 64)
			a2, _ := strconv.ParseInt(line[1], 10, 64)
			a3, _ := strconv.ParseInt(line[2], 10, 64)
			linesCSV = append(linesCSV, a1, a2, a3)

			if len(line) > 3 {
				a4, _ := strconv.ParseInt(line[3], 10, 64)
				a5, _ := strconv.ParseInt(line[4], 10, 64)
				linesCSV = append(linesCSV, a4, a5)
			} else {
				linesCSV = append(linesCSV, 0, 0)
			}

			dataCSV = append(dataCSV, CsvFile{
				linesCSV[0],
				linesCSV[1],
				linesCSV[2],
				linesCSV[3],
				linesCSV[4]})
		}
		nl++
	}

	return dataCSV, nil
}
