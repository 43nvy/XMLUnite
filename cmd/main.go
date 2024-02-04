package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/43nvy/XMLUnite/internal/service"
	"github.com/43nvy/XMLUnite/internal/ui"
)

const (
	currentDirMsg string = "Текущая папка: "
	pathMsg       string = "Введите путь до папки"
	fileNameMsg   string = "Введите название файла для обьединения"
)

func main() {
	consoleUI := ui.New()

	currDir, _ := os.Getwd()
	consoleUI.OutputData(currentDirMsg, fmt.Sprintf("%v\n", currDir))

	// var dirPath string
	// consoleUI.OutputData(pathMsg)
	// consoleUI.InputData(&dirPath)
	dirPath := "/home/shiva/Desktop/proj/Go/data"
	// var fileName string
	// consoleUI.OutputData(fileNameMsg)
	// consoleUI.InputData(&fileName)
	fileName := "ssss"

	srv := service.New(consoleUI, dirPath, fileName)
	files, err := srv.FindFiles()
	if err != nil {
		consoleUI.OutputData(fmt.Sprintf("Возникла ошибка при поиске файлов: %s", err.Error()))
		os.Exit(1)
	}

	consoleUI.OutputData(fmt.Sprintf("Найденные %d файлы:\n", len(files)))

	var wgRead sync.WaitGroup
	errChan := make(chan error, len(files))

	xmlDataList := make([]*service.XMLData, len(files))
	for i, filePath := range files {
		wgRead.Add(1)
		tempFilePath := filePath
		tempIndex := i

		go func() {
			defer wgRead.Done()
			data, err := srv.ReadXMLFile(tempFilePath)
			if err != nil {
				errChan <- fmt.Errorf("Возникла ошибка при чтении файла [%s]: %w", tempFilePath, err)
			}

			xmlDataList[tempIndex] = data
		}()
	}

	wgRead.Wait()

	close(errChan)
	if len(errChan) == len(files) {
		consoleUI.OutputData("Программа не смогла прочитать ни один файл, проверьте целостность файлов.")
		os.Exit(1)
	}
	for err := range errChan {
		consoleUI.OutputData(err.Error())
	}

	var wgParse sync.WaitGroup

	xlsxDataList := make([]*service.XLSXData, len(xmlDataList))

	for i, xmlData := range xmlDataList {
		wgParse.Add(1)
		tempData := xmlData
		tempIndex := i

		go func() {
			defer wgParse.Done()
			xlsxData := srv.ParseXMLFile(tempData)
			xlsxDataList[tempIndex] = xlsxData
		}()
	}

	wgParse.Wait()

	err = srv.ExtractToXLSX(xlsxDataList)
	if err != nil {
		consoleUI.OutputData(fmt.Sprintf("Возникла ошибка при создании итогового файла: %s", err.Error()))
		os.Exit(1)
	}
}
