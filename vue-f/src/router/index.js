import {
    createRouter,
    createWebHashHistory
} from 'vue-router'

import Index from "~/pages/index.vue"
import NotFound from "~/pages/404.vue"
import Login from "~/pages/login.vue"
import Register from "~/pages/register.vue"
import User from "~/pages/user.vue"
import UserEdit from "~/pages/edituser.vue"
import Doing from "~/pages/doing.vue"
import Addmsg from "~/pages/addmsg.vue"
import Addmsga from "~/pages/addmsga.vue"
import UserList from "~/pages/userlist.vue"
import Timing from "~/pages/timing.vue"
import Thing from "~/pages/things.vue"
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
        path:'/user',
        name:'User',
        component: User,
    },
    {
        path: '/edituser',
        name: 'EditUser',
        component:UserEdit,
    },
    {
        path: '/doing',
        name: 'Doing',
        component: Doing,
    },
    {
        path: '/addmsg',
        name: 'addmsg',
        component: Addmsg,
    },
    {
        path: '/addmsga',
        name: 'addmsga',
        component: Addmsga,
    },
    {
        path: '/userlist',
        name: 'userlist',
        component: UserList,
    },
    {
        path: '/timing',
        name: 'timing',
        component: Timing,
    },
    {
        path: '/thing',
        name: 'thing',
        component: Thing,
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