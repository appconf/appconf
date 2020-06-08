package appconf

import (
	"github.com/appconf/engine"
	"github.com/appconf/log"
	"github.com/appconf/storage"
)

var c *engine.Core

//Run 启动appconf
func Run(cfg Config) error {
	logger := log.GetLogger()
	logger.SetLevel(cfg.Log.Level)

	backend, err := storage.New(cfg.Storage.Type, logger, cfg.Storage.Config)
	if err != nil {
		return err
	}

	c, err = engine.New(logger, backend, cfg.Templates)
	if err != nil {
		return err
	}

	return c.Run()
}

//Stop 停止appconf
func Stop() error {
	if c == nil {
		return nil
	}
	return c.Stop()
}
