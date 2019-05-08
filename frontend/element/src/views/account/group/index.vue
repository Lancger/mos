<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-input v-model="listQuery.search" size="small" style="" placeholder="输入用户组搜索"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="small" icon="el-icon-search" @click="handleSearch()">搜索</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="small" icon="el-icon-plus" @click="handleAdd()">添加</el-button>
        </el-form-item>
      </el-form>
      <el-table
        :data="tableData.filter(data => !listQuery.search || data.group.toLowerCase().includes(listQuery.search.toLowerCase()))"
        stripe
        border
        style="width: 100%">
        <el-table-column
          label="ID"
          prop="id" />
        <el-table-column
          label="用户组名称"
          prop="group_name" />  
        <el-table-column
          label="中文名称"
          prop="nick_name" />
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
            <el-button
              size="mini"
              type="primary"
              icon="el-icon-edit"
              @click="handlePermMap(scope.$index, scope.row)">关联权限</el-button>
            <el-button
              size="mini"
              type="primary"
              icon="el-icon-edit"
              @click="handleEdit(scope.$index, scope.row)">修改</el-button>
            <el-button
              size="mini"
              type="danger"
              icon="el-icon-delete"
              @click="handleDelete(scope.$index, scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <pagination v-show="total>0" :total="total" :current_page.sync="listQuery.current_page" :page_size.sync="listQuery.page_size" align="right" @pagination="getGroupList" />
      <el-dialog :visible.sync="dialogGroupFormVisible" title="新增/修改用户组">
        <el-form ref="accountForm" :model="group_form" label-width="90px">
          <el-form-item label="ID">
            <el-input v-model="group_form.id" disabled />
          </el-form-item>
          <el-form-item label="用户组名称">
            <el-input v-model="group_form.group_name" :disabled=isGroupNameEditAble />
          </el-form-item>
          <el-form-item label="用户组中文名称">
            <el-input v-model="group_form.nick_name" />
          </el-form-item>
          <el-form-item label="组员">
            <el-transfer
              v-model="group_form.user_selected"
              :data="userArray"
              :props="{key: 'id', label: 'nick_name'}"
              :titles="['未选组员', '已选组员']"
              filterable
              filter-placeholder="请输入关键字" />
          </el-form-item>
          <el-form-item align="right" style="padding-right: 48px">
            <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">确定</el-button>
            <el-button @click="dialogGroupFormVisible = false">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
      <el-dialog :visible.sync="dialogGroupMapFormVisible" title="用户组关联权限">
        <el-form ref="mapForm" :model="group_map_form" label-width="90px">
          <el-form-item label="ID">
            <el-input v-model="group_map_form.id" disabled />
          </el-form-item>
          <el-form-item label="用户组名称">
            <el-input v-model="group_map_form.name" disabled/>
          </el-form-item>
          <el-form-item label="权限">
            <el-transfer
              v-model="group_map_form.permSelected"
              :data="permArray"
              :props="{key: 'id', label: 'perm_name'}"
              :titles="['未选权限', '已选权限']" />
          </el-form-item>
          <el-form-item align="right" style="padding-right: 48px">
            <el-button type="primary" @click="mapSourceData">确定</el-button>
            <el-button @click="dialogGroupMapFormVisible = false">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { getGroupList, addGroup, deleteGroup, getGroupMsg, updateGroup } from '@/api/group'
import { getAccountMsg } from '@/api/user'
// import { setGroupTicketSourceMap, setGroupTicketPermMap, getGroupTicketWorkMsg, getGroupTicketWorkMap } from '@/api/ticket'
import Pagination from '@/components/Pagination'

export default {
  name: 'Group',
  components: { Pagination },
  data() {
    return {
      tableData: [],
      userArray: [],
      permArray: [],
      total: 0,
      listQuery: {
        current_page: 1,
        page_size: 20,
        search: ''
      },
      // 用户组名是否可修改
      isGroupNameEditAble: false,
      // 新增，更新用户组状态
      dialogStatus: '',
      // 新增用户组dialog是否显示
      dialogGroupFormVisible: false,
      // 用户组关联权限 dialog是否显示
      dialogGroupMapFormVisible: false,

      group_form: {
        id: undefined,
        group_name: '',
        nick_name : '',
        user_selected: []
      },
      group_map_form: {
        id: undefined,
        name: '',
        permSelected: []
      }
    }
  },
  created() {
    this.getGroupList()
  },
  methods: {
    resetTemp() {
      this.group_form = {
        id: undefined,
        group_name: '',
        nick_name : '',
        user_selected: []
      }
      this.group_map_form = {
        id: undefined,
        name: '',
        permSelected: []
      }
    },
    handlePermMap(index, row) {
      getGroupPermMap(row.id).then(response => {
        this.$set(this.group_map_form, 'permSelected', response.data)
        this.$set(this.group_map_form, 'id', row.id)
        // console.log(this.group_map_form.ticketSourceSelected)
      })
      this.dialogGroupMapFormVisible = true
    },

    mapSourceData() {
      const This = this
      setGroupTicketSourceMap(this.group_map_form).then(response => {
        this.$message({
          message: response.message,
          type: 'success',
          duration: 2000,
          onClose: function refresh() {
            This.dialogGroupMapFormVisible = false
          }
        })
      })
    },
    getGroupList() {
      getGroupList(this.listQuery).then(response => {
        this.tableData = response.data
        this.total = response.total
      })
    },
    handleSearch() {
      getGroupList(this.listQuery).then(response => {
        this.tableData = response.data
        this.total = response.total
      })
    },
    handleAdd() {
      this.resetTemp()
      this.dialogStatus = 'create'
       this.isGroupNameEditAble = false
      this.dialogGroupFormVisible = true
      getAccountMsg().then(response => {
        this.userArray = response.data
      })
    },
    handleEdit(index, row) {
      this.group_form = Object.assign({}, row)
      this.dialogStatus = 'update'
      this.isGroupNameEditAble = true
      getAccountMsg().then(response => {
        this.userArray = response.data
      })
      getGroupMsg(this.group_form.id).then(response => {
        this.$set(this.group_form, 'user_selected', response.data)
        // console.log(this.group_form.userSelected)
      })
      // console.log(index, row)
      this.dialogGroupFormVisible = true
    },
    handleDelete(index, row) {
      const This = this
      this.$confirm('是否删除组?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteGroup(row.id).then(response => {
          this.$message({
            message: response.message,
            type: 'success',
            duration: 2000,
            onClose: function refresh() {
              This.getGroupList()
            }
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
          // console.log(row)
        })
      })
    },
    createData() {
      const This = this
      addGroup(this.group_form).then(response => {
        this.$message({
          message: response.message,
          type: 'success',
          duration: 2000,
          onClose: function refresh() {
            This.dialogGroupFormVisible = false
            This.getGroupList()
          }
        })
      })
    },
    updateData() {
      const This = this
      updateGroup(this.group_form).then(response => {
        this.$message({
          message: response.message,
          type: 'success',
          duration: 2000,
          onClose: function refresh() {
            This.dialogGroupFormVisible = false
            This.getGroupList()
          }
        })
      })
    }
  }
}
</script>
