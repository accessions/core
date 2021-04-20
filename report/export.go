package report

import (
	"fmt"
	"github.com/accessions/core/ali"
	"github.com/tealeg/xlsx"
	"os"
	"reflect"

	"time"
)

var ReportService *Service

func init() {
	ReportService = &Service{}
}

type Options struct {
	Name   string
	Title  []interface{}
	Data   [][]interface{}
	Suffix string
}
type Option func(options *Options)

type Service struct {
	option   Options
	fileName string
	oss      bool
}

/**
lib.ReportService.New(func(options *lib.Options) {
	options.Name = fmt.Sprintf("广告%s信息", cpx)
	options.Title = []interface{}{"A","B"}
	options.Data = [][]interface{}{{"1","2"},{"A","B"}}
	options.Suffix = string(cpx)
}).
Report().
Oss()

*/
func (repo *Service) New(opts ...Option) *Service {
	opt := Options{
		Name:   "",
		Title:  nil,
		Data:   nil,
		Suffix: "",
	}
	for _, o := range opts {
		o(&opt)
	}
	repo.option = opt
	return repo
}

func (repo *Service) Buffer() (string, error) {
	// return repo.Report().fileName, nil
	return repo.fileName, nil
}

func (repo *Service) Oss() (string, error) {
	// oss, err := os.Open(repo.Report().fileName)
	oss, err := os.Open(repo.fileName)
	if err != nil {
		return "", err
	}
	defer oss.Close()
	return ali.AliOss(repo.fileName, oss), nil

}

// Report
func (repo *Service) Report() *Service {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error
	file = xlsx.NewFile()
	sheet, _ = file.AddSheet(repo.option.Name)
	row = sheet.Row(0)
	row.WriteSlice(&repo.option.Title, -1)
	for i, data := range repo.option.Data {
		row = sheet.Row(i + 1)
		_ = row.WriteSlice(&data, -1)
	}
	fileName := "path" + time.Now().Format("20060102030405") + "_" + repo.option.Suffix + ".xlsx"
	if err = file.Save(fileName); err != nil {
		return repo
	}
	repo.fileName = fileName
	return repo
}

// StructToSlice
func StructToSlice(structure interface{}) []interface{} {
	var v reflect.Value
	v = reflect.ValueOf(structure)
	effect := make([]interface{}, v.NumField())
	for i := range effect {
		effect[i] = fmt.Sprintf("%v", v.Field(i))
	}
	return effect
}
