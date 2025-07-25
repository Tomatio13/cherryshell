package main

import (
	"fmt"
	"os"

	"cherrysh/config"
	"cherrysh/i18n"
	"cherrysh/shell"
)

func main() {
	// 設定を読み込む
	cfg := config.NewConfig()
	if err := cfg.LoadConfigFile(); err != nil {
		fmt.Printf("Warning: Could not load config file: %v\n", err)
	}

	// 言語を取得
	language := cfg.GetLanguage(os.Args)

	// 国際化を初期化
	if err := i18n.Init(language); err != nil {
		fmt.Printf("Warning: Could not initialize i18n: %v\n", err)
		// フォールバックとして英語で初期化
		i18n.Init("en")
	}

	// アプリケーション情報を表示
	// fmt.Println(i18n.T("app.title"))
	// fmt.Println(i18n.T("app.description"))
	// fmt.Println(i18n.T("app.exit_instruction"))
	// fmt.Println(i18n.T("shell.runtime_separator"))
	// fmt.Println(i18n.T("shell.runtime_info"))
	// fmt.Printf(i18n.T("shell.runtime_os")+"\n", runtime.GOOS)
	// fmt.Printf(i18n.T("shell.runtime_arch")+"\n", runtime.GOARCH)
	// fmt.Println(i18n.T("shell.runtime_separator"))

	// シェルを開始
	s := shell.NewShell(cfg)
	s.Start()
}
