package goutils

import (
	"math/rand"
)

type IntList []int

func (list IntList) Equal(list2 IntList) bool {
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

func (list IntList) Join(separator string) string {
	return Int64ListToString(list, separator)
}
func (list IntList) ToStringList() (strList []string) {
	strList = make([]string, len(list), len(list))
	for idx, _ := range list {
		strList[idx] = Int2Str(list[idx])
	}
	return
}

func (list IntList) LenInt() int {
	return len(list)
}

func (list IntList) LenInt8() int8 {
	return int8(len(list))
}
func (list IntList) Len() int {
	return int(len(list))
}

func (list IntList) Clone() (newList IntList) {
	newList = make(IntList, list.LenInt())
	copy(newList, list)
	return
}
func (list IntList) HasDup() bool {
	// return 是否有重复id
	return len(list) != len(list.RemoveDup())
}
func (list IntList) IsInList(ele int) bool {
	return IsInSlice(ele, list)
}
func (list IntList) RemoveZero() IntList {
	return list.Remove(0)
}
func (list IntList) Remove(ele int) (ret IntList) {
	ret = make(IntList, len(list), len(list))
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
func (list *IntList) RemoveBySetZero(ele int) (changed bool) {
	// 如果list 中有ele, 则将其所在位置设置成0
	for idx, _ := range *list {
		if (*list)[idx] == ele {
			(*list)[idx] = 0
			changed = true
		}
	}
	return
}
func (list *IntList) RemoveSelf(ele int) {
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
func (list IntList) RemoveDup() (newList IntList) {
	// 去重
	newList = make(IntList, len(list), len(list))
	found := make(map[int]bool)
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
func (list IntList) Swap(fromPos, toPos int) bool {
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
func (list *IntList) ShuffleSelf() {
	// 乱序
	for i := range *list {
		j := rand.Intn(i + 1)
		(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
	}
}
