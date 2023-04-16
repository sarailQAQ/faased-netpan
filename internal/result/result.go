package result

import "cloud-disk/internal/types"

// OK 成功返回结果
func OK(data ...interface{}) types.Result {
	resp := types.Result{
		Code:    200,
		Message: "操作成功",
	}
	if len(data) > 0 {
		for _, datum := range data {
			if interfaceAssert(datum) == "string" {
				if datum.(string) != "" {
					resp.Message = datum.(string)
				}
			} else if interfaceAssert(datum) == "object" {
				resp.Data = datum
			}
		}
	}
	return resp
}

// ERROR 失败返回结果
func ERROR(msg string) types.Result {
	return types.Result{
		Code:    500,
		Message: msg,
	}
}

func interfaceAssert(unknow interface{}) (retType string) {
	switch unknow.(type) {
	case string:
		return "string"
	case int:
		return "int"
	default:
		return "object"
	}
}
