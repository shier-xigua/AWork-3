package main

import (
	"AWork-3/zfunc"
	"AWork-3/zhttp"
	"AWork-3/zvar"
	"log"
	"time"
)

func main() {
	//日志
	logfile := zfunc.LogFunc()
	defer logfile.Close()
	//打印标识
	zfunc.PrintName()

	var voice int = zfunc.Voice(zvar.Lineto)
	for {
		//==============================第一阶段： 目的：获取Form的map==============================
		var body, bodyErr = zhttp.Zhttp1("POST", zvar.URL, zvar.Payload)
		//println(body)

		//将工单输出 //Form工单摘要表
		var Form = zhttp.InfoMap(body)
		if bodyErr == nil {
			log.Printf(" 工单系统共%d张工单\n", len(Form))
			for i, oder := range Form {
				log.Printf("%v 工单号：%v 工单标题：%v", i+1, oder["processKey"], oder["processTitle"])
			}
		}
		//============================= 第一阶段： 目的：获取Form的map =============================

		//

		//如果Form为0不会执行循环 跳过
		if len(Form) != 0 {
			//***************************** 第二阶段： 目的：将需要接单的内容返回 *********************
			orderSummary := zhttp.Zhttp2(Form, zvar.Payload)
			log.Println("测试[][]interface{}返回值： ", orderSummary)

			//***************************** 第二阶段： 目的：将需要接单的内容返回 **********************

			//

			//############################# 第三阶段: 目的： 处理接单 ###############################
			zhttp.Zhttp3(voice, orderSummary)
			//############################# 第三阶段: 目的： 处理接单 ###############################
		}

		log.Println("=========== = sleep = ===========")
		time.Sleep(12 * time.Second)
	}

}
