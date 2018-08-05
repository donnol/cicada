import Vue from 'vue'
import Router from 'vue-router'
import Expense from '@/page/expense'
import Note from '@/page/note'
import Markdown from '@/page/markdown'
import MarkdownList from '@/page/markdown_list'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'

Vue.use(ElementUI)
Vue.use(mavonEditor)
Vue.use(Router)

export default new Router({
  routes: [
    // eslint-disable-next-line
    {
      path: '/expense',
      name: 'Expense',
      component: Expense
    },
    {
      path: '/note',
      name: 'Note',
      component: Note
    },
    {
      path: '/markdown',
      name: 'Markdown',
      component: Markdown,
      meta: {
        title: '笔记-留住美好'
      }
    },
    {
      path: '/markdownList',
      name: 'MarkdownList',
      component: MarkdownList,
      meta: {
        title: '笔记列表-留住美好'
      }
    }
  ]
})
