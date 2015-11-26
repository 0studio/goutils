package goutils

import (
	"math/rand"
	"strconv"
	"strings"
)

type Int8List []int8

func (list Int8List) Equal(list2 Int8List) bool {
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

func (list Int8List) Join(separator string) string {
	strLi := make([]string, len(list))
	for index, v := range list {
		strValue := strconv.Itoa(int(v))
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}
func (list Int8List) ToStringList() (strList []string) {
	strList = make([]string, len(list), len(list))
	for idx, _ := range list {
		strList[idx] = Int82Str(list[idx])
	}
	return
}

func (list Int8List) LenInt() int {
	return len(list)
}

func (list Int8List) LenInt8() int8 {
	return int8(len(list))
}
func (list Int8List) Len() int8 {
	return int8(len(list))
}

func (list Int8List) Clone() (newList Int8List) {
	newList = make(Int8List, list.LenInt())
	copy(newList, list)
	return
}
func (list Int8List) HasDup() bool {
	// return 是否有重复id
	return len(list) != len(list.RemoveDup())
}
func (list Int8List) IsInList(ele int8) bool {
	return IsInSlice8(ele, list)
}
func (list Int8List) RemoveZero() Int8List {
	return list.Remove(0)
}
func (list Int8List) Remove(ele int8) (ret Int8List) {
	ret = make(Int8List, len(list), len(list))
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
func (list *Int8List) RemoveBySetZero(ele int8) (changed bool) {
	// 如果list 中有ele, 则将其所在位置设置成0
	for idx, _ := range *list {
		if (*list)[idx] == ele {
			(*list)[idx] = 0
			changed = true
		}
	}
	return
}
func (list *Int8List) RemoveSelf(ele int8) {
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
func (list Int8List) RemoveDup() (newList Int8List) {
	// 去重
	newList = make(Int8List, len(list), len(list))
	found := make(map[int8]bool)
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
func (list Int8List) Swap(fromPos, toPos int8) bool {
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
func (list *Int8List) ShuffleSelf() {
	// 乱序
	for i := range *list {
		j := rand.Intn(i + 1)
		(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
	}
}
func (list Int8List) Random() int8 {
	if len(list) == 0 {
		return 0
	}
	return list[rand.Intn(len(list))]
}
