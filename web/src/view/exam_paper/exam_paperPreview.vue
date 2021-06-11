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
    <div class="main" v-if="(!isExpired && isNotPublish) || isAdmin">

      aaaa
      <div class="header">
        aaa
        <h1>{{tableData.paper_title}}</h1>
      </div>

    </div>
<!--    <ParperQuestion :tableChildData ="this.tableData.question"></ParperQuestion>>-->
  </div>
</template>
<script>
import {
  findExam_paperQuestion,
} from "@/api/exam_paper";  //  此处请自行替换地址

// import ParperQuestion from './components/question.vue'
export default {
  components: {
    // ParperQuestion
  },
  props: [],
  data() {
    return {
      tableData:'',
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
  computed: {
    isAdmin(){
        return ''
    },
    // 判断是否过期 过期1 未过期0
    isExpired() {
      alert(this.time_end)
    },
    // 判断是否公开
    isNotPublish() {
      // alert(this.tableData.paper_status)
      // return this.tableData.paper_status
      return 1
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
  padding: 30px 20px;
  height: auto;
  min-height: 33px;
  border-bottom: 2px dotted #eee;
}

.view-layout .header h1 {
  /*width: 500px;*/
  width: 100%;
  margin: 0 auto;
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
