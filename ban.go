package goutils

import (
	"github.com/0studio/cachemap"
	"github.com/0studio/logger"
	"net"
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

func (ban *Ban) IsBannedIP(ip net.IP, now time.Time) (isBanned bool) {
	intip := IP2Int(ip)
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
	if reason != "" {
		ban.log.Warn("ban_uin", uin, reason)
	}
	ban.uins.Put(uin, cachemap.NewCacheObject(true, now, seconds))
}
func (ban *Ban) AddDefaultBinUin(uin uint64, now time.Time, reason string) {
	ban.AddBanUin(uin, now, (DEFAULT_UIN_BAN_SECONDS + int(LCG())%DEFAULT_UIN_BAN_SECONDS), reason)
}
func (ban *Ban) AddBanIP(ip net.IP, now time.Time, seconds int, reason string) {
	if ban.log != nil {
		ban.log.Warn("ban_ip_for_reason : ", reason, ip.String())
	}

	intip := IP2Int(ip)
	ban.ips.Put(intip, cachemap.NewCacheObject(true, now, seconds))
}
func (ban *Ban) AddDosBan(ip net.IP, now time.Time, reason string) {
	ban.AddBanIP(ip, now, (DOS_BAN_SECONDS + int(LCG())%DOS_BAN_SECONDS), reason)
}
