package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"myblog/core/metrics"
	"net/http"
)

func (r *Runner) ListenGinServer() error {
	var err error
	engine := gin.Default()
	// 开发健康检测探针
	metrics.InnerHandler(engine)
	if r.app.GinHandler != nil {
		err = r.app.GinHandler(engine)
		if err != nil {
			return fmt.Errorf("r.app.GinHandler err %v", err)
		}
	}
	s := http.Server{ // server 每个参数值定义和设置优化 TODO
		Addr:    r.app.Addr,
		Handler: engine,
		//TLSConfig:         nil,
		//ReadTimeout:       0,
		//ReadHeaderTimeout: 0,
		//WriteTimeout:      0,
		//IdleTimeout:       0,
		//MaxHeaderBytes:    0,
		//TLSNextProto:      nil,
		//ConnState:         nil,
		//ErrorLog:          nil,
		//BaseContext:       nil,
		//ConnContext:       nil,
	}
	log.Printf(".ListenAndServe %v",r.app.Addr)
	err = s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
