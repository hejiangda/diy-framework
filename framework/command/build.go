package command

import (
	"fmt"
	"github.com/hejiangda/diy-framework/framework/cobra"
	"github.com/hejiangda/diy-framework/framework/contract"
	"log"
	"os/exec"
	"path/filepath"
)

var buildCommand = &cobra.Command{
	Use:   "build",
	Short: "编译相关命令",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			c.Help()
		}
		return nil
	},
}

// build相关的命令
func initBuildCommand() *cobra.Command {
	buildCommand.AddCommand(buildFrontendCommand)
	return buildCommand
}

// 打印前端的命令
var buildFrontendCommand = &cobra.Command{
	Use:   "frontend",
	Short: "使用npm编译前端",
	RunE: func(c *cobra.Command, args []string) error {
		// 获取path路径下的npm命令
		path, err := exec.LookPath("npm")
		if err != nil {
			log.Fatalln("请安装npm在你的PATH路径下")
		}

		// 执行npm run build
		cmd := exec.Command(path, "run", "build")
		app := c.GetContainer().MustMake(contract.AppKey).(contract.App)

		cmd.Dir = filepath.Join(app.BaseFolder(), "frontend")
		// 将输出保存在out中
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("=============  前端编译失败 ============")
			fmt.Println(string(out))
			fmt.Println("=============  前端编译失败 ============")
			return err
		}
		// 打印输出
		fmt.Print(string(out))
		fmt.Println("=============  前端编译成功 ============")
		return nil
	},
}
