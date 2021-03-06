package search

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	m := map[Term]int{}
	mp := map[*Term]int{}
	t1 := Term{"a", "a"}
	t2 := Term{"a", "a"}
	m[t1] = 1
	m[t2] = 2
	if len(m) != 1 {
		t.Error("not equal")
	}
	mp[&t1] = 1
	mp[&t2] = 2
	if len(mp) != 2 {
		t.Error("not equal")
	}
}

func TestSearch(t *testing.T) {
	idName := "id"
	termsName := "terms"
	doc1 := &Document{
		[]Field{
			&IntField{BaseField{true, idName}, 1}, &StrSliceField{BaseField{true, termsName}, []string{"中国", "北京","江苏","南京"}},
		},
	}
	doc2 := &Document{
		[]Field{
			&IntField{BaseField{true, idName}, 2}, &StrSliceField{BaseField{true, termsName}, []string{"台湾", "北京", "上海"}},
		},
	}
	doc3 := &Document{
		[]Field{
			&IntField{BaseField{true, idName}, 3}, &StrSliceField{BaseField{true, termsName}, []string{"南京","湖北", "武汉", "上海"}},
		},
	}
	searcher := NewSearcher()
	searcher.Add(doc1)
	searcher.Add(doc2)
	searcher.Add(doc3)
	q1 := &TermQuery{&Term{termsName, "中国"}}
	q2 := &TermQuery{&Term{termsName, "北京"}}
	q21 := &TermQuery{&Term{termsName, "上海"}}
	q3 := &BooleanQuery{q1, q21, SHOULD}
	q4 := &BooleanQuery{q1, q2, MUST}
	fmt.Println("search:中国")
	printDocs(searcher.Find(q1))
	fmt.Println("search:北京")
	printDocs(searcher.Find(q2))
	fmt.Println("search:中国 || 上海")
	printDocs(searcher.Find(q3))
	fmt.Println("search:中国 && 北京")
	printDocs(searcher.Find(q4))
}

func printDocs(docs []*Document) {
	for _, d := range docs {
		printDoc(d)
	}
}

func printDoc(doc *Document) {
	for _, f := range doc.Fields {
		fmt.Printf("%s:%s ", f.GetName(), f.GetValue())
	}
	fmt.Println()
}
