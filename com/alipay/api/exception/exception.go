package exception

import "fmt"

type AlipaySDKError struct {
	Message string
}

func (p *AlipaySDKError) Error() string {
	return fmt.Sprintf("AlipaySDKErrore=%s", p.Message)
}
