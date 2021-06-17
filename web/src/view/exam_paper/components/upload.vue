<template>
  <div>
    <el-upload
        :action="`${path}/extend/upload`"
        list-type="picture-card"
        :on-preview="handlePictureCardPreview"
        :on-success="handleImageSuccess"
        :on-remove="handleRemove">
      <i class="el-icon-plus"></i>
    </el-upload>
    <el-dialog :visible.sync="dialogVisible">
      <img width="100%" :src="dialogImageUrl" alt="">
    </el-dialog>
  </div>
</template>
<script>
import {mapGetters} from "vuex";

const path = process.env.VUE_APP_BASE_API;
export default {
  data() {
    return {
      dialogImageUrl: '',
      dialogVisible: false,
      path: path,
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo", "token"])
  },
  methods: {
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePictureCardPreview(file) {
      console.log("file:",file)
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },

    handleImageSuccess(res) {
    //  console.log(res);
      // this.imageUrl = URL.createObjectURL(file.raw);
      const { data } = res;
      if (data.file) {
        this.$emit("extend_data", data.file.url);
      }
    },

  }
}
</script>