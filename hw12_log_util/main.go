package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

type appConfig struct {
	File   string
	Level  string
	Output string
}

type status struct {
	ip     string
	level  string
	method string
	code   string
	engine string
}

func (cfg *appConfig) ConfigFile(value string) {
	// fmt.Println(value)
	switch {
	case value == "":
		cfg.File = os.Getenv("LOG_ANALYZER_FILE")
	default:
		cfg.File = value
	}
}

func (cfg *appConfig) ConfigLevel(value string) {
	switch value {
	case "":
		cfg.Level = os.Getenv("LOG_ANALYZER_LEVEL")
	default:
		cfg.Level = value
	}
}

func (cfg *appConfig) ConfigOutput(value string) {
	switch value {
	case "":
		cfg.Output = os.Getenv("LOG_ANALYZER_OUTPUT")
	default:
		cfg.Output = value
	}
}

func ReadFile(logfile string) ([]string, error) {
	words := []string{}
	file, err := os.Open(logfile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, nil
}

func (s *status) Add(ip, level, method, code, engine string) {
	s.ip = ip
	s.level = level
	s.method = method
	s.code = code
	s.engine = engine
}

func sort(sl []status, level string) map[string]map[string]int64 {
	m := make(map[string]map[string]int64)
	m["ip"] = make(map[string]int64)
	m["method"] = make(map[string]int64)
	m["code"] = make(map[string]int64)
	m["engine"] = make(map[string]int64)
	for s := range sl {
		if sl[s].level != strings.ToUpper(level) {
			continue
		}
		_, ok := m["ip"][sl[s].ip]
		if ok {
			m["ip"][sl[s].ip]++
		} else {
			m["ip"][sl[s].ip] = 1
		}
		_, okmethod := m["method"][sl[s].method]

		if okmethod {
			m["method"][sl[s].method]++
		} else {
			m["method"][sl[s].method] = 1
		}
		_, okengine := m["engine"][sl[s].engine]
		if okengine {
			m["engine"][sl[s].engine]++
		} else {
			m["engine"][sl[s].engine] = 1
		}

		_, okstatus := m["code"][sl[s].code]
		if okstatus {
			m["code"][sl[s].code]++
		} else {
			m["code"][sl[s].code] = 1
		}
	}
	return m
}

func writeFile(input map[string]map[string]int64, pathfile string) error {
	file, err := os.Create(pathfile)
	if err != nil {
		return err
	}

	defer file.Close()
	for k, v := range input {
		for kk, vv := range v {
			_, err1 := file.WriteString(fmt.Sprintf("[%s][%s][%d]\n", k, kk, vv))
			if err1 != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	var (
		logAnalyzerFile   string
		logAnalyzerLevel  string
		logAnalyzerOutput string
		showHelp          bool
	)

	pflag.StringVarP(&logAnalyzerFile, "file", "i", "",
		"Input file")
	pflag.StringVarP(&logAnalyzerLevel, "level", "l", "",
		"log level: DEBUG, INFO, ERROR")
	pflag.StringVarP(&logAnalyzerOutput, "output", "o", "",
		"Out file")
	pflag.BoolVarP(&showHelp, "help", "h", false,
		"Show help message")
	pflag.Parse()
	if showHelp {
		pflag.Usage()
		return
	}

	c := appConfig{}
	c.ConfigFile(logAnalyzerFile)
	c.ConfigLevel(logAnalyzerLevel)
	c.ConfigOutput(logAnalyzerOutput)

	if c.File == "" {
		panic("not log file")
	}

	if c.Level == "" {
		c.Level = "info"
	}

	// var st []status
	struc := status{}
	stringlog, err := ReadFile(c.File)
	st := make([]status, 0, len(stringlog))
	if err != nil {
		fmt.Println(err)
	}
	for s := range stringlog {
		// fmt.Println(stringlog[s])
		splitstring := strings.Fields(stringlog[s])
		ip := splitstring[0]
		level := splitstring[1]
		method := splitstring[4]
		code := splitstring[5]
		engine := splitstring[6]

		struc.Add(ip, level, method, code, engine)
		st = append(st, struc)
	}
	rr := sort(st, c.Level)
	if c.Output == "" {
		for k, v := range rr {
			for kk, vv := range v {
				fmt.Println(k, kk, vv)
			}
		}
	} else {
		writeFile(rr, c.Output)
	}
}
