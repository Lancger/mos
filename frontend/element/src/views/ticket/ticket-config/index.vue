<template>
    <div class="app-container">
        <div class="filter-container">
            <el-tabs style="margin-top:15px;">
                <!-- 工单类型 -->
                <el-tab-pane label="工单类型">
                    <el-form :inline="true" class="demo-form-inline">
                        <el-form-item>
                        <el-input v-model="ticketTypeQuery.search" size="small" style="" placeholder="输入工单类型搜索"/>
                        </el-form-item>
                        <el-form-item>
                        <el-button type="primary" size="small" icon="el-icon-search" @click="handleTicketTypeSearch()">搜索</el-button>
                        </el-form-item>
                        <el-form-item>
                        <el-button type="success" size="small" icon="el-icon-plus" @click="handleTicketTypeAdd()">添加</el-button>
                        </el-form-item>
                    </el-form>
                    <el-table
                        :data="dataTicketType.filter(data => !ticketTypeQuery.search || dataTicketType.name.toLowerCase().includes(ticketTypeQuery.search.toLowerCase()))"
                        stripe
                        border
                        style="width: 100%">
                        <el-table-column
                        type="index"
                        width="50" />
                        <el-table-column
                        label="工单类型"
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
                            <!-- <span @click="handleEdit(scope.$index, scope.row)" class="el-tag el-tag--success">修改</span> -->
                            <span class="el-tag el-tag--danger"><a @click="handleTicketTypeDelete(scope.$index, scope.row)" class="a_underline">删除</a></span>
                            <span class="el-tag el-tag--primary"><a @click="handleTicketTypeEdit(scope.$index, scope.row)" class="a_underline">修改</a></span>
                        </template>
                        </el-table-column>
                    </el-table>
                    <pagination v-show="total.ticketType>0" :total="total.ticketType" :current_page.sync="ticketTypeQuery.current_page" :page_size.sync="ticketTypeQuery.page_size" align="right" @pagination="getTicketTypeList" />
                    <el-dialog :visible.sync="dialogTypeFormVisible" title="新增/修改工单类型">
                        <el-form ref="typeForm" :model="type_form" label-width="120px">
                        <el-form-item label="ID">
                            <el-input v-model="type_form.id" disabled />
                        </el-form-item>
                        <el-form-item label="工单类型">
                            <el-input v-model="type_form.name" />
                        </el-form-item>
                        <el-form-item align="right" style="padding-right: 48px">
                            <el-button type="primary" @click="dialogTypeStatus==='create'?createTicketTypeData():updateTicketTypeData()">确定</el-button>
                            <el-button @click="dialogTypeFormVisible = false">取消</el-button>
                        </el-form-item>
                        </el-form>
                    </el-dialog>
                </el-tab-pane>
                <!-- 工单等级 -->
                <el-tab-pane label="工单等级">
                    <el-form :inline="true" class="demo-form-inline">
                        <el-form-item>
                        <el-button type="success" size="small" icon="el-icon-plus" @click="handleTicketLevelAdd()">添加</el-button>
                        </el-form-item>
                    </el-form>
                    <el-table
                        :data="dataTicketLevel"
                        stripe
                        border
                        style="width: 100%">
                        <el-table-column
                        type="index"
                        width="50" />
                        <el-table-column
                        label="等级名称"
                        prop="name" />
                        <el-table-column
                        label="工单等级"
                        prop="level" />
                        <el-table-column
                        label="工单时效"
                        prop="time" />
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
                            <span class="el-tag el-tag--danger"><a @click="handleTicketLevelDelete(scope.$index, scope.row)" class="a_underline">删除</a></span>
                            <span class="el-tag el-tag--primary"><a @click="handleTicketLevelEdit(scope.$index, scope.row)" class="a_underline">修改</a></span>
                        </template>
                        </el-table-column>
                    </el-table>
                    <el-dialog :visible.sync="dialogLevelFormVisible" title="新增/修改工单等级">
                        <el-form ref="levelForm" :model="level_form" label-width="120px">
                        <el-form-item label="ID">
                            <el-input v-model="level_form.id" disabled />
                        </el-form-item>
                        <el-form-item label="工单中文名称">
                            <el-input v-model="level_form.name" />
                        </el-form-item>
                        <el-form-item label="工单等级">
                            <el-input-number v-model="level_form.level" :min="0" :max="8" label="等级" />
                        </el-form-item>
                        <el-form-item label="工单时效">
                            <el-input-number v-model="level_form.time" :min="0" :max="128" label="天" />
                        </el-form-item>
                        <el-form-item align="right" style="padding-right: 48px">
                            <el-button type="primary" @click="dialogLevelStatus==='create'?createTicketLevelData():updateTicketLevelData()">确定</el-button>
                            <el-button @click="dialogLevelFormVisible = false">取消</el-button>
                        </el-form-item>
                        </el-form>
                    </el-dialog>
                </el-tab-pane>
                 <!-- 工单来源 -->
                <el-tab-pane label="工单来源">
                    <el-form :inline="true" class="demo-form-inline">
                        <el-form-item>
                        <el-button type="success" size="small" icon="el-icon-plus" @click="handleTicketSourceAdd()">添加</el-button>
                        </el-form-item>
                    </el-form>
                    <el-table
                        :data="dataTicketSource"
                        stripe
                        border
                        style="width: 100%">
                        <el-table-column
                        type="index"
                        width="50" />
                        <el-table-column
                        label="来源名称"
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
                            <span class="el-tag el-tag--danger"><a @click="handleTicketSourceDelete(scope.$index, scope.row)" class="a_underline">删除</a></span>
                            <span class="el-tag el-tag--primary"><a @click="handleTicketSourceEdit(scope.$index, scope.row)" class="a_underline">修改</a></span>
                        </template>
                        </el-table-column>
                    </el-table>
                    <el-dialog :visible.sync="dialogSourceFormVisible" title="新增/修改工单来源">
                        <el-form ref="sourceForm" :model="source_form" label-width="120px">
                        <el-form-item label="ID">
                            <el-input v-model="source_form.id" disabled />
                        </el-form-item>
                        <el-form-item label="工单来源名称">
                            <el-input v-model="source_form.name" />
                        </el-form-item>
                        <el-form-item align="right" style="padding-right: 48px">
                            <el-button type="primary" @click="dialogSourceStatus==='create'?createTicketSourceData():updateTicketSourceData()">确定</el-button>
                            <el-button @click="dialogSourceFormVisible = false">取消</el-button>
                        </el-form-item>
                        </el-form>
                    </el-dialog>
                </el-tab-pane>
            </el-tabs>
        </div>
    </div>
</template>

<script>
import { getTicketLevelList, addTicketLevel, updateTicketLevel, deleteTicketLevel } from '@/api/ticket'
import { getTicketTypeList, addTicketType, updateTicketType, deleteTicketType } from '@/api/ticket'
import { getTicketSourceList, addTicketSource, updateTicketSource, deleteTicketSource } from '@/api/ticket'
import Pagination from '@/components/Pagination'

export default {
  name: 'TicketConfig',
  components: { Pagination },
  data() {
    return {
      // 工单类型tableData
      dataTicketType: [],
      // 工单等级tableData
      dataTicketLevel: [],
      // 工单来源tableData
      dataTicketSource: [],
      total: {
        ticketType: 0
      },
      ticketTypeQuery: {
        current_page: 1,
        page_size: 20,
        search: ''
      },
      dialogTypeStatus: '',
      dialogTypeFormVisible: false,
      dialogLevelStatus: '',
      dialogLevelFormVisible: false,
      dialogSourceStatus: '',
      dialogSourceFormVisible: false,
      // 工单类型dialog Form数据
      type_form: {
        id: undefined,
        name: ''
      },
      // 工单level_form:
      level_form: {
        id: undefined,
        name: '',
        level: 0,
        time: 0
      },
      source_form: {
        id: undefined,
        name: ''
      },
    }
  },
  created() {
    this.getTicketTypeList()
    this.getTicketSourceList()
    this.getTicketLevelList()
  },
  methods: {
    resetTemp() {
      this.type_form = {
        id: undefined,
        name: ''
      }
      this.level_form = {
        id: undefined,
        name: '',
        level: 0,
        time: 0
      }
      this.source_form = {
        id: undefined,
        name: ''
      }
    },
    // 获取工单类型列表
    getTicketTypeList() {
      const This = this
      getTicketTypeList(This.ticketTypeQuery).then(response => {
        This.dataTicketType = response.data
        This.$set(This.total, 'ticketType', response.total)
      })
    },
    // 搜索工单类型
    handleTicketTypeSearch() {
      const This = this
      getTicketTypeList(This.ticketTypeQuery).then(response => {
          This.tableData = response.data
          This.$set(This.total, 'ticketType', response.total)
      })
    },
    // 触发添加工单类型dialog
    handleTicketTypeAdd() {
      this.resetTemp()
      this.dialogTypeStatus = 'create'
      this.dialogTypeFormVisible = true
    },
    // 触发修改工单类型dialog
    handleTicketTypeEdit(index, row) {
      this.resetTemp()
      this.type_form = Object.assign({}, row)
      this.dialogTypeStatus = 'update'
      this.dialogTypeFormVisible = true
    },
    // 触发删除工单类型dialog
    handleTicketTypeDelete(index, row) {
      const This = this
      This.$confirm('是否删除工单类型?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteTicketType(row.id).then(response => {
          This.$message({
            message: response.message,
            type: 'success',
            duration: 1000,
            onClose: function refresh() {
              This.getTicketTypeList()
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
    // 创建工单类型
    createTicketTypeData() {
      const This = this
      addTicketType(this.type_form).then(response => {
        this.$message({
          message: response.message,
          duration: 1000,
          type: 'success',
          onClose: function refresh() {
            This.dialogTypeFormVisible = false
            This.getTicketTypeList()
          }
        })
      })
    },
    // 更新工单类型
    updateTicketTypeData() {
      const This = this
      updateTicketType(this.type_form).then(response => {
        this.$message({
          message: response.message,
          duration: 1000,
          type: 'success',
          onClose: function refresh() {
            This.dialogTypeFormVisible = false
            This.getTicketTypeList()
          }
        })
      })
    },
    // 获取工单等级列表
    getTicketLevelList() {
      const This = this
      getTicketLevelList(This.tickeLevelQuery).then(response => {
          This.dataTicketLevel = response.data
      })
    },
    // 触发添加工单等级dialog
    handleTicketLevelAdd() {
      this.resetTemp()
      this.dialogLevelStatus = 'create'
      this.dialogLevelFormVisible = true
    },
     // 创建工单等级
    createTicketLevelData() {
      const This = this
      addTicketLevel(this.level_form).then(response => {
        this.$message({
          message: response.message,
          duration: 1000,
          type: 'success',
          onClose: function refresh() {
            This.dialogLevelFormVisible = false
            This.getTicketLevelList()
          }
        })
      })
    },
     // 触发修改工单等级dialog
    handleTicketLevelEdit(index, row) {
      this.resetTemp()
      this.level_form = Object.assign({}, row)
      this.dialogLevelStatus = 'update'
      this.dialogLevelFormVisible = true
    },
     // 更新工单等级
    updateTicketLevelData() {
      const This = this
      updateTicketLevel(this.level_form).then(response => {
        this.$message({
          message: response.message,
          duration: 1000,
          type: 'success',
          onClose: function refresh() {
            This.dialogLevelFormVisible = false
            This.getTicketLevelList()
          }
        })
      })
    },
    // 触发删除工单等级dialog
    handleTicketLevelDelete(index, row) {
      const This = this
      This.$confirm('是否删除工单等级?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteTicketLevel(row.id).then(response => {
          This.$message({
            message: response.message,
            type: 'success',
            duration: 1000,
            onClose: function refresh() {
              This.getTicketLevelList()
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
     // 获取工单来源列表
    getTicketSourceList() {
        const This = this
        getTicketSourceList().then(response => {
            This.dataTicketSource = response.data
        })
    },
    // 触发添加工单来源dialog
    handleTicketSourceAdd () {
      this.resetTemp()
      this.dialogSourceStatus = 'create'
      this.dialogSourceFormVisible = true
    },
     // 创建工单来源
    createTicketSourceData() {
      const This = this
      addTicketSource(this.source_form).then(response => {
        this.$message({
          message: response.message,
          duration: 1000,
          type: 'success',
          onClose: function refresh() {
            This.dialogLevelFormVisible = false
            This.getTicketSourceList()
          }
        })
      })
      This.dialogSourceFormVisible = false
    },
     // 触发修改工单来源dialog
    handleTicketSourceEdit(index, row) {
      this.resetTemp()
      this.source_form = Object.assign({}, row)
      this.dialogSourceStatus = 'update'
      this.dialogSourceFormVisible = true
      
    },
     // 更新工单等级
    updateTicketSourceData() {
      const This = this
      updateTicketSource(This.source_form).then(response => {
        This.$message({
          message: response.message,
          duration: 1000,
          type: 'success',
          onClose: function refresh() {
            This.getTicketSourceList()
          }
        })
      })
      This.dialogSourceFormVisible = false
    },
    // 触发删除工单来源dialog
    handleTicketSourceDelete(index, row) {
      const This = this
      This.$confirm('是否删除工单来源?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteTicketSource(row.id).then(response => {
          This.$message({
            message: response.message,
            type: 'success',
            duration: 1000,
            onClose: function refresh() {
              This.getTicketSourceList()
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
  }
}
</script>
