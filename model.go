package main

import (
	"encoding/json"
	"fmt"
)

var (
	DomainID = "mynews" // DomainID 는 외부에 domain id 를 참조시키기 위해 사용
)

// Domain 뉴스의 DomainObject interface 구현체
type Domain struct {
	Rule     string   `json:"rule"`
	Keywords []string `json:"keywords"`
}

func (n Domain) Match(source json.RawMessage) (bool, string, error) {
	fmt.Printf("mynews Match called! %s\n", source)
	return true, "", nil
}

func (n Domain) Crawling(q chan interface{}) error {
	fmt.Printf("mynews Crawling called! %s\n", q)
	return nil
}

func MakeDomain(jsonText []byte) interface{} {
	var n Domain
	json.Unmarshal(jsonText, &n)
	return n
}
