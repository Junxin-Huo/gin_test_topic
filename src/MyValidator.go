package src

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func TopicUrl(fl validator.FieldLevel) bool {
	//fmt.Println(fl)
	//fmt.Println(fl.Field().Interface())
	//fmt.Println(fl.Parent().Interface())
	//fmt.Println(fl.Param())
	//fmt.Println(fl.FieldName())
	//fmt.Println(fl.GetTag())
	//fmt.Println(fl.StructFieldName())
	//fmt.Println(fl.Top().Interface())

	_, ok1 := fl.Top().Interface().(*Topic)
	_, ok2 := fl.Top().Interface().(*Topics)
	if !ok1 && !ok2 {
		return false
	}

	if url, ok := fl.Field().Interface().(string); ok {
		//fmt.Println(url)
		if matched, err:= regexp.MatchString(`^\w{4,10}$`, url); matched {
			return true
		} else {
			fmt.Println(err)
		}
	}

	return false
}

func TopicsValidate(fl validator.FieldLevel) bool {
	topics, ok := fl.Top().Interface().(*Topics)
	fmt.Println(topics)
	if !ok {
		return false
	}

	if topics.TopicListSize == len(topics.TopicList) {
		return true
	}
	return false
}
