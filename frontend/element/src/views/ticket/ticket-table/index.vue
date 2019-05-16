<template>
  <div class="app-container">
    <div class="filter-container">
      <!-- 表头搜索 -->
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-input v-model="listQuery.search" placeholder="输入内容搜索" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="listQuery.ticket_number" placeholder="输入工单编号" class="form-select-width" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="listQuery.author" placeholder="输入建单人搜索" class="form-select-width" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="listQuery.dealer" placeholder="输入处理人搜索" class="form-select-width" />
        </el-form-item>
        <el-form-item>
          <el-select v-model="listQuery.stage" multiple placeholder="选择工单状态" class="form-select-width">
            <el-option
              v-for="(stage_item, index) in stageOptions"
              :key="stage_item.value+index"
              :label="stage_item.label"
              :value="stage_item.label" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select v-model="listQuery.project" multiple placeholder="选择项目" size="mini">
            <el-option
              v-for="(project_item, index) in projectOptions"
              :key="project_item.name+index"
              :label="project_item.name"
              :value="project_item.name" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select v-model="listQuery.type" multiple placeholder="选择工单类型" >
            <el-option
              v-for="item in typeOptions"
              :key="item.name"
              :label="item.name"
              :value="item.name" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-date-picker
            v-model="time_picker.start_time"
            type="datetime"
            size="small"
            placeholder="开始时间"
            @change="date_time_change('start')"/>
          <el-date-picker
            v-model="time_picker.end_time"
            type="datetime"
            size="small"
            placeholder="结束时间"
            @change="date_time_change('end')"/>
        </el-form-item>
        <el-form-item>
          <el-button type="info" icon="el-icon-search" class="pan-btn gray-btn" @click="handleSearch()">搜索</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="danger" icon="el-icon-refresh" class="pan-btn pink-btn" @click="handleTicketRefresh()">重置</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="el-icon-plus" class="pan-btn tiffany-btn" @click="redirectTicketAdd()">新增工单</el-button>
        </el-form-item>
        <el-form-item>
          <el-switch
            v-model="myself_ticket"
            active-text="我的工单"
            active-color="#13ce66" />
        </el-form-item>
        <el-form-item>
          <el-switch
            v-model="ticket_unfinish"
            active-text="未完成工单"
            active-color="#13ce66" />
        </el-form-item>
      </el-form>
      <el-table
        ref="multipleTable"
        :data="tableData"
        :v-loading="true"
        border
        stripe
        style="width: 100%; font-size: 14px"
        @sort-change="sortChange"
        @selection-change="handleMultiSelectionChange">
        <el-table-column
          type="selection"
          width="55" />
        <el-table-column
          :sort-orders="sort_order_arr"
          label="工单编号"
          prop="number"
          width="160"
          sortable="custom" />
        <el-table-column
          label="工单标题"
          prop="title"
          width="160">
          <template slot-scope="scope">
            <el-popover trigger="hover" placement="top" width="1280">
              <h3 style="color: #3A71A8;">历史记录 / {{ scope.row.number }} / ({{ scope.row.title }}) </h3>
              <hr size=3>
              <h4>创建时间: {{ scope.row.create_time }}</h4>
              <h4>创建人: {{ scope.row.author }}</h4>
              <h4>现处理人: {{ scope.row.dealer }}</h4>
              <h4>工单状态: {{ scope.row.stage }}</h4>
              <hr size=3>
              <el-table
                ref="multipleTable"
                :data="scope.row.history"
                border
                stripe
                style="width: 100%; font-size: 14px">
                <el-table-column
                  label="动作"
                  prop="action"
                  width=180 />
                <el-table-column
                  label="内容"
                  prop="content" />
                <el-table-column
                  label="操作人"
                  prop="from"
                  width=180 />
                <!-- <el-table-column
                  label="目标"
                  prop="to"
                  width=180 /> -->
                <el-table-column
                  label="时间"
                  prop="create_time"
                  width=180 />  
              </el-table>
              <div slot="reference" class="name-wrapper">
                <span class="el-tag el-tag-blue">{{ scope.row.title }}</span>
              </div>
            </el-popover>
          </template>
        </el-table-column>  
        <el-table-column
          label="所属项目"
          prop="project"
          align="center" />
        <el-table-column
          label="工单类型"
          prop="type"
          align="center" />
        <el-table-column
          label="处理阶段"
          prop="stage"
          align="center" />
        <el-table-column
          :sort-orders="sort_order_arr"
          label="创建时间"
          prop="create_time"
          align="center" />
        <el-table-column
          label="创建人"
          prop="author"
          align="center">
        </el-table-column>
        <el-table-column
          label="当前处理人"
          prop="dealer"
          align="center">
        </el-table-column>
        <el-table-column
          align="center"
          width="360"
          label="操作">
          <template slot-scope="scope">
            <span class="el-tag el-tag--success"><a @click="redirectTicketMsg(scope.row.number)" class="a_underline a_green">详情</a></span>
            <span class="el-tag el-tag--danger"><a @click="handleTicketDelete(scope.$index, scope.row)" class="a_underline a_red">删除</a></span>
            <span class="el-tag"><a @click="redirectTicketEdit(scope.row.number)" class="a_underline a_blue">修改</a></span>
            <el-dropdown>
              <el-button class="pan-btn light-blue-btn">
                操作<i class="el-icon-arrow-down el-icon--right" />
              </el-button>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item>
                  <div>
                    <span v-if="checkPermStage(['send_ticket'], scope.row.stage, ['新建工单'])" class="el-tag el-tag--success">
                      <a @click="handleTicketSend(scope.$index, scope.row)" class="a_underline a_green">派发工单</a>
                    </span>
                  </div>
                  <div>
                    <span v-if="checkPermNotStage(['change_ticket'], scope.row.stage, ['结单', '处理完毕'])" class="el-tag el-tag--success">
                      <a @click="handleTicketChange(scope.$index, scope.row)" class="a_underline a_green">转发工单</a>
                    </span>
                  </div>
                  <div>
                    <span v-if="checkPermNotStage(['deal_ticket'], scope.row.stage, ['结单', '处理完毕', '处理中'])" class="el-tag el-tag--info">
                      <a @click="handleTicketDeal(scope.row.id)" class="a_underline a_blue">开始处理</a>
                    </span>
                  </div>
                  <div>
                    <span v-if="checkPermNotStage(['deal_ticket'], scope.row.stage, ['结单', '处理完毕'])" class="el-tag el-tag--warning">
                      <a @click="handleTicketCtl(scope.row, 'mark')" class="a_underline a_blue">添加记录</a>
                    </span>
                  </div>
                  <div>
                    <span v-if="checkPermStage(['deal_ticket'], scope.row.stage, ['处理中'])" class="el-tag el-tag--warning">
                      <a @click="handleTicketCtl(scope.row, 'finish')" class="a_underline a_blue">处理完毕</a>
                    </span>
                  </div>
                  <div>
                    <span v-if="checkPermStage(['close_ticket'], scope.row.stage, ['处理完毕'])" class="el-tag el-tag--danger">
                      <a @click="handleTicketCtl(scope.row, 'close')" class="a_underline a_blue">结单</a>
                    </span>
                  </div>
                  <!-- <el-button
                    size="mini"
                    type="danger"
                    @click="handleTicketSend(scope.$index, scope.row)">派发工单
                  </el-button> -->
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </template>
        </el-table-column>
      </el-table>
      <pagination v-show="total>0" :total="total" :current_page.sync="listQuery.current_page" :page_size.sync="listQuery.page_size" align="right" @pagination="getList" />
      <!-- 转单及派单dialog -->
      <el-dialog :visible.sync="dialogTicketChangeVisible" :title="ticketChangeTitle" :close-on-click-modal="false">
        <el-form ref="ticketChangeForm" :model="ticketChangeForm" label-width="120px">
          <el-form-item label="工单编号">
            <el-input v-model="ticketChangeForm.number" class="form-item-width" disabled />
          </el-form-item>
          <el-form-item label="工单标题">
            <el-input v-model="ticketChangeForm.title" disabled />
          </el-form-item>
          <el-form-item label="处理人:" class="postInfo-container-item">
            <el-cascader
              v-model="selected_user"
              :options="user_options"
              placeholder="选择用户"
              change-on-select
              filterable
              class="form-item-width"
              @change="user_change"/>
          </el-form-item>
          <el-form-item label="交接内容">
            <el-input v-model="ticketChangeForm.content" :rows="6" type="textarea" class="form-item-width" />
          </el-form-item>
          <el-form-item align="right" style="padding-right: 48px">
            <el-button type="primary" @click="dialogSendStatus==='send'?ctlTicketSend('send'):ctlTicketSend('change')">确定</el-button>
            <el-button @click="dialogTicketChangeVisible = false">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
      <!-- 工单操作dialog -->
      <el-dialog :visible.sync="dialogTicketCtlVisible" :title="ticketCtlTitle" :close-on-click-modal="false">
        <el-form ref="ticketCtlForm" :model="ticketCtlForm" label-width="120px">
          <el-form-item label="工单编号">
            <el-input v-model="ticketCtlForm.number" class="form-item-width" disabled />
          </el-form-item>
          <el-form-item label="工单标题">
            <el-input v-model="ticketCtlForm.title" disabled />
          </el-form-item>
          <el-form-item label="内容">
            <el-input v-model="ticketCtlForm.content" :rows="6" type="textarea" class="form-item-width" />
          </el-form-item>
          <el-form-item align="right" style="padding-right: 48px">
            <el-button type="primary" @click="ctlTicketManger()">确定</el-button>
            <el-button @click="dialogTicketCtlVisible = false">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { getTicketList, getTicketTypeMsg, deleteTicket, sendTicket, ctlTicket } from '@/api/ticket'
import { getGroupUserOptions } from '@/api/group'
import { getProjectMsg } from '@/api/project'
import { mapGetters } from 'vuex'
import global_ from '../../../utils/global'
import { checkPermStage, checkTicketWorkPerm, checkPermNotStage } from '@/utils/permission'
import Pagination from '@/components/Pagination'

export default {
  name: 'TicketList',
  components: { Pagination },
  data() {
    return {
      selected_user: [],
      queryUser: this.$store.getters.nick_name,
      ticketChangeTitle: '派发工单',
      time_picker: {
        start_time: undefined,
        end_time: undefined
      },
      select_author: [],
      levelOptions: [],
      typeOptions: [],
      sourceOptions: [],
      projectOptions: [],
      user_options: [],
      stageOptions: global_.stage_options,
      // 派发工单dialog 默认派发
      dialogSendStatus: 'send',
      // 派发工单dialog是否显示
      dialogTicketChangeVisible: false,
      // 工单操作dialog是否显示
      dialogTicketCtlVisible: false,
      ticketChangeForm: {
        id: '',
        number: '',
        title: '',
        content: '',
        to: '',
        from: ''
      },
      ticketCtlForm: {
        id: '',
        number: '',
        title: '',
        from: '',
        to: '',
        content: '',
        ctl: ''
      },
      ticketCtlTitle: '',
      // queryUser: {
      //   nick_name: this.$store.getters.nick_name
      // },
      sort_order_arr: global_.sort_order_arr,
      tableData: [],
      total: 0,
      myself_ticket: true,
      ticket_unfinish: true,
      multipleSelectionList: [],

      listQuery: {
        current_page: 1,
        page_size: 20,
        search: '',
        ticket_number: '',
        dealer: '',
        author: '',
        project: '',
        type: '',
        stage: '',
        unfinish: 1,
        start_time: 0,
        end_time: 0
      }
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'nick_name',
      'roles'
    ])
  },
  watch: {
    // 监听我的工单switch切换
    myself_ticket(newval, oldval) {
      const This = this
      if (newval === true) {
        This.$set(This.listQuery, 'dealer', This.$store.getters.nick_name)
      } else {
        This.$set(This.listQuery, 'dealer', '')
      }
      getTicketList(This.listQuery).then(response => {
        This.tableData = response.data
        This.total = response.total
      })
    },
    // 监听未完成工单switch切换
    ticket_unfinish(newval, oldval) {
      const This = this
      if (This.myself_ticket === false) {
        console.log("ok")
        This.$set(This.listQuery, 'dealer', '')
      } else {
        This.$set(This.listQuery, 'dealer', This.queryUser)
      }
      if (newval === true) {
        This.$set(This.listQuery, 'unfinish', 1)
      } else {
        This.$set(This.listQuery, 'unfinish', 0)
      }
      This.getList()
    }
  },
  created() {
    this.getTicketTypeMsg()
    this.getProjectMsg()
    this.getList()
    this.getGroupUserOptions()
  },
  methods: {
    checkTicketWorkPerm,
    checkPermStage,
    checkPermNotStage,
    redirectTicketEdit(ticket_number) {
      this.$router.push({
        name: 'TicketEdit',
        query: {
          id: ticket_number
        }
      })
    },
    redirectTicketMsg(ticket_number) {
      this.$router.push({
        name: 'TicketMsg',
        query: {
          id: ticket_number
        }
      })
    },
    redirectTicketAdd() {
      this.$router.push({
        name: 'TicketAdd'
      })
    },
    date_time_change(val) {
      const This = this
      if (val === 'start') {
        This.$set(This.listQuery, 'start_time', Date.parse(new Date(This.time_picker.start_time)))
      } else if (val === 'end') {
        This.$set(This.listQuery, 'end_time', Date.parse(new Date(This.time_picker.end_time)))
      } else {
        console.log('selected')
      }
    },
    // 多选选中工单号ID数组
    handleMultiSelectionChange(val) {
      const This = this
      This.multipleSelectionList = val
    },
    // table自定义排序
    sortChange(column, prop, order) {
      const This = this
      if (column.prop === null) {
        column.prop = ''
      }
      if (column.order === null) {
        column.order = ''
      }
      This.$set(This.listQuery, 'order_props', column.prop)
      This.$set(This.listQuery, 'order_type', column.order)
      This.getList()
    },
    // 清空数据
    resetTemp() {
      this.time_picker.start_time = undefined
      this.time_picker.end_time = undefined
      this.ticketForm = {
        id: undefined,
        number: '',
        title: '',
        project: '',
        level: '',
        type: '',
        stage: '',
        sla_time: '',
        create_time: '',
        author: '',
        dealer: '',
        unfinish: 1,
        start_time: 0,
        end_time: 0
      }
    },
    // 刷新工单列表，清空查询参数
    handleTicketRefresh() {
      const This = this
      This.ticket_unfinish = true
      This.myself_ticket = true
      This.time_picker.start_time = undefined
      This.time_picker.end_time = undefined
      This.listQuery = {
        current_page: 1,
        page_size: 20,
        search: '',
        ticket_number: '',
        dealer: '',
        author: '',
        project: [],
        type: [],
        stage: []
      }
      if (This.myself_ticket === true) {
        This.$set(This.listQuery, 'dealer', This.queryUser)
      } else {
        This.$set(This.listQuery, 'dealer', '')
      }
      This.getList()
    },
    // 搜索工单
    handleSearch() {
      const This = this
      console.log(This.listQuery)
      // console.log(This.listQuery)
      This.getList()
    },
    // 获取类型options
    getTicketTypeMsg() {
      getTicketTypeMsg().then(response => {
        this.typeOptions = response.data
      })
    },
    // 获取项目options
    getProjectMsg() {
      getProjectMsg().then(response => {
        this.projectOptions = response.data
      })
    },
    // 获取用户cas框信息
    getGroupUserOptions() {
      const This = this
      getGroupUserOptions().then(response => {
        This.user_options = response.data
      })
    },
    formatTicketRate(timestamp) {
      if (timestamp === 0) {
        var ts = (new Date()).getTime() / 1000
      } else {
        ts = timestamp
      }
      return ts
    },
    setInvisible() {
      this.dialogTicketHistoryVisible = false
      this.dialogMultiCtlVisible = false
    },
    // 派单
    handleTicketSend(index, row) {
      const This = this
      This.ticketChangeTitle = '派发工单'
      This.ticketChangeForm = {
        id: row.id,
        number: row.number,
        title: row.title,
        content: '',
        to: '',
        from: This.queryUser,
      }
      This.dialogSendStatus = 'send'
      This.dialogTicketChangeVisible = true
    },
    // 转单
    handleTicketChange(index, row) {
      const This = this
      This.ticketChangeTitle = '转发工单'
      This.ticketChangeForm = {
        id: row.id,
        number: row.number,
        title: row.title,
        content: '',
        to: '',
        from: This.queryUser,
      }
      This.dialogSendStatus = 'change'
      This.dialogTicketChangeVisible = true
    },
    handleTicketDeal(id) {
      const This = this
      ctlTicket(id, 'deal', This.queryUser).then(response => {
        This.$message({
          message: response.message,
          type: 'success',
          duration: 2000,
          onClose: function refresh() {
            This.getList()
          }
        })
      })
    },
    // 触发删除工单来源dialog
    handleTicketDelete(index, row) {
      const This = this
      This.$confirm('是否删除工单?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteTicket(row.id).then(response => {
          This.$message({
            message: response.message,
            type: 'success',
            duration: 1000,
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
    // 工单操作dialog触发
    handleTicketCtl(row, ctl) {
      const This = this
      if (ctl === 'mark') {
        This.ticketCtlTitle = "添加处理记录"
      } else if (ctl === 'finish') {
        This.ticketCtlTitle = "处理完毕"
      } else if (ctl === 'close') {
        This.ticketCtlTitle = "结单"
      } else {
        This.$message({
          message: response.message,
          type: 'success',
          duration: 2000,
          onClose: function refresh() {
            This.getList()
          }
        })
        return
      }
      This.ticketCtlForm = {
        id: row.id,
        number: row.number,
        title: row.title,
        from: row.from,
        to: row.to,
        content: '',
        ctl: ctl
      }
      This.dialogTicketCtlVisible = true
    },
    // 获取工单列表
    getList() {
      const This = this
      if (This.myself_ticket === true) {
        This.$set(This.listQuery, 'dealer', This.queryUser)
      }
      // } else {
      //   This.$set(This.listQuery, 'dealer', '')
      // }
      getTicketList(This.listQuery).then(response => {
        This.tableData = response.data
        This.total = response.total
      })
    },
    user_change(sval) {
      const This = this
      if (sval.length < 2) {
        This.$message('请选择用户')
      } else {
        This.$set(This.ticketChangeForm, 'to', sval[1])
      }
    },
    // 执行派单
    ctlTicketSend() {
      const This = this
      sendTicket(This.ticketChangeForm, This.dialogSendStatus).then(response => {
        This.$message({
          message: response.message,
          type: 'success',
          duration: 2000,
          onClose: function refresh() {
            This.getList()
          }
        })
      })
      This.dialogTicketChangeVisible = false
    },
    ctlTicketManger() {
      const This = this
      var id = This.ticketCtlForm.id
      var ctl = This.ticketCtlForm.ctl
      var content = This.ticketCtlForm.content
      ctlTicket(id, ctl, This.queryUser, content).then(response => {
        This.$message({
          message: response.message,
          type: 'success',
          duration: 2000,
          onClose: function refresh() {
            This.getList()
          }
        })
      })
      This.dialogTicketCtlVisible = false
    }
  }
}

</script>
<style lang="scss" scoped>
@import "~@/styles/mixin.scss";
  .form-item-width {
    width: 100%
  }
  .form-select-width {
    width: 160px
  }
  .el-table .danger-row {
    background: indianred;
  }
  .el-table .warning-row {
    background: oldlace;
  }
  .el-table .success-row {
    background: #f0f9eb;
  }
  .pre_message {
    display: block;
    overflow: auto;
    background: #f4f4f4;
    padding: 5px 0px;
    border: 1px solid #eee;
    whitewhite-space:pre-wrap; /* css-3 */
    whitewhite-space:-moz-pre-wrap; /* Mozilla, since 1999 */
    whitewhite-space:-pre-wrap; /* Opera 4-6 */
    whitewhite-space:-o-pre-wrap; /* Opera 7 */
    word-wrap:break-word; /* Internet Explorer 5.5+ */
    white-space: pre-wrap; /* Firefox */
  }
  a, a:link {
    text-decoration: blink;
  }
  .a_red {
    color: red
  }
  .a_blue {
    color: #3A71A8;
  }
  .a_green {
    color:olivedrab;
  }
  .tag-blue {
    color: #3A71A8;
  }
  .popover {
    background: #909399;
  }
  
</style>

