import { createRouter, createWebHistory } from 'vue-router'
import HomeView from './routes/HomeView.vue'

const router = createRouter({
	//@ts-ignore
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			name: 'index',
			component: HomeView
		},
		{
			path: '/login',
			name: 'login',
			// route level code-splitting
			// this generates a separate chunk (About.[hash].js) for this route
			// which is lazy-loaded when the route is visited.
			component: () => import('./routes/LoginView.vue')
		}
	]
})

export default router
