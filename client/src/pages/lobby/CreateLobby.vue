<script setup lang="ts">
import { computed, ref } from 'vue'
import TavernShell from '../../components/TavernShell.vue'
import UiButton from '../../components/UiButton.vue'
import { gameApi, type Room } from '../../api'
import { useRouter } from 'vue-router'

const router = useRouter()
const maxPlayers = ref<2 | 4 | 6>(4)

const loading = ref(false)
const error = ref('')
const room = ref<Room | null>(null)

const canCreate = computed(() => !loading.value)

const createRoom = async () => {
  error.value = ''
  loading.value = true
  try {
    room.value = await gameApi.createRoom({ maxPlayers: maxPlayers.value })

    await router.push(`/lobby/${room.value.code}`)
  } catch (e: any) {
    error.value = e?.message ?? 'Failed to create room'
  } finally {
    loading.value = false
  }
}

const reset = () => {
  room.value = null
  error.value = ''
}

const copyCode = async () => {
  if (!room.value) return
  try {
    await navigator.clipboard.writeText(room.value.code)
  } catch {
    // fallback: not blocking UI for now
  }
}
</script>

<template>
  <TavernShell>
    <div class="panel">
      <div class="flex flex-col gap-1 sm:flex-row sm:items-end sm:justify-between">
        <div>
          <div class="font-display text-xl">Create game</div>
          <div class="text-sm text-ink-900/70">
            Host a room and share the code with friends.
          </div>
        </div>

        <RouterLink to="/" class="text-sm text-ink-900/60 hover:text-ink-900">
          ← Back to menu
        </RouterLink>
      </div>

      <div v-if="error" class="mt-4 rounded-xl border border-danger-500/40 bg-danger-500/10 p-3 text-sm">
        {{ error }}
      </div>

      <!-- BEFORE CREATED -->
      <div v-if="!room" class="mt-6 grid grid-cols-1 gap-6 lg:grid-cols-12">
        <div class="lg:col-span-7">
          <div class="font-display text-lg">Room settings</div>

          <div class="mt-3">
            <div class="text-sm font-semibold">Max players</div>

            <div class="mt-2 grid grid-cols-3 gap-2">
              <button
                class="rounded-xl border px-4 py-3 text-left transition"
                :class="maxPlayers === 2
                  ? 'border-candle-400/60 bg-candle-300/30'
                  : 'border-wood-700/35 bg-parchment-50/50 hover:bg-parchment-50/70'"
                @click="maxPlayers = 2"
              >
                <div class="font-semibold">2</div>
                <div class="text-sm text-ink-900/70">Duel</div>
              </button>

              <button
                class="rounded-xl border px-4 py-3 text-left transition"
                :class="maxPlayers === 4
                  ? 'border-candle-400/60 bg-candle-300/30'
                  : 'border-wood-700/35 bg-parchment-50/50 hover:bg-parchment-50/70'"
                @click="maxPlayers = 4"
              >
                <div class="font-semibold">4</div>
                <div class="text-sm text-ink-900/70">Party</div>
              </button>

              <button
                class="rounded-xl border px-4 py-3 text-left transition"
                :class="maxPlayers === 6
                  ? 'border-candle-400/60 bg-candle-300/30'
                  : 'border-wood-700/35 bg-parchment-50/50 hover:bg-parchment-50/70'"
                @click="maxPlayers = 6"
              >
                <div class="font-semibold">6</div>
                <div class="text-sm text-ink-900/70">Tavern</div>
              </button>
            </div>

            <div class="mt-3 text-sm text-ink-900/70">
              You can expand this later with rules (bet limits, timer, dice mode).
            </div>
          </div>

          <div class="mt-5 flex gap-2">
            <UiButton variant="primary" :disabled="!canCreate" @click="createRoom">
              {{ loading ? 'Creating…' : 'Create room' }}
            </UiButton>

            <RouterLink to="/">
              <UiButton variant="ghost">Back</UiButton>
            </RouterLink>
          </div>
        </div>

        <div class="lg:col-span-5">
          <div class="font-display text-lg">Preview</div>

          <div class="mt-3 rounded-[var(--radius-tavern)] border border-wood-700/30 bg-tavern-900/50 p-4">
            <div class="text-sm text-parchment-50/70">Room will be created as:</div>

            <div class="mt-3 space-y-2">
              <div class="chip">👥 Max players: {{ maxPlayers }}</div>
              <div class="chip">🕯️ Mode: Tavern Dice</div>
              <div class="chip">🔒 Access: Code</div>
            </div>

            <div class="mt-4 text-parchment-50/70 text-sm">
              After creation you’ll get a code to share.
            </div>
          </div>
        </div>
      </div>

      <!-- AFTER CREATED -->
      <div v-else class="mt-6 grid grid-cols-1 gap-6 lg:grid-cols-12">
        <div class="lg:col-span-7">
          <div class="font-display text-lg">Room created</div>
          <div class="mt-1 text-sm text-ink-900/70">
            Share this code with players to join your room.
          </div>

          <div class="mt-4 rounded-[var(--radius-tavern)] border border-wood-700/35 bg-parchment-50/60 p-4">
            <div class="text-sm font-semibold">Room code</div>

            <div class="mt-2 flex flex-col gap-2 sm:flex-row sm:items-center">
              <div
                class="flex-1 rounded-xl border border-wood-700/35 bg-parchment-50 px-4 py-3 font-display text-xl tracking-[0.08em]"
              >
                {{ room.code }}
              </div>

              <UiButton variant="primary" @click="copyCode">Copy</UiButton>
            </div>

            <div class="mt-3 text-sm text-ink-900/70">
              Max players: <span class="font-semibold">{{ room.maxPlayers }}</span> •
              Status: <span class="font-semibold">{{ room.status }}</span>
            </div>
          </div>

          <div class="mt-5 flex flex-wrap gap-2">
            <!-- Пізніше заміниш на перехід у реальну кімнату /room/:code -->
            <RouterLink to="/join">
              <UiButton variant="ghost">Go to Join page</UiButton>
            </RouterLink>

            <UiButton variant="primary" @click="reset">Create another</UiButton>

            <RouterLink to="/">
              <UiButton variant="ghost">Back</UiButton>
            </RouterLink>
          </div>
        </div>

        <div class="lg:col-span-5">
          <div class="font-display text-lg">Next steps</div>

          <div class="mt-3 space-y-3">
            <div class="rounded-[var(--radius-tavern)] border border-wood-700/30 bg-tavern-900/50 p-4">
              <div class="text-parchment-50 font-semibold">Invite players</div>
              <div class="text-parchment-50/70 text-sm mt-1">
                Send the code. Later you’ll replace this with REST + WebSocket room flow.
              </div>
            </div>

            <div class="rounded-[var(--radius-tavern)] border border-wood-700/30 bg-tavern-900/50 p-4">
              <div class="text-parchment-50 font-semibold">Start match</div>
              <div class="text-parchment-50/70 text-sm mt-1">
                On MVP, you can navigate to a `/room/:code` route and plug Three.js there.
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </TavernShell>
</template>
