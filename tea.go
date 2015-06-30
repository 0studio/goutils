package goutils

// tea 加密解密
import (
	"math/rand"
)

func encipher(vv []byte, kk []byte) (encryptcode []byte) {
	var v = make([]uint, 8)
	for i := 0; i < 8; i++ {
		v[i] = uint(vv[i])
	}
	var y, z uint = v[0]<<24 + v[1]<<16 + v[2]<<8 + v[3], v[4]<<24 + v[5]<<16 + v[6]<<8 + v[7]
	var sum, delta, i uint = 0, 0x9e3779b9, 0
	var k = make([]uint, 16)
	for i := 0; i < 16; i++ {
		k[i] = uint(kk[i])
	}
	var a, b, c, d uint = k[0]<<24 + k[1]<<16 + k[2]<<8 + k[3], k[4]<<24 + k[5]<<16 + k[6]<<8 + k[7], k[8]<<24 + k[9]<<16 + k[10]<<8 + k[11], k[12]<<24 + k[13]<<16 + k[14]<<8 + k[15]
	for i = 0; i < 16; i++ {
		sum = sum & 0xffffffff
		delta = delta & 0xffffffff
		sum += delta
		y += ((z << 4) + a) ^ (z + sum) ^ ((z >> 5) + b)
		y &= 0xffffffff
		z += ((y << 4) + c) ^ (y + sum) ^ ((y >> 5) + d)
		z &= 0xffffffff
	}
	encryptcode = make([]byte, 8)
	encryptcode[0] = byte(y >> 24)
	encryptcode[1] = byte((y << 8) >> 24)
	encryptcode[2] = byte((y << 16) >> 24)
	encryptcode[3] = byte((y << 24) >> 24)
	encryptcode[4] = byte(z >> 24)
	encryptcode[5] = byte((z << 8) >> 24)
	encryptcode[6] = byte((z << 16) >> 24)
	encryptcode[7] = byte((z << 24) >> 24)
	return encryptcode
}

func TeaEncrypt(in string, k string) (encryptString string) {
	return TeaEncryptByteArray([]byte(in), k)
}
func TeaEncryptByteArray(inBuff []byte, k string) (encryptString string) {
	var encryptcode []byte
	var key = []byte(k)
	if key == nil {
		return encryptString
	}
	random := rand.New(rand.NewSource(99))
	var plain = make([]byte, 8)
	var preplain = make([]byte, 8)
	var pos, padding int = 1, 0
	var mlen = len(inBuff)
	pos = (mlen + 0x0A) % 8
	if pos != 0 {
		pos = 8 - pos
	}
	plain[0] = (byte)((random.Int() & 0xF8) | pos)
	var i int
	for i = 1; i <= pos; i++ {
		plain[i] = (byte)(random.Int() & 0xFF)
	}
	pos = pos + 1
	for i = 0; i < 8; i++ {
		preplain[i] = 0x0
	}
	var ivCrypt = make([]byte, 8)
	copy(ivCrypt, preplain)
	padding = 1
	for padding <= 2 {
		if pos < 8 {
			plain[pos] = (byte)(random.Int() & 0xFF)
			pos++
			padding++
		}
		if pos == 8 {
			for j := 0; j < 8; j++ {
				plain[j] ^= ivCrypt[j]
			}
			outBuff := encipher(plain, key)
			for j := 0; j < 8; j++ {
				outBuff[j] ^= preplain[j]
			}
			encryptcode = append(encryptcode, outBuff...)
			copy(preplain, plain)
			copy(ivCrypt, outBuff)
			pos = 0
		}
	}
	i = 0
	for mlen > 0 {
		if pos < 8 {
			plain[pos] = inBuff[i]
			pos++
			i++
			mlen--
		}
		if pos == 8 {
			for j := 0; j < 8; j++ {
				plain[j] ^= ivCrypt[j]
			}
			outBuff := encipher(plain, key)
			for j := 0; j < 8; j++ {
				outBuff[j] ^= preplain[j]
			}
			encryptcode = append(encryptcode, outBuff...)
			copy(preplain, plain)
			copy(ivCrypt, outBuff)
			pos = 0
		}
	}

	padding = 0
	for padding <= 7 {
		if pos < 8 {
			plain[pos] = 0x0
			pos++
			padding++
		}
		if pos == 8 {
			for j := 0; j < 8; j++ {
				plain[j] ^= ivCrypt[j]
			}
			outBuff := encipher(plain, key)
			for j := 0; j < 8; j++ {
				outBuff[j] ^= preplain[j]
			}
			encryptcode = append(encryptcode, outBuff...)
			copy(preplain, plain)
			copy(ivCrypt, outBuff)
			pos = 0
		}
	}
	encryptString = string(encryptcode)
	return encryptString
}

func decipher(vv []byte, kk []byte) (decryptcode []byte) {
	var v = make([]uint, 8)
	for i := 0; i < 8; i++ {
		v[i] = uint(vv[i])
	}
	var y, z uint = v[0]<<24 + v[1]<<16 + v[2]<<8 + v[3], v[4]<<24 + v[5]<<16 + v[6]<<8 + v[7]
	var sum, delta, i uint = 0xe3779b90, 0x9e3779b9, 0
	var k = make([]uint, 16)
	for i := 0; i < 16; i++ {
		k[i] = uint(kk[i])
	}
	var a, b, c, d uint = k[0]<<24 + k[1]<<16 + k[2]<<8 + k[3], k[4]<<24 + k[5]<<16 + k[6]<<8 + k[7], k[8]<<24 + k[9]<<16 + k[10]<<8 + k[11], k[12]<<24 + k[13]<<16 + k[14]<<8 + k[15]
	for i = 0; i < 16; i++ {
		sum = sum & 0xffffffff
		delta = delta & 0xffffffff
		z -= ((y << 4) + c) ^ (y + sum) ^ ((y >> 5) + d)
		z &= 0xffffffff
		y -= ((z << 4) + a) ^ (z + sum) ^ ((z >> 5) + b)
		y &= 0xffffffff
		sum -= delta
	}
	decryptcode = make([]byte, 8)
	decryptcode[0] = byte(y >> 24)
	decryptcode[1] = byte((y << 8) >> 24)
	decryptcode[2] = byte((y << 16) >> 24)
	decryptcode[3] = byte((y << 24) >> 24)
	decryptcode[4] = byte(z >> 24)
	decryptcode[5] = byte((z << 8) >> 24)
	decryptcode[6] = byte((z << 16) >> 24)
	decryptcode[7] = byte((z << 24) >> 24)
	return decryptcode
}

func TeaDecrypt(in string, k string) (plain string) {
	return TeaDecryptByteArray([]byte(in), k)
}
func TeaDecryptByteArray(inBuff []byte, k string) (plain string) {
	//检查密钥
	var decryptcode []byte
	var key = []byte(k)
	if key == nil {
		return plain
	}
	//检查消息字节数是8的倍数，且至少是16字节
	var mlen = len(inBuff)
	if mlen%8 != 0 || mlen < 16 {
		return plain
	}
	//得到消息头部，得到明文开始位置，信息存在第一个字节里面，将解密得到的信息的第一个字节与7做与运算
	var headBuf = make([]byte, 8)
	for i := 0; i < 8; i++ {
		headBuf[i] = inBuff[i]
	}
	var preplain = make([]byte, 8)
	for i := 0; i < 8; i++ {
		preplain[i] = 0x0
	}
	var decryptBuf = decipher(headBuf, key)
	for j := 0; j < 8; j++ {
		decryptBuf[j] ^= 0x0
	}

	var pos int = int(decryptBuf[0] & 0x7)
	//得到明文的长度，不小于0
	var count = mlen - pos - 10
	//用于最后解密明文数组切片
	var posSlice, countSlice = pos + 3, count
	if count < 0 {
		return plain
	}
	var plainBuff = make([]byte, 8)
	var ivCrypt = make([]byte, 8)
	var temp = make([]byte, 8)
	var preCrypt int = 0
	pos++
	var padding = 1
	for padding <= 2 {
		if pos < 8 {
			pos++
			padding++
		}
		if pos == 8 {
			for j := 0; j < 8; j++ {
				plainBuff[j] = inBuff[preCrypt+j]
				temp[j] = inBuff[preCrypt+j]
			}
			for j := 0; j < 8; j++ {
				plainBuff[j] ^= preplain[j]
			}
			decryptBuf = decipher(plainBuff, key)
			if decryptBuf == nil {
				return plain
			}
			copy(preplain, decryptBuf)
			for j := 0; j < 8; j++ {
				decryptBuf[j] ^= ivCrypt[j]
			}
			copy(ivCrypt, temp)
			preCrypt += 8
			pos = 0
			decryptcode = append(decryptcode, decryptBuf...)
		}
	}

	var i int = 0
	for count != 0 {
		if pos < 8 {
			i++
			count--
			pos++
		}
		if pos == 8 {
			for j := 0; j < 8; j++ {
				plainBuff[j] = inBuff[preCrypt+j]
				temp[j] = inBuff[preCrypt+j]
			}
			for j := 0; j < 8; j++ {
				plainBuff[j] ^= preplain[j]
			}
			decryptBuf = decipher(plainBuff, key)
			if decryptBuf == nil {
				return plain
			}
			copy(preplain, decryptBuf)
			for j := 0; j < 8; j++ {
				decryptBuf[j] ^= ivCrypt[j]
			}

			copy(ivCrypt, temp)
			preCrypt += 8
			pos = 0
			decryptcode = append(decryptcode, decryptBuf...)
		}
	}
	for padding = 1; padding < 8; padding++ {
		if pos < 8 {
			pos++
		}
		if pos == 8 {
			for j := 0; j < 8; j++ {
				plainBuff[j] = inBuff[preCrypt+j]
				temp[j] = inBuff[preCrypt+j]
			}
			for j := 0; j < 8; j++ {
				plainBuff[j] ^= preplain[j]
			}
			decryptBuf = decipher(plainBuff, key)
			if decryptBuf == nil {
				return plain
			}
			copy(preplain, decryptBuf)
			for j := 0; j < 8; j++ {
				decryptBuf[j] ^= ivCrypt[j]
			}
			copy(ivCrypt, temp)
			preCrypt += 8
			pos = 0
			decryptcode = append(decryptcode, decryptBuf...)
		}
	}
	decryptcode = decryptcode[posSlice : countSlice+posSlice]
	plain = string(decryptcode)

	return plain
}
