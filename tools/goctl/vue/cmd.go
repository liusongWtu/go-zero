package vue

import (
	"github.com/zeromicro/go-zero/tools/goctl/internal/cobrax"
	"github.com/zeromicro/go-zero/tools/goctl/vue/command"
)

var (
	Cmd                 = cobrax.NewCommand("vue")
	mysqlCmd            = cobrax.NewCommand("mysql")
	ddlCmd              = cobrax.NewCommand("ddl", cobrax.WithRunE(command.MysqlDDL))
	datasourceGenVueCmd = cobrax.NewCommand("datasource-gen-vue", cobrax.WithRunE(command.MySqlDataSourceGenVue))
)

func init() {
	var (
		ddlCmdFlags              = ddlCmd.Flags()
		datasourceGenVueCmdFlags = datasourceGenVueCmd.Flags()
	)

	ddlCmdFlags.StringVarP(&command.VarStringSrc, "src", "s")
	ddlCmdFlags.StringVarP(&command.VarStringDir, "dir", "d")
	ddlCmdFlags.StringVar(&command.VarStringStyle, "style")
	ddlCmdFlags.BoolVarP(&command.VarBoolCache, "cache", "c")
	ddlCmdFlags.BoolVar(&command.VarBoolIdea, "idea")
	ddlCmdFlags.StringVar(&command.VarStringDatabase, "database")
	ddlCmdFlags.StringVar(&command.VarStringHome, "home")
	ddlCmdFlags.StringVar(&command.VarStringRemote, "remote")
	ddlCmdFlags.StringVar(&command.VarStringBranch, "branch")

	datasourceGenVueCmdFlags.StringVar(&command.VarStringURL, "url")
	datasourceGenVueCmdFlags.StringSliceVarP(&command.VarStringSliceTable, "table", "t")
	datasourceGenVueCmdFlags.BoolVarP(&command.VarBoolCache, "cache", "c")
	datasourceGenVueCmdFlags.StringVarP(&command.VarStringDir, "dir", "d")
	datasourceGenVueCmdFlags.StringVar(&command.VarStringStyle, "style")
	datasourceGenVueCmdFlags.BoolVar(&command.VarBoolIdea, "idea")
	datasourceGenVueCmdFlags.StringVar(&command.VarStringHome, "home")
	datasourceGenVueCmdFlags.StringVar(&command.VarStringRemote, "remote")
	datasourceGenVueCmdFlags.StringVar(&command.VarStringBranch, "branch")
	datasourceGenVueCmdFlags.StringVar(&command.VarStringRequestPath, "requestpath")

	mysqlCmd.PersistentFlags().BoolVar(&command.VarBoolStrict, "strict")
	mysqlCmd.PersistentFlags().StringSliceVarPWithDefaultValue(&command.VarStringSliceIgnoreColumns, "ignore-columns", "i", []string{"create_at", "created_at", "create_time", "update_at", "updated_at", "update_time"})

	mysqlCmd.AddCommand(datasourceGenVueCmd, ddlCmd)
	Cmd.AddCommand(mysqlCmd)
}
