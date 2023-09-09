export{}
import VueRouter, { Router } from 'vue-router';
declare module 'vue/types/vue' {
  interface Vue {
    $router: Router
  }
}