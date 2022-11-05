package xx

import "reflect"

func GetFieldDiffOne(a, o interface{}) (keyList []string, v1List []string, v2List []string) {
	if o == nil {
		return
	} else if Str(a) == Str(o) || Str(o) == "" {
		return
	}

	oElem := reflect.ValueOf(o).Elem()
	aElem := reflect.ValueOf(a).Elem()
	typeOfT := aElem.Type()
	for i := 0; i < aElem.NumField(); i++ {
		name := typeOfT.Field(i).Name
		switch name {
		case "TimeMs": // exclude
			continue
		}
		v1, v2 := Str(oElem.Field(i).Interface()), Str(aElem.Field(i).Interface())
		if v1 != v2 {
			keyList = append(keyList, name)
			v1List = append(v1List, SubStr(v1, 0, 4096))
			v2List = append(v2List, SubStr(v2, 0, 4096))
		}
	}
	return
}
