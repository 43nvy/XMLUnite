package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/43nvy/XMLUnite/internal/config"
	"github.com/43nvy/XMLUnite/internal/repository"
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

	cfgPath := filepath.Join(currDir, "config.txt")
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		consoleUI.OutputData(fmt.Sprintf("Возникла ошибка при чтении файла конфигурации: %s", err.Error()))
		os.Exit(1)
	}

	consoleUI.OutputData(currentDirMsg, fmt.Sprintf("%v\n", currDir))
	dataDirPath := filepath.Join(currDir, "data")

	var outputFileName string
	// consoleUI.OutputData(fileNameMsg)
	// consoleUI.InputData(&outputFileName)
	outputFileName = "ssa"

	repos := repository.New(dataDirPath)

	files, err := repos.FindFiles()
	if err != nil {
		consoleUI.OutputData(fmt.Sprintf("Возникла ошибка при поиске файлов: %s", err.Error()))
		os.Exit(1)
	}

	consoleUI.OutputData(fmt.Sprintf("Найдено %d XML файлов.\n", len(files)))

	var wgRead sync.WaitGroup
	var mu sync.Mutex
	var counter int
	errChan := make(chan error, len(files))

	consoleUI.OutputData("Началось чтение файлов...")
	xmlDataList := make([]*service.XMLData, len(files))

	for i, file := range files {
		wgRead.Add(1)
		counter++

		go func(index int, filePath string) {
			defer wgRead.Done()

			data, err := repos.ReadXMLFile(filePath)
			if err != nil {
				errChan <- fmt.Errorf("Возникла ошибка при чтении файла [%s]: %w", filepath.Base(filePath), err)
			}
			mu.Lock()
			xmlDataList[index] = data
			mu.Unlock()
		}(i, file)

		if counter == 8 {
			wgRead.Wait()
			counter = 0
		}
	}

	wgRead.Wait()
	counter = 0

	close(errChan)
	if len(errChan) == len(files) {
		consoleUI.OutputData("Программа не смогла прочитать ни один файл, проверьте целостность файлов.")
		os.Exit(1)
	}
	for err := range errChan {
		consoleUI.OutputData(err.Error())
	}

	srv := service.New(outputFileName, cfg.Fields, cfg.TagFileds)
	var wgParse sync.WaitGroup

	consoleUI.OutputData("Чтение .xml файлов окончено, начинается создание итогового файла...")
	xlsxDataList := make([]*service.XLSXData, len(xmlDataList))

	for i, value := range xmlDataList {
		wgParse.Add(1)
		counter++

		go func(index int, xmlData *service.XMLData) {
			defer wgParse.Done()

			data := srv.ParseXMLFile(xmlData)

			mu.Lock()
			xlsxDataList[index] = data
			mu.Unlock()
		}(i, value)

		if counter == 8 {
			wgParse.Wait()
			counter = 0
		}
	}

	wgParse.Wait()

	err = srv.ExtractToXLSX(xlsxDataList)
	if err != nil {
		consoleUI.OutputData(fmt.Sprintf("Возникла ошибка при создании итогового файла: %s", err.Error()))
		os.Exit(1)
	}

	consoleUI.OutputData(fmt.Sprintf("Выполнение программы завершено успешно, файл %s.xlsx создан.\n", outputFileName))
}
