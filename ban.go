package goutils

import (
	"github.com/0studio/cachemap"
	"github.com/0studio/logger"
	"net"
	"strings"
	"time"
)

const (
	// 为了应对DoS DDos 攻击，自动临时封ip 的时间
	DOS_BAN_SECONDS         = 5
	DEFAULT_UIN_BAN_SECONDS = 10
)

type Ban struct {
	uins *cachemap.Uint64SafeCacheMap
	ips  *cachemap.Uint32SafeCacheMap
	log  logger.Logger
}

func NewBan(log logger.Logger) (ban Ban) {
	ban = Ban{
		uins: cachemap.NewUint64SafeCacheMap(time.Second),
		ips:  cachemap.NewUint32SafeCacheMap(time.Second),
		log:  log,
	}
	return
}

func (ban *Ban) IsBannedIP(addr net.Addr, now time.Time) (isBanned bool) {
	ip := net.ParseIP(strings.Split(addr.String(), ":")[0])
	intip := IP2Int(ip)
	_, isBanned = ban.ips.Get(intip, now)
	return
}

func (ban *Ban) IsBannedIPByString(ip string, now time.Time) (isBanned bool) {
	intip := IPStr2Int(ip)
	_, isBanned = ban.ips.Get(intip, now)
	return
}
func (ban *Ban) IsBannedUin(Uin uint64, now time.Time) (isBanned bool) {
	if Uin == 0 {
		return false
	}
	_, isBanned = ban.uins.Get(Uin, now)
	return
}
func (ban *Ban) AddBanUin(uin uint64, now time.Time, seconds int, reason string) {
	if reason != "" && ban.log != nil {
		ban.log.Warn("ban_uin", uin, reason)
	}
	ban.uins.Put(uin, cachemap.NewCacheObject(true, now, seconds))
}
func (ban *Ban) AddDefaultBinUin(uin uint64, now time.Time, reason string) {
	ban.AddBanUin(uin, now, (DEFAULT_UIN_BAN_SECONDS + int(LCG())%DEFAULT_UIN_BAN_SECONDS), reason)
}
func (ban *Ban) AddBanIP(addr net.Addr, now time.Time, seconds int, reason string) {
	if ban.log != nil {
		ban.log.Warn("ban_ip_for_reason : ", reason, addr.String())
	}

	intip := IPStr2Int(strings.Split(addr.String(), ":")[0])
	ban.ips.Put(intip, cachemap.NewCacheObject(true, now, seconds))
}
func (ban *Ban) AddBanIPByString(ipStr string, now time.Time, seconds int, reason string) {
	if ban.log != nil {
		ban.log.Warn("ban_ip_for_reason : ", reason, ipStr)
	}
	intip := IPStr2Int(ipStr)
	ban.ips.Put(intip, cachemap.NewCacheObject(true, now, seconds))
}
func (ban *Ban) AddDosBan(addr net.Addr, now time.Time, reason string) {
	ban.AddBanIP(addr, now, (DOS_BAN_SECONDS + int(LCG())%DOS_BAN_SECONDS), reason)
}
