import { defineStore } from "pinia";

export const useUsersStore = defineStore("user", {
    state: () =>
    ({
        username: "",
        Id:"",
        isAdmin: false,
        isLogin: false
    }),
    actions: {
        Login(thename,theIsAdmin,theId){
            this.Id=theId
            this.username=thename
            this.isAdmin=theIsAdmin
            this.isLogin=true
        },
    }
})