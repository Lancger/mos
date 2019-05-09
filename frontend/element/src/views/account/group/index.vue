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
          type="index"
          width="50" />
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
            <span class="el-tag el-tag--danger"><a @click="handleDelete(scope.$index, scope.row)" class="a_underline">删除</a></span>
            <span class="el-tag el-tag--primary"><a @click="handleEdit(scope.$index, scope.row)" class="a_underline">修改</a></span>
            <span class="el-tag el-tag--info"><a @click="handlePermMap(scope.$index, scope.row)" class="a_underline">关联权限</a></span>
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
      <el-dialog :visible.sync="dialogGroupPermFormVisible" title="用户组关联权限">
        <el-form ref="mapForm" :model="group_map_form" label-width="90px">
          <el-form-item label="ID">
            <el-input v-model="group_map_form.id" disabled />
          </el-form-item>
          <el-form-item label="用户组名称">
            <el-input v-model="group_map_form.name" disabled/>
          </el-form-item>
          <el-form-item label="权限">
            <el-transfer
              v-model="group_map_form.perm_selected"
              :data="permArray"
              :props="{key: 'id', label: 'nick_name'}"
              :titles="['未选权限', '已选权限']" />
          </el-form-item>
          <el-form-item align="right" style="padding-right: 48px">
            <el-button type="primary" @click="updateGroupPerm">确定</el-button>
            <el-button @click="dialogGroupPermFormVisible = false">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { getGroupList, addGroup, deleteGroup, getGroupMsg, updateGroup, getGroupPermMap, updateGroupPerm } from '@/api/group'
import { getAccountMsg } from '@/api/user'
import { getPermList } from '@/api/perm'
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
      dialogGroupPermFormVisible: false,

      group_form: {
        id: undefined,
        group_name: '',
        nick_name : '',
        user_selected: []
      },
      group_map_form: {
        id: undefined,
        name: '',
        perm_selected: []
      }
    }
  },
  created() {
    this.getGroupList()
    this.getPermList()
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
        perm_selected: []
      }
    },
    handlePermMap(index, row) {
        this.$set(this.group_map_form, 'id', row.id)
        this.$set(this.group_map_form, 'name', row.nick_name)
      getGroupPermMap(row.id).then(response => {
        console.log(response.data)
        this.$set(this.group_map_form, 'perm_selected', response.data)
        // console.log(this.group_map_form.ticketSourceSelected)
      })
      this.dialogGroupPermFormVisible = true
    },
    // 获取权限列表
    getPermList() {
      const This = this
      getPermList().then(response => {
        This.permArray = response.data
      })
    },
    updateGroupPerm() {
      const This = this
      updateGroupPerm(this.group_map_form).then(response => {
        this.$message({
          message: response.message,
          type: 'success',
          duration: 1000,
          onClose: function refresh() {
            This.dialogGroupPermFormVisible = false
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

