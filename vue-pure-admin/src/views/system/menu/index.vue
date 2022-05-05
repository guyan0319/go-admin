<script lang="ts">
export default {
  name: "menu"
};
</script>

<script setup lang="ts">
import dayjs from "dayjs";
import { handleTree } from "/@/utils/tree";
import { getMenuList } from "/@/api/menu";
import { FormInstance } from "element-plus";
import { reactive, ref, onMounted } from "vue";
import { EpTableProBar } from "/@/components/ReTable";
import { useRenderIcon } from "/@/components/ReIcon/src/hooks";
import {VxeFormPropTypes} from "vxe-table";

const dictData = reactive({
  submitLoading: false,
  showEdit: false,
  selectRow: null,
  filterName: "",
  formData: {
    name: "",
    alias: "",
    path: "",
    component: "",
    redirect: "",
    url: "",
    metaTitle: "",
    metaIcon: "",
    metaI18n: 0,
    metaShowlink: 0,
    metaRank: 0,
    metaKeepalive: 0,
    type: 2,
    metaFramesrc: "",
    transitionName: "",
    transitionEnter: "",
    transitionLeave: "",
    dynamiclevel: 0,
    refreshredirect: "",
    extraiconSvg: 0,
    extraiconName: "",
    pid: 0,
    state: 0,
    level: "",
  },
  formItems: [
    {
      field: "name",
      title: "名称",
      span: 24,
      itemRender: {
        name: "$input",
        props: { placeholder: "请输入名称" }
      }
    },
    {
      field: "alias",
      title: "别名",
      span: 24,
      itemRender: {
        name: "$input",
        props: {
          placeholder: "请输入别名",
        }
      }
    },
    {
      field: "path",
      title: "路径",
      span: 24,
      itemRender: {
        name: "$input",
        props: { placeholder: "请输入路径",}
      }
    },
   {
      field: "redirect",
      title: "重定向",
      span: 24,
      itemRender: {
        name: "$input",
        props: { placeholder: "请输入重定向",}
      }
    },
 {
      field: "url",
      title: "url",
      span: 24,
      itemRender: {
        name: "$input",
        props: { placeholder: "请输入url",}
      }
    },
 {
      field: "metaTitle",
      title: "meta标题",
      span: 24,
      itemRender: {
        name: "$input",
        props: { placeholder: "请输入meta标题",}
      }
    },
 {
      field: "meta_icon",
      title: "meta icon",
      span: 24,
      itemRender: {
        name: "$input",
        props: { placeholder: "请输入meta icon",}
      }
    },
{
      field: "meta_i18n",
      title: "是否国际化",
      span: 24,
      itemRender: {
        name: "$select",

        // props: { name: "请输入字典名称",label:1 }
      },
    },

    {
      align: "right",
      span: 24,
      itemRender: {
        name: "$buttons",
        children: [
          { props: { type: "submit", content: "提交", status: "primary" } },
          { props: { type: "reset", content: "重置" } }
        ]
      }
    }
  ] as VxeFormPropTypes.Items
});
let dataList = ref([]);
let loading = ref(true);

const formRef = ref<FormInstance>();
const tableRef = ref();
// 新增
function onAdd() {
  commonFn(null, false);
}

// 新增子类型
function onAddChild(row?: object) {
  console.log("onAddChild", row);
  commonFn(null, false);
}

function handleUpdate(row) {

  commonFn(row, true);
  // VXETable.modal.message({
  //   content: "测试数据，不可编辑",
  //   status: "error"
  // });
  console.log(row);
}
function commonFn(value, disabled) {
  dictData.selectRow = value;
  dictData.showEdit = true;
  dictData.formItems[1].itemRender.props.disabled = disabled;
}

function handleDelete(row) {
  console.log(row);
}

function handleSelectionChange(val) {
  console.log("handleSelectionChange", val);
}

async function onSearch() {
  loading.value = true;
  let { data } = await getMenuList();
  dataList.value = handleTree(data);
  setTimeout(() => {
    loading.value = false;
  }, 500);
}

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
  onSearch();
};

onMounted(() => {
  onSearch();
});
</script>

<template>
  <div class="main">
    <el-form
      ref="formRef"
      :inline="true"
      :model="dictData.formData"
      class="bg-white w-99/100 pl-8 pt-4"
    >
      <el-form-item label="菜单名称：" prop="user">
        <el-input v-model="dictData.formData.name" placeholder="请输入菜单名称" clearable />
      </el-form-item>
      <el-form-item label="状态：" prop="state">
        <el-select v-model="dictData.formData.state" placeholder="请选择状态" clearable>
          <el-option label="开启" value="1" />
          <el-option label="关闭" value="0" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          :icon="useRenderIcon('search')"
          :loading="loading"
          @click="onSearch"
        >
          搜索
        </el-button>
        <el-button :icon="useRenderIcon('refresh')" @click="resetForm(formRef)">
          重置
        </el-button>
      </el-form-item>
    </el-form>

    <EpTableProBar
      title="菜单列表"
      :loading="loading"
      :tableRef="tableRef"
      :dataList="dataList"
      @refresh="onSearch"
    >
      <template #buttons>
        <el-button type="primary" :icon="useRenderIcon('add')"  @click="onAdd">
          新增菜单
        </el-button>
      </template>
      <template v-slot="{ size, checkList }">
        <el-table
          ref="tableRef"
          border
          row-key="id"
          table-layout="auto"
          default-expand-all
          :size="size"
          :data="dataList"
          :header-cell-style="{ background: '#fafafa', color: '#606266' }"
          @selection-change="handleSelectionChange"
        >
          <el-table-column
            v-if="checkList.includes('勾选列')"
            type="selection"
            align="center"
            width="55"
          />
          <el-table-column
            v-if="checkList.includes('序号列')"
            type="index"
            align="center"
            label="序号"
            width="60"
          />
          <el-table-column label="菜单名称" prop="name" width="180" />
          <el-table-column label="路径" align="center" prop="path" width="160" />
          <el-table-column label="重定向" align="center" prop="redirect" width="160" />
          <el-table-column label="标题" align="center" prop="metaTitle" width="160" />
          <el-table-column label="状态" align="center" prop="state" width="80">
            <template #default="scope">
              <el-tag
                :size="size"
                :type="scope.row.state === 0 ? 'danger' : 'success'"
                effect="plain"
              >
                {{ scope.row.state === 0 ? "关闭" : "开启" }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column
            label="创建时间"
            align="center"
            width="180"
            prop="ctime"
            :formatter="
              ({ ctime }) => {
                return dayjs(ctime).format('YYYY-MM-DD HH:mm:ss');
              }
            "
          />
          <el-table-column
            fixed="right"
            label="操作"
            align="center"
            width="140"
          >
            <template #default="scope">
              <el-button
                class="reset-margin"
                type="text"
                :size="size"
                @click="handleUpdate(scope.row)"
                :icon="useRenderIcon('edits')"
              >
                修改
              </el-button>
              <el-popconfirm title="是否确认删除?">
                <template #reference>
                  <el-button
                    class="reset-margin"
                    type="text"
                    :size="size"
                    :icon="useRenderIcon('delete')"
                    @click="handleDelete(scope.row)"
                  >
                    删除
                  </el-button>
                </template>
              </el-popconfirm>
            </template>
          </el-table-column>
        </el-table>
      </template>
    </EpTableProBar>
    <!-- 修改、添加弹框 -->
    <vxe-modal
      resize
      width="450"
      v-model="dictData.showEdit"
      :title="dictData.selectRow ? '编辑' : '新增'"
      :loading="dictData.submitLoading"
      @hide="$refs.xForm.reset()"
    >
      <template #default>
        <vxe-form
          ref="xForm"
          :data="dictData.formData"
          :items="dictData.formItems"
          title-align="right"
          title-width="100"
          @submit="submitEvent"
        />
      </template>
    </vxe-modal>

  </div>
</template>
