package main

import (
    "reflect"
    "fmt"
)

type ControllerInterface interface {
    Init(action string, method string)
}

type Controller struct {
    Action string
    Method string
    Tag string `json:"tag"`
}

func (c *Controller) Init(action string, method string){
    c.Action = action
    c.Method = method
}

func main(){
    //初始化
    runController := &Controller{
        Action:"Run1",
        Method:"GET",
    }

    //Controller实现了ControllerInterface方法,因此它就实现了ControllerInterface接口
    var i ControllerInterface
    i = runController

    // 得到实际的值,通过v我们获取存储在里面的值,还可以去改变值
    v := reflect.ValueOf(i)
    fmt.Println("value:",v)

    // 得到类型的元数据,通过t我们能获取类型定义里面的所有元素
    t := reflect.TypeOf(i)
    fmt.Println("type:",t)

    // 转化为reflect对象之后我们就可以进行一些操作了,也就是将reflect对象转化成相应的值
	// Elem返回值v包含的接口
	controllerValue := v.Elem() 
	controllerType := t.Elem()  
	fmt.Println("controllerType(reflect.Value):",controllerType)
	//获取存储在第一个字段里面的值
	fmt.Println("Action:", controllerValue.Field(0).String())
}