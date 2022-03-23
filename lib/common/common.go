package common

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

func StrJoin(args ...string) string {
	var buffer bytes.Buffer
	//接受的参数放在args数组中
	for _, e := range args {
		buffer.WriteString(e)
	}
	return buffer.String()
}

//截取字符
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//获取上一级目录
func GetParentDirectory(dirctory string) string {
	return Substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

//获取当前目录
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
func GetAbsDir() string {
	dir, err := filepath.Abs("")
	if err != nil {
		return ""
	}
	return dir
}

//转换为大驼峰命名法则
//首字母大写，“_” 忽略后大写
func StrFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}

//查找某值是否在数组中
func InArrayString(v string, m *[]string) bool {
	for _, value := range *m {
		if value == v {
			return true
		}
	}
	return false
}
func WriteFile(path string, content string) (string, bool) {
	s := `^data:\s*image\/(\w+);base64,`
	b, _ := regexp.MatchString(s, content)
	if !b {
		return "", false
	}
	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	allData := re.FindAllSubmatch([]byte(content), 2)
	fileType := string(allData[0][1])
	base64Str := re.ReplaceAllString(content, "")
	date := time.Now().Format("20060102")
	if ok := FileExists(path + "/" + date); !ok {
		_ = os.Mkdir(path+"/"+date, 0666)
	}
	relative := date + "/" + GetRandomBoth(32) + "." + fileType
	buffer, _ := base64.StdEncoding.DecodeString(base64Str)

	err := ioutil.WriteFile(path+"/"+relative, buffer, 0666)
	if err != nil {
		log.Println(err)
	}
	return relative, true
}
func MkDir(path string, perm os.FileMode) error {
	var err error
	if ok := FileExists(path); !ok {
		err = os.MkdirAll(path, perm)
	}
	return err
}
func Base64Content(url, path, content string) string {
	reg := regexp.MustCompile(`data:\s*image\/(\w+);base64,[\w\d+/=]*[=|==]`)
	imageArr := reg.FindAllString(content, -1)

	for _, v := range imageArr {
		imgPath, res := WriteFile(path, v)
		if !res {
			continue
		}
		content = strings.Replace(content, v, url+imgPath, 1)
	}
	return content
}
func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("not in array")
}

//获取随机数 纯文字
func GetRandomString(n int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//获取随机数  数字和文字
func GetRandomBoth(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//获取随机数  纯数字
func GetRandomNum(n int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//获取随机数  base32
func GetRandomBase32(n int) string {
	str := "234567abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//生成区间随机数
func RandInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

//sha1加密
func Sha1En(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//对字符串进行MD5哈希
func Md5En(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//自定义唯一id
func GetUniqueId() string {
	cur := time.Now()
	timestamps := cur.UnixNano()
	uid := strconv.FormatInt(timestamps, 10) + GetRandomNum(5)
	return Md5En(uid)
}

//自定义唯一id
func OrderUniqueId() string {
	cur := time.Now()
	timestamps := cur.UnixNano() / 1000000 //获取毫秒
	return strconv.FormatInt(timestamps, 10) + GetRandomNum(5)
}

//获取程序运行路径
func GetRunDirectory() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return strings.Replace(dir, "\\", "/", -1), nil
}

//验证手机
func MobileVerify(data map[string]interface{}) bool {
	if data["country"] == "86" {
		reg := `^1\d{10}$`
		rgx := regexp.MustCompile(reg)
		if !rgx.MatchString(data["mobile"].(string)) {
			return false
		}
	} else {
		reg := `^(00){1}\d+`
		rgx := regexp.MustCompile(reg)
		if !rgx.MatchString(data["mobile"].(string)) {
			return false
		}
	}
	return true
}

//验证邮箱
func MailVerify(email string) bool {
	reg := `^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+`
	rgx := regexp.MustCompile(reg)
	if !rgx.MatchString(email) {
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径文件/文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//替换名字为星号
func ReplaceName(str string) string {
	if str == "" {
		return ""
	}
	return string([]rune(str)[:1]) + "**"
}

//浮点数位数
func DecimalValue(value float64, num string) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%."+num+"f", value), 64)
	return value
}

//获取地区
func GetCityByIp(ip string) string {
	return ""
}

//时间戳转字符
func GetDate(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02")
}

//时间戳转字符 带具体时间
func GetDatetime(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(TIME_LAYOUT)
}

//字符转时间戳
func StrToDatetime(dates string) int64 {
	tm2, _ := time.Parse("2006-01-02", dates)
	return tm2.Unix()
}

//字符转时间戳   带详细时间
func StrToTime(dates string) int64 {
	tm2, _ := time.Parse(TIME_LAYOUT, dates)
	return tm2.Unix()
}

//字符转时间格式   2020-04-19T16:00:00.000Z
func StrToTimes(dates string) time.Time {
	//layout := "2006-01-02T15:04:05.000Z"
	//t, err := time.Parse(layout, str)
	t, _ := time.Parse(time.RFC3339, dates)
	return t
}

//时间格式转字符   2022-01-12 20:28:09
func TimeToStr(dates time.Time) string {
	return dates.Format(TIME_LAYOUT)
}

//目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//f:需要处理的浮点数，n：要保留小数的位数
//Pow10（）返回10的n次方，最后一位四舍五入，对ｎ＋１位加０．５后四舍五入
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

//获取本地ip
func GetLocalIp() (IpAddr string) {
	IpAddr = "127.0.0.1"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic("Get local IP addr failed!!!")
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				IpAddr = ipnet.IP.String()
			}
		}
	}
	return
}
func CheckErr(err error) {
	if err != nil {
		//panic(err)
		//return fmt.Errorf("Operation error: %s \n", err)
		fmt.Printf("Operation error: %s \n", err)
		os.Exit(0)
	}
}
func ParseFile(path string) map[string]interface{} {
	content, err := ioutil.ReadFile(path)
	CheckErr(err)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(content), &data)
	CheckErr(err)
	return data
}

//是否是文件
func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

func TypeOfV(v interface{}) string {
	return reflect.TypeOf(v).String()
}
func ShowMsg(v interface{}) {
	fmt.Printf("Operation error: %s \n", v)
	os.Exit(0)
	//panic(v)
}

//加载包
func LoadPackage(p string) error {
	var c *exec.Cmd
	if runtime.GOOS != "linux" {
		c = exec.Command("cmd", "/C", "go", "get", p)
	} else {
		c = exec.Command("go", "get", p)
	}
	if err := c.Run(); err != nil {
		return err
	}
	return nil
}

//截取指定字符子串
func SubstrContains(s, substr string) string {
	n := strings.Index(s, substr)
	return s[n:]
}

//截取指定字符子串,不包含子串
func SubstrPos(s, substr string) string {
	n := strings.Index(s, substr)
	return s[n+len(substr):]
}
func SubstrrPos(s, substr string) string {
	n := strings.Index(s, substr)
	if n < 1 {
		return s
	}
	return s[:n]
}

func Fwrite(fname, content string) error {
	var d1 = []byte(content)
	return ioutil.WriteFile(fname, d1, 0666) //写入文件(字节数组)
}

//返回子串str在字符串s中第一次出现的位置。
//如果找不到则返回-1；如果str为空，则返回0

func StrPos(s, need string) int {
	return strings.Index(s, need)
}

func LogWriter(name string, con ...interface{}) {
	t := time.Now()
	path := "./log/" + t.Format("20060102")
	err := MkDir(path, 0766)
	if err != nil {
		return
	}
	fileFullName := path + "/" + name
	// 按照所需读写权限创建文件
	f, err := os.OpenFile(fileFullName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	// 完成后延迟关闭
	defer f.Close()

	//设置日志输出到 f
	log.SetOutput(f)
	//写入日志内容
	//log.Println(con)
	log.Printf("%+v", con)
}

//根据layout 20060102150405  或 "2006-01-02 15:04:05"
func GetDateByLayout(layout string) string {
	return time.Now().Format(layout)
}
