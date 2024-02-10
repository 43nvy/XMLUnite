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
	currentDirMsg      string = "Текущая папка: "
	fileNameMsg        string = "Введите название файла для обьединения"
	pressEnterMsg      string = "Нажмите Enter для выхода..."
	dataDirNotFoundMsg string = "папка data не была найден в директории с программой, она был создана заново."
	startReadMsg       string = "Началось чтение файлов..."
	manyReadErrors     string = "Программа не смогла прочитать ни один файл, проверьте целостность файлов."
	readXmlFinalMsh    string = "Чтение .xml файлов окончено, начинается создание итогового файла..."
	finalMsg           string = "Выполнение программы завершено без критических ошибок, файл %s.xlsx создан.\n"
)

func main() {
	consoleUI := ui.New()

	currDir, _ := os.Getwd()

	cfgPath := filepath.Join(currDir, "config.txt")
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		consoleUI.OutputData(fmt.Sprintf("Возникла ошибка при чтении файла конфигурации: %s", err.Error()))
		fmt.Println(pressEnterMsg)
		fmt.Scanln()
		os.Exit(0)
	}

	consoleUI.OutputData(currentDirMsg, fmt.Sprintf("%v\n", currDir))

	dataDirPath := filepath.Join(currDir, "data")
	err = haveDataDir(dataDirPath)
	if err != nil {
		consoleUI.OutputData(fmt.Sprintf("Возникла ошибка при поиске директории data: %s", err.Error()))
		fmt.Println(pressEnterMsg)
		fmt.Scanln()
		os.Exit(0)
	}

	var outputFileName string
	// consoleUI.OutputData(fileNameMsg)
	// consoleUI.InputData(&outputFileName)
	outputFileName = "sas"

	repos := repository.New(dataDirPath)

	files, err := repos.FindFiles()
	if err != nil {
		consoleUI.OutputData(fmt.Sprintf("Возникла ошибка при поиске файлов: %s", err.Error()))
		fmt.Println(pressEnterMsg)
		fmt.Scanln()
		os.Exit(1)
	}

	consoleUI.OutputData(fmt.Sprintf("Найдено %d XML файлов.\n", len(files)))

	var wgRead sync.WaitGroup
	var mu sync.Mutex
	var counter int
	errChan := make(chan error, len(files))

	consoleUI.OutputData(startReadMsg)
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

	if len(errChan) == len(files) {
		consoleUI.OutputData(manyReadErrors)
		fmt.Println(pressEnterMsg)
		fmt.Scanln()
		os.Exit(1)
	} else if len(errChan) > 0 {
		for err := range errChan {
			consoleUI.OutputData(err.Error())
		}
	}

	srv := service.New(outputFileName, cfg.Tags, cfg.TagNames)
	var wgParse sync.WaitGroup

	consoleUI.OutputData(readXmlFinalMsh)
	xlsxDataList := make([]*service.XLSXData, len(xmlDataList))

	for i, value := range xmlDataList {
		wgParse.Add(1)
		counter++

		go func(index int, xmlData *service.XMLData) {
			defer wgParse.Done()

			data, err := srv.ParseXMLFile(xmlData)
			if err != nil {
				errChan <- fmt.Errorf("Возникла ошибка при парсинге данных в файле [%s]: %w", data.FileName, err)
			}

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
	close(errChan)

	if len(errChan) == len(files) {
		consoleUI.OutputData(manyReadErrors)
		fmt.Println(pressEnterMsg)
		fmt.Scanln()
		os.Exit(1)
	} else if len(errChan) > 0 {
		for err := range errChan {
			consoleUI.OutputData(err.Error())
		}
	}

	err = srv.ExtractToXLSX(xlsxDataList)
	if err != nil {
		consoleUI.OutputData(fmt.Sprintf("Возникла ошибка при создании итогового файла: %s", err.Error()))
		fmt.Println(pressEnterMsg)
		fmt.Scanln()
		os.Exit(1)
	}

	consoleUI.OutputData(fmt.Sprintf(finalMsg, outputFileName))
}

func haveDataDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("haveDateDir mkdir error: %w", err)
		}

		return fmt.Errorf(dataDirNotFoundMsg)
	}

	return nil
}
