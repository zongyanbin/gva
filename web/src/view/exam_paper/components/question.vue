<template>
  <div>

    <template>
      <el-table
          :data="tableChildData.Question"
          style="width: 100%">
        <el-table-column
            prop="ID"
            label="ID"
            width="180">
        </el-table-column>
        <el-table-column
            prop="question_name"
            label="question_name"
            width="180">
        </el-table-column>
        <el-table-column
            prop="direction"
            label="direction">
        </el-table-column>
      </el-table>
    </template>
    <p v-if="tableChildData.length == 0" style="margin: 10px 0">一点东西都没有，赶快点击上方按钮添加题目吧！</p>
    <el-form-item label="活动名称">
      <el-input v-model="form.name"></el-input>
    </el-form-item>


<!--    {{tableChildData.Question}}-->
    <li v-for="(item ,index) in tableChildData.Question" :key="index">{{item}}</li>
     zi zujian
    <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
      <el-form-item label="单行文本" prop="field101">
        <el-input v-model="formData.field101" placeholder="请输入单行文本" clearable :style="{width: '100%'}">
        </el-input>
      </el-form-item>
      <el-form-item label="多行文本" prop="field102">
        <el-input v-model="formData.field102" type="textarea" placeholder="请输入多行文本"
                  :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item size="large">
        <el-button type="primary" @click="submitForm">提交</el-button>
        <el-button @click="resetForm">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  components: {},
  props:['tableChildData'],
  data() {
    return {
      list:'',
      tableData: [{
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1517 弄'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1519 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }],
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
    submitForm() {
      this.$refs['elForm'].validate(valid => {
        if (!valid) return
        // TODO 提交表单
      })
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
  }
}

</script>
<style>
</style>
