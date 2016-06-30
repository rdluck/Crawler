package analysis

import (
	"github.com/opesun/goquery"
	"fmt"
	"net/http"
)

func Analysis(w http.ResponseWriter,r *http.Request)  {
	var url="http://ygyobs.baijia.baidu.com/article/519535"
	p,err:=goquery.ParseUrl(url)
	if err !=nil{
		panic(err)
	}else{
		pTitle:=p.Find("title").Text()
		fmt.Println(pTitle)
		fmt.Println(p.Html())
		fmt.Println(p.Find("h2").Text())

		t:=p.Find(".entry-title a")
		for i:=0;i<t.Length();i++{
			d:=t.Eq(i).Attr("href")
			fmt.Println(d)
		}
		fmt.Println("######################################")
		//获取所有href
		fmt.Println(p.Find("").Attr("href"))
		//判断元素释放存在
		fmt.Println(p.Find("div").HasClass("entry-content"))
	}
}
