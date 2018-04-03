<template>
    <div id="expense">
        <el-table
            :data="data"
            border
            style="width: 100%"
            :row-class-name="tableRowClassName">
            <el-table-column
            prop="ID"
            label="编号"
            width="180">
            </el-table-column>
            <el-table-column
            prop="UserID"
            label="用户编号"
            width="180">
            </el-table-column>
            <el-table-column
            prop="Pay"
            label="支出(元)">
            </el-table-column>
            <el-table-column
            prop="Thing"
            label="事情/物品">
            </el-table-column>
            <el-table-column
            prop="CreatedAt"
            label="时间">
            </el-table-column>
            <el-table-column
            prop="CreatedOn"
            label="地点">
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Expense',
  methods: {
    tableRowClassName({ row, rowIndex }) {
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
        .get('/ExpenseList', {})
        .then(response => {
          that.data = response.data
        })
        .catch(error => {
          console.log(error)
        })
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
