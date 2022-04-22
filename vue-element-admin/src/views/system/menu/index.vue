<template>
  <div class="app-container">
    <el-button type="primary" @click="handleAddMenu">
      {{ $t('menu.addRoot') }}
    </el-button>

    <el-table :data="MenusList" style="width: 100%;margin-top:30px;" border>
      <el-table-column align="center" label="主键" width="80">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="left" label="菜单名称" width="220">
        <template slot-scope="scope">
          <span v-for="c in scope.row.level">|--</span>{{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column align="header-center" label="路径">
        <template slot-scope="scope">
          {{ scope.row.path }}
        </template>
      </el-table-column>
           <el-table-column align="header-center" label="组件">
        <template slot-scope="scope">
          {{ scope.row.component }}
        </template>
      </el-table-column>
           <el-table-column align="header-center" label="后端接口url" width="180">
        <template slot-scope="scope">
          {{ scope.row.url }}
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="状态" width="80">
        <template slot-scope="{row}">
          <el-tag :type="row.status | statusFilter">
            {{ row.status | statusNameFilter }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作">
        <template slot-scope="scope">
          <el-button @click="handleAddMenu(scope)">
            {{ $t('buttons.add') }}
          </el-button>
          <el-button type="primary" size="small" @click="handleEdit(scope)">
            {{ $t('buttons.edit') }}
          </el-button>
          <el-button type="danger" size="small" @click="handleDelete(scope)">
            {{ $t('buttons.delete') }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :visible.sync="dialogVisible" :title="dialogType==='edit'?'Edit Menu':'New Menu'">
      <el-form ref="formData" :model="Menu" :rules="rules" label-width="80px" label-position="left">
        <el-form-item label="名称" prop="name">
          <el-input v-model="Menu.name" placeholder="Menu Name" clearable/>
        </el-form-item>
        <el-form-item label="路径" prop="path">
          <el-input v-model="Menu.path" placeholder="Menu Path" clearable/>
        </el-form-item>
        <el-form-item label="组件" prop="component">
          <el-input v-model="Menu.component" placeholder="Menu component" clearable/>
        </el-form-item>
        <el-form-item label="跳转">
          <el-input v-model="Menu.redirect" placeholder="Menu redirect" clearable/>
        </el-form-item>
        <el-form-item label="api url" prop="url">
          <el-input v-model="Menu.url" placeholder="Menu url" clearable/>
        </el-form-item>
        <el-form-item label="meta icon">
          <el-input v-model="Menu.meta_icon" placeholder="Menu meta_icon" clearable/>
        </el-form-item>
        <el-form-item label="sort">
          <el-input v-model="Menu.sort" placeholder="Menu sort"/>
        </el-form-item>

        <el-form-item label="是否缓存">
          <el-switch
            v-model="Menu.meta_nocache"
            :on-value="true"
            :off-value="false"
          ></el-switch>
        </el-form-item>
        <el-form-item label="总是展示">
          <el-switch
            v-model="Menu.alwaysshow"
            :on-value="true"
            :off-value="false"
          ></el-switch>
        </el-form-item>

        <el-form-item label="是否隐藏">
          <el-switch
            v-model="Menu.hidden"
            :on-value="true"
            :off-value="false"
          ></el-switch>
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="Menu.status" :on-value="true" :off-value="false"></el-switch>
        </el-form-item>
        <el-input v-model="Menu.pid" type="hidden"/>

      </el-form>
      <div style="text-align:right;">
        <el-button type="danger" @click="dialogVisible=false">
          {{ $t('buttons.cancel') }}
        </el-button>
        <el-button type="primary" @click="confirmMenu">
          {{ $t('buttons.confirm') }}
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>

import { deepClone } from '@/utils'
import { getMenus, addMenu, updateMenu, deleteMenu } from '@/api/menu'
const defaultMenu = {
  id: '',
  name: '',
  path: '',
  component: '',
  redirect: '',
  url: '',
  meta_title: '',
  meta_icon: '#',
  meta_nocache: true,
  alwaysshow: true,
  meta_affix: false,
  hidden: false,
  pid: 0,
  sort: '0',
  status: true
}
export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        1: 'success',
        0: 'info'
      }
      return statusMap[status]
    },
    statusNameFilter(status) {
      const statusMap = {
        1: '启动',
        0: '停止'
      }
      return statusMap[status]
    }
  },
  data() {
    const validateRequire = (rule, value, callback) => {
      if (value === '') {
        this.$message({
          message: rule.field + '为必填项',
          type: 'error'
        })
        callback(new Error(rule.field + '为必填项'))
      } else {
        callback()
      }
    }
    return {
      Menu: Object.assign({}, defaultMenu),
      routes: [],
      MenusList: [],
      dialogVisible: false,
      dialogType: 'new',
      checkStrictly: false,
      rules: {
        name: [{ validator: validateRequire }],
        path: [{ validator: validateRequire }],
        component: [{ validator: validateRequire }],
        url: [{ validator: validateRequire }]
      },
      defaultProps: {
        children: 'children',
        label: 'title'
      }
    }
  },
  computed: {
    routesData() {
      return this.routes
    }
  },
  created() {
    // Mock: get all routes and Menus list from server
    this.getMenus()
  },
  methods: {
    async getMenus() {
      const res = await getMenus()
      this.MenusList = res.data
    },
    handleAddMenu(scope) {
      this.dialogType = 'new'
      this.dialogVisible = true
      this.Menu.pid = scope.row ? scope.row.id : 0
      this.$nextTick(() => {
        this.$refs['formData'].clearValidate()
        this.$refs['formData'].resetFields()
        this.Menu.status = true
      })
    },
    handleEdit(scope) {
      this.dialogType = 'edit'
      this.dialogVisible = true
      this.checkStrictly = true
      this.Menu = deepClone(scope.row)
      this.Menu.status = (this.Menu.status === 1) ? true : false
      this.Menu.meta_nocache = (this.Menu.meta_nocache ===1) ? true : false
      this.Menu.hidden = (this.Menu.hidden ===1) ? true : false
      this.Menu.alwaysshow = (this.Menu.alwaysshow ===1) ? true : false
      this.$nextTick(() => {
        this.$refs['formData'].clearValidate()
      })
    },
    handleDelete({ $index, row }) {
      this.$confirm('Confirm to remove the Menu?', 'Warning', {
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
        type: 'warning'
      })
        .then(async() => {
          await deleteMenu(row.id)
          for (let index = 0; index < this.MenusList.length; index++) {
            if (this.MenusList[index].id === row.id) {
              row.status = 0
              this.MenusList.splice(index, 1, Object.assign({}, row))
              break
            }
          }
          // this.MenusList.splice($index, 1)
          this.$message({
            type: 'success',
            message: 'Delete succed!'
          })
        })
        .catch(err => {
          console.error(err)
        })
    },
    confirmMenu() {
      const isEdit = this.dialogType === 'edit'
      this.$refs['formData'].validate((valid) => {
        if (!valid) {
          return false
        }
        if (isEdit) {
          this.Menu.sort = String(this.Menu.sort)
          updateMenu(this.Menu.id, this.Menu).then(response => {
            for (let index = 0; index < this.MenusList.length; index++) {
              if (this.MenusList[index].id === this.Menu.id) {
                this.Menu.status = this.Menu.status ? 1 : 0
                this.MenusList.splice(index, 1, Object.assign({}, this.Menu))
                break
              }
            }
            const { path, id, name } = this.Menu
            this.dialogVisible = false
            this.$notify({
              title: 'Success',
              dangerouslyUseHTMLString: true,
              message: `
          <div>Menu Id: ${id}</div>
          <div>Menu Name: ${name}</div>
          <div>path: ${path}</div>
        `,
              type: 'success'
            })
          }).catch(err => {
            console.log(err)
          })
        } else {
          addMenu(this.Menu).then(response => {
            this.Menu.id = response.data.id
            this.Menu.status = (response.data.status === 1) ? 1 : 0
            this.Menu.level = this.Menu.level + 1
            this.MenusList.push(this.Menu)
            const { path, id, name } = this.Menu
            this.dialogVisible = false
            this.$notify({
              title: 'Success',
              dangerouslyUseHTMLString: true,
              message: `
          <div>Menu Id: ${id}</div>
          <div>Menu Name: ${name}</div>
          <div>path: ${path}</div>
        `,
              type: 'success'
            })
          }).catch(err => {
            console.log(err)
          })
        }
      })
    }

  }
}
</script>

<style lang="scss" scoped>
  .app-container {
    .Menus-table {
      margin-top: 30px;
    }
    .permission-tree {
      margin-bottom: 30px;
    }
  }
</style>
