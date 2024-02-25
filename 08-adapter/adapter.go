package adapter

import "fmt"

// 由于第三方支付的SDK方法都不一样，我们可以将它们抽象为 IPay 接口来适配我们的系统（原理与代理模式类似）
type IPay interface {
	Pay()
}

// 模拟第三方 SDK
type TencenSDK struct {
}

func (t *TencenSDK) GetToken() {
	fmt.Println("TencenSDK.GetToken")
}

func (t *TencenSDK) PayOrder() {
	fmt.Println("TencenSDK.PayOrder")
}

type AliSDK struct {
}

func (a *AliSDK) Check() {
	fmt.Println("AliSDK.Check")
}

func (a *AliSDK) PayGoods() {
	fmt.Println("AliSDK.PayGoods")
}

// 适配器
type TencenAdapter struct {
	TSDK *TencenSDK
}

func (t *TencenAdapter) Pay() {
	t.TSDK.GetToken()
	t.TSDK.PayOrder()
}

type AliAdapter struct {
	ASDK *AliSDK
}

func (t *AliAdapter) Pay() {
	t.ASDK.Check()
	t.ASDK.PayGoods()
}
