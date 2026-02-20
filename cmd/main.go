package main

import (
	"flag"
	"os"
	"poster/config"
	"poster/parser"
	"poster/sender"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Аргументы командной строки
type Args struct {
	config  *string // Путь к файлу настроек
	file    *string // Путь к файлу сообщения
	verbose bool    // Отправлять больше логов
	senders []string // В какие источники отправить сообщения
}

// Парсит аргументы командной строки
func parseArgs() Args {
	config := flag.String("config", "config.toml", "Config file for poster")
	file := flag.String("file", "post.toml", "message file to load")
	verbose := flag.Bool("verbose", false, "Send debug logs")

	flag.Parse()
	args := flag.Args()
	return Args{config, file, *verbose, args}
}


func main() {
	// Setup pretty console writer
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	args := parseArgs()
	// Load project config
	cfg, err := config.LoadFile(*args.config)
	if err != nil {
		log.Fatal().Err(err).Msg("Load config file")
	}

	// Set logging level
	if args.verbose {
		log.Logger = log.Level(zerolog.DebugLevel)
		log.Debug().Msg("Set debug log level")
	}

	// Выбор отправителей
	senders := cfg.Senders
	if len(args.senders) > 0 {
		senders = &args.senders
	}
	if senders == nil {
		log.Fatal().Msg("Senders is not specify")
	}

	// Подготовка всех отправителей
	manager := sender.NewManager(cfg)

	// Получение сообщения
	message, err := parser.LoadMessage(*args.file)
	if err != nil {
		log.Fatal().Err(err).Msg("Load message from file")
	}

	// Отправка сообщения
	manager.Send(message, *senders)
}
