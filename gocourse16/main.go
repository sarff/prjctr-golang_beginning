/*
Виконати пошук телефонних номерів у файлі з даними контактів. Задача полягає в тому, щоб створити регулярний вираз,
який можна використовувати для знаходження телефонних номерів, записаних у різних форматах. Наприклад, ви можете почати
з використання виразу, який знаходить номери телефонів, що складаються з 10 цифр, а потім розширити його, додавши
підтримку різних форматів, наприклад, номери з круглими дужками, пробілами та дефісами. Для відпрацювання регулярного
виразу дивіться https://regex101.com/. Далі треба перезаписати цей файл із іменами та номерами телефонів, щоб усі
номери були в одному форматі, а саме (000) 11-22-333.
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var regexPhone = regexp.MustCompile(`[^\d]`)

func FormatNumber(s string) (string, error) {
	cleanNumber := regexPhone.ReplaceAllString(s, "")
	if strings.HasPrefix(cleanNumber, "0") {
		cleanNumber = "380" + cleanNumber[1:]
	} else if strings.HasPrefix(cleanNumber, "+380") {
		cleanNumber = strings.TrimPrefix(cleanNumber, "+")
	}
	if len(cleanNumber) != 12 {
		return "", errors.New("invalid phone number")
	}

	formattedPhone := fmt.Sprintf("(%s) %s-%s-%s", cleanNumber[2:5], cleanNumber[5:7], cleanNumber[7:9], cleanNumber[9:])

	return formattedPhone, nil
}

func main() {
	file, err := os.Open("tn.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	outputFile, err := os.Create("phones_normalized.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	err = PhoneNormalize(file, outputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Phone numbers normalized and saved to phones_normalized.txt")
}

func PhoneNormalize(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " - ")
		if len(parts) != 2 {
			continue
		}

		name := parts[0]
		phone, err := FormatNumber(parts[1])
		if err != nil {
			return err
		}

		_, err = writer.WriteString(fmt.Sprintf("%s - %s\n", name, phone))
		if err != nil {
			return err
		}

	}
	if err := scanner.Err(); err != nil {
		return err
	}

	err := writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
