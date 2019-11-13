package main

import (
	"encoding/json"
	"fmt"
	"go-admin/models"
	"time"
)

func main()  {
	constantRoutes:="[{\"path\":\"/redirect\",\"component\":\"layout/Layout\",\"hidden\":true,\"children\":[{\"path\":\"/redirect/:path*\",\"component\":\"views/redirect/index\"}]},{\"path\":\"/login\",\"component\":\"views/login/index\",\"hidden\":true},{\"path\":\"/auth-redirect\",\"component\":\"views/login/auth-redirect\",\"hidden\":true},{\"path\":\"/404\",\"component\":\"views/error-page/404\",\"hidden\":true},{\"path\":\"/401\",\"component\":\"views/error-page/401\",\"hidden\":true},{\"path\":\"\",\"component\":\"layout/Layout\",\"redirect\":\"dashboard\",\"children\":[{\"path\":\"dashboard\",\"component\":\"views/dashboard/index\",\"name\":\"Dashboard\",\"meta\":{\"title\":\"dashboard\",\"icon\":\"dashboard\",\"affix\":true}}]},{\"path\":\"/documentation\",\"component\":\"layout/Layout\",\"children\":[{\"path\":\"index\",\"component\":\"views/documentation/index\",\"name\":\"Documentation\",\"meta\":{\"title\":\"documentation\",\"icon\":\"documentation\",\"affix\":true}}]},{\"path\":\"/guide\",\"component\":\"layout/Layout\",\"redirect\":\"/guide/index\",\"children\":[{\"path\":\"index\",\"component\":\"views/guide/index\",\"name\":\"Guide\",\"meta\":{\"title\":\"guide\",\"icon\":\"guide\",\"noCache\":true}}]}]"
	ansyncRoutes:="[{\"path\":\"/permission\",\"component\":\"layout/Layout\",\"redirect\":\"/permission/index\",\"alwaysShow\":true,\"meta\":{\"title\":\"permission\",\"icon\":\"lock\",\"roles\":[\"admin\",\"editor\"]},\"children\":[{\"path\":\"page\",\"component\":\"views/permission/page\",\"name\":\"PagePermission\",\"meta\":{\"title\":\"pagePermission\",\"roles\":[\"admin\"]}},{\"path\":\"directive\",\"component\":\"views/permission/directive\",\"name\":\"DirectivePermission\",\"meta\":{\"title\":\"directivePermission\"}},{\"path\":\"role\",\"component\":\"views/permission/role\",\"name\":\"RolePermission\",\"meta\":{\"title\":\"rolePermission\",\"roles\":[\"admin\"]}}]},{\"path\":\"/icon\",\"component\":\"layout/Layout\",\"children\":[{\"path\":\"index\",\"component\":\"views/icons/index\",\"name\":\"Icons\",\"meta\":{\"title\":\"icons\",\"icon\":\"icon\",\"noCache\":true}}]},{\"path\":\"/components\",\"component\":\"layout/Layout\",\"redirect\":\"noRedirect\",\"name\":\"ComponentDemo\",\"meta\":{\"title\":\"components\",\"icon\":\"component\"},\"children\":[{\"path\":\"tinymce\",\"component\":\"views/components-demo/tinymce\",\"name\":\"TinymceDemo\",\"meta\":{\"title\":\"tinymce\"}},{\"path\":\"markdown\",\"component\":\"views/components-demo/markdown\",\"name\":\"MarkdownDemo\",\"meta\":{\"title\":\"markdown\"}},{\"path\":\"json-editor\",\"component\":\"views/components-demo/json-editor\",\"name\":\"JsonEditorDemo\",\"meta\":{\"title\":\"jsonEditor\"}},{\"path\":\"split-pane\",\"component\":\"views/components-demo/split-pane\",\"name\":\"SplitpaneDemo\",\"meta\":{\"title\":\"splitPane\"}},{\"path\":\"avatar-upload\",\"component\":\"views/components-demo/avatar-upload\",\"name\":\"AvatarUploadDemo\",\"meta\":{\"title\":\"avatarUpload\"}},{\"path\":\"dropzone\",\"component\":\"views/components-demo/dropzone\",\"name\":\"DropzoneDemo\",\"meta\":{\"title\":\"dropzone\"}},{\"path\":\"sticky\",\"component\":\"views/components-demo/sticky\",\"name\":\"StickyDemo\",\"meta\":{\"title\":\"sticky\"}},{\"path\":\"count-to\",\"component\":\"views/components-demo/count-to\",\"name\":\"CountToDemo\",\"meta\":{\"title\":\"countTo\"}},{\"path\":\"mixin\",\"component\":\"views/components-demo/mixin\",\"name\":\"ComponentMixinDemo\",\"meta\":{\"title\":\"componentMixin\"}},{\"path\":\"back-to-top\",\"component\":\"views/components-demo/back-to-top\",\"name\":\"BackToTopDemo\",\"meta\":{\"title\":\"backToTop\"}},{\"path\":\"drag-dialog\",\"component\":\"views/components-demo/drag-dialog\",\"name\":\"DragDialogDemo\",\"meta\":{\"title\":\"dragDialog\"}},{\"path\":\"drag-select\",\"component\":\"views/components-demo/drag-select\",\"name\":\"DragSelectDemo\",\"meta\":{\"title\":\"dragSelect\"}},{\"path\":\"dnd-list\",\"component\":\"views/components-demo/dnd-list\",\"name\":\"DndListDemo\",\"meta\":{\"title\":\"dndList\"}},{\"path\":\"drag-kanban\",\"component\":\"views/components-demo/drag-kanban\",\"name\":\"DragKanbanDemo\",\"meta\":{\"title\":\"dragKanban\"}}]},{\"path\":\"/charts\",\"component\":\"layout/Layout\",\"redirect\":\"noRedirect\",\"name\":\"Charts\",\"meta\":{\"title\":\"charts\",\"icon\":\"chart\"},\"children\":[{\"path\":\"keyboard\",\"component\":\"views/charts/keyboard\",\"name\":\"KeyboardChart\",\"meta\":{\"title\":\"keyboardChart\",\"noCache\":true}},{\"path\":\"line\",\"component\":\"views/charts/line\",\"name\":\"LineChart\",\"meta\":{\"title\":\"lineChart\",\"noCache\":true}},{\"path\":\"mixchart\",\"component\":\"views/charts/mixChart\",\"name\":\"MixChart\",\"meta\":{\"title\":\"mixChart\",\"noCache\":true}}]},{\"path\":\"/nested\",\"component\":\"layout/Layout\",\"redirect\":\"/nested/menu1/menu1-1\",\"name\":\"Nested\",\"meta\":{\"title\":\"nested\",\"icon\":\"nested\"},\"children\":[{\"path\":\"menu1\",\"component\":\"views/nested/menu1/index\",\"name\":\"Menu1\",\"meta\":{\"title\":\"menu1\"},\"redirect\":\"/nested/menu1/menu1-1\",\"children\":[{\"path\":\"menu1-1\",\"component\":\"views/nested/menu1/menu1-1\",\"name\":\"Menu1-1\",\"meta\":{\"title\":\"menu1-1\"}},{\"path\":\"menu1-2\",\"component\":\"views/nested/menu1/menu1-2\",\"name\":\"Menu1-2\",\"redirect\":\"/nested/menu1/menu1-2/menu1-2-1\",\"meta\":{\"title\":\"menu1-2\"},\"children\":[{\"path\":\"menu1-2-1\",\"component\":\"views/nested/menu1/menu1-2/menu1-2-1\",\"name\":\"Menu1-2-1\",\"meta\":{\"title\":\"menu1-2-1\"}},{\"path\":\"menu1-2-2\",\"component\":\"views/nested/menu1/menu1-2/menu1-2-2\",\"name\":\"Menu1-2-2\",\"meta\":{\"title\":\"menu1-2-2\"}}]},{\"path\":\"menu1-3\",\"component\":\"views/nested/menu1/menu1-3\",\"name\":\"Menu1-3\",\"meta\":{\"title\":\"menu1-3\"}}]},{\"path\":\"menu2\",\"name\":\"Menu2\",\"component\":\"views/nested/menu2/index\",\"meta\":{\"title\":\"menu2\"}}]},{\"path\":\"/example\",\"component\":\"layout/Layout\",\"redirect\":\"/example/list\",\"name\":\"Example\",\"meta\":{\"title\":\"example\",\"icon\":\"example\"},\"children\":[{\"path\":\"create\",\"component\":\"views/example/create\",\"name\":\"CreateArticle\",\"meta\":{\"title\":\"createArticle\",\"icon\":\"edit\"}},{\"path\":\"edit/:id(\\\\d+)\",\"component\":\"views/example/edit\",\"name\":\"EditArticle\",\"meta\":{\"title\":\"editArticle\",\"noCache\":true},\"hidden\":true},{\"path\":\"list\",\"component\":\"views/example/list\",\"name\":\"ArticleList\",\"meta\":{\"title\":\"articleList\",\"icon\":\"list\"}}]},{\"path\":\"/tab\",\"component\":\"layout/Layout\",\"children\":[{\"path\":\"index\",\"component\":\"views/tab/index\",\"name\":\"Tab\",\"meta\":{\"title\":\"tab\",\"icon\":\"tab\"}}]},{\"path\":\"/error\",\"component\":\"layout/Layout\",\"redirect\":\"noRedirect\",\"name\":\"ErrorPages\",\"meta\":{\"title\":\"errorPages\",\"icon\":\"404\"},\"children\":[{\"path\":\"401\",\"component\":\"views/error-page/401\",\"name\":\"Page401\",\"meta\":{\"title\":\"page401\",\"noCache\":true}},{\"path\":\"404\",\"component\":\"views/error-page/404\",\"name\":\"Page404\",\"meta\":{\"title\":\"page404\",\"noCache\":true}}]},{\"path\":\"/error-log\",\"component\":\"layout/Layout\",\"redirect\":\"noRedirect\",\"children\":[{\"path\":\"log\",\"component\":\"views/error-log/index\",\"name\":\"ErrorLog\",\"meta\":{\"title\":\"errorLog\",\"icon\":\"bug\"}}]},{\"path\":\"/excel\",\"component\":\"layout/Layout\",\"redirect\":\"/excel/export-excel\",\"name\":\"Excel\",\"meta\":{\"title\":\"excel\",\"icon\":\"excel\"},\"children\":[{\"path\":\"export-excel\",\"component\":\"views/excel/export-excel\",\"name\":\"ExportExcel\",\"meta\":{\"title\":\"exportExcel\"}},{\"path\":\"export-selected-excel\",\"component\":\"views/excel/select-excel\",\"name\":\"SelectExcel\",\"meta\":{\"title\":\"selectExcel\"}},{\"path\":\"export-merge-header\",\"component\":\"views/excel/merge-header\",\"name\":\"MergeHeader\",\"meta\":{\"title\":\"mergeHeader\"}},{\"path\":\"upload-excel\",\"component\":\"views/excel/upload-excel\",\"name\":\"UploadExcel\",\"meta\":{\"title\":\"uploadExcel\"}}]},{\"path\":\"/zip\",\"component\":\"layout/Layout\",\"redirect\":\"/zip/download\",\"alwaysShow\":true,\"meta\":{\"title\":\"zip\",\"icon\":\"zip\"},\"children\":[{\"path\":\"download\",\"component\":\"views/zip/index\",\"name\":\"ExportZip\",\"meta\":{\"title\":\"exportZip\"}}]},{\"path\":\"/pdf\",\"component\":\"layout/Layout\",\"redirect\":\"/pdf/index\",\"children\":[{\"path\":\"index\",\"component\":\"views/pdf/index\",\"name\":\"PDF\",\"meta\":{\"title\":\"pdf\",\"icon\":\"pdf\"}}]},{\"path\":\"/pdf/download\",\"component\":\"views/pdf/download\",\"hidden\":true},{\"path\":\"/theme\",\"component\":\"layout/Layout\",\"redirect\":\"noRedirect\",\"children\":[{\"path\":\"index\",\"component\":\"views/theme/index\",\"name\":\"Theme\",\"meta\":{\"title\":\"theme\",\"icon\":\"theme\"}}]},{\"path\":\"/clipboard\",\"component\":\"layout/Layout\",\"redirect\":\"noRedirect\",\"children\":[{\"path\":\"index\",\"component\":\"views/clipboard/index\",\"name\":\"ClipboardDemo\",\"meta\":{\"title\":\"clipboardDemo\",\"icon\":\"clipboard\"}}]},{\"path\":\"/i18n\",\"component\":\"layout/Layout\",\"children\":[{\"path\":\"index\",\"component\":\"views/i18n-demo/index\",\"name\":\"I18n\",\"meta\":{\"title\":\"i18n\",\"icon\":\"international\"}}]},{\"path\":\"external-link\",\"component\":\"layout/Layout\",\"children\":[{\"path\":\"https://github.com/PanJiaChen/vue-element-admin\",\"meta\":{\"title\":\"externalLink\",\"icon\":\"link\"}}]}]"
	special:="[{\"path\":\"*\",\"redirect\":\"/404\",\"hidden\":true}]"
	//处理constant
	add(constantRoutes,1)
	add(ansyncRoutes,2)
	add(special,3)
	fmt.Println("执行完成")

}
func add(jsonStr string ,typev int){
	var data []interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println(err)
	}
	//处理固定地址
	for _,value:=range data{
		//var children []models.SystemMenu
		menu :=models.SystemMenu{Type:typev,Status:1,Ctime:time.Now()}
		var roles []interface{}
		var children []interface{}

		for k,v:=range value.(map[string]interface{}){
			switch k {
			case "path":
				menu.Path=v.(string)
			case "component":
				menu.Component=v.(string)
			case "redirect":
				menu.Redirect=v.(string)
			case "hidden":
				if v.(bool) {
					menu.Hidden=1
				}
			case "alwaysShow":
				if v.(bool) {
					menu.Alwaysshow=1
				}
			case "meta":
				if v.(map[string]interface{})["title"] !=nil {
					menu.MetaTitle = v.(map[string]interface{})["title"].(string)
				}
				if  v.(map[string]interface{})["icon"] !=nil{
					menu.MetaIcon = v.(map[string]interface{})["icon"].(string)
					//fmt.Println(menu.MetaIcon)
				}
				if v.(map[string]interface{})["noCache"]!=nil {
					if v.(map[string]interface{})["noCache"].(bool){
						menu.MetaNocache = 1
					}
				}
				if  v.(map[string]interface{})["affix"]!=nil {
					if  v.(map[string]interface{})["affix"].(bool){
						menu.MetaAffix = 1
					}
				}
				if  v.(map[string]interface{})["roles"]!=nil {
					roles = v.(map[string]interface{})["roles"].([]interface{})
				}
			}
			if k=="children" {
				children=v.([]interface{})
			}
		}

		//插入menu
		_,err :=menu.Add()
		if err!=nil {
			return
		}
		//处理roles
		if len(roles)>0  {
			srm :=models.SystemRoleMenu{SystemMenuId:menu.Id}
			if !addrole(roles,srm){
				fmt.Println("插入SystemRole失败")
				return
			}
		}
		//处理childen
		if len(children)>0 {
			for _,v:=range children{
				menuChildren:=models.SystemMenu{Type:typev,Status:1,Pid:menu.Id,Ctime:time.Now()}
				if  v.(map[string]interface{})["component"]!=nil {
					menuChildren.Component = v.(map[string]interface{})["component"].(string)
				}
				if  v.(map[string]interface{})["name"]!=nil {
					menuChildren.Name = v.(map[string]interface{})["name"].(string)
				}
				if  v.(map[string]interface{})["path"]!=nil {
					menuChildren.Path = v.(map[string]interface{})["path"].(string)
				}
				if  v.(map[string]interface{})["meta"]!=nil {
					meta:=v.(map[string]interface{})["meta"].(map[string]interface{})
					if meta["title"]!=nil {
						menuChildren.MetaTitle = meta["title"].(string)
					}
					if meta["roles"]!=nil {
						roles = meta["roles"].([]interface{})
					}
				}
				_,err :=menuChildren.Add()
				fmt.Println(menuChildren)
				if err!=nil {
					return
				}
				if len(roles)>0  {
					srm :=models.SystemRoleMenu{SystemMenuId:menuChildren.Id}
					if !addrole(roles,srm){
						fmt.Println("插入SystemRole失败")
						return
					}
					roles=nil
				}
			}
		}
	}

}
func addrole(roles []interface{},srm models.SystemRoleMenu)bool{
	for _,r:=range roles {
		role := models.SystemRole{Name:r.(string),AliasName:r.(string),Ctime:time.Now()}
		res := role.Add()
		if !res {
			return false
		}
		srm.SystemRoleId=role.Id
		srm.Id=0
		rmbool:=srm.Add()
		if !rmbool {
			return false
		}
	}
	return true
}