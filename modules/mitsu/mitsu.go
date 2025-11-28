package mitsu

import (
	"errors"
	"fmt"
	"goMH/core"
	"goMH/tui"
	"strings"
)

type Module struct{}

func (m *Module) ID() string {
	return "MITSU"
}

func (m *Module) MenuText() string {
	return "Установить Mitsu Utilities Pack"
}

func (m *Module) Run(am core.AssetManager, wu core.WinUtils) error {
	cfg := am.Cfg().MitsuConfig

	tui.Title(fmt.Sprintf("\n--- Начало установки: %s ---", m.MenuText()))

	if cfg.AssetID == "" {
		return errors.New("в секции 'mitsu_config' не указан asset_id")
	}

	tui.Info("Получение установщика через AssetManager...")
	installerPath, err := am.DownloadToCache(cfg.AssetID)
	if err != nil {
		return fmt.Errorf("не удалось получить ассет '%s': %w", cfg.AssetID, err)
	}

	tui.Info("Запуск установки...")
	tui.InfoF("Аргументы: %s", cfg.InstallArgs)
	args := strings.Fields(cfg.InstallArgs)

	_, err = wu.RunCommand(installerPath, args...)
	if err != nil {
		return fmt.Errorf("ошибка при установке Mitsu Utilities Pack: %w", err)
	}

	return nil
}
