package command

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func RunNetworkQualityOneLine() string {
	out := exec.Command("networkQuality")
	stdout, err := out.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	out.Start()
	scanner := bufio.NewScanner(stdout)
	text := ""
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		text += scanner.Text() + "\n"
	}
	out.Wait()

	return text
}

func RunNetworkQualityAllLine() string {
	/*
		networkQualityコマンド実行
		Output()で全行取得
	*/
	out, err := exec.Command("networkQuality").Output()

	// エラー処理
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	text := string(out)

	return text
}
