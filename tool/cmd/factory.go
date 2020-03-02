package cmd

import (
	"github.com/spf13/cobra"
	//"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jukylin/esim/log"
	"github.com/jukylin/esim/tool/factory"
)

var factoryCmd = &cobra.Command{
	Use:   "factory",
	Short: "对model进行优化",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := factory.HandleModel(v)
		if err != nil {
			log.Log.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(factoryCmd)

	factoryCmd.Flags().BoolP("sort", "s", true, "按照内存对齐排序")

	factoryCmd.Flags().BoolP("new", "n", false, "生成New方法")

	factoryCmd.Flags().BoolP("star", "", false, "返回指针变量")

	factoryCmd.Flags().BoolP("option", "o", false, "New with option")

	factoryCmd.Flags().BoolP("gen_logger_option", "", false, "generate logger option")

	factoryCmd.Flags().BoolP("gen_conf_option", "", false, "generate conf option")

	factoryCmd.Flags().BoolP("pool", "p", false, "生成临时对象池")

	//modelCmd.Flags().BoolP("print", "", false, "扩展print方法")

	//modelCmd.Flags().BoolP("Print", "", false, "打印到终端")

	factoryCmd.Flags().StringP("modelname", "m", "", "模型的名称")

	factoryCmd.Flags().BoolP("plural", "", false, "支持复数")

	v.BindPFlags(factoryCmd.Flags())
}