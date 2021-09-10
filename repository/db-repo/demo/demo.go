package demo

import (
	"fmt"
	"time"

	db_repo "gitlab.weimiaocaishang.com/weimiao/go-basic/repository/db-repo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewDemo() *WmAbout {
	return new(WmAbout)
}

func NewDemoQueryBuilder() *demoQueryBuilder {
	return new(demoQueryBuilder)
}

func (t *WmAbout) Create(db *gorm.DB) (id int, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type demoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *demoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *demoQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&WmAbout{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *demoQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&WmAbout{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *demoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&WmAbout{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *demoQueryBuilder) First(db *gorm.DB) (*WmAbout, error) {
	ret := &WmAbout{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *demoQueryBuilder) QueryOne(db *gorm.DB) (*WmAbout, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *demoQueryBuilder) QueryAll(db *gorm.DB) ([]*WmAbout, error) {
	var ret []*WmAbout
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *demoQueryBuilder) Limit(limit int) *demoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *demoQueryBuilder) Offset(offset int) *demoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *demoQueryBuilder) WhereId(p db_repo.Predicate, value int32) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereIdIn(value []int32) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereIdNotIn(value []int32) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) OrderById(asc bool) *demoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *demoQueryBuilder) WhereUsername(p db_repo.Predicate, value string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", p),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereUsernameIn(value []string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", "IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereUsernameNotIn(value []string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", "NOT IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) OrderByUsername(asc bool) *demoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "username "+order)
	return qb
}

func (qb *demoQueryBuilder) WhereNickname(p db_repo.Predicate, value string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", p),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereNicknameIn(value []string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", "IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereNicknameNotIn(value []string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", "NOT IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) OrderByNickname(asc bool) *demoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "nickname "+order)
	return qb
}

func (qb *demoQueryBuilder) WhereMobile(p db_repo.Predicate, value string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "mobile", p),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereMobileIn(value []string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "mobile", "IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereMobileNotIn(value []string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "mobile", "NOT IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) OrderByMobile(asc bool) *demoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "mobile "+order)
	return qb
}

func (qb *demoQueryBuilder) WhereIsUsed(p db_repo.Predicate, value int32) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", p),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereIsUsedIn(value []int32) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereIsUsedNotIn(value []int32) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "NOT IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) OrderByIsUsed(asc bool) *demoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_used "+order)
	return qb
}

func (qb *demoQueryBuilder) WhereIsDeleted(p db_repo.Predicate, value int32) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereIsDeletedIn(value []int32) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereIsDeletedNotIn(value []int32) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) OrderByIsDeleted(asc bool) *demoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}

func (qb *demoQueryBuilder) WhereCreatedAt(p db_repo.Predicate, value time.Time) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereCreatedAtIn(value []time.Time) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) OrderByCreatedAt(asc bool) *demoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *demoQueryBuilder) WhereCreatedUser(p db_repo.Predicate, value string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereCreatedUserIn(value []string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereCreatedUserNotIn(value []string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) OrderByCreatedUser(asc bool) *demoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *demoQueryBuilder) WhereUpdatedAt(p db_repo.Predicate, value time.Time) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereUpdatedAtIn(value []time.Time) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) OrderByUpdatedAt(asc bool) *demoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *demoQueryBuilder) WhereUpdatedUser(p db_repo.Predicate, value string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", p),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereUpdatedUserIn(value []string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) WhereUpdatedUserNotIn(value []string) *demoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *demoQueryBuilder) OrderByUpdatedUser(asc bool) *demoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}
