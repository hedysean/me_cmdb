import { createRouter, createWebHistory } from 'vue-router'
// import ShowCenter from "../views/ShowCenter.vue"
import Base from "../views/Base.vue"
import Login from "@/views/Login.vue"
import Provider from "@/views/Provider.vue"
import { ConfigProvider } from 'ant-design-vue'
import { provide } from 'vue';

const routes = [
  {
    meta:{
      title: 'MeCMDB'
    }
  },
  {
    path: '/',
    alias: '/',
    name: 'home',
    component: Base,
    children: [
      {
        path: 'assets/provider',
        name: 'provider',
        component: Provider,
        meta: {
          title: "供应商管理",
          authorization: false
        }
      }
    ],
  },
  {
    meta: {
      title: '账户登陆',
      authorization: false
    },
    path: '/login',
    name: 'Login',
    component: Login // 快捷键：Alt+Enter快速导包，
  },
  {
    // path: '/about',
    // name: 'about',
    // // route level code-splitting
    // // this generates a separate chunk (about.[hash].js) for this route
    // // which is lazy-loaded when the route is visited.
    // component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

// 导航守卫
router.beforeEach((to, from, next) => {
  // console.log("to", to)
  // console.log("from", from)
  // console.log("nextß", next)

  document.title = to.meta.title

  // 登录状态验证
  console.log("to meta", to.meta)
  if (to.meta.authorization && !storage.getUserInfo()) {
    next({"name": "Login"})
  } else {
    next()
  }
})

export default router
