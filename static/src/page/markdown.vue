<template>
    <div id="markdown" align="left">
        <el-button @click="goback()" type="text" icon="el-icon-arrow-left">返回</el-button>
        <el-input
          type="textarea"
          autosize
          placeholder="请输入标题"
          v-model="title">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <mavon-editor v-model="detail" @save="saveNote"/>
    </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Markdown',
  methods: {
    saveNote() {
      var noteID = this.$route.query.noteID

      var title = this.title
      var detail = this.detail
      console.log(title, detail)
      axios.defaults.baseURL = 'http://localhost:8520'
      var params = new URLSearchParams() // 不是simple request的话，会将post方法变为options。使用 URLSearchParams 可以避免
      params.append('Title', title)
      params.append('Detail', detail)

      if (noteID === undefined) {
        // 添加
        axios
          .post('/AddNote', params)
          .then(response => {
            this.open('文章', '提交成功')
            console.log(response)
          })
          .catch(error => {
            this.open('文章', '提交失败: ' + error)
            console.log(error)
          })
      } else {
        // 修改
        params.append('ID', noteID)
        axios
          .post('/ModifyNote', params)
          .then(response => {
            this.open('文章', '提交成功')
            console.log(response)
          })
          .catch(error => {
            this.open('文章', '提交失败: ' + error)
            console.log(error)
          })
      }
    },
    open(title, message) {
      const h = this.$createElement

      this.$notify({
        title: title,
        message: h('i', { style: 'color: teal' }, message)
      })
    },
    initData() {
      var noteID = this.$route.query.noteID
      if (noteID === undefined) {
        return
      }
      const that = this // 没有这句，下面给 data 赋值会失败

      var instance = axios.create({
        baseURL: 'http://localhost:8520',
        timeout: 1000
      })
      console.log(noteID)
      instance
        .get('/GetNote', { params: { ID: noteID } })
        .then(response => {
          console.log(response.data)
          that.detail = response.data.Detail
          that.title = response.data.Title
        })
        .catch(error => {
          console.log(error)
        })
    },
    goback() {
      this.$router.go(-1)
    }
  },
  created() {
    // `this` 指向 vm 实例
    this.initData()
  },
  data() {
    return {
      detail: '# 今天又是美好的一天', // 这里可以设置默认值
      title: ''
    }
  }
}
</script>
