package sql

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"web-demo/lib/common"
)

const (
	//分离表
	tablePattern = `CREATE TABLE([\s\S]*?);`
	// 获取表名称
	tableNamePattern = "CREATE TABLE `(.*)`"
	// 获取字段名称
	namePattern = "`(.*?)` ([\\S]{1,}).* [COMMENT]{0,} ('(.*?)'){0,}"
	//获取表详情
	tableComment = `COMMENT=\'(.*?)\'`
	// 获取字段类型
	namePatternSub = `(.*)\((\d+)\)`
)

var (
	regTableContent     = regexp.MustCompile(tablePattern)
	regTableNamePattern = regexp.MustCompile(tableNamePattern)
	regNamePattern      = regexp.MustCompile(namePattern)
	regNamePatternSub   = regexp.MustCompile(namePatternSub)
	regTableComment     = regexp.MustCompile(tableComment)
)

type field struct {
	Pattern      string
	PatternWidth int
	Comment      string
}

type TableMap struct {
	Table           map[string]map[string]field
	TableCommentMap map[string]string
}

func ParseSql(fileName string) *TableMap {
	tableContentList := splitTable(fileName)
	//tableList := &TableMap{}
	tableList := parseTable(tableContentList)
	return tableList
}

type ApiJs struct {
	Model  string
	Action string
}
type Inventory struct {
	OperationName string
	Name          string
}
type NewInventory struct {
	Fields     []Input
	Computed   []ComputedDetail
	Form       []DefaultForm
	Rule       []RuleDetail
	FunNameApi string
	ImportFile string
	Name       string
}
type RuleDetail struct {
	Content string
}
type DefaultForm struct {
	Col string
}
type Input struct {
	Width       string
	LabelName   string
	PropName    string
	Name        string
	Placeholder string
	Pattern     string
	FunName     string
}

type ComputedDetail struct {
	Type    string
	Name    string
	FunName string
}

//解析table
func parseTable(tableContentList [][]byte) *TableMap {
	//var tableList *TableMap
	tableList := &TableMap{}

	//存储表名：备注
	tableCommentMap := make(map[string]string)
	tableMap := make(map[string]map[string]field)
	for _, tableContent := range tableContentList {
		contentTmp := bytes.Split(tableContent, []byte("\n"))
		tableName := regTableNamePattern.FindSubmatch(contentTmp[0])
		tableComment := regTableComment.FindSubmatch(contentTmp[len(contentTmp)-1])

		var tableCommentStr string
		var keyArr [][]string
		if len(tableComment) > 1 {
			tableCommentStr = string(tableComment[1])
			keyArr = [][]string{{string(tableComment[1])}} // 第一个位置留给table comment
		} else {
			tableCommentStr = ""
			keyArr = [][]string{{}} // 第一个位置留给table comment
		}
		//tableNameNew := strFirstToUpper(string(tableName[1]))
		tableNameNew := string(tableName[1])
		tableCommentMap[tableNameNew] = tableCommentStr

		tableFieldMap := make(map[string]field)
		//  key  key_type key_comment(include)
		for _, contentLine := range contentTmp[1 : len(contentTmp)-1] {
			f := field{}
			keyName := regNamePattern.FindSubmatch(contentLine)
			if len(keyName) > 2 {
				patternSub := regNamePatternSub.FindSubmatch(keyName[2])
				if len(patternSub) > 2 {
					f.Pattern = string(patternSub[1])
					b, _ := strconv.Atoi(string(patternSub[2]))
					f.PatternWidth = b
				} else {
					f.Pattern = string(keyName[2])
				}

				//for ps := range patternSub {
				//	fmt.Println(string(ps))
				//}
				keyTmp := []string{string(keyName[1]), string(keyName[2])}
				if len(keyName) > 3 || string(keyName[3]) != "" {
					//fmt.Println(string(keyName[3]))
					//fmt.Println(string(keyName[4]))
					f.Comment = string(keyName[4])
					keyTmp = append(keyTmp, string(keyName[4]))
				}
				tableFieldMap[string(keyName[1])] = f
				keyArr = append(keyArr, keyTmp)
			}
		}
		tableMap[tableNameNew] = tableFieldMap
	}
	tableList.Table = tableMap
	tableList.TableCommentMap = tableCommentMap
	//fmt.Println(tableList)
	return tableList
}

//拆分表
func splitTable(fileName string) [][]byte {
	readBy, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return regTableContent.FindAll(readBy, -1)
}

//生成views模板 tableMap 表解析数据 tplPath 模板地址
func GenViews(tableMap *TableMap, tplPath, viewPath string) {
	//生成创建view
	GenCreateViews(tableMap, tplPath, viewPath)
	GenEditViews(tableMap, tplPath, viewPath)
	GenListViews(tableMap, tplPath, viewPath)
}
func GenListViews(tableMap *TableMap, tplPath, viewPath string) {
	operation := "CreateForm"
	for k, v := range tableMap.Table {
		fmt.Printf(k, v)
	}
	fmt.Println(operation)

}
func GenCreateViews(tableMap *TableMap, tplPath, viewPath string) {
	operation := "CreateForm"
	for k, v := range tableMap.Table {
		//表名为空跳过
		if k == "" {
			continue
		}
		dirName1 := common.SubstrrPos(k, "_")
		viewPathTemp := viewPath + "/views"
		if dirName1 != "" {
			viewPathTemp = viewPath + "/views/" + dirName1
		}
		dirName2 := common.SubstrPos(k, "_")
		if dirName2 != "" {
			viewPathTemp = viewPathTemp + "/" + dirName2
		}

		createPath := viewPathTemp + "/" + "create"
		//创建create目录
		_ = common.MkDir(createPath, 0755)
		//创建components
		componentsPath := viewPathTemp + "/" + "components"
		_ = common.MkDir(componentsPath, 0755)

		//生成具体操作views
		sweaters := Inventory{operation, common.StrFirstToUpper(k)}
		tmpl, _ := template.ParseFiles(tplPath + "/create/index.vue")

		file, err := os.OpenFile(createPath+"/index.vue", os.O_CREATE|os.O_WRONLY, 0755)
		_ = tmpl.Execute(file, sweaters)
		//生成components
		tmplComp, _ := template.ParseFiles(tplPath + "/components/CreateDetail.vue")

		dirName2New := common.StrFirstToUpper(dirName2)
		sweatersComponents := NewInventory{FunNameApi: dirName2New + "Create", ImportFile: k + ".js", Name: dirName2New}

		for key, value := range v {
			if value.Pattern == "timestamp" {
				continue
			}
			defaultForm := DefaultForm{}
			switch {
			case key == "id":
				defaultForm.Col = "id: undefined,"
			case key == "status":
				defaultForm.Col = "status: true,"
			default:
				defaultForm.Col = key + ": '',"
			}
			//添加默认defaultform
			sweatersComponents.Form = append(sweatersComponents.Form, defaultForm)
			input := Input{Name: key, Width: "10", Pattern: value.Pattern, PropName: key}
			if value.Pattern == "varchar" && value.PatternWidth > 100 {
				input.Width = "20"
			} else if value.Pattern == "datetime" {
				detail := ComputedDetail{Name: key, Type: "datetime", FunName: common.StrFirstToUpper(key)}
				input.FunName = detail.FunName
				sweatersComponents.Computed = append(sweatersComponents.Computed, detail)
			}
			input.LabelName = common.SubstrrPos(value.Comment, "（")
			input.Placeholder = dirName2 + " " + key
			//fmt.Println(key, value)
			sweatersComponents.Fields = append(sweatersComponents.Fields, input)
			//添加rules
			if common.StrPos(value.Comment, "required") >= 0 {
				ruleDetail := RuleDetail{Content: key + ": [{ validator: validateRequire }],"}
				sweatersComponents.Rule = append(sweatersComponents.Rule, ruleDetail)
			}
		}

		file, err = os.OpenFile(componentsPath+"/CreateDetail.vue", os.O_CREATE|os.O_WRONLY, 0755)
		err = tmplComp.Execute(file, sweatersComponents)
		//生成api文件
		GenApiJs(tplPath, viewPath, dirName2New, k)
		fmt.Println(v)
		fmt.Println(err)
		return
	}
}
func GenApiJs(tplPath, viewPath, model, fileName string) {
	//创建api目录
	apiPath := viewPath + "/api"
	_ = common.MkDir(apiPath, 0755)
	//生成具体操作views
	sweaters := ApiJs{Model: model, Action: strings.ToLower(model)}
	tmpl, _ := template.ParseFiles(tplPath + "/user.js")
	fileFullName := apiPath + "/" + fileName + ".js"

	if ok := common.FileExists(fileFullName); ok {
		return
	}
	file, _ := os.OpenFile(fileFullName, os.O_CREATE|os.O_WRONLY, 0755)
	_ = tmpl.Execute(file, sweaters)

}
func GenEditViews(tableMap *TableMap, tplPath, viewPath string) {
	operation := "EditForm"
	for k, v := range tableMap.Table {
		//表名为空跳过
		if k == "" {
			continue
		}
		dirName1 := common.SubstrrPos(k, "_")
		viewPathTemp := viewPath + "/views"
		if dirName1 != "" {
			viewPathTemp = viewPath + "/views/" + dirName1
		}
		dirName2 := common.SubstrPos(k, "_")
		if dirName2 != "" {
			viewPathTemp = viewPathTemp + "/" + dirName2
		}

		createPath := viewPathTemp + "/" + "create"
		//创建create目录
		_ = common.MkDir(createPath, 0755)
		//创建components
		componentsPath := viewPathTemp + "/" + "components"
		_ = common.MkDir(componentsPath, 0755)

		//生成具体操作views
		sweaters := Inventory{operation, common.StrFirstToUpper(k)}
		tmpl, _ := template.ParseFiles(tplPath + "/edit/index.vue")

		file, err := os.OpenFile(createPath+"/index.vue", os.O_CREATE|os.O_WRONLY, 0755)
		_ = tmpl.Execute(file, sweaters)
		//生成components
		tmplComp, _ := template.ParseFiles(tplPath + "/components/EditDetail.vue")

		dirName2New := common.StrFirstToUpper(dirName2)
		sweatersComponents := NewInventory{FunNameApi: dirName2New + "Fetch, " + dirName2New + "Edit", ImportFile: k + ".js", Name: dirName2New}

		for key, value := range v {
			if value.Pattern == "timestamp" {
				continue
			}
			defaultForm := DefaultForm{}
			switch {
			case key == "id":
				defaultForm.Col = "id: undefined,"
			case key == "status":
				defaultForm.Col = "status: true,"
			default:
				defaultForm.Col = key + ": '',"
			}
			//添加默认defaultform
			sweatersComponents.Form = append(sweatersComponents.Form, defaultForm)
			input := Input{Name: key, Width: "10", Pattern: value.Pattern, PropName: key}
			if value.Pattern == "varchar" && value.PatternWidth > 100 {
				input.Width = "20"
			} else if value.Pattern == "datetime" {
				detail := ComputedDetail{Name: key, Type: "datetime", FunName: common.StrFirstToUpper(key)}
				input.FunName = detail.FunName
				sweatersComponents.Computed = append(sweatersComponents.Computed, detail)
			}
			input.LabelName = common.SubstrrPos(value.Comment, "（")
			input.Placeholder = dirName2 + " " + key
			//fmt.Println(key, value)
			sweatersComponents.Fields = append(sweatersComponents.Fields, input)
			//添加rules
			if common.StrPos(value.Comment, "required") >= 0 {
				ruleDetail := RuleDetail{Content: key + ": [{ validator: validateRequire }],"}
				sweatersComponents.Rule = append(sweatersComponents.Rule, ruleDetail)
			}
		}
		file, err = os.OpenFile(componentsPath+"/EditDetail.vue", os.O_CREATE|os.O_WRONLY, 0755)
		err = tmplComp.Execute(file, sweatersComponents)
		//生成api文件
		GenApiJs(tplPath, viewPath, dirName2New, k)
		fmt.Println(v)
		fmt.Println(err)
		return
	}
}
