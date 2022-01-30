package gormmysql

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
	"time"
)

//CREATE TABLE `demo` (
//`id` bigint NOT NULL AUTO_INCREMENT,
//`title` varchar(300) NOT NULL DEFAULT '' COMMENT '标题',
//`price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '价格',
//`create_time` int NOT NULL DEFAULT '0' COMMENT '创建时间',
//`modify_time` int NOT NULL DEFAULT '0' COMMENT '修改时间',
//`delete_time` int NOT NULL DEFAULT '0' COMMENT '删除时间',
//`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除@：0否@：1是',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='demo测试表';

//func TestNewMysqlConn(t *testing.T) {
//	cfg := &MysqlCfg{
//		User:      "root",
//		Pass:      "123456",
//		Host:      "192.168.56.104",
//		Port:      3306,
//		DbName:    "myblog",
//		Charset:   "utf8mb4",
//		ParseTime: "true",
//		Loc:       "Local",
//	}
//	con, err := NewMysqlConn(cfg)
//	if err != nil {
//		t.Fatalf("err %v", err)
//	}
//
//	//d := map[string]interface{}{
//	//	"title":       "标题",
//	//	"price":       11,
//	//	"create_time": time.Now().Unix(),
//	//	"modify_time": time.Now().Unix(),
//	//}
//	//err = con.Table("demo").Create(d).Error
//	var ret interface{}
//	err = con.Table("demo").Find(&ret).Error
//	if err != nil {
//		t.Fatalf("find err %v", err)
//	}
//}


//CREATE TABLE `demo` (
//`id` bigint NOT NULL AUTO_INCREMENT,
//`title` varchar(300) NOT NULL DEFAULT '' COMMENT '标题',
//`price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '价格',
//`create_time` int NOT NULL DEFAULT '0' COMMENT '创建时间',
//`modify_time` int NOT NULL DEFAULT '0' COMMENT '修改时间',
//`delete_time` int NOT NULL DEFAULT '0' COMMENT '删除时间',
//`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除@：0否@：1是',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='demo测试表';


type Demo struct{
	Id int64 `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	CreateTime int64 `json:"create_time"`
	ModifyTime int64 `json:"modify_time"`
	DeleteTime int64 `json:"delete_time"`
	IsDel int32 `json:"is_del"`
}

//func ReplaceCreateId(scope *gorm.){
//
//}

func TestSetBeforeCreateCallback(t *testing.T){
	//// SetBeforeCreateCallback is set before create callback
	//func (r *mysqlConn) SetBeforeCreateCallback(db *gorm.DB, callback func(scope *gorm.Scope)) {
	//	db.Callback().Create().Replace("gorm:before_create", callback)
	//}
	cfg := &MysqlCfg{
		User:      "root",
		Pass:      "123456",
		Host:      "192.168.56.104",
		Port:      3306,
		DbName:    "myblog",
		Charset:   "utf8mb4",
		ParseTime: "true",
		Loc:       "Local",
	}
	con, err := NewMysqlConn(cfg)
	if err != nil {
		t.Fatalf("err %v", err)
	}
	err = con.Callback().Create().Replace("gorm:before_create", func(db *gorm.DB) {
		//db.Scopes()
		idField := db.Statement.Schema.LookUpField("Id")
		db.Statement.SetColumn("Id",12222)
		v,z :=idField.ValueOf(db.Statement.ReflectValue)
		fmt.Printf("v %v z; %v \n",v,z)
		//fmt.Printf("idField %v,ok %v \n",idField.ValueOf(db.Statement.ReflectValue))
		//if !ok {
		//	return
		//}
		//fmt.Printf("idField:%v",idField)
		//m := db.Migrator().HasColumn()
		//idField.S

	})

	if err != nil{
		t.Fatalf("con.Callback().Create().Replace err %v",err)
	}

	//d := map[string]interface{}{
	//	"id":333,
	//	"title":       "标题11",
	//	"price":       11,
	//	"create_time": time.Now().Unix(),
	//	"modify_time": time.Now().Unix(),
	//}
	d := &Demo{
		Id:         555,
		Title:      "111",
		Price:      12,
		CreateTime: time.Now().Unix(),
		ModifyTime: time.Now().Unix(),
	}
	err = con.Table("demo").Create(d).Error

	if err != nil{
		t.Fatalf("create err %v",err)
	}




}
