package test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T)  {
	matched, _ := regexp.MatchString(`https://myxy99.cn/posts/[^\s]`,"https://myxy99.cn/posts/a/a")
	fmt.Println(matched)
}