package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yaji1122/goDict/internal/tools"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ICIBA_API_URL = "https://dict-mobile.iciba.com/interface/index.php?c=word&m=getsuggest&nums=10&is_need_mean=1&word=%s"
)

func Query(c *gin.Context) {
	v := c.Param("word")
	s := getWord(v)
	c.String(http.StatusOK, s)
}

func getWord(word string) string {
	var d string
	queryUrl := fmt.Sprintf(ICIBA_API_URL, word)
	fmt.Println("Query Url : " + queryUrl)
	resp, err := http.Get(queryUrl)
	if err != nil {
		log.Println(err)
		return d
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return d
	}

	return tools.ToZhTW(string(body))
}
