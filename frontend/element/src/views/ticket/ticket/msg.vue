<template>
  <div class="createPost-container">
    <el-form ref="ticketForm" :model="ticketForm" class="form-container">
      <div class="createPost-main-container">
        <sticky :z-index="10" :class-name="'sub-navbar draft'">
          <span style="font-weight:bold">工单编号: {{ ticketForm.number }}</span>
          <el-button @click="redirectTicketList()" icon="el-icon-back" class="pan-btn tiffany-btn">返回</el-button> 
        </sticky>
        <el-card class="box-card custom_card">
          <el-row>
            <el-col :span="12">
              <el-form-item label-width="100px" label="工单标题:" class="postInfo-container-item">
                <span style="font-weight:bold">{{ ticketForm.title }}</span>
              </el-form-item>
            </el-col>
          </el-row>
        </el-card>
        <el-card class="box-card custom_card">
          <el-row>
            <el-col :span="8">
              <el-form-item label-width="100px" label="创建人:" class="postInfo-container-item">
                  <span style="font-weight:bold">{{ ticketForm.author }}</span>
              </el-form-item>
            </el-col>
            <el-col :span="8">
                <el-form-item label-width="100px" label="现处理人:" class="postInfo-container-item">
                    <span style="font-weight:bold">{{ ticketForm.dealer }}</span>
                </el-form-item>
            </el-col>
              <el-col :span="8">
                <el-form-item label-width="100px" label="前处理人:" class="postInfo-container-item">
                  <span style="border:2px solid #ffff">{{ ticketForm.pre_dealer }}</span>
                </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="8">
              <el-form-item label-width="100px" label="工单状态:" class="postInfo-container-item">
                <span style="font-weight:bold">{{ ticketForm.stage }}</span>
              </el-form-item>
            </el-col>
            <el-col :span="8">
            <el-form-item label-width="100px" label="工单等级:" class="postInfo-container-item">
              <span style="font-weight:bold">{{ ticketForm.level }}</span>
            </el-form-item>
            </el-col>
            <el-col :span="8">
            <el-form-item label-width="100px" label="工单类型:" class="postInfo-container-item">
              <span style="font-weight:bold">{{ ticketForm.type }}</span>
            </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="8">
              <el-form-item label-width="100px" label="所属项目:" class="postInfo-container-item">
                <span style="font-weight:bold">{{ ticketForm.project }}</span>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label-width="100px" label="改进方:" class="postInfo-container-item">
                <span style="font-weight:bold">{{ ticketForm.improve }}</span>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label-width="100px" label="工单来源:" class="postInfo-container-item">
                <span style="font-weight:bold">{{ ticketForm.source }}</span>
              </el-form-item>
            </el-col>
          </el-row>
        </el-card>
        <el-card class="box-card custom_card" >  
          <el-row>
            <el-form-item label-width="100px" prop="content" label="工单内容:" style="margin-bottom: 30px; margin-right:25px">
              <div v-html="ticketForm.content" />
            </el-form-item>
          </el-row>
        </el-card>
        <el-card class="box-card custom_card">  
          <el-row>
            <el-form-item label-width="100px" prop="content" label="解决方案:" style="margin-bottom: 30px; margin-right:25px">
              <div v-html="ticketForm.soultion" />
            </el-form-item>
          </el-row>
        </el-card>
      </div>
    </el-form>
  </div>
</template>

<script>
import { fetchTicket } from '@/api/ticket'
import Sticky from '@/components/Sticky' // 粘性header组件
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
  name: 'TicketMsg',
  components: { Sticky },
  data() {
    return {
      ticketForm: Object.assign({}, defaultForm),
      tempRoute: {}
    }
  },
  created() {
    this.fetchTicket()
    this.tempRoute = Object.assign({}, this.$route)
  },
  methods: {
    redirectTicketList() {
      this.$router.push({
        name: 'TicketList'
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
    redirectTicketList() {
      this.$router.push({
        name: 'TicketList'
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
</style>
