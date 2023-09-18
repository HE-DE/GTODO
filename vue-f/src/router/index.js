import {
    createRouter,
    createWebHashHistory
} from 'vue-router'

import Index from "~/pages/index.vue"
import NotFound from "~/pages/404.vue"
import Login from "~/pages/login.vue"
import Register from "~/pages/register.vue"
import {useUsersStore} from '../store/user';

const routes = [
    {
        path: '/',
        redirect: '/login'
    },
    {
        path: '/information',
        name: 'Index',
        component: Index,
        beforeEnter: (to, from, next) => {
            // 判断是否登录
            const user = useUsersStore();
            if (user.isLogin) {
                next();
            }else{
                router.push('/login');
            }
        }
    },
    {
        path: '/login',
        name: 'Login',
        component: Login
    },
    {
        path: '/register',
        name: 'Register',
        component: Register,
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: NotFound
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})

export default router