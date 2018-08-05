<template>
    <div id="markdown_list" align="left">
        <!-- 笔记列表： -->
        <el-button @click="redirectNote()" type="text" size="small">添加</el-button>
        <el-table
            :data="data"
            border
            style="width: 100%"
            :row-class-name="tableRowClassName">
            <el-table-column
            prop="ID"
            label="编号"
            width="180"
            align="center">
            </el-table-column>
            <el-table-column
            prop="Title"
            label="标题"
            align="center">
            </el-table-column>
            <el-table-column
            prop="CreatedAt"
            label="时间"
            width="180"
            align="center">
            </el-table-column>
            <el-table-column
              fixed="right"
              label="操作"
              width="130"
              align="center">
              <template slot-scope="scope">
                <el-button @click="redirectNote(scope.row)" type="text" size="small">查看</el-button>
                <!-- <el-button type="text" size="small">编辑</el-button>
                <el-button type="text" size="small">删除</el-button> -->
              </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'MarkdownList',
  methods: {
    tableRowClassName(row, rowIndex) {
      if (rowIndex === 1) {
        return 'warning-row'
      } else if (rowIndex === 3) {
        return 'success-row'
      }
      return ''
    },
    initData() {
      const that = this // 没有这句，下面给 data 赋值会失败

      var instance = axios.create({
        baseURL: 'http://localhost:8520',
        timeout: 1000
        //   headers: { 'X-Custom-Header': 'foobar' }
      })
      instance
        .get('/GetNoteList', {})
        .then(response => {
          that.data = response.data.Data
        })
        .catch(error => {
          console.log(error)
        })
    },
    redirectNote(row) {
      console.log(row)
      let note
      if (row === undefined) {
        note = {
          path: 'markdown'
        }
      } else {
        note = {
          path: 'markdown',
          query: { noteID: row.ID }
        }
      }

      // 携带noteID跳转到markdown页面
      this.$router.push(note)
    }
  },
  created() {
    // `this` 指向 vm 实例
    this.initData()
  },
  data() {
    return {
      data: []
    }
  }
}
</script>

<style scoped>
.el-table .warning-row {
  background: rgb(148, 116, 58);
}

.el-table .success-row {
  background: #79b658;
}
</style>
