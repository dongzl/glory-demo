package main

import (
	"log"

	"github.com/glory-go/glory/mysql/extend_type"

	"github.com/glory-go/glory/mysql"
	_ "github.com/glory-go/glory/registry/nacos"
)

type PayLoadSubStruct struct {
	Name string `json:"name"`
}

type PayLoadStruct struct {
	Token     string           `json:"token"`
	TTL       string           `json:"ttl"`
	Age       int64            `json:"age"`
	SubStruct PayLoadSubStruct `json:"substruct"`
}

// 自定义mysqlModel
type myMysqlModel struct {
	ID       int64
	UserName string
	Payload  *extend_type.JSONTagedGORMType
}

func newMyMysqlModel() *myMysqlModel {
	return &myMysqlModel{
		ID:       126,
		UserName: "laurence",
		Payload: extend_type.NewJSONTagedGORMType(PayLoadStruct{
			Token: "xxx",
			TTL:   "ttl",
			Age:   13,
			SubStruct: PayLoadSubStruct{
				Name: "lizhixin",
			},
		}),
	}
}

func (m *myMysqlModel) TableName() string {
	return "glory_test"
}

var tablePtr *mysql.MysqlTable

func main() {
	table, err := mysql.RegisterModel("dev-mysql", &myMysqlModel{})
	if err != nil {
		panic(err)
	}

	tablePtr = table
	model := newMyMysqlModel()
	if err := tablePtr.Insert(model); err != nil {
		tablePtr.Delete(model)
	}
	newModel := &myMysqlModel{}
	if err := tablePtr.First("id = ?", newModel, 126); err != nil {
		panic(err)
	}
	log.Printf("get from db model = %+v\n", newModel.Payload)
	tablePtr.Delete(newModel)

}
