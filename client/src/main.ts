import {createApp} from 'vue'
import './index.css'
import App from './App.vue'
import router from "./routes";
import {createPinia} from 'pinia'
import {useAuthStore} from "@/stores/auth.ts";

const app = createApp(App)
  .use(router)
  .use(createPinia())

const auth = useAuthStore()

auth.bootstrap().finally(() => app.mount('#app'))
