// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"

	"github.com/go-cinch/layout/internal/data/model"
)

func newGame(db *gorm.DB, opts ...gen.DOOption) game {
	_game := game{}

	_game.gameDo.UseDB(db, opts...)
	_game.gameDo.UseModel(&model.Game{})

	tableName := _game.gameDo.TableName()
	_game.ALL = field.NewAsterisk(tableName)
	_game.ID = field.NewUint64(tableName, "id")
	_game.CreatedAt = field.NewField(tableName, "created_at")
	_game.UpdatedAt = field.NewField(tableName, "updated_at")
	_game.Name = field.NewString(tableName, "name")

	_game.fillFieldMap()

	return _game
}

type game struct {
	gameDo gameDo

	ALL       field.Asterisk
	ID        field.Uint64 // auto increment id
	CreatedAt field.Field  // create time
	UpdatedAt field.Field  // update time
	Name      field.String // name

	fieldMap map[string]field.Expr
}

func (g game) Table(newTableName string) *game {
	g.gameDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g game) As(alias string) *game {
	g.gameDo.DO = *(g.gameDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *game) updateTableName(table string) *game {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewUint64(table, "id")
	g.CreatedAt = field.NewField(table, "created_at")
	g.UpdatedAt = field.NewField(table, "updated_at")
	g.Name = field.NewString(table, "name")

	g.fillFieldMap()

	return g
}

func (g *game) WithContext(ctx context.Context) *gameDo { return g.gameDo.WithContext(ctx) }

func (g game) TableName() string { return g.gameDo.TableName() }

func (g game) Alias() string { return g.gameDo.Alias() }

func (g game) Columns(cols ...field.Expr) gen.Columns { return g.gameDo.Columns(cols...) }

func (g *game) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *game) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 4)
	g.fieldMap["id"] = g.ID
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
	g.fieldMap["name"] = g.Name
}

func (g game) clone(db *gorm.DB) game {
	g.gameDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g game) replaceDB(db *gorm.DB) game {
	g.gameDo.ReplaceDB(db)
	return g
}

type gameDo struct{ gen.DO }

// SELECT * FROM `@@table` WHERE `id` = @id LIMIT 1
func (g gameDo) GetByID(id uint64) (result model.Game) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM `game` WHERE `id` = ? LIMIT 1 ")

	var executeSQL *gorm.DB
	executeSQL = g.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	_ = executeSQL

	return
}

// SELECT * FROM `@@table`
// {{where}}
//
//	{{if val != ""}}
//	  {{if strings.HasPrefix(val, "%") && strings.HasSuffix(val, "%")}}
//	    @@col LIKE CONCAT('%', TRIM(BOTH '%' FROM @val), '%')
//	  {{else if strings.HasPrefix(val, "%")}}
//	    @@col LIKE CONCAT('%', TRIM(BOTH '%' FROM @val))
//	  {{else if strings.HasSuffix(val, "%")}}
//	    @@col LIKE CONCAT(TRIM(BOTH '%' FROM @val), '%')
//	  {{else}}
//	    @@col = @val
//	  {{end}}
//	{{end}}
//
// {{end}}
// LIMIT 1
func (g gameDo) GetByCol(col string, val string) (result model.Game) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM `game` ")
	var whereSQL0 strings.Builder
	if val != "" {
		if strings.HasPrefix(val, "%") && strings.HasSuffix(val, "%") {
			params = append(params, val)
			whereSQL0.WriteString(g.Quote(col) + " LIKE CONCAT('%', TRIM(BOTH '%' FROM ?), '%') ")
		} else if strings.HasPrefix(val, "%") {
			params = append(params, val)
			whereSQL0.WriteString(g.Quote(col) + " LIKE CONCAT('%', TRIM(BOTH '%' FROM ?)) ")
		} else if strings.HasSuffix(val, "%") {
			params = append(params, val)
			whereSQL0.WriteString(g.Quote(col) + " LIKE CONCAT(TRIM(BOTH '%' FROM ?), '%') ")
		} else {
			params = append(params, val)
			whereSQL0.WriteString(g.Quote(col) + " = ? ")
		}
	}
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)
	generateSQL.WriteString("LIMIT 1 ")

	var executeSQL *gorm.DB
	executeSQL = g.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	_ = executeSQL

	return
}

// SELECT * FROM `@@table`
// {{if len(cols) == len(vals)}}
// {{where}}
//
//	  {{for i, col := range cols}}
//	    {{for j, val := range vals}}
//	      {{if i == j}}
//	        {{if val != ""}}
//	          {{if strings.HasPrefix(val, "%") && strings.HasSuffix(val, "%")}}
//	            @@col LIKE CONCAT('%', TRIM(BOTH '%' FROM @val), '%') AND
//	          {{else if strings.HasPrefix(val, "%")}}
//	            @@col LIKE CONCAT('%', TRIM(BOTH '%' FROM @val)) AND
//	          {{else if strings.HasSuffix(val, "%")}}
//	            @@col LIKE CONCAT(TRIM(BOTH '%' FROM @val), '%') AND
//	          {{else}}
//	            @@col = @val AND
//	          {{end}}
//	        {{end}}
//	      {{end}}
//	    {{end}}
//	  {{end}}
//	{{end}}
//
// {{end}}
// LIMIT 1
func (g gameDo) GetByCols(cols []string, vals []string) (result model.Game) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM `game` ")
	if len(cols) == len(vals) {
		var whereSQL0 strings.Builder
		for i, col := range cols {
			for j, val := range vals {
				if i == j {
					if val != "" {
						if strings.HasPrefix(val, "%") && strings.HasSuffix(val, "%") {
							params = append(params, val)
							whereSQL0.WriteString(g.Quote(col) + " LIKE CONCAT('%', TRIM(BOTH '%' FROM ?), '%') AND ")
						} else if strings.HasPrefix(val, "%") {
							params = append(params, val)
							whereSQL0.WriteString(g.Quote(col) + " LIKE CONCAT('%', TRIM(BOTH '%' FROM ?)) AND ")
						} else if strings.HasSuffix(val, "%") {
							params = append(params, val)
							whereSQL0.WriteString(g.Quote(col) + " LIKE CONCAT(TRIM(BOTH '%' FROM ?), '%') AND ")
						} else {
							params = append(params, val)
							whereSQL0.WriteString(g.Quote(col) + " = ? AND ")
						}
					}
				}
			}
		}
		helper.JoinWhereBuilder(&generateSQL, whereSQL0)
	}
	generateSQL.WriteString("LIMIT 1 ")

	var executeSQL *gorm.DB
	executeSQL = g.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	_ = executeSQL

	return
}

// SELECT * FROM `@@table`
// {{where}}
//
//	{{if val != ""}}
//	  {{if strings.HasPrefix(val, "%") && strings.HasSuffix(val, "%")}}
//	    @@col LIKE CONCAT('%', TRIM(BOTH '%' FROM @val), '%')
//	  {{else if strings.HasPrefix(val, "%")}}
//	    @@col LIKE CONCAT('%', TRIM(BOTH '%' FROM @val))
//	  {{else if strings.HasSuffix(val, "%")}}
//	    @@col LIKE CONCAT(TRIM(BOTH '%' FROM @val), '%')
//	  {{else}}
//	    @@col = @val
//	  {{end}}
//	{{end}}
//
// {{end}}
func (g gameDo) FindByCol(col string, val string) (result []model.Game) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM `game` ")
	var whereSQL0 strings.Builder
	if val != "" {
		if strings.HasPrefix(val, "%") && strings.HasSuffix(val, "%") {
			params = append(params, val)
			whereSQL0.WriteString(g.Quote(col) + " LIKE CONCAT('%', TRIM(BOTH '%' FROM ?), '%') ")
		} else if strings.HasPrefix(val, "%") {
			params = append(params, val)
			whereSQL0.WriteString(g.Quote(col) + " LIKE CONCAT('%', TRIM(BOTH '%' FROM ?)) ")
		} else if strings.HasSuffix(val, "%") {
			params = append(params, val)
			whereSQL0.WriteString(g.Quote(col) + " LIKE CONCAT(TRIM(BOTH '%' FROM ?), '%') ")
		} else {
			params = append(params, val)
			whereSQL0.WriteString(g.Quote(col) + " = ? ")
		}
	}
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = g.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	_ = executeSQL

	return
}

// SELECT * FROM `@@table`
// {{if len(cols) == len(vals)}}
// {{where}}
//
//	  {{for i, col := range cols}}
//	    {{for j, val := range vals}}
//	      {{if i == j}}
//	        {{if val != ""}}
//	          {{if strings.HasPrefix(val, "%") && strings.HasSuffix(val, "%")}}
//	            @@col LIKE CONCAT('%', TRIM(BOTH '%' FROM @val), '%') AND
//	          {{else if strings.HasPrefix(val, "%")}}
//	            @@col LIKE CONCAT('%', TRIM(BOTH '%' FROM @val)) AND
//	          {{else if strings.HasSuffix(val, "%")}}
//	            @@col LIKE CONCAT(TRIM(BOTH '%' FROM @val), '%') AND
//	          {{else}}
//	            @@col = @val AND
//	          {{end}}
//	        {{end}}
//	      {{end}}
//	    {{end}}
//	  {{end}}
//	{{end}}
//
// {{end}}
func (g gameDo) FindByCols(cols []string, vals []string) (result []model.Game) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM `game` ")
	if len(cols) == len(vals) {
		var whereSQL0 strings.Builder
		for i, col := range cols {
			for j, val := range vals {
				if i == j {
					if val != "" {
						if strings.HasPrefix(val, "%") && strings.HasSuffix(val, "%") {
							params = append(params, val)
							whereSQL0.WriteString(g.Quote(col) + " LIKE CONCAT('%', TRIM(BOTH '%' FROM ?), '%') AND ")
						} else if strings.HasPrefix(val, "%") {
							params = append(params, val)
							whereSQL0.WriteString(g.Quote(col) + " LIKE CONCAT('%', TRIM(BOTH '%' FROM ?)) AND ")
						} else if strings.HasSuffix(val, "%") {
							params = append(params, val)
							whereSQL0.WriteString(g.Quote(col) + " LIKE CONCAT(TRIM(BOTH '%' FROM ?), '%') AND ")
						} else {
							params = append(params, val)
							whereSQL0.WriteString(g.Quote(col) + " = ? AND ")
						}
					}
				}
			}
		}
		helper.JoinWhereBuilder(&generateSQL, whereSQL0)
	}

	var executeSQL *gorm.DB
	executeSQL = g.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	_ = executeSQL

	return
}

func (g gameDo) Debug() *gameDo {
	return g.withDO(g.DO.Debug())
}

func (g gameDo) WithContext(ctx context.Context) *gameDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gameDo) ReadDB() *gameDo {
	return g.Clauses(dbresolver.Read)
}

func (g gameDo) WriteDB() *gameDo {
	return g.Clauses(dbresolver.Write)
}

func (g gameDo) Session(config *gorm.Session) *gameDo {
	return g.withDO(g.DO.Session(config))
}

func (g gameDo) Clauses(conds ...clause.Expression) *gameDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gameDo) Returning(value interface{}, columns ...string) *gameDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g gameDo) Not(conds ...gen.Condition) *gameDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gameDo) Or(conds ...gen.Condition) *gameDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gameDo) Select(conds ...field.Expr) *gameDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gameDo) Where(conds ...gen.Condition) *gameDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gameDo) Order(conds ...field.Expr) *gameDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gameDo) Distinct(cols ...field.Expr) *gameDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gameDo) Omit(cols ...field.Expr) *gameDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gameDo) Join(table schema.Tabler, on ...field.Expr) *gameDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gameDo) LeftJoin(table schema.Tabler, on ...field.Expr) *gameDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gameDo) RightJoin(table schema.Tabler, on ...field.Expr) *gameDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gameDo) Group(cols ...field.Expr) *gameDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gameDo) Having(conds ...gen.Condition) *gameDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gameDo) Limit(limit int) *gameDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gameDo) Offset(offset int) *gameDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gameDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *gameDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gameDo) Unscoped() *gameDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gameDo) Create(values ...*model.Game) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gameDo) CreateInBatches(values []*model.Game, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gameDo) Save(values ...*model.Game) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gameDo) First() (*model.Game, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Game), nil
	}
}

func (g gameDo) Take() (*model.Game, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Game), nil
	}
}

func (g gameDo) Last() (*model.Game, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Game), nil
	}
}

func (g gameDo) Find() ([]*model.Game, error) {
	result, err := g.DO.Find()
	return result.([]*model.Game), err
}

func (g gameDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Game, err error) {
	buf := make([]*model.Game, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gameDo) FindInBatches(result *[]*model.Game, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gameDo) Attrs(attrs ...field.AssignExpr) *gameDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gameDo) Assign(attrs ...field.AssignExpr) *gameDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gameDo) Joins(fields ...field.RelationField) *gameDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g gameDo) Preload(fields ...field.RelationField) *gameDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g gameDo) FirstOrInit() (*model.Game, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Game), nil
	}
}

func (g gameDo) FirstOrCreate() (*model.Game, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Game), nil
	}
}

func (g gameDo) FindByPage(offset int, limit int) (result []*model.Game, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g gameDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g gameDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g gameDo) Delete(models ...*model.Game) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *gameDo) withDO(do gen.Dao) *gameDo {
	g.DO = *do.(*gen.DO)
	return g
}
