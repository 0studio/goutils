package goutils

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func Str2uint64(str string, defaultvalue uint64) (value uint64) {
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		value = defaultvalue
	}
	return
}

func Str2int64(str string, defaultvalue int64) (value int64) {
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		value = defaultvalue
	}
	return
}

func Int2Str(v int) string {
	return strconv.Itoa(v)
}

func Int82Str(v int8) string {
	return fmt.Sprintf("%d", v)
}

func Int322Str(v int32) string {
	return fmt.Sprintf("%d", v)
}

func Int642Str(v int64) string {
	return fmt.Sprintf("%d", v)
}

func Uint642Str(v uint64) string {
	return strconv.FormatUint(v, 10)
}

func Float642Str(v float64) string {
	return strconv.FormatFloat(v, 'E', -1, 64)
}
func Str2UInt32(str string, defaultvalue uint32) (value uint32) {
	return uint32(Str2Int(str, int(defaultvalue)))
}

func Str2Int32(str string, defaultvalue int32) (value int32) {
	return int32(Str2Int(str, int(defaultvalue)))
}
func Uint642StrWithZeroEmpty(v uint64) string {
	if v == 0 {
		return ""
	}

	return fmt.Sprintf("%d", v)
}
func Str2Int(str string, defaultvalue int) (value int) {
	value, err := strconv.Atoi(str)
	if err != nil {
		value = defaultvalue
	}
	return
}

/*
* RAND UTIL [min,max]
 */
func RandInRange(min int, max int) int {
	return rand.Intn((max-min)+1) + min
}
func Rand32InRange(min int32, max int32) int32 {
	return rand.Int31n((max-min)+1) + min
}
func RandFloatInRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func IsFileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)

	// _, err := os.Stat(name)
	// if os.IsNotExist(err) {
	// 	return false, nil
	// }
	// return err != nil, err
}
func IsInSlice64(elem uint64, slic []uint64) (ret bool) {
	for _, value := range slic {
		if elem == value {
			return true
		}
	}
	return false
}
func IsInStringList(elem string, slic []string) (ret bool) {
	for _, value := range slic {
		if elem == value {
			return true
		}
	}
	return false
}

func StringListToIntList64(strLi []string) (intLi []uint64) {
	intLi = make([]uint64, len(strLi))
	for index, v := range strLi {
		intValue, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			intLi[index] = 0
		}
		intLi[index] = intValue

	}
	return intLi
}

func SplitAsFloat64Arr(str string, sep string) (ret []float64) {
	ret = make([]float64, 0)
	for _, elem := range strings.Split(str, sep) {
		value, err := strconv.ParseFloat(strings.TrimSpace(elem), 64)
		if err == nil {
			ret = append(ret, value)
		}
	}
	return

}
func SplitAsInt32Arr(str string, sep string) (ret []int32) {
	ret = make([]int32, 0)
	for _, elem := range strings.Split(str, sep) {
		value, err := strconv.Atoi(strings.TrimSpace(elem))
		if err == nil {
			ret = append(ret, int32(value))
		}
	}
	return

}

func SplitAsInt8Arr(str string, sep string) (ret []int8) {
	ret = make([]int8, 0)
	for _, elem := range strings.Split(str, sep) {
		value, err := strconv.Atoi(strings.TrimSpace(elem))
		if err == nil {
			ret = append(ret, int8(value))
		}
	}
	return

}
func UInt64ListToStringListWithSep(intLi []uint64, separator string) string {
	strLi := make([]string, len(intLi))
	for index, v := range intLi {
		strValue := strconv.Itoa(int(v))
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}

func IntListToString(intLi []int32, separator string) string {
	strLi := make([]string, len(intLi))
	for index, v := range intLi {
		strValue := strconv.Itoa(int(v))
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}

func Int64ListToString(intLi []int, separator string) string {
	strLi := make([]string, len(intLi))
	for index, v := range intLi {
		strValue := strconv.Itoa(v)
		strLi[index] = strValue
	}
	return strings.Join(strLi, separator)
}

// do some thing at 4:00 am.
// 获取从现在到 某一时间点 需要多少秒
func GetDurToNextDeadlineTime(now time.Time, hour int, min int, seconds int) (dur int) {
	deadlineTime := time.Date(now.Year(), now.Month(), now.Day(), hour, min, seconds, 0, now.Location())
	if deadlineTime.Before(now) {
		tomorrow := now.Add(time.Hour * 24)
		deadlineTime = time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), hour, min, seconds, 0, now.Location())
	}
	dur = int(deadlineTime.Sub(now).Seconds())
	return
}
func GetDurToNextDeadlineTimeDuration(now time.Time, hour int, min int, seconds int) (dur time.Duration) {
	deadlineTime := time.Date(now.Year(), now.Month(), now.Day(), hour, min, seconds, 0, now.Location())
	if deadlineTime.Before(now) {
		tomorrow := now.Add(time.Hour * 24)
		deadlineTime = time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), hour, min, seconds, 0, now.Location())
	}
	dur = deadlineTime.Sub(now)
	return
}
func IsBeforeHMS(now time.Time, hour int, min int, seconds int) bool {
	deadlineTime := time.Date(now.Year(), now.Month(), now.Day(), hour, min, seconds, 0, now.Location())
	return now.Before(deadlineTime)

}
func IsAfterHMS(now time.Time, hour int, min int, seconds int) bool {
	deadlineTime := time.Date(now.Year(), now.Month(), now.Day(), hour, min, seconds, 0, now.Location())
	return now.After(deadlineTime)

}

/* 生成随机战斗卡片 */
func ShuffleArray(slice []int32) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func ShuffleArrayInt(slice []int) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func ShuffleArray64(slice []uint64) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
func GetRandomNFromList64(slice []uint64, n int) []uint64 {
	ShuffleArray64(slice)
	if len(slice) < n {
		return slice
	}
	return slice[0:n]

}

func RemoveDuplicate32(slis *[]int32) {
	found := make(map[int32]bool)
	j := 0
	for i, val := range *slis {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slis)[j] = (*slis)[i]
			j++
		}
	}
	*slis = (*slis)[:j]
}
func RemoveDuplicate(slis *[]int) {
	found := make(map[int]bool)
	j := 0
	for i, val := range *slis {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slis)[j] = (*slis)[i]
			j++
		}
	}
	*slis = (*slis)[:j]
}

func RemoveDuplicate64(slis *[]uint64) {
	found := make(map[uint64]bool)
	j := 0
	for i, val := range *slis {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slis)[j] = (*slis)[i]
			j++
		}
	}
	*slis = (*slis)[:j]
}
func RemoveDuplicateString(slis *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, val := range *slis {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slis)[j] = (*slis)[i]
			j++
		}
	}
	*slis = (*slis)[:j]
}

func RemoveFromSlice64(elem uint64, list *[]uint64) {
	var index int
	for idx, _ := range *list {
		if elem != (*list)[idx] {
			(*list)[index] = (*list)[idx]
			index++
		}
	}
	(*list) = (*list)[:index]
}
func IsInSlice(elem int, slic []int) (ret bool) {
	for _, value := range slic {
		if elem == value {
			return true
		}
	}
	return false
}

func IsInSlice32(elem int32, slic []int32) (ret bool) {
	for _, value := range slic {
		if elem == value {
			return true
		}
	}
	return false
}

func IsInSlice8(elem int8, slic []int8) (ret bool) {
	for _, value := range slic {
		if elem == value {
			return true
		}
	}
	return false
}
func StringListToIntList(strLi []string) (intLi []int32) {
	intLi = make([]int32, len(strLi))
	for index, v := range strLi {
		intValue, _ := strconv.Atoi(v)
		intLi[index] = int32(intValue)
	}
	return intLi
}

func UInt64ListToStringList(intLi []uint64) (strLi []string) {
	strLi = make([]string, len(intLi))
	for index, v := range intLi {
		strValue := Uint642Str(v)
		strLi[index] = strValue
	}
	return
}

func UInt64ListToStringListWithZeroEmpty(intLi []uint64) (strLi []string) {
	strLi = make([]string, len(intLi))
	for index, v := range intLi {
		strValue := Uint642StrWithZeroEmpty(v)
		strLi[index] = strValue
	}
	return
}

func IsRandomed(ratio float64) bool { //  ratio 一个小于1的小数
	if rand.Int31n(100) <= int32(100*ratio) {
		return true
	}
	return false
}

func IsRandomedInt(ratio int32) bool { //  ratio 一个小于100的小数
	if rand.Int31n(100) <= ratio {
		return true
	}
	return false
}

/* translate连线卡片位置 */
func GetLocalAttackPos(posLi []int32) (posLiLocal []int) {
	posLiLocal = make([]int, len(posLi))
	for index, value := range posLi {
		posLiLocal[index] = int(value)
	}
	return
}

/* 坐标转ID */
func XY2ID(x int, y int) int {
	return x*3 + y
}

/* ID转坐标 */
func ID2XY(pos int) (x int, y int) {
	return int(math.Floor(float64(pos) / 3)), int(math.Mod(float64(pos), 3))
}

/*
* TIME UTIL
 */

func IsSameDay(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return (y1 == y2 && m1 == m2 && d1 == d2)
}

func IsSameMonth(t1 time.Time, t2 time.Time) bool {
	y1, m1, _ := t1.Date()
	y2, m2, _ := t2.Date()
	return (y1 == y2 && m1 == m2)
}

// 1=monday  7 =sunday
func GetWeekDay(t time.Time) int32 {
	if t.Weekday() == time.Sunday {
		return 7
	}
	return int32(t.Weekday())
}

// check("1.2.3.4")
//    check("216.14.49.185")
//    check("1::16")

func CheckIP(ip string) bool {
	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		fmt.Printf("%v is not an IPv4 address\n", trial)
		return false
	}
	if trial.To4().Equal(net.ParseIP("127.0.0.1")) {
		return false
	}

	return true
}

func SplitTrim(str string, sep string) (ret []string) {
	ret = strings.Split(str, sep)
	for idx, _ := range ret {
		ret[idx] = strings.TrimSpace(ret[idx])
	}
	return

}

func SplitAsIntArr(str string, sep string) (ret []int) {
	ret = make([]int, 0)
	for _, elem := range strings.Split(str, sep) {
		value, err := strconv.Atoi(strings.TrimSpace(elem))
		if err == nil {
			ret = append(ret, value)
		}
	}
	return

}

func SplitAsUInt64Arr(str string, sep string) (ret []uint64) {
	ret = make([]uint64, 0)
	for _, elem := range strings.Split(str, sep) {
		value, err := strconv.ParseUint(strings.TrimSpace(elem), 10, 64)
		if err == nil {
			ret = append(ret, value)
		}
	}
	return

}
func AddValueAndPercent(value int32, addvalue int32, addpercent float64) (ret int32) {
	ret = value + addvalue + int32(float64(value)*(1+addpercent))
	return
}

func IPStr2Int(ipStr string) uint32 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0
	}
	return IP2Int(ip)
}
func IP2Int(_ip net.IP) uint32 {
	ip := _ip.To4()
	if ip != nil {
		return binary.BigEndian.Uint32(ip)
	}

	return 0
}
func ReverseASCII(s string) string {
	b := make([]byte, len(s))
	var j int = len(s) - 1
	for i := 0; i <= j; i++ {
		b[j-i] = s[i]
	}
	return string(b)
}

func GetFileMd5(filePath string) string {
	file, inerr := os.Open(filePath)
	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h, file)
		file.Close()
		return hex.EncodeToString(md5h.Sum(nil))
	}
	return ""
}
func GetSumN1(n int32) (sum int32) {
	// n+(n-1)+(n-2)+...+1
	if n <= 0 {
		return 0
	}
	var i int32 = 0
	for i = 1; i <= n; i++ {
		sum += i
	}
	return
}

func CopyFile(dst, src string) (int64, error) {
	sf, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer sf.Close()
	df, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer df.Close()
	return io.Copy(df, sf)
}

func GetFriendPrimaryKey(Uin uint64, friendUin uint64) (smallUin uint64, bigUin uint64) {
	smallUin = Uin
	bigUin = friendUin
	if smallUin > bigUin {
		smallUin = friendUin
		bigUin = Uin
	}
	return
}

func GetFormatDate(time time.Time) (timeStr string) {
	year := strconv.Itoa(time.Year())
	month := time.Month()
	day := time.Day()
	monthStr := ""
	dayStr := ""
	if month < 10 {
		monthStr = "0" + strconv.Itoa(int(month))
	} else {
		monthStr = strconv.Itoa(int(month))
	}
	if day < 10 {
		dayStr = "0" + strconv.Itoa(day)
	} else {
		dayStr = strconv.Itoa(day)
	}
	strs := []string{year, monthStr, dayStr}
	timeStr = strings.Join(strs, "-")
	return
}
