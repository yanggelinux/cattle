package errors

import "fmt"

func DoApi() {

}

func DoService() error {
	err := DoModel()
	fmt.Println("print", err)
	if err != nil {

		return err
	}
	return nil
}

func DoModel() error {
	err := fmt.Errorf("数据不存在")
	err = Wrap(err, "jjjjjjj错误")
	err = Wrap(err, "jbbbbbj错误")
	err = WithCode(2004, "操作数据库失败:%w", err)
	return err
}
