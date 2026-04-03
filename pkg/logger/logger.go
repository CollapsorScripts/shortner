package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"shortner/internal/config"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

const Console = false

var mutex sync.Mutex
var fullPathToFile string

var blue = color.New(color.FgBlue).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

func colorizeJSON(in string) string {
	var buf bytes.Buffer
	dec := json.NewDecoder(strings.NewReader(in))
	dec.UseNumber()

	for {
		t, err := dec.Token()
		if err != nil {
			break
		}
		switch v := t.(type) {
		case json.Delim: // { } [ ]
			buf.WriteString(color.HiWhiteString("%s", v))
		case string:
			// Определяем ключ или значение
			if dec.More() { // значит это ключ
				buf.WriteString(blue(fmt.Sprintf("\"%s\"", v)))
			} else {
				buf.WriteString(green(fmt.Sprintf("\"%s\"", v)))
			}
		case json.Number:
			buf.WriteString(yellow(v.String()))
		case bool:
			buf.WriteString(red(fmt.Sprintf("%v", v)))
		case nil:
			buf.WriteString(red("null"))
		default:
			buf.WriteString(fmt.Sprintf("%v", v))
		}
		// Добавляем пробел после токена
		buf.WriteByte(' ')
	}

	return buf.String()
}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "   ")
	if err != nil {
		return in
	}
	return out.String()
}

// ToJSON - конвертирует объект в JSON строку
func toJSON(object any) string {
	jsonByte, err := json.Marshal(object)
	if err != nil {
		return ""
	}
	n := len(jsonByte)             //Find the length of the byte array
	result := string(jsonByte[:n]) //convert to string

	return colorizeJSON(jsonPrettyPrint(result))
}

// getFuncName - получаем имя функции из которой вызван логгер
func getFuncName() string {
	// 2 уровня вверх от текущей точки: 0 — runtime.Caller, 1 — getFuncName,
	// 2 — нужная нам функция
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "undefined"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "undefined"
	}

	// Получаем полное имя, например:
	//  github.com/user/project/pkg/sub.(*Type).Method-fm
	full := fn.Name()

	// Обрезаем до последнего элемента пути (последнего '/')
	// Теперь short может быть: "pkg/sub.(*Type).Method-fm"
	short := path.Base(full)

	// Если у нас есть суффикс "-fm" (метод-замыкание), убираем его
	if i := strings.Index(short, "-fm"); i >= 0 {
		short = short[:i]
	}

	// Разбиваем по последней точки, отделяя функцию/метод от пакета/типа
	lastDot := strings.LastIndex(short, ".")
	if lastDot < 0 {
		return fmt.Sprintf("[?]=>[%s]", short)
	}

	pkgPart := short[:lastDot]    // "sub.(*Type)" или просто "sub"
	funcPart := short[lastDot+1:] // "Method" или "Function"

	return fmt.Sprintf("[%s]=>[%s]", pkgPart, funcPart)
}

func writeToLog(str string, a ...any) error {
	mutex.Lock()
	// Открываем файл с логом
	file, err := os.OpenFile(fullPathToFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}

	defer func() {
		mutex.Unlock()
		file.Close()
	}()

	// Дозаписываем данные в файл
	data := fmt.Appendf(nil, str+"\n", a...)
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	fmt.Printf(str+"\n", a...)

	return nil
}

// New - инициализация логгера, создание папки и файла с логом
func New(cfg *config.Config) error {
	//Полный путь к файлу с логом
	fullPathToFile = filepath.Join(cfg.Paths.LogDir, cfg.Paths.LogName)

	// Создаем папку с правами доступа для текущего пользователя
	if err := os.MkdirAll(cfg.Paths.LogDir, os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}

	// Создаем файл с логом
	file, err := os.OpenFile(fullPathToFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	mutex = sync.Mutex{}

	return nil
}

// Info - лог с пометкой info
func Info(format string, a ...any) {
	funcName := fmt.Sprintf("%s ->", getFuncName())
	go func() {
		loc, err := time.LoadLocation("Europe/Moscow")
		if err != nil {
			fmt.Printf("Ошибка при загрузке локации: %v\n", err)
			return
		}
		date := time.Now().In(loc)
		dateFormat := fmt.Sprintf("[%s]", date.Format("02.01.2006 15.04"))
		str := fmt.Sprintf("%s %s [%s]: %s", dateFormat, funcName, blue("INFO"), format)
		if Console {
			fmt.Printf(str, a...)
			fmt.Println()
		} else {
			err := writeToLog(str, a...)
			if err != nil {
				fmt.Printf("Ошибка при записи в лог: %v\n", err)
				return
			}
		}
	}()
}

// Error - лог с пометкой error
func Error(format string, a ...any) {
	funcName := fmt.Sprintf("%s ->", getFuncName())
	go func() {
		loc, err := time.LoadLocation("Europe/Moscow")
		if err != nil {
			fmt.Printf("Ошибка при загрузке локации: %v\n", err)
			return
		}
		date := time.Now().In(loc)
		dateFormat := fmt.Sprintf("[%s]", date.Format("02.01.2006 15.04"))
		str := fmt.Sprintf("%s %s [%s]: %s", dateFormat, funcName, red("ERROR"), format)
		if Console {
			fmt.Printf(str, a...)
			fmt.Println()
		} else {
			err := writeToLog(str, a...)
			if err != nil {
				fmt.Printf("Ошибка при записи в лог: %v\n", err)
				return
			}
		}
	}()
}

// Warn - лог с пометкой warn
func Warn(format string, a ...any) {
	funcName := fmt.Sprintf("%s ->", getFuncName())
	go func() {
		loc, err := time.LoadLocation("Europe/Moscow")
		if err != nil {
			fmt.Printf("Ошибка при загрузке локации: %v\n", err)
			return
		}
		date := time.Now().In(loc)
		dateFormat := fmt.Sprintf("[%s]", date.Format("02.01.2006 15.04"))
		str := fmt.Sprintf("%s %s [%s]: %s", dateFormat, funcName, yellow("WARN"), format)
		if Console {
			fmt.Printf(str, a...)
			fmt.Println()
		} else {
			err := writeToLog(str, a...)
			if err != nil {
				fmt.Printf("Ошибка при записи в лог: %v\n", err)
				return
			}
		}
	}()
}
