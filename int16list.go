package goutils

import (
	"math/rand"
	"strconv"
	"strings"
)

type Int16List []int16

func (list Int16List) Equal(list2 Int16List) bool {
	if list == nil && list2 == nil {
		return true
	}
	if len(list) != len(list2) {
		return false
	}
	for idx, _ := range list {
		if list[idx] != list2[idx] {
			return false
		}
	}
	return true
}

func (list Int16List) Join(separator string) string {
	strLi := make([]string, len(list))
	for index, v := range list {
		strValue := strconv.Itoa(int(v))
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}
func (list Int16List) ToStringList() (strList []string) {
	strList = make([]string, len(list), len(list))
	for idx, _ := range list {
		strList[idx] = strconv.Itoa(int(list[idx]))
	}
	return
}

func (list Int16List) LenInt() int {
	return len(list)
}

func (list Int16List) LenInt16() int16 {
	return int16(len(list))
}
func (list Int16List) Len() int16 {
	return int16(len(list))
}

func (list Int16List) Clone() (newList Int16List) {
	newList = make(Int16List, list.LenInt())
	copy(newList, list)
	return
}
func (list Int16List) HasDup() bool {
	// return 是否有重复id
	return len(list) != len(list.RemoveDup())
}
func (list Int16List) IsInList(ele int16) bool {
	for _, e := range list {
		if e == ele {
			return true
		}
	}
	return false
}
func (list Int16List) RemoveZero() Int16List {
	return list.Remove(0)
}
func (list Int16List) Remove(ele int16) (ret Int16List) {
	ret = make(Int16List, len(list), len(list))
	var index int = 0
	for _, value := range list {
		if value != ele {
			ret[index] = value
			index++
		}
	}
	ret = ret[:index]
	return
}
func (list *Int16List) RemoveBySetZero(ele int16) (changed bool) {
	// 如果list 中有ele, 则将其所在位置设置成0
	for idx, _ := range *list {
		if (*list)[idx] == ele {
			(*list)[idx] = 0
			changed = true
		}
	}
	return
}
func (list *Int16List) RemoveSelf(ele int16) {
	// 从list 中移除ele, 直接在list 本身上操作， 不再返回新的list
	var index int = 0
	for _, value := range *list {
		if value != ele {
			(*list)[index] = value
			index++
		}
	}
	(*list) = (*list)[:index]
	return
}
func (list Int16List) RemoveDup() (newList Int16List) {
	// 去重
	newList = make(Int16List, len(list), len(list))
	found := make(map[int16]bool)
	j := 0
	for i, val := range list {
		if _, ok := found[val]; !ok {
			found[val] = true
			newList[j] = list[i]
			j++
		}
	}
	newList = newList[:j]
	return
}
func (list Int16List) Swap(fromPos, toPos int16) bool {
	// 直接在list 上进行交换
	Len := list.Len()
	if fromPos >= Len || toPos >= Len {
		return false
	}
	tmp := list[fromPos]
	list[fromPos] = list[toPos]
	list[toPos] = tmp
	return true
}
func (list *Int16List) ShuffleSelf() {
	// 乱序
	for i := range *list {
		j := rand.Intn(i + 1)
		(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
	}
}
func (list Int16List) Random() int16 {
	if len(list) == 0 {
		return 0
	}
	return list[rand.Intn(len(list))]
}
