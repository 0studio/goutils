package goutils

import (
	"math/rand"
)

type Int32List []int32

func (list Int32List) Equal(list2 Int32List) bool {
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

func (list Int32List) Join(separator string) string {
	return IntListToString(list, separator)
}
func (list Int32List) ToStringList() (strList []string) {
	strList = make([]string, len(list), len(list))
	for idx, _ := range list {
		strList[idx] = Int322Str(list[idx])
	}
	return
}

func (list Int32List) LenInt() int {
	return len(list)
}

func (list Int32List) LenInt8() int8 {
	return int8(len(list))
}
func (list Int32List) Len() int32 {
	return int32(len(list))
}

func (list Int32List) Clone() (newList Int32List) {
	newList = make(Int32List, list.LenInt())
	copy(newList, list)
	return
}
func (list Int32List) HasDup() bool {
	// return 是否有重复id
	return len(list) != len(list.RemoveDup())
}
func (list Int32List) IsInList(ele int32) bool {
	return IsInSlice32(ele, list)
}
func (list Int32List) RemoveZero() Int32List {
	return list.Remove(0)
}
func (list Int32List) Random() int32 {
	if len(list) == 0 {
		return 0
	}
	return list[rand.Intn(len(list))]
}
func (list Int32List) Remove(ele int32) (ret Int32List) {
	ret = make(Int32List, len(list), len(list))
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
func (list *Int32List) RemoveBySetZero(ele int32) (changed bool) {
	// 如果list 中有ele, 则将其所在位置设置成0
	for idx, _ := range *list {
		if (*list)[idx] == ele {
			(*list)[idx] = 0
			changed = true
		}
	}
	return
}
func (list *Int32List) RemoveSelf(ele int32) {
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
func (list Int32List) RemoveDup() (newList Int32List) {
	// 去重
	newList = make(Int32List, len(list), len(list))
	found := make(map[int32]bool)
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
func (list Int32List) Swap(fromPos, toPos int32) bool {
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
func (list *Int32List) Shuffle() {
	// 乱序
	for i := range *list {
		j := rand.Intn(i + 1)
		(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
	}
}
