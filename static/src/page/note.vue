<template>
    <div id="note">
        <el-input v-model="input" placeholder="请输入内容"></el-input>
        <el-button v-on:click="addNote">提交</el-button>
    </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Note',
  methods: {
    addNote() {
      var detail = this.input
      console.log(detail)
      axios.defaults.baseURL = 'http://localhost:8520'
      var params = new URLSearchParams()
      params.append('Detail', detail)
      axios
        .post('/AddNote', params) // 不是simple request的话，会将post方法变为options。使用 URLSearchParams 可以避免
        .then(response => {
          this.open('笔记', '提交成功')
          console.log(response)
        })
        .catch(error => {
          this.open('笔记', '提交失败: ' + error)
          console.log(error)
        })
    },
    open(title, message) {
      const h = this.$createElement

      this.$notify({
        title: title,
        message: h('i', { style: 'color: teal' }, message)
      })
    }
  },
  data() {
    return {
      input: '哈哈' // 这里可以设置默认值
    }
  }
}
</script>
