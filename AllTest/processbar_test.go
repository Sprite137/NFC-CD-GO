package AllTest

import (
	"fmt"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"
)

func TestProcessBar(t *testing.T) {
	cmd := exec.Command("tput", "cols")
	cols, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting terminal width:", err)
		return
	}
	_, _ = strconv.Atoi(string(cols))

	// 创建进度条
	bar := progressbar.NewOptions(100,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionFullWidth(),
		progressbar.OptionShowDescriptionAtLineEnd(),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]-[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: "[red]-[reset]",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		progressbar.OptionSetPredictTime(false),
	)

	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(time.Millisecond * 100)

		// 移动光标到进度条开始位置
		moveCursorCmd := exec.Command("tput", "sc")
		moveCursorCmd.Stdout = os.Stdout
		moveCursorCmd.Run()

		// 清除当前行
		clearLineCmd := exec.Command("tput", "el")
		clearLineCmd.Stdout = os.Stdout
		clearLineCmd.Run()

		// 移动光标到进度条结束位置
		moveCursorToEndCmd := exec.Command("tput", "rc")
		moveCursorToEndCmd.Stdout = os.Stdout
		moveCursorToEndCmd.Run()
	}

	bar.Finish()

}
