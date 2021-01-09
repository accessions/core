package test

import (
	"fmt"
	"github.com/accessions/core/sql"
	"testing"
)

type TestReports struct {
	Id         int64     `gorm:"column:id;" json:"id"`
	Uin       int64     `gorm:"column:uin" json:"uin"`
	Amount     float64   `gorm:"column:amount" json:"amount"`
}

func TestHelp_BuildBatchUpdateSqlArray(t *testing.T)  {
	editMod := []*TestReports{
		{1,100, 100.21},
		{2,200, 200.20},
	}
	fmt.Println(sql.BuildBatchUpdateSqlArray("table", editMod))
}
