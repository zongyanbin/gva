<template>
  <div class="view-layout">
    <div class="main" v-if="!isNotPublish">
      <div class="header">
        <h1>试卷状态已禁止！</h1>
      </div>
      <div class="content">
          <p>您所填写的问卷未发布，暂不能填写。</p>
      </div>
    </div>

    <div class="main" v-else-if="isExpired">
        <div class="header">
          <h1>问卷已过期！</h1>
        </div>
        <div class="content">
          <p>您所填写的问卷已到截止日期，暂不能填写。</p>
        </div>
      </div>

    <div class="main" v-if="(!isExpired && isNotPublish)">
      <div class="header">
        <h1>{{tableData.paper_title}}</h1>
      </div>

      <div class="content">
        <div class="intro">
          <p>{{tableData.intro}}</p>
          <p>截止日期：{{ tableData.end_at | renderFormatDate }}</p>
        </div>
      </div>
  <!--ParperQuestion 试卷组件-->
  <ParperQuestion :tableChildData ="this.tableData">
     <ElRow>
       <div style=" text-align:center; margin: 15px;">
         <el-button type="success"
                    v-if="isAdmin"
                    @click="goBack">返回试卷列表
         </el-button>

         <el-button type="primary"
                    @click="submitQuestion">提交问卷
         </el-button>
       </div>
     </ElRow>
  </ParperQuestion>
    </div>
  </div>
</template>
<script>
import {
  findExam_paperQuestion,
} from "@/api/exam_paper";  //  此处请自行替换地址

import ParperQuestion from './components/question.vue'
export default {
  components: {
     ParperQuestion
  },
  props: [],
  data() {
    return {
      exam_paper:{
        topic:[]
      },
      tableData:'',
      finished: false,
      time_end:'',
      formData: {
        field101: undefined,
        field102: undefined,
      },
      rules: {
        field101: [{
          required: true,
          message: '请输入单行文本',
          trigger: 'blur'
        }],
        field102: [{
          required: true,
          message: '请输入多行文本',
          trigger: 'blur'
        }],
      },
    }
  },
  filters: {  // 前端时间格式2020-02-11T12:24:18.000+0000转化成正常格式
    renderFormatDate: function (time) {
      var dateee = new Date(time).toJSON();
      return new Date(+new Date(dateee) + 8 * 3600 * 1000).toISOString().replace(/T/g, ' ').replace(/\.[\d]{3}Z/, '')
    },
  },
  computed: {
    isAdmin () {
      return 1
      // return this.$store.getters.isAdmin
    },

    // true过期  false未过期
    isExpired () {
      const timestamp = parseInt(new Date().getTime()/1000);    // 当前时间戳
      console.log("isExpired:"+this.getTimestamp(this.time_end) , timestamp)
      return Number(this.getTimestamp(this.time_end) < timestamp)
    },

    // 判断是否 公开1 禁止0
    isNotPublish () {
      return this.tableData.paper_status
    },
  },
  watch: {},
  created() {
    this.tableData=this.getTableData()
  },
  mounted() {},
  methods: {
    submitForm() {
      this.$refs['elForm'].validate(valid => {
        if (!valid) return
        // TODO 提交表单
      })
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
    // GET QUESTION ALL
    async getTableData() {

      var  exam_paper_id = this.$route.query.exam_paper_id
      const  table = await findExam_paperQuestion({ exam_paper_id: exam_paper_id })
      if (table.code == 0) {
        console.log(table.data.list)
        this.tableData = table.data.list
        this.time_end = table.data.list.end_at
      }
    },

    getTimestamp(time) { //把时间日期转成时间戳
      return (new Date(time)).getTime() / 1000
    },

    // 前端时间格式2020-02-11T12:24:18.000+0000转化成正常格式
    renderTime(date) {
      var dateee = new Date(date).toJSON();
      return new Date(+new Date(dateee) + 8 * 3600 * 1000).toISOString().replace(/T/g, ' ').replace(/\.[\d]{3}Z/, '')
    },
    goBack () {
      this.$router.push('/layout/exam_paperP/exam_paperTable')
    },
    // 提交问题
    submitQuestion(){
      this.tableData.question.forEach((question,index)=>{
          console.log(question,index)
      })
      alert(1)
    }
  }
}

</script>
<style scoped>

.view-layout {
  background-color: rgb(237, 240, 248);
  min-height: 100%;
  max-height: 100%;
  height: 100%;
  width: 100%;
  padding: 20px 10px;
  overflow-y: scroll;
}

.view-layout .main {
  max-width: 800px;
  width: 100%;
  height: auto;
  margin: 0 auto;
  padding-bottom: 30px;
  background-color: #fff;
  -webkit-box-shadow: 0 2px 5px 1px rgba(124, 124, 124, .2);
  box-shadow: 0 2px 5px 1px rgba(124, 124, 124, .2);
}

.view-layout .header {
  padding: 30px 10px;
  height: auto;
  min-height: 10px;
  border-bottom: 2px dotted #eee;
}

.view-layout .header h1 {
  /*width: 500px;*/
  width: 100%;
  margin: 0 auto;
  margin-top: 5px;
  margin-bottom: -10px;
  text-align: center;
}

.view-layout .content {
  padding: 20px;
}

.view-layout .intro {
  font-size: 14px;
  text-indent: 2em;
}

.view-layout .footer {
  margin-top: 10px;
  padding-top: 10px;
  text-align: center;
  border-top: 1px dotted #eee;
}

.code-row-bg button {
  margin: 0 10px;
}

.user-info {
  width: 100%;
  padding: 30px 30px 0 30px;
}
</style>
<style>
</style>
