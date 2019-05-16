<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-input v-model="listQuery.search" size="small" style="" placeholder="输入项目名搜索"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="small" icon="el-icon-search" @click="handleSearch()">搜索</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="small" icon="el-icon-plus" @click="handleAdd()">添加</el-button>
        </el-form-item>
      </el-form>
      <el-table
        :data="tableData.filter(data => !listQuery.search || data.name.toLowerCase().includes(listQuery.search.toLowerCase()))"
        stripe
        border
        style="width: 100%">
        <el-table-column
          type="index"
          width="50" />
        <el-table-column
          label="项目名称"
          prop="name" />
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
            <span class="el-tag el-tag--danger"><a @click="handleDelete(scope.$index, scope.row)" class="a_underline">删除</a></span>
            <span class="el-tag el-tag--primary"><a @click="handleEdit(scope.$index, scope.row)" class="a_underline">修改</a></span>
          </template>
        </el-table-column>
      </el-table>
      <pagination v-show="total>0" :total="total" :current_page.sync="listQuery.current_page" :page_size.sync="listQuery.page_size" align="right" @pagination="getList" />
      <el-dialog :visible.sync="dialogProjectFormVisible" title="新增/修改项目">
        <el-form ref="typeForm" :model="project_form" label-width="120px">
          <el-form-item label="ID">
            <el-input v-model="project_form.id" disabled />
          </el-form-item>
          <el-form-item label="项目名称">
            <el-input v-model="project_form.name" />
          </el-form-item>
          <el-form-item align="right" style="padding-right: 48px">
            <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">确定</el-button>
            <el-button @click="dialogProjectFormVisible = false">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { getProjectList,  addProject, updateProject, deleteProject } from '@/api/project'
import Pagination from '@/components/Pagination'

export default {
  name: 'Project',
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
      dialogProjectFormVisible: false,
      project_form: {
        id: undefined,
        name: ''
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    resetTemp() {
      this.project_form = {
        id: undefined,
        name: ''
      }
    },
    getList() {
      getProjectList(this.listQuery).then(response => {
        this.tableData = response.data
        this.total = response.total
      })
    },
    handleSearch() {
      getProjectList(this.listQuery).then(response => {
        this.tableData = response.data
        this.total = response.total
      })
    },
    handleAdd() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogProjectFormVisible = true
    },
    handleEdit(index, row) {
      this.resetTemp()
      this.project_form = Object.assign({}, row)
      this.dialogStatus = 'update'
      this.dialogProjectFormVisible = true
    },
    handleDelete(index, row) {
      const This = this
      This.$confirm('是否删除项目?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteProject(row.id).then(response => {
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
      addProject(this.project_form).then(response => {
        this.$message({
          message: response.message,
          duration: 2000,
          type: 'success',
          onClose: function refresh() {
            This.dialogProjectFormVisible = false
            This.getList()
          }
        })
      })
    },
    updateData() {
      const This = this
      updateProject(this.project_form).then(response => {
        this.$message({
          message: response.message,
          duration: 2000,
          type: 'success',
          onClose: function refresh() {
            This.dialogProjectFormVisible = false
            This.getList()
          }
        })
      })
    }
  }
}
</script>
