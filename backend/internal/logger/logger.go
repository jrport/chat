package logger

import (
	"fmt"
	"log/slog"
	"os"
)

func SetupLogger(filepath string) (*slog.Logger, error){
	fh, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("Error creating/opening filepath: %v", err.Error())
	}

	opts := slog.HandlerOptions{AddSource: true}
	applog := slog.New(slog.NewJSONHandler(fh, &opts))
	return applog, nil
}
