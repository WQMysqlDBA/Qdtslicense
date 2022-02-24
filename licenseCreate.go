package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
	"time"
)

var CertificateUsagePeriod string

func init() {
	flag.StringVar(&CertificateUsagePeriod, "Duration", "30", "the Duration of the certificate")
}
func main() {
	flag.Parse()
	timeObj := time.Now()
	year := timeObj.Year()

	var aliasmonth string
	month := timeObj.Month()

	if timeObj.Month() < 10 {
		aliasmonth = "0" + strconv.Itoa(int(month))
	} else {
		aliasmonth = strconv.Itoa(int(month))
	}
	var aliasday string
	day := timeObj.Day()

	if timeObj.Day() < 10 {
		aliasday = "0" + strconv.Itoa(day)
	} else {
		aliasday = strconv.Itoa(day)
	}

	formatdata := strconv.Itoa(year) + "-" + aliasmonth + "-" + aliasday

	//formatdata:="2021-11-17"
	//byte_data := []byte(formatdata)
	// 将 byte 装换为 16进制的字符串
	//hex_string_data := hex.EncodeToString(byte_data)
	// 如何生成日期的 16进制数字
	// byte 转 16进制 的结果
	hexStringDataKey := key2secret(formatdata)
	hexStringDataPeriod := key2secret(CertificateUsagePeriod)
	lic := hexStringDataKey + "@Wsq12qa1Q&" + hexStringDataPeriod
	fmt.Println("生成的license: ", lic)
	howtouse := "echo -e \"" + lic + "\\c\" > .QFUSION_DTS_LICENSE"
	fmt.Printf("执行%s\n", howtouse)

	/* ====== 分割线 ====== */
	/* ====== 解密license ====== */
	// 将 16进制的字符串 转换 byte
	hex_data, _ := hex.DecodeString(hexStringDataKey)
	// 将 byte 转换 为字符串 输出结果
	period, _ := hex.DecodeString(hexStringDataPeriod)
	fmt.Printf("证书生成日期: %s,有效期: %s\n", string(hex_data), period)
	//checklicense()
}

func key2secret(k string) (s string) {
	// firstly , we put the string to a type of byte
	tmp := []byte(k)
	s = hex.EncodeToString(tmp)
	return s
}
func secret2key() {}

//func checklicense() {
//	timeObj := time.Now()
//	year := timeObj.Year()
//	month := timeObj.Month()
//	day := timeObj.Day()
//	formatdata:= strconv.Itoa(year) +"-"+ strconv.Itoa(int(month)) +"-"+ strconv.Itoa(day)
//	byte_data := []byte(formatdata)
//	// 将 byte 装换为 16进制的字符串
//	hex_string_data := hex.EncodeToString(byte_data)
//	// 如何生成日期的 16进制数字
//	// byte 转 16进制 的结果
//	println("生成的license: ",hex_string_data)

/* ====== 分割线 ====== */

// 将 16进制的字符串 转换 byte
//hex_data, _ := hex.DecodeString(hex_string_data)
//// 将 byte 转换 为字符串 输出结果
//println(string(hex_data))
//
//loc, _ := time.LoadLocation("Local")
//timeLayout := "2006-01-02 00:00:00"  //转化所需模板
//time1:="2021-10-26 00:00:00"
//time2:="2021-11-17 00:00:00"
//tmp1, _ := time.ParseInLocation(timeLayout, time1, loc)
//tmp2, _ := time.ParseInLocation(timeLayout, time2, loc)
//timestamp1 := tmp1.Unix()
//timestamp2 := tmp2.Unix()
//date :=	(timestamp2- timestamp1)/ 86400
//fmt.Println(date)

//}
