import {createWebHistory, createRouter, type RouteRecordRaw} from "vue-router";
import {useAuthStore} from "@/stores/auth.ts";

const routes: RouteRecordRaw[] = [
  { path: '/', name: 'home', component: () => import('../pages/home/Index.vue') },
  { path: '/create', name: 'create', component: () => import('../pages/lobby/CreateLobby.vue') },
  { path: '/join', name: 'join', component: () => import('../pages/lobby/JoinLobby.vue') },
  { path: '/leaderboard', name: 'leaderboard', component: () => import('../pages/leaderboard/LeaderboardList.vue') },
  { path: '/login', name: 'login', component: () => import('../pages/auth/Login.vue') },
  { path: '/register', name: 'register', component: () => import('../pages/auth/Registration.vue') },
  { path: '/lobby/:code', name: 'lobby', component: () => import('../pages/lobby/LobbyRoom.vue') },
  { path: '/room/:code', name: 'room', component: () => import('../pages/game/GameRoom.vue') },
]

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
})

router.beforeEach((to, from, next) => {
  const { isLoggedIn } = useAuthStore();
  const isLoginRoutes = ['login', 'register'].includes(to.name as string)

  if (isLoggedIn && isLoginRoutes) {
    next({ name: 'home' })
  } else if (!isLoggedIn && !isLoginRoutes) {
    next({ name: 'login' })
  } else {
    next()
  }
})

export default router
