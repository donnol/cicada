<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <basic-table> abc </basic-table>
    <basicTable> abc </basicTable>
    <BasicTable> abc </BasicTable>
    <basictable> abc </basictable>
    <table class="exp_table" border=1>
      <thead>
      <tr>
      <th v-for='key in headers' v-bind:key="sortBy(key)">
        {{ key }}
      </th>
      <!-- <th>编号</th>
      <th>用户</th>
      <th>支出（元）</th>
      <th>物品/事情</th>
      <th>时间</th>
      <th>地点</th>
      <th>操作</th> -->
      </tr>
      </thead>

      <tbody>
      <tr>
      <td>1</td>
      <td>用户1</td>
      <td>100</td>
      <td>充车卡</td>
      <td>2018-03-27</td>
      <td>南桂路地铁站自助充值机</td>
      <td><a>修改</a>|<a>删除</a></td>
      </tr>
      </tbody>

    </table>
  </div>
</template>

<script>
import BasicTable from '@/components/BasicTable'

export default {
  name: 'ExpensesTable',
  components: { BasicTable },
  data() {
    return {
      msg: '支出明细表',
      header: [
        '编号',
        '用户',
        '支出（元）',
        '物品/事情',
        '时间',
        '地点',
        '操作'
      ]
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  font-weight: normal;
}
</style>
