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
          console.log(response)
        })
        .catch(error => {
          console.log(error)
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
