<script setup lang="ts">
import UiButton from './UiButton.vue'
import {useAuthStore} from "@/stores/auth.ts";

const props = withDefaults(defineProps<{ withHeader?: boolean, withFooter?: boolean }>(), {
  withHeader: true,
  withFooter: true,
})

const {user, logout} = useAuthStore()

console.log(user)

</script>

<template>
  <div
    v-if="props.withHeader"
    class="min-h-dvh w-full
           flex flex-col justify-between
           px-3 py-4 sm:px-6 sm:py-6
           bg-[radial-gradient(circle_at_50%_20%,rgba(212,145,57,.22),transparent_60%)]
           "
  >
    <header class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
      <div @click="$router.push('/')" class="cursor-pointer">
        <div class="text-candle-200 font-display tracking-[0.06em] text-2xl">
          Tavern Dice
        </div>
        <div class="text-parchment-50/70">
          Multiplayer dice in a tavern mood.
        </div>
      </div>

      <div class="flex items-center gap-2 flex-wrap justify-start sm:justify-end">
        <span class="chip" v-if="user">👤 {{ user.username }}</span>

        <RouterLink v-if="!user" to="/login">
          <UiButton variant="ghost">Login</UiButton>
        </RouterLink>

        <RouterLink v-if="!user" to="/register">
          <UiButton variant="primary">Register</UiButton>
        </RouterLink>

        <UiButton v-if="user" variant="ghost" @click="logout()">
          Logout
        </UiButton>
      </div>
    </header>

    <main class="mt-6">
      <slot />
    </main>

    <footer v-if="props.withFooter" class="mt-8 text-parchment-50/50 text-sm text-center">
      <span>© {{ new Date().getFullYear() }} Tavern Dice. All Rights Reserved</span>
    </footer>
  </div>
</template>
