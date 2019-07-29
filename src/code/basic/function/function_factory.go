package function

import "strings"

//MakeAddSuffix 添加后缀
func MakeAddSuffix(suffix string) func(string) string{
	return func(name string)string{
		if strings.HasSuffix(name,suffix){
			return name + suffix;
		}
		return name;
	}
}