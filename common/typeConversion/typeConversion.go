/**
 * @Description: 类型转换
 * @Author Lee
 * @Date 2023/12/8 13:33
 **/

package tc

import "encoding/json"

// JsonToMap /*
func JsonToMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}
