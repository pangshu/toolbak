package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestStripTags String/Global.go String/StripTags.go String/StripTags_test.go
func TestStripTags(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.StripTags("<ul class=\"post_side_mod_list\">\n<li class=\"post_side_mod_item\">\n<h3><a href=\"https://www.163.com/dy/article/FVOTPDCC0514R9OJ.html\">石家庄全市人员不得出市！检测241万人 新增11例阳性</a></h3>\n            <!-- 环球网资讯 -->\n        </li>\n                            <li class=\"post_side_mod_item\">\n            <h3><a href=\"https://www.163.com/dy/article/FVOORDB20519D4UH.html\">美国会确认拜登胜选！4条人命为代价 特朗普彻底失败</a></h3>\n<!-- 和讯网 -->\n</li>\n<li class=\"post_side_mod_item\">\n<h3><a href=\"https://www.163.com/dy/article/FVOP7NCB05504DOQ.html\">美国会遭冲击 伊万卡发推称抗议者是爱国者 随后秒删</a></h3>\n<!-- 环球时报国际 -->\n</li>\n<li class=\"post_side_mod_item\">\n<h3><a href=\"https://www.163.com/dy/article/FVOPRAFI0530PDK6.html\">河北两地病例在一处轨迹重叠 中疾控研究有重大发现</a></h3>\n<!-- 新京报北京知道 -->\n</li>\n</ul>")
	fmt.Println(res1)
}

// go test -v -run TestStripTags -bench=BenchmarkStripTags -count=5 String/Global.go String/StripTags.go String/StripTags_test.go
func BenchmarkStripTags(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.StripTags("<ul class=\"post_side_mod_list\">\n<li class=\"post_side_mod_item\">\n<h3><a href=\"https://www.163.com/dy/article/FVOTPDCC0514R9OJ.html\">石家庄全市人员不得出市！检测241万人 新增11例阳性</a></h3>\n            <!-- 环球网资讯 -->\n        </li>\n                            <li class=\"post_side_mod_item\">\n            <h3><a href=\"https://www.163.com/dy/article/FVOORDB20519D4UH.html\">美国会确认拜登胜选！4条人命为代价 特朗普彻底失败</a></h3>\n<!-- 和讯网 -->\n</li>\n<li class=\"post_side_mod_item\">\n<h3><a href=\"https://www.163.com/dy/article/FVOP7NCB05504DOQ.html\">美国会遭冲击 伊万卡发推称抗议者是爱国者 随后秒删</a></h3>\n<!-- 环球时报国际 -->\n</li>\n<li class=\"post_side_mod_item\">\n<h3><a href=\"https://www.163.com/dy/article/FVOPRAFI0530PDK6.html\">河北两地病例在一处轨迹重叠 中疾控研究有重大发现</a></h3>\n<!-- 新京报北京知道 -->\n</li>\n</ul>")
	}
}