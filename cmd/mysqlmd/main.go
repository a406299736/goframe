package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/a406299736/goframe/configs"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/a406299736/goframe/cmd/mysqlmd/mysql"

	"gorm.io/gorm"
)

type tableInfo struct {
	Name    string         `db:"table_name"`    // name
	Comment sql.NullString `db:"table_comment"` // comment
}

type tableColumn struct {
	OrdinalPosition uint16         `db:"ORDINAL_POSITION"` // position
	ColumnName      string         `db:"COLUMN_NAME"`      // name
	ColumnType      string         `db:"COLUMN_TYPE"`      // column_type
	DataType        string         `db:"DATA_TYPE"`        // data_type
	ColumnKey       sql.NullString `db:"COLUMN_KEY"`       // key
	IsNullable      string         `db:"IS_NULLABLE"`      // nullable
	Extra           sql.NullString `db:"EXTRA"`            // extra
	ColumnComment   sql.NullString `db:"COLUMN_COMMENT"`   // comment
	ColumnDefault   sql.NullString `db:"COLUMN_DEFAULT"`   // default value
}

var (
	dbAddr    string
	dbUser    string
	dbPass    string
	dbName    string
	genTables string
)

func init() {
	flag.StringVar(&dbName, "name", "databaseName", "请输入 db 名称\n")
	flag.StringVar(&genTables, "tables", "", "请输入 table 名称，默认为“ ”，如需生成所有表请输入“*”，多个表可用“,”分割\n")
	flag.Parse()

	conf := configs.Get().MySQL
	switch dbName {
	case "conn":
		dbAddr = conf.Read.Addr
		dbUser = conf.Read.User
		dbPass = conf.Read.Pass
	case "conn1":
		dbAddr = conf.Conn1read.Addr
		dbUser = conf.Conn1read.User
		dbPass = conf.Conn1read.Pass
	default:
		dbAddr = conf.Read.Addr
		dbUser = conf.Read.User
		dbPass = conf.Read.Pass
	}
}

func main() {
	// 初始化 DB
	db, err := mysql.New(dbAddr, dbUser, dbPass, dbName)
	if err != nil {
		log.Fatal("new db err", err)
	}

	defer func() {
		if err := db.DbClose(); err != nil {
			log.Println("db close err", err)
		}
	}()

	tables, err := queryTables(db.GetDb(), dbName, genTables)
	if err != nil {
		log.Println("query tables of database err", err)
		return
	}

	for _, table := range tables {

		filepath := "./repository/dbrepo/" + table.Name
		_ = os.Mkdir(filepath, 0766)
		fmt.Println("create dir : ", filepath)

		modelName := fmt.Sprintf("%s/gen_model.go", filepath)
		modelFile, err := os.OpenFile(modelName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0766)
		if err != nil {
			fmt.Printf("create and open model file error %v\n", err.Error())
			return
		}
		fmt.Println("  └── file : ", table.Name+"/gen_model.go")

		columnInfo, columnInfoErr := queryTableColumn(db.GetDb(), dbName, table.Name)
		if columnInfoErr != nil {
			continue
		}

		modelContent := fmt.Sprintf("package %s\n", table.Name)
		for _, info := range columnInfo {
			if textType(info.DataType) == "time.Time" {
				modelContent += fmt.Sprintf(`import "time"`)
				break
			}
		}
		modelContent += fmt.Sprintf("\n\n// %s %s \n", capitalize(table.Name), table.Comment.String)
		modelContent += fmt.Sprintf("//go:generate gormgen -structs %s -input . \n", capitalize(table.Name))
		modelContent += fmt.Sprintf("type %s struct {\n", capitalize(table.Name))

		for _, info := range columnInfo {
			if textType(info.DataType) == "time.Time" {
				modelContent += fmt.Sprintf("%s %s `%s json:\"%s\" form:\"%s\"` // %s\n", capitalize(info.ColumnName), textType(info.DataType), "gorm:\"time\"", info.ColumnName, info.ColumnName, info.ColumnComment.String)
			} else {
				modelContent += fmt.Sprintf("%s %s `json:\"%s\" form:\"%s\"` // %s\n", capitalize(info.ColumnName), textType(info.DataType), info.ColumnName, info.ColumnName, info.ColumnComment.String)
			}
		}

		modelContent += "}\n"
		modelFile.WriteString(modelContent)
		modelFile.Close()

	}

}

func queryTables(db *gorm.DB, dbName string, tableName string) ([]tableInfo, error) {
	var tableCollect []tableInfo
	var tableArray []string
	var commentArray []sql.NullString

	sqlTables := fmt.Sprintf("SELECT `table_name`,`table_comment` FROM `information_schema`.`tables` WHERE `table_schema`= '%s'", dbName)
	rows, err := db.Raw(sqlTables).Rows()
	if err != nil {
		return tableCollect, err
	}
	defer rows.Close()

	for rows.Next() {
		var info tableInfo
		err = rows.Scan(&info.Name, &info.Comment)
		if err != nil {
			fmt.Printf("execute query tables action error,had ignored, detail is [%v]\n", err.Error())
			continue
		}

		tableCollect = append(tableCollect, info)
		tableArray = append(tableArray, info.Name)
		commentArray = append(commentArray, info.Comment)
	}

	// filter tables when specified tables params
	if tableName != "*" {
		tableCollect = nil
		chooseTables := strings.Split(tableName, ",")
		indexMap := make(map[int]int)
		for _, item := range chooseTables {
			subIndexMap := fullMatchIndexMap(tableArray, item)
			for k, v := range subIndexMap {
				if _, ok := indexMap[k]; ok {
					continue
				}
				indexMap[k] = v
			}
		}

		if len(indexMap) != 0 {
			for _, v := range indexMap {
				var info tableInfo
				info.Name = tableArray[v]
				info.Comment = commentArray[v]
				tableCollect = append(tableCollect, info)
			}
		}
	}

	return tableCollect, err
}

func queryTableColumn(db *gorm.DB, dbName string, tableName string) ([]tableColumn, error) {
	// 定义承载列信息的切片
	var columns []tableColumn

	sqlTableColumn := fmt.Sprintf("SELECT `ORDINAL_POSITION`,`COLUMN_NAME`,`COLUMN_TYPE`,`DATA_TYPE`,`COLUMN_KEY`,`IS_NULLABLE`,`EXTRA`,`COLUMN_COMMENT`,`COLUMN_DEFAULT` FROM `information_schema`.`columns` WHERE `table_schema`= '%s' AND `table_name`= '%s' ORDER BY `ORDINAL_POSITION` ASC",
		dbName, tableName)

	rows, err := db.Raw(sqlTableColumn).Rows()
	if err != nil {
		fmt.Printf("execute query table column action error, detail is [%v]\n", err.Error())
		return columns, err
	}
	defer rows.Close()

	for rows.Next() {
		var column tableColumn
		err = rows.Scan(
			&column.OrdinalPosition,
			&column.ColumnName,
			&column.ColumnType,
			&column.DataType,
			&column.ColumnKey,
			&column.IsNullable,
			&column.Extra,
			&column.ColumnComment,
			&column.ColumnDefault)
		if err != nil {
			fmt.Printf("query table column scan error, detail is [%v]\n", err.Error())
			return columns, err
		}
		columns = append(columns, column)
	}

	return columns, err
}

// 模糊匹配表名
func getTargetIndexMap(tableNameArr []string, item string) map[int]int {
	indexMap := make(map[int]int)
	for i := 0; i < len(tableNameArr); i++ {
		if match, _ := regexp.MatchString(item, tableNameArr[i]); match {
			if _, ok := indexMap[i]; ok {
				continue
			}
			indexMap[i] = i
		}
	}
	return indexMap
}

// 全匹配表名
func fullMatchIndexMap(tableNameArr []string, item string) map[int]int {
	indexMap := make(map[int]int)
	for i := 0; i < len(tableNameArr); i++ {
		if tableNameArr[i] == item {
			if _, ok := indexMap[i]; ok {
				continue
			}
			indexMap[i] = i
		}
	}
	return indexMap
}

func capitalize(s string) string {
	var upperStr string
	chars := strings.Split(s, "_")
	for _, val := range chars {
		vv := []rune(val)
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				if vv[i] >= 97 && vv[i] <= 122 {
					vv[i] -= 32
					upperStr += string(vv[i])
				}
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}

func textType(s string) string {
	var mysqlTypeToGoType = map[string]string{
		"tinyint":    "int",
		"smallint":   "int",
		"mediumint":  "int",
		"int":        "int",
		"integer":    "int64",
		"bigint":     "int64",
		"float":      "float64",
		"double":     "float64",
		"decimal":    "float64",
		"date":       "string",
		"time":       "string",
		"year":       "string",
		"datetime":   "time.Time",
		"timestamp":  "time.Time",
		"char":       "string",
		"varchar":    "string",
		"tinyblob":   "string",
		"tinytext":   "string",
		"blob":       "string",
		"text":       "string",
		"mediumblob": "string",
		"mediumtext": "string",
		"longblob":   "string",
		"longtext":   "string",
	}
	return mysqlTypeToGoType[s]
}
