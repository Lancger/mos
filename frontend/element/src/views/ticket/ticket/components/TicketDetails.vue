<template>
  <div class="createPost-container">
    <el-form ref="ticketForm" :model="ticketForm" :rules="rules" class="form-container">
      <div class="createPost-main-container">
        <sticky :z-index="10" :class-name="'sub-navbar draft'">
          <span>工单编号: {{ ticketForm.number }}</span>
          <el-button @click="redirectTicketList()" icon="el-icon-back" class="pan-btn tiffany-btn">返回</el-button>
          <el-button v-if="isCreate" icon="el-icon-document" @click="submitForm()" class="pan-btn blue-btn">创建</el-button>
          <el-button v-else icon="el-icon-document" @click="updateForm()" class="pan-btn blue-btn">更新</el-button>
        </sticky>
        <el-card class="box-card"> 
          <el-row>
            <el-col :span="24">
              <el-form-item style="margin-bottom: 40px;">
                  <MDinput v-model="ticketForm.title" :maxlength="100" name="ticket_title" required>
                      工单标题
                  </MDinput>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="8">
                <!-- <el-form-item label-width="80px" label="创建人:" class="postInfo-container-item">
                    <el-cascader
                      v-model="select_author"
                      :options="user_options"
                      placeholder="选择用户"
                      change-on-select
                      filterable
                      class="form-item-width"
                      @change="author_change"/>
                </el-form-item> -->
              <el-form-item label-width="100px" label="创建人:" class="postInfo-container-item">
                  <span>{{ ticketForm.author }}</span>
              </el-form-item>
            </el-col>
            <el-col :span="8">
                <el-form-item label-width="100px" label="当前处理人:" class="postInfo-container-item">
                    <span>{{ ticketForm.dealer }}</span>
                </el-form-item>
            </el-col>
              <el-col :span="8">
                <el-form-item label-width="100px" label="上一处理人:" class="postInfo-container-item">
                    <span>{{ ticketForm.pre_dealer }}</span>
                </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="8">
            <el-form-item label-width="100px" label="工单状态:" class="postInfo-container-item">
                <el-select v-model="ticketForm.stage"  placeholder="选择工单状态" class="postInfo-container-item" :disabled="disableStage">
                <el-option
                  v-for="item in stageOptions"
                  :key="item.label"
                  :label="item.label"
                  :value="item.value" />
                </el-select>
            </el-form-item>
            </el-col>
            <el-col :span="8">
            <el-form-item label-width="100px" label="工单等级:" class="postInfo-container-item">
                <el-select v-model="ticketForm.level" placeholder="选择工单等级" class="postInfo-container-item">
                <el-option
                  v-for="item in levelOptions"
                  :key="item.name"
                  :label="item.name"
                  :value="item.level" />
                </el-select>
            </el-form-item>
            </el-col>
            <el-col :span="8">
            <el-form-item label-width="100px" label="工单类型:" class="postInfo-container-item">
                <el-select v-model="ticketForm.type" placeholder="选择工单类型" class="postInfo-container-item">
                <el-option
                  v-for="item in typeOptions"
                  :key="item.name"
                  :label="item.name"
                  :value="item.name" />
                </el-select>
            </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="8">
                <el-form-item label-width="100px" label="所属项目:" class="postInfo-container-item">
                <el-select v-model="ticketForm.project" placeholder="选择项目" class="postInfo-container-item">
                <el-option
                  v-for="item in projectOptions"
                  :key="item.name"
                  :label="item.name"
                  :value="item.name" />
                </el-select>
            </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label-width="100px" label="改进方:" class="postInfo-container-item">
                <el-select v-model="ticketForm.improve" placeholder="选择改进方" class="postInfo-container-item">
                  <el-option
                    v-for="item in improveOptions"
                    :key="item.label"
                    :label="item.label"
                    :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label-width="100px" label="工单来源:" class="postInfo-container-item">
                  <el-select v-model="ticketForm.source" placeholder="选择工单来源" class="postInfo-container-item" >
                  <el-option
                      v-for="item in sourceOptions"
                      :key="item.name"
                      :label="item.name"
                      :value="item.name" />
                  </el-select>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-form-item label-width="100px" prop="content" label="工单内容:" style="margin-bottom: 30px; margin-right:25px">
              <Tinymce ref="ticket_content" v-model="ticketForm.content" />
            </el-form-item>
          </el-row>
          <el-row>
            <el-form-item label-width="100px" prop="content" label="解决方案:" style="margin-bottom: 30px; margin-right:25px">
              <Tinymce ref="ticket_soultion" v-model="ticketForm.soultion" />
            </el-form-item>
          </el-row>
        </el-card>
      </div>
    </el-form>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import global_ from '../../../../utils/global'
import { getTicketLevelList, getTicketTypeMsg, getTicketSourceList, addTicket, updateTicket, fetchTicket } from '@/api/ticket'
import { getGroupUserOptions } from '@/api/group'
import { getProjectMsg } from '@/api/project'
import Tinymce from '@/components/Tinymce'
import MDinput from '@/components/MDinput'
import Sticky from '@/components/Sticky' // 粘性header组件
// import svgIcons from '@/components/icons/svg-icons'
// import elementIcons from '@/components/icons/element-icons'

// import { validURL } from '@/utils/validate'
// import Warning from './Warning'

const defaultForm = {
  id: '',
  number: '',
  title: '', // 工单标题
  level: '',
  type: '',
  stage: '',
  source: '',
  improve: '',
  content: '',
  solution: '',
  start_time: '',
  end_time: '',
  dealer: '',
  pre_dealer: '',
  author: '',
  project: ''
}

export default {
  name: 'TicketDetails',
  components: { Tinymce, MDinput, Sticky },
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
          message: rule.field + '为必传项',
          type: 'error'
        })
        callback(new Error(rule.field + '为必传项'))
      } else {
        callback()
      }
    }
    return {
      isCreate: false,
      id: '',
      select_author: [],
      myself_ticket: true,
      ticket_unfinish: true,
      ticketForm: Object.assign({}, defaultForm),
      loading: false,
      levelOptions: [],
      typeOptions: [],
      improveOptions: global_.improve_options,
      sourceOptions: [],
      projectOptions: [],
      user_options: [],
      stageOptions: global_.stage_options,
      disableStage: false,
      rules: {
        title: [{ validator: validateRequire }],
        content: [{ validator: validateRequire }]
      },
      tempRoute: {}
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'nick_name'
    ])
  },
  created() {
    if (this.isEdit) {
      this.fetchTicket()
      this.disableStage = true
      this.isCreate = false
    } else {
      this.isCreate = true
      this.createData()
    }
    this.initOptions()
    this.tempRoute = Object.assign({}, this.$route)
  },
  methods: {
    initOptions() {
      this.getSourceOptions()
      this.getLevelOptions()
      this.getTypeOptions()
      this.getProjectOptions()
      // this.getGroupUserOptions()
    },
    redirectTicketList() {
      this.$router.push({
        name: 'TicketList'
      })
    },
    createData() {
      this.ticketForm = Object.assign({}, defaultForm)
      this.$set(this.ticketForm, 'author', this.nick_name)
      this.$set(this.ticketForm, 'dealer', this.nick_name)
      this.$set(this.ticketForm, 'pre_dealer', this.nick_name)
      this.$set(this.ticketForm, 'stage', '新建工单')
      this.disableStage = true
    },
    author_change(sval) {
      const This = this
      if (sval.length < 2) {
        This.$message('请选择用户')
      } else {
        This.$set(This.ticketForm, 'author', sval[1])
      }
    },
    getProjectOptions() {
      const This = this
      getProjectMsg().then(response => {
        This.projectOptions = response.data
      })
    },
    getGroupUserOptions() {
      const This = this
      getGroupUserOptions().then(response => {
        This.user_options = response.data
      })
    },
    // 获取工单等级
    getLevelOptions() {
      const This = this
      getTicketLevelList().then(response => {
        This.levelOptions = response.data
      })
    },
    getSourceOptions() {
      const This = this
      getTicketSourceList().then(response => {
        This.sourceOptions = response.data
      })
    },
    // 获取工单类型
    getTypeOptions() {
      const This = this
      getTicketTypeMsg().then(response => {
        This.typeOptions = response.data
      })
    },
    fetchTicket() {
      const id = this.$route.query.id
      fetchTicket(id).then(response => {
        this.ticketForm = response.data
      }).catch(err => {
        console.log(err)
      })
    },
    checkForm() {
      const This = this
      if (This.ticketForm.title === '') {
        This.$message({
          message: '请填写必要的标题',
          type: 'warning'
        })
        return false
      }
      if (This.ticketForm.level === '') {
        This.$message({
          message: '请选择工单等级',
          type: 'warning'
        })
        return false
      }
      if (This.ticketForm.content === '') {
        This.$message({
          message: '请填写必要的内容',
          type: 'warning'
        })
        return false
      }
      if (This.ticketForm.type === '') {
        This.$message({
          message: '请填写工单类型',
          type: 'warning'
        })
        return false
      }
      if (This.ticketForm.project === '') {
        This.$message({
          message: '请填写所属项目',
          type: 'warning'
        })
        return false
      }
      return true
    },
    submitForm() {
      const This = this
      if (This.checkForm() === false) {
        return
      }
      addTicket(This.ticketForm).then(response => {
        This.$message({
          message: response.message,
          type: 'success',
          duration: 2000,
          onClose: function refresh() {
            This.redirectTicketList()
          }
        })
      })
    },
    updateForm() {
      const This = this
      if (This.checkForm() === false) {
        return
      }
      updateTicket(This.ticketForm).then(response => {
        This.$message({
          message: response.message,
          type: 'success',
          duration: 2000,
          onClose: function refresh() {
            This.redirectTicketList()
          }
        })
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import "~@/styles/mixin.scss";

.createPost-container  {
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
