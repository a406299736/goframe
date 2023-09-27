package pkg

import "text/template"

// Make sure that the template compiles during package initialization
func parseTemplateOrPanic(t string) *template.Template {
	tpl, err := template.New("output_template").Parse(t)
	if err != nil {
		panic(err)
	}
	return tpl
}

var outputTemplate = parseTemplateOrPanic(`
///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package {{.PkgName}}

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"time"

	"github.com/a406299736/goframe/repository/dbrepo"
	e "github.com/a406299736/goframe/pkg/errors"
	"github.com/a406299736/goframe/app/pkg/code"

	"gorm.io/gorm"
)

func NewModel() *{{.StructName}} {
	return new({{.StructName}})
}

func NewQueryBuilder() *{{.QueryBuilderName}} {
	return new({{.QueryBuilderName}})
}

func (t *{{.StructName}}) Create(db *gorm.DB) (id int32, er e.Er) {
	if err := db.Create(t).Error; err != nil {
		var sqlError *mysql.MySQLError
		if errors.As(err, &sqlError) && sqlError.Number == 1062 {
			return 0, e.NewErr(code.UniqueKeyConflict, err.Error())
		}
		return 0, e.NewErr(code.MySQLExecError, err.Error())
	}
	return t.Id, nil
}

type {{.QueryBuilderName}} struct {
	fields []string
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *{{.QueryBuilderName}}) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	if len(qb.fields) > 0 {
		ret = ret.Select(qb.fields)
	}
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	if qb.limit > 0 {
		ret = ret.Limit(qb.limit)
	}
	ret = ret.Offset(qb.offset)
	return ret
}

func (qb *{{.QueryBuilderName}}) Select(fields []string) *{{.QueryBuilderName}} {
	qb.fields = fields
	return qb
}

func (qb *{{.QueryBuilderName}}) Updates(db *gorm.DB, m map[string]interface{}) (er e.Er) {
	db = db.Model(&{{.StructName}}{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if qb.limit > 0 {
		db.Limit(qb.limit)
	}

	for _, order := range qb.order {
		db = db.Order(order)
	}

	if err := db.Updates(m).Error; err != nil {
		return e.NewErr(code.MySQLExecError, err.Error())
	}
	return nil
}

func (qb *{{.QueryBuilderName}}) Delete(db *gorm.DB) (er e.Er) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if qb.limit > 0 {
		db.Limit(qb.limit)
	}

	for _, order := range qb.order {
		db = db.Order(order)
	}

	if err := db.Delete(&{{.StructName}}{}).Error; err != nil {
		return e.NewErr(code.MySQLExecError, err.Error())
	}
	return nil
}

func (qb *{{.QueryBuilderName}}) Count(db *gorm.DB) (int64, e.Er) {
	var c int64
	res := qb.buildQuery(db).Model(&{{.StructName}}{}).Count(&c)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return 0, e.NewErr(code.QueryNotExist, res.Error.Error())
		} else {
			return 0, e.NewErr(code.MySQLExecError, res.Error.Error())
		}
	}
	return c, nil
}

func (qb *{{.QueryBuilderName}}) First(db *gorm.DB) (*{{.StructName}}, e.Er) {
	ret := &{{.StructName}}{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, e.NewErr(code.QueryNotExist, res.Error.Error())
		} else {
			return nil, e.NewErr(code.MySQLExecError, res.Error.Error())
		}
	}
	return ret, nil
}

func (qb *{{.QueryBuilderName}}) QueryOne(db *gorm.DB) (*{{.StructName}}, e.Er) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *{{.QueryBuilderName}}) QueryAll(db *gorm.DB) ([]*{{.StructName}}, e.Er) {
	var ret []*{{.StructName}}
	err := qb.buildQuery(db).Find(&ret).Error
	if err != nil {
		return nil, e.NewErr(code.MySQLExecError, err.Error())
	}
	return ret, nil
}

func (qb *{{.QueryBuilderName}}) Limit(limit int) *{{.QueryBuilderName}} {
	qb.limit = limit
	return qb
}

func (qb *{{.QueryBuilderName}}) Offset(offset int) *{{.QueryBuilderName}} {
	qb.offset = offset
	return qb
}

{{$queryBuilderName := .QueryBuilderName}}
{{range .OptionFields}}
func (qb *{{$queryBuilderName}}) Where{{call $.Helpers.Titelize .FieldName}}(p dbrepo.Predicate, value {{.FieldType}}) *{{$queryBuilderName}} {
	 qb.where = append(qb.where, struct {
		prefix string
		value interface{}
	}{
		fmt.Sprintf("%v %v ?", "{{.ColumnName}}", p),
		value,
	})
	return qb
}

func (qb *{{$queryBuilderName}}) Where{{call $.Helpers.Titelize .FieldName}}In(value []{{.FieldType}}) *{{$queryBuilderName}} {
	 qb.where = append(qb.where, struct {
		prefix string
		value interface{}
	}{
		fmt.Sprintf("%v %v ?", "{{.ColumnName}}", "IN"),
		value,
	})
	return qb
}

func (qb *{{$queryBuilderName}}) Where{{call $.Helpers.Titelize .FieldName}}NotIn(value []{{.FieldType}}) *{{$queryBuilderName}} {
	 qb.where = append(qb.where, struct {
		prefix string
		value interface{}
	}{
		fmt.Sprintf("%v %v ?", "{{.ColumnName}}", "NOT IN"),
		value,
	})
	return qb
}

func (qb *{{$queryBuilderName}}) OrderBy{{call $.Helpers.Titelize .FieldName}}(asc bool) *{{$queryBuilderName}} {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "{{.ColumnName}} " + order)
	return qb
}
{{end}}
`)
