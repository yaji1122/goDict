package tools

import (
	"github.com/longbridgeapp/opencc"
	"log"
)

func ToZhTW(s string) string{
	s2t, err := opencc.New("s2t")
	if err != nil {
		log.Fatal(err)
	}
	out, err := s2t.Convert(s)
	if err != nil {
		log.Fatal(err)
	}
	return out
}
