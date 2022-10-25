package cfg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type GlobalCfg struct {
	DeviceListSize int
	Token          string
	Secret         string
}

// 設定ファイルのスキーマ
type cfgFile struct {
	DeviceListSize int    `json:"deviceListSize"`
	Token          string `json:"token"`
	Secret         string `json:"secret"`
}

func NewGlobalConfig() GlobalCfg {
	return GlobalCfg{}
}

func (g *GlobalCfg) Prepare(path string) error {
	if path != "" {
		if err := g.setConfigFileContentVal(path); err == nil {
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

	if err := g.setConfigFileContentVal(defaultCfgFilePath); err != nil {
		return fmt.Errorf("設定ファイルの読み込みに失敗しました。設定ファイルが存在しない または ファイルのスキーマが不正です。実際のエラー: %v", err)
	}

	return nil
}

func (g *GlobalCfg) setConfigFileContentVal(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var cfgFileContent cfgFile
	if err := json.Unmarshal(b, &cfgFileContent); err != nil {
		return err
	}

	g.DeviceListSize = cfgFileContent.DeviceListSize
	g.Token = cfgFileContent.Token
	g.Secret = cfgFileContent.Secret

	return nil
}
