<template>
  <div class="app-container">
    <el-button type="primary" @click="handleAddMenu">
      {{ $t('menu.addRoot') }}
    </el-button>

    <el-table :data="MenusList" style="width: 100%;margin-top:30px;" border>
      <el-table-column align="center" label="Menu Key" width="220">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="Menu Name" width="220">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column align="header-center" label="Description">
        <template slot-scope="scope">
          {{ scope.row.path }}
        </template>
      </el-table-column>
      <el-table-column align="header-center" label="Description">
        <template slot-scope="scope">
          {{ scope.row.status }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="Operations">
        <template slot-scope="scope">
          <el-button type="primary" size="small" @click="handleEdit(scope)">
            {{ $t('permission.editPermission') }}
          </el-button>
          <el-button type="danger" size="small" @click="handleDelete(scope)">
            {{ $t('permission.delete') }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :visible.sync="dialogVisible" :title="dialogType==='edit'?'Edit Menu':'New Menu'">
      <el-form :model="Menu" label-width="80px" label-position="left">
        <el-form-item label="Name">
          <el-input v-model="Menu.name" placeholder="Menu Name" />
        </el-form-item>
        <el-form-item label="Desc">
          <el-input
            v-model="Menu.description"
            :autosize="{ minRows: 2, maxRows: 4}"
            type="textarea"
            placeholder="Menu Description"
          />
        </el-form-item>
        <el-form-item label="Menus">
          <el-tree ref="tree" :check-strictly="checkStrictly" :data="routesData" :props="defaultProps" show-checkbox node-key="path" class="permission-tree" />
        </el-form-item>
      </el-form>
      <div style="text-align:right;">
        <el-button type="danger" @click="dialogVisible=false">
          {{ $t('permission.cancel') }}
        </el-button>
        <el-button type="primary" @click="confirmMenu">
          {{ $t('permission.confirm') }}
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import path from 'path'
import { deepClone } from '@/utils'
import { getMenus } from '@/api/menu'
// import { getMenus, addMenu, deleteMenu, updateMenu } from '@/api/menu'
const defaultMenu = {
  key: '',
  name: '',
  description: ''
}
export default {
  data() {
    return {
      Menu: Object.assign({}, defaultMenu),
      routes: [],
      MenusList: [],
      dialogVisible: false,
      dialogType: 'new',
      checkStrictly: false,
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
    handleAddMenu() {
      this.Menu = Object.assign({}, defaultMenu)
      if (this.$refs.tree) {
        this.$refs.tree.setCheckedNodes([])
      }
      this.dialogType = 'new'
      this.dialogVisible = true
    },
    handleEdit(scope) {
      this.dialogType = 'edit'
      this.dialogVisible = true
      this.checkStrictly = true
      this.Menu = deepClone(scope.row)
    },
    handleDelete({ $index, row }) {
      this.$confirm('Confirm to remove the Menu?', 'Warning', {
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
        type: 'warning'
      })
        .then(async() => {
          await deleteMenu(row.key)
          this.MenusList.splice($index, 1)
          this.$message({
            type: 'success',
            message: 'Delete succed!'
          })
        })
        .catch(err => { console.error(err) })
    },
    generateTree(routes, basePath = '/', checkedKeys) {
      const res = []

      for (const route of routes) {
        const routePath = path.resolve(basePath, route.path)

        // recursive child routes
        if (route.children) {
          route.children = this.generateTree(route.children, routePath, checkedKeys)
        }

        if (checkedKeys.includes(routePath) || (route.children && route.children.length >= 1)) {
          res.push(route)
        }
      }
      return res
    },
    async confirmMenu() {
      const isEdit = this.dialogType === 'edit'

      const checkedKeys = this.$refs.tree.getCheckedKeys()
      this.Menu.routes = this.generateTree(deepClone(this.serviceRoutes), '/', checkedKeys)

      if (isEdit) {
        await updateMenu(this.Menu.key, this.Menu)
        for (let index = 0; index < this.MenusList.length; index++) {
          if (this.MenusList[index].key === this.Menu.key) {
            this.MenusList.splice(index, 1, Object.assign({}, this.Menu))
            break
          }
        }
      } else {
        const { data } = await addMenu(this.Menu)
        this.Menu.key = data.key
        this.MenusList.push(this.Menu)
      }

      const { description, key, name } = this.Menu
      this.dialogVisible = false
      this.$notify({
        title: 'Success',
        dangerouslyUseHTMLString: true,
        message: `
            <div>Menu Key: ${key}</div>
            <div>Menu Name: ${name}</div>
            <div>Description: ${description}</div>
          `,
        type: 'success'
      })
    },
    // reference: src/view/layout/components/Sidebar/SidebarItem.vue
    onlyOneShowingChild(children = [], parent) {
      let onlyOneChild = null
      const showingChildren = children.filter(item => !item.hidden)

      // When there is only one child route, the child route is displayed by default
      if (showingChildren.length === 1) {
        onlyOneChild = showingChildren[0]
        onlyOneChild.path = path.resolve(parent.path, onlyOneChild.path)
        return onlyOneChild
      }

      // Show parent if there are no child route to display
      if (showingChildren.length === 0) {
        onlyOneChild = { ... parent, path: '', noShowingChildren: true }
        return onlyOneChild
      }

      return false
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
