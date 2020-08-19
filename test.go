package main

import (
	"fmt"
	"strings"
)

func main() {
	//urlSlug := "Hôm nay trời đẹp quá í ì á à ạ ă â ế"
	urlSlug := "Go là một ngôn ngữ lập trình được thiết kế dựa trên tư duy lập trình hệ thống"
	fmt.Print(urlgen(urlSlug))
}

func urlgen(slug string) string {
	slug = strings.ToLower(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "đ", "d")

	var equalA = []string{"á", "à", "ạ", "ả", "ã", "â", "ấ", "ầ", "ậ", "ẩ", "ẫ", "ă", "ắ", "ằ", "ặ", "ẳ", "ẵ"}
	a := "a"
	for i := 0; i < len(equalA); i++ {
		if a != equalA[i] {
			slug = strings.ReplaceAll(slug, equalA[i], a)
		}
	}

	var equalE = []string{"é", "è", "ẹ", "ẻ", "ẽ", "ê", "ế", "ề", "ệ", "ể", "ễ"}
	e := "e"
	for i := 0; i < len(equalE); i++ {
		if e != equalE[i] {
			slug = strings.ReplaceAll(slug, equalE[i], e)
		}
	}

	var equalI = []string{"í", "ì", "ị", "ỉ", "ĩ"}
	I := "i"
	for i := 0; i < len(equalI); i++ {
		if I != equalI[i] {
			slug = strings.ReplaceAll(slug, equalI[i], I)
		}
	}

	var equalO = []string{"ó", "ò", "ọ", "ỏ", "õ", "ô", "ơ", "ố", "ồ", "ộ", "ổ", "ỗ", "ớ", "ờ", "ở", "ợ", "ỡ"}
	o := "o"
	for i := 0; i < len(equalO); i++ {
		if o != equalO[i] {
			slug = strings.ReplaceAll(slug, equalO[i], o)
		}
	}

	var equalU = []string{"ú", "ù", "ụ", "ủ", "ũ", "ư", "ứ", "ừ", "ự", "ử", "ữ"}
	u := "u"
	for i := 0; i < len(equalU); i++ {
		if u != equalU[i] {
			slug = strings.ReplaceAll(slug, equalU[i], u)
		}
	}

	return slug
}
