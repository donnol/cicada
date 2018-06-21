<template>
    <div id="markdown">
        <mavon-editor v-model="input" @save="addNote"/>
    </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Markdown',
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
      input: '# 今天又是美好的一天' // 这里可以设置默认值
    }
  }
}
</script>
