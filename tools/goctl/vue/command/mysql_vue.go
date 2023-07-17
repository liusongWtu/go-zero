package command

import (
	"errors"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/vue/command/migrationnotes"
	"github.com/zeromicro/go-zero/tools/goctl/vue/gen"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/model"
	file "github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

var (
	// VarStringRequestPath describes whether the request path.
	VarStringRequestPath string
)

// MySqlDataSource generates model code from datasource
func MySqlDataSourceGenVue(_ *cobra.Command, _ []string) error {
	migrationnotes.BeforeCommands(VarStringDir, VarStringStyle)
	url := strings.TrimSpace(VarStringURL)
	dir := strings.TrimSpace(VarStringDir)
	cache := VarBoolCache
	idea := VarBoolIdea
	style := VarStringStyle
	home := VarStringHome
	remote := VarStringRemote
	branch := VarStringBranch
	requestPath := VarStringRequestPath
	if len(remote) > 0 {
		repo, _ := file.CloneIntoGitHome(remote, branch)
		if len(repo) > 0 {
			home = repo
		}
	}
	if len(home) > 0 {
		pathx.RegisterGoctlHome(home)
	}

	tableValue := VarStringSliceTable
	patterns := parseTableList(tableValue)
	cfg, err := config.NewConfig(style)
	if err != nil {
		return err
	}

	arg := dataSourceArg{
		url:           url,
		dir:           dir,
		tablePat:      patterns,
		cfg:           cfg,
		cache:         cache,
		idea:          idea,
		strict:        VarBoolStrict,
		ignoreColumns: mergeColumns(VarStringSliceIgnoreColumns),
		requestPath:   requestPath,
	}
	return fromMysqlDataSourceGenVue(arg)
}

func fromMysqlDataSourceGenVue(arg dataSourceArg) error {
	log := console.NewConsole(arg.idea)
	if len(arg.url) == 0 {
		log.Error("%v", "expected data source of mysql, but nothing found")
		return nil
	}

	if len(arg.tablePat) == 0 {
		log.Error("%v", "expected table or table globbing patterns, but nothing found")
		return nil
	}

	dsn, err := mysql.ParseDSN(arg.url)
	if err != nil {
		return err
	}

	logx.Disable()
	databaseSource := strings.TrimSuffix(arg.url, "/"+dsn.DBName) + "/information_schema"
	db := sqlx.NewMysql(databaseSource)
	im := model.NewInformationSchemaModel(db)

	tables, err := im.GetAllTables(dsn.DBName)
	if err != nil {
		return err
	}

	matchTables := make(map[string]*model.Table)
	for _, item := range tables {
		if !arg.tablePat.Match(item) {
			continue
		}

		columnData, err := im.FindColumns(dsn.DBName, item)
		if err != nil {
			return err
		}

		table, err := columnData.Convert()
		if err != nil {
			return err
		}

		matchTables[item] = table
	}

	if len(matchTables) == 0 {
		return errors.New("no tables matched")
	}

	generator, err := gen.NewDefaultGenerator(arg.dir, arg.cfg,
		gen.WithConsoleOption(log),
		gen.WithIgnoreColumns(arg.ignoreColumns),
		gen.WithRequestPath(arg.requestPath),
	)
	if err != nil {
		return err
	}

	return generator.StartFromInformationSchema(matchTables, arg.cache, arg.strict)
}
