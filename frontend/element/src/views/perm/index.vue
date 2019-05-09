<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-input v-model="listQuery.search" size="small" style="" placeholder="输入权限搜索"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="small" icon="el-icon-search" @click="handleSearch()">搜索</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="small" icon="el-icon-plus" @click="handleAdd()">添加</el-button>
        </el-form-item>
      </el-form>
      <el-table
        :data="tableData.filter(data => !listQuery.search || data.nick_name.toLowerCase().includes(listQuery.search.toLowerCase()))"
        stripe
        border
        style="width: 100%">
        <el-table-column
          label="权限名"
          prop="name" />
        <el-table-column
          label="权限中文"
          prop="nick_name" />
        <el-table-column
          label="类型"
          prop="type" />
        <el-table-column
          align="center"
          label="操作">
          <template slot-scope="scope">
            <!-- <span @click="handleEdit(scope.$index, scope.row)" class="el-tag el-tag--success">修改</span> -->
            <span class="el-tag el-tag--danger"><a @click="handleDelete(scope.$index, scope.row)">删除</a></span>
            <span class="el-tag el-tag--primary"><a @click="handleEdit(scope.$index, scope.row)">修改</a></span>
          </template>
        </el-table-column>
      </el-table>
      <pagination v-show="total>0" :total="total" :current_page.sync="listQuery.current_page" :page_size.sync="listQuery.page_size" align="right" @pagination="getPermMsg" />
      <el-dialog :visible.sync="dialogPermFormVisible" :close-on-click-modal="false" title="新增/修改权限">
        <el-form ref="permForm" :model="perm_form" label-width="120px">
          <el-form-item label="ID">
            <el-input v-model="perm_form.id" disabled />
          </el-form-item>
          <el-form-item label="权限名">
            <el-input v-model="perm_form.name" />
          </el-form-item>
          <el-form-item label="权限中文">
            <el-input v-model="perm_form.nick_name" />
          </el-form-item>
          <el-form-item label="权限类型">
            <el-input v-model="perm_form.type" />
          </el-form-item>
          <el-form-item align="right" style="padding-right: 48px">
            <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">确定</el-button>
            <el-button @click="dialogPermFormVisible = false">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { deletePerm, addPerm, updatePerm, getPermMsg } from '@/api/perm'
import Pagination from '@/components/Pagination'

export default {
  name: 'Perm',
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
      dialogStatus: '',
      dialogPermFormVisible: false,
      perm_form: {
        id: undefined,
        name: '',
        nick_name: '',
        type: '',
      }
    }
  },
  created() {
    this.getPermMsg()
  },
  methods: {
    resetTemp() {
      this.perm_form = {
        id: undefined,
        name: '',
        nick_name: '',
        type: '',
      }
    },
    getPermMsg() {
      getPermMsg(this.listQuery).then(response => {
        this.tableData = response.data
        this.total = response.total
      })
    },
    handleSearch() {
      getPermMsg(this.listQuery).then(response => {
        this.tableData = response.data
        this.total = response.total
      })
    },
    handleAdd() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogPermFormVisible = true
    },
    handleEdit(index, row) {
      this.resetTemp()
      this.perm_form = Object.assign({}, row)
      this.dialogStatus = 'update'
      this.dialogPermFormVisible = true
    },
    handleDelete(index, row) {
      const This = this
      This.$confirm('是否删除权限?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deletePerm(row.id).then(response => {
          This.$message({
            message: response.message,
            type: 'success',
            duration: 2000,
            onClose: function refresh() {
              This.getPermMsg()
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
      addPerm(this.perm_form).then(response => {
        this.$message({
          message: response.message,
          duration: 2000,
          type: 'success',
          onClose: function refresh() {
            This.dialogPermFormVisible = false
            This.getPermMsg()
          }
        })
      })
    },
    updateData() {
      const This = this
      updatePerm(this.perm_form).then(response => {
        this.$message({
          message: response.message,
          duration: 2000,
          type: 'success',
          onClose: function refresh() {
            This.dialogPermFormVisible = false
            This.getPermMsg()
          }
        })
      })
    }
  }
}
</script>
