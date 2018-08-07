<template>
    <div id="markdown_list" align="left">
        <!-- 笔记列表： -->
        <el-button @click="redirectNote()" type="text" size="small">添加</el-button>
        <el-button @click="initData()" type="text" size="small">刷新</el-button>
        <el-button @click="resetData()" type="text" size="small">重置</el-button>
        <el-input placeholder="请输入内容" v-model="searchInput" class="input-with-select">
          <el-select v-model="select" slot="prepend" placeholder="请选择">
            <el-option label="编号" value="1"></el-option>
            <el-option label="标题" value="2"></el-option>
          </el-select>
          <el-button @click="initData()" slot="append" icon="el-icon-search"></el-button>
        </el-input>
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
        <el-button-group align="center">
          <el-button @click="initData()" type="text" icon="el-icon-arrow-left">上一页</el-button>
          <el-button @click="initData()" type="text">下一页<i class="el-icon-arrow-right el-icon--right"></i></el-button>
        </el-button-group>
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
      // TODO 分页
      var param = {}
      var searchInput = this.searchInput
      if (searchInput !== '') {
        let values = {}
        if (that.select === '1') {
          values['ID'] = searchInput
        } else if (that.select === '2') {
          values['Title'] = searchInput
        } else {
          values['Title'] = searchInput
        }
        param = { params: values }
      }
      console.log(that.select)
      console.log(param)
      instance
        .get('/GetNoteList', param)
        .then(response => {
          that.data = response.data.Data
        })
        .catch(error => {
          console.log(error)
        })
    },
    resetData() {
      const that = this

      that.searchInput = ''
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
      data: [],
      searchInput: '',
      select: []
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

.el-select {
  width: 130px;
}

.input-with-select .el-input-group__prepend {
  background-color: #fff;
}
</style>
