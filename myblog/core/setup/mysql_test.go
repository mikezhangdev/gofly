package setup

import (
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

func TestNewMysqlConn(t *testing.T) {
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

	d := map[string]interface{}{
		"title":       "标题",
		"price":       11,
		"create_time": time.Now().Unix(),
		"modify_time": time.Now().Unix(),
	}
	err = con.Table("demo").Create(d).Error
	//var ret interface{}
	//err = con.Table("demo").Find(&ret).Error
	//if err != nil {
	//	t.Fatalf("find err %v", err)
	//}
}
