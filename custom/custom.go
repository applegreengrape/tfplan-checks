package custom

import (
	"fmt"
	"reflect"
)


func S3_Tag_Check_01(typ []interface{}, value []interface{}) (string, string){
	var s string
	var r string
	var keys []string
	// mandatory checks for tags
	var tags = []string{"key_1", "key_2", "key_3"}
	for i:= range typ {
		if typ[i] == "aws_s3_bucket"{
			for k, _ :=range value[i].(map[string]interface{})["tags"].(map[string]interface{}){
				keys = append(keys, string(k))
			}
		}
	}
	
	if reflect.DeepEqual(tags, keys){
		s = "okay"
		r = fmt.Sprintf(`
	✅[Custom Checks][S3_Tag_Check_01] All Mandatory Tags found
		- Requrired Tags: %s
		- Current Tags: %s
		  `,tags, keys)
		return s,r
	}else{
		s = "nope"
		r = fmt.Sprintf(`
	❌[Custom Checks][S3_Tag_Check_01] Mandatory Tags Missing
		- Requrired Tags: %s
		- Current Tags: %s
		  `,tags, keys)
		return s,r
	}

}

func S3_Tag_Check_02(typ []interface{}, value []interface{}) (string, string){
	var s string
	var r string
	var keys_2 []string
	// mandatory checks for tags
	var tags_2 = []string{"key_1", "key_2", "key_3", "key_4"}
	for i:= range typ {
		if typ[i] == "aws_s3_bucket"{
			for k, _ :=range value[i].(map[string]interface{})["tags"].(map[string]interface{}){
				keys_2 = append(keys_2, string(k))
			}
		}
	}
	
	if reflect.DeepEqual(tags_2, keys_2){
		s = "okay"
		r = fmt.Sprintf(`
	✅[Custom Checks][S3_Tag_Check_02] All Mandatory Tags found
		- Requrired Tags: %s
		- Current Tags: %s
		  `,tags_2, keys_2)
		return s,r
	}else{
		s = "nope"
		r = fmt.Sprintf(`
	❌[Custom Checks][S3_Tag_Check_02] Mandatory Tags Missing
		- Requrired Tags: %s
		- Current Tags: %s
		  `,tags_2, keys_2)
		return s,r
	}

}