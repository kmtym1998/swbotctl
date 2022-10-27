package cfg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type GlobalCfg struct {
	DeviceListSize int    `json:"deviceListSize"`
	Token          string `json:"token"`
	Secret         string `json:"secret"`
}

func NewGlobalConfig() GlobalCfg {
	return GlobalCfg{}
}

func (g *GlobalCfg) Prepare(path string) error {
	if path != "" {
		if err := g.setCfgFileContent(path); err == nil {
			return nil
		}
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	cfgDirName := ".swbotctl"
	cfgFileName := "config.json"
	defaultCfgFilePath := filepath.Join(homeDir, cfgDirName, cfgFileName)

	if err := g.setCfgFileContent(defaultCfgFilePath); err != nil {
		return fmt.Errorf("設定ファイルの読み込みに失敗しました。設定ファイルが存在しない または ファイルのスキーマが不正です。実際のエラー: %v", err)
	}

	return nil
}

func (g *GlobalCfg) setCfgFileContent(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("⚠️ 設定ファイルが見つかりません。`swbotctl init` を実行してください。")
			return nil
		}
		return err
	}

	var cfgFileContent GlobalCfg
	if err := json.Unmarshal(b, &cfgFileContent); err != nil {
		return err
	}

	g.DeviceListSize = cfgFileContent.DeviceListSize
	g.Secret = cfgFileContent.Secret
	g.Token = cfgFileContent.Token

	return nil
}
