<template>
  <ElRow class="question-wrapper">

    <el-col :span="24" class="question-list">
      <p v-if="tableChildData.length == 0" style="margin: 10px 0">一点东西都没有，赶快点击上方按钮添加题目吧！</p>
      <ElRow type="flex" justify="start" align="top" v-for="(question, index) in tableChildData.Question" class="question-item" :key="question.ID">
        <el-col :span="6" style="width: 60px; text-align: center">
              <h2>Q{{ index + 1 }}:</h2>
              <div class="question-action" v-show="isPreview" @click="delQuestion(index)">
              </div>
        </el-col>

        <el-col :span="18">
          <h3>[{{ question.topic_id }}] {{ question.question_name}} {{question.answer ? "（必填）" : "（选填）"}} <em style="color: #f00;" v-if="question.answer">*</em></h3>
          <p class="question-desc" v-if="question.direction !== ''">说明：{{ question.direction }}</p>

          <div class="question-options">
                <!--判断是否有图片-->
                <div v-if="question.topic_id ===3">

                </div>
                <!--判断多选-->

                <!--判断文本-->
                <div v-if="question.topic_id ==='4 '">
                    <el-input v-model="question.selectContent" placeholder="请输入..." ></el-input>
                </div>
                <!--判断图片-->
          </div>
        </el-col>
       </ElRow>
      <!--卡槽 传递按钮-->
      <div class="question-btns">
        <slot></slot>
      </div>

    </el-col>

  </ElRow>

</template>
<script>
export default {
  components: {},
  props:['tableChildData'],
  data() {
    return {
      addTextarea_form:{},
      environmentForm: {
        items: [{
          name: '',
          variable: '',
          description: ''
        }]
      },
      list:'',
      isPreview:'',
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
  computed: {},
  watch: {},
  created() {},
  mounted() {},
  methods: {
    // submitForm1() {
    //   this.$refs['elForm'].validate(valid => {
    //     if (!valid) return
    //     // TODO 提交表单
    //   })
    // },
    //提交事件
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          console.log('提交',this.environmentForm);
          alert('submit!');
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    //重置事件
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    //移除表单项事件
    removeEnvironmentForm(item) {
      var index = this.environmentForm.items.indexOf(item)
      if (index !== -1) {
        this.environmentForm.items.splice(index, 1)
      }
    },
    //添加表单项事件
    addEnvironmentForm() {
      this.environmentForm.items.push({
        name: '',
        variable: '',
        description: '',
        key: Date.now()
      });
    }
  },
}

</script>

<style scoped>
.question-list {
  padding: 30px 0;
  overflow: visible;
}

.checkbox-list label {
  display: block;
}

.question-options {
  padding: 5px 0;
}

.option-item {
  margin: 5px 0;
}

.question-desc {
  padding: 5px 0;
}

.question-action {
  display: none;
  color: #a9afb2;
}

.question-item {
  padding: 20px 0;
  min-width: 300px;
}

.question-item:hover {
  background: rgba(238, 238, 238, 0.47);
}

.question-item:hover .question-action {
  display: block;
  margin-top: 8px;
}

.question-action:hover {
  color: #018fe5;
  cursor: pointer;
}

.option-item {

}

.question-btns {
  margin-top: 20px;
}

.option-action {
  display: none;
  margin-left: 100px;
  padding: 0 10px;
  text-align: right;
  color: #a9afb2;
}

.option-action:hover {
  color: #018fe5;
}

.option-item:hover .option-action {
  display: inline-block;
  cursor: pointer;
}
</style>
