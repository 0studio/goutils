package goutils

import (
	"crypto/md5"
	"encoding/hex"
	// "github.com/cihub/seelog"
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
func GetHttpResponseAsJson(urlStr string, now time.Time, timeout int) (data []byte, err error) {
	client := HttpWithTimeOut(now, timeout)
	response, err := client.Get(urlStr)
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	return body, nil
}

func UrlEncode(urlStr string) string {
	return url.QueryEscape(urlStr)
}

func GetHexMd5UpCase(str string) string {
	return strings.ToUpper(GetHexMd5(str))
}
func GetHexMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))

}
func GetSha1(str string) string {
	t := sha1.New()
	io.WriteString(t, str)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// func postHttpResponseAsJson(urlStr string, values url.Values) (data []byte, err error) {
// 	response, err := http.PostForm(urlStr, values)
// 	if err != nil {
// 		return
// 	}

// 	defer response.Body.Close()
// 	body, _ := ioutil.ReadAll(response.Body)
// 	return body, nil

// }
func PostHttpResponse(urlStr string, content []byte, now time.Time, timeout int) (data []byte, err error) {
	client := HttpWithTimeOut(now, timeout)
	response, err := client.Post(urlStr, "text/html", bytes.NewReader(content))
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body, nil

}
func PostHttpResponseWithCookie(urlStr string, content []byte, now time.Time, timeout int, cookies []*http.Cookie) (data []byte, err error) {
	// expire := time.Now().AddDate(0, 0, 1)
	//    cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/", Expires: expire, MaxAge: 86400}
	// cookie := &http.Cookie{
	//       Name:  http.CanonicalHeaderKey("uid-test"), //Name值为Uid-Test
	//       Value: "1234",
	//   }
	//   r.AddCookie(cookie)

	client := HttpWithTimeOut(now, timeout)
	req, err := http.NewRequest("POST", urlStr, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	response, err := client.Do(req)
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body, nil

}
func PostFormHttpResponse(urlStr string, v url.Values, now time.Time, timeout int) (data []byte, err error) {
	client := HttpWithTimeOut(now, timeout)
	respose, err := client.PostForm(urlStr, v)
	if err != nil {
		return
	}

	defer respose.Body.Close()
	body, _ := ioutil.ReadAll(respose.Body)
	return body, nil
}

func HttpWithTimeOut(now time.Time, timeoutMillSeconds int) http.Client {
	timeoutDur := time.Millisecond * time.Duration(timeoutMillSeconds)
	// 在拨号回调中，使用DialTimeout来支持连接超时，当连接成功后，利用SetDeadline来让连接支持读写超时。
	fun := func(network, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(network, addr, timeoutDur)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(now.Add(timeoutDur))
		return conn, nil
	}
	transport := &http.Transport{Dial: fun, ResponseHeaderTimeout: timeoutDur}

	client := http.Client{
		Transport: transport,
	}
	return client
}
