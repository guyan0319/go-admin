<template>
  <div class="createPost-container">
    <el-form ref="postForm" :model="postForm" :rules="rules" class="form-container">
      <sticky :z-index="10" :class-name="'sub-navbar draft'">
        <el-button v-loading="loading" style="margin-left: 10px;" type="success" @click="submitForm">
          提交
        </el-button>
      </sticky>

      <div class="createPost-main-container">
        <el-row>
          <el-col :span="10">
            <el-form-item label="名称" prop="name">
              <el-input v-model="postForm.name" placeholder="user Name" clearable/>
            </el-form-item>
          </el-col>

        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label="昵称" prop="nickname">
              <el-input v-model="postForm.nickname" placeholder="user nickName" clearable/>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label="手机号" prop="phone">
              <el-input v-model="postForm.phone" placeholder="user phone" clearable/>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="10">
            <el-form-item label="密码" prop="password">
              <el-input type="password" v-model="postForm.password"  clearable/>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label="确认密码" prop="repassword">
              <el-input type="password" v-model="postForm.repassword"  clearable/>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="状态">
          <el-switch v-model="postForm.status"
                     :on-value="true"
                     :off-value="false"
          ></el-switch>
        </el-form-item>
      </div>
    </el-form>
  </div>
</template>

<script>
import Sticky from '@/components/Sticky' // 粘性header组件
import { createUser } from '@/api/user'

const defaultForm = {
  name: '', // 姓名
  nickname: '', // 昵称
  password: '', // 密码
  phone: '', // 手机号
  id: undefined,
  status: true
}

export default {
  name: 'UserDetail',
  components: { Sticky },
  props: {
    isEdit: {
      type: Boolean,
      default: false
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
    const validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.postForm.password) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      postForm: Object.assign({}, defaultForm),
      loading: false,
      userListOptions: [],
      rules: {
        name: [{ validator: validateRequire }],
        nickname: [{ validator: validateRequire }],
        password: [{ validator: validateRequire }],
        repassword: [{ validator: validatePass2 }]
      },
      tempRoute: {}
    }
  },
  computed: {
    lang() {
      return this.$store.getters.language
    }
  },
  created() {
    // Why need to make a copy of this.$route here?
    // Because if you enter this page and quickly switch tag, may be in the execution of the setTagsViewTitle function, this.$route is no longer pointing to the current page
    // https://github.com/PanJiaChen/vue-element-admin/issues/1221
    this.tempRoute = Object.assign({}, this.$route)
  },
  methods: {
    submitForm() {
      // console.log(this.postForm)
      this.$refs.postForm.validate(valid => {
        if (!valid) {
          this.loading = false
          return false
        }
        createUser(this.postForm).then(() => {
          this.loading = true
          this.$notify({
            title: 'Success',
            dangerouslyUseHTMLString: true,
            message: '添加成功',
            type: 'success',
            duration: 2000
          })
        })
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import "~@/styles/mixin.scss";

.createPost-container {
  position: relative;

  .createPost-main-container {
    padding: 40px 45px 20px 50px;

    .postInfo-container {
      position: relative;
      @include clearfix;
      margin-bottom: 10px;

      .postInfo-container-item {
        float: left;
      }
    }
  }

  .word-counter {
    width: 40px;
    position: absolute;
    right: 10px;
    top: 0px;
  }
}

.article-textarea /deep/ {
  textarea {
    padding-right: 40px;
    resize: none;
    border: none;
    border-radius: 0px;
    border-bottom: 1px solid #bfcbd9;
  }
}
</style>
