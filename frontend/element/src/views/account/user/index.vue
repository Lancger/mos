<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-input v-model="listQuery.search" size="small" style="" placeholder="输入用户名搜索"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="small" icon="el-icon-search" @click="handleSearch()">搜索</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="small" icon="el-icon-plus" @click="handleAdd()">添加</el-button>
        </el-form-item>
      </el-form>
      <el-table
        :data="tableData.filter(data => !listQuery.search || data.username.toLowerCase().includes(listQuery.search.toLowerCase()))"
        stripe
        border
        style="width: 100%">
        <el-table-column
          label="用户名"
          prop="username" />
        <el-table-column
          label="中文名"
          prop="nick_name" />
        <el-table-column
          label="邮箱"
          prop="email" />
        <el-table-column
          label="手机号"
          prop="phone" />
        <el-table-column
          align="center"
          label="创建时间"
          prop="create_time">
          <template slot-scope="scope">
            <span>{{ scope.row.create_time }}</span>
          </template>
        </el-table-column>  
        <el-table-column
          align="center"
          label="操作">
          <template slot-scope="scope">
            <!-- <span @click="handleEdit(scope.$index, scope.row)" class="el-tag el-tag--success">修改</span> -->
            <el-button
              size="mini"
              type="primary"
              class="pan-btn blue-btn"
              @click="handleEdit(scope.$index, scope.row)">修改</el-button>
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <pagination v-show="total>0" :total="total" :current_page.sync="listQuery.current_page" :page_size.sync="listQuery.page_size" align="right" @pagination="getList" />
      <el-dialog :visible.sync="dialogAccountFormVisible" title="新增用户">
        <el-form ref="accountForm" :model="account_form" label-width="120px">
          <el-form-item label="ID">
            <el-input v-model="account_form.id" disabled />
          </el-form-item>
          <el-form-item label="用户名">
            <el-input v-model="account_form.username" />
          </el-form-item>
          <el-form-item label="密码">
            <el-input v-model="account_form.password" type="password" />
          </el-form-item>
          <el-form-item label="中文名">
            <el-input v-model="account_form.nick_name" />
          </el-form-item>
          <el-form-item label="邮箱">
            <el-input
              v-model="account_form.email"
              :rules="[
                { required: true, message: '请输入账号名', trigger: 'blur' },
                { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }
            ]"/>
          </el-form-item>
          <el-form-item label="手机号">
            <el-input v-model="account_form.phone" />
          </el-form-item>
          <el-form-item align="right" style="padding-right: 48px">
            <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">确定</el-button>
            <el-button @click="dialogAccountFormVisible = false">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { getAccountList, deleteAccount, addAccount, updateAccount } from '@/api/user'
import Pagination from '@/components/Pagination'

export default {
  name: 'User',
  components: { Pagination },
  data() {
    return {
      tableData: [],
      total: 0,
      listQuery: {
        current_page: 1,
        page_size: 20,
        search: ''
      },
      password_disable: true,
      dialogStatus: '',
      dialogAccountFormVisible: false,
      account_form: {
        id: undefined,
        username: '',
        password: '',
        nick_name: '',
        email: '',
        phone: ''
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    resetTemp() {
      this.account_form = {
        id: undefined,
        username: '',
        password: '',
        nick_name: '',
        email: '',
        phone: ''
      }
    },
    getList() {
      getAccountList(this.listQuery).then(response => {
        this.tableData = response.data
        this.total = response.total
      })
    },
    handleSearch() {
      getAccountList(this.listQuery).then(response => {
        this.tableData = response.data
        this.total = response.total
      })
    },
    handleAdd() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogAccountFormVisible = true
    },
    handleEdit(index, row) {
      this.resetTemp()
      this.account_form = Object.assign({}, row)
      this.dialogStatus = 'update'
      this.password_disable = true
      this.dialogAccountFormVisible = true
    },
    handleDelete(index, row) {
      const This = this
      This.$confirm('是否删除用户?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteAccount(row.id).then(response => {
          This.$message({
            message: response.message,
            type: 'success',
            duration: 2000,
            onClose: function refresh() {
              This.getList()
            }
          })
        }).catch(() => {
          This.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
      })
    },
    createData() {
      const This = this
      addAccount(this.account_form).then(response => {
        this.$message({
          message: response.message,
          duration: 2000,
          type: 'success',
          onClose: function refresh() {
            This.dialogAccountFormVisible = false
            This.getList()
          }
        })
      })
    },
    updateData() {
      const This = this
      updateAccount(this.account_form).then(response => {
        this.$message({
          message: response.message,
          duration: 2000,
          type: 'success',
          onClose: function refresh() {
            This.dialogAccountFormVisible = false
            This.getList()
          }
        })
      })
    }
  }
}
</script>
