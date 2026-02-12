<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import TavernShell from '../../components/TavernShell.vue'
import UiButton from '../../components/UiButton.vue'
import { useToast } from '../../composables/useToast'

type LobbyPlayer = {
  id: string
  name: string
  isHost?: boolean
  isReady?: boolean
}

const route = useRoute()
const router = useRouter()
const toast = useToast()

const roomCode = computed(() => String(route.params.code ?? 'TAV-AA22').toUpperCase())

// ---- MOCK "session" ----
const myId = 'me'
const myName = 'Guest'
const isHost = ref(true) // перемкни на false щоб глянути guest-UI

// ---- MOCK lobby state ----
const maxPlayers = ref<2 | 4 | 6>(4)
const players = ref<LobbyPlayer[]>([
  { id: 'me', name: myName, isHost: true, isReady: true },
  { id: 'p2', name: 'Bohdan', isReady: true },
  { id: 'p3', name: 'Yulia', isReady: true },
  { id: 'p3', name: 'Yulia', isReady: true },
])

const me = computed(() => players.value.find((p) => p.id === myId))
const filled = computed(() => players.value.length)
const emptySlots = computed(() => Math.max(0, maxPlayers.value - players.value.length))

const allReady = computed(() => players.value.length >= 2 && players.value.every((p) => p.isReady))
const canStart = computed(() => isHost.value && allReady.value)

const toggleReady = () => {
  const p = me.value
  if (!p) return
  p.isReady = !p.isReady
}

const copyCode = async () => {
  try {
    await navigator.clipboard.writeText(roomCode.value)
    toast.push({ kind: 'success', title: 'Copied', message: `Room code ${roomCode.value} copied.` })
  } catch {
    toast.push({ kind: 'error', title: 'Copy failed', message: 'Clipboard is not available.' })
  }
}

const copyInviteLink = async () => {
  const link = `${window.location.origin}/lobby/${roomCode.value}`
  try {
    await navigator.clipboard.writeText(link)
    toast.push({ kind: 'success', title: 'Copied', message: 'Invite link copied.' })
  } catch {
    toast.push({ kind: 'error', title: 'Copy failed', message: 'Clipboard is not available.' })
  }
}

const leaveLobby = () => {
  // пізніше тут буде API call
  router.push('/')
}

const startGame = () => {
  if (!canStart.value) {
    toast.push({ kind: 'error', title: 'Cannot start', message: 'Need at least 2 players, and everyone must be ready.' })
    return
  }

  toast.push({ kind: 'success', title: 'Starting', message: 'Game is starting… (mock)' })
  router.push(`/room/${roomCode.value}`)
}
</script>

<template>
  <TavernShell>
    <div class="panel">
      <!-- HEADER -->
      <div class="flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between">
        <div>
          <div class="font-display text-xl">Lobby</div>
          <div class="text-sm text-ink-900/70">
            Room <span class="font-semibold">{{ roomCode }}</span> • {{ filled }}/{{ maxPlayers }} players
          </div>
        </div>

        <div class="flex flex-wrap items-center gap-2">
          <UiButton variant="ghost" @click="copyCode">Copy code</UiButton>
          <UiButton variant="ghost" @click="copyInviteLink">Copy invite link</UiButton>
          <UiButton variant="ghost" @click="leaveLobby">Leave</UiButton>
        </div>
      </div>

      <div class="mt-6 grid grid-cols-1 gap-6 lg:grid-cols-12">
        <!-- LEFT: PLAYERS -->
        <section class="lg:col-span-7">
          <div class="font-display text-lg">Players</div>

          <div class="mt-3 grid grid-cols-1 gap-2 sm:grid-cols-2">
            <div
              v-for="p in players"
              :key="p.id"
              class="flex items-center justify-between rounded-xl border border-wood-700/35 bg-tavern-900/50 px-4 py-3"
            >
              <div class="flex items-center gap-3 min-w-0">
                <div class="h-9 w-9 rounded-full bg-candle-300/35 flex items-center justify-center shrink-0">
                  🎲
                </div>

                <div class="min-w-0">
                  <div class="font-semibold text-parchment-50 truncate">
                    {{ p.name }}
                    <span v-if="p.isHost" class="ml-1 text-xs text-candle-300">(host)</span>
                    <span v-if="p.id === myId" class="ml-1 text-xs text-parchment-50/60">(you)</span>
                  </div>
                  <div class="text-xs text-parchment-50/60">
                    {{ p.isReady ? 'Ready' : 'Not ready' }}
                  </div>
                </div>
              </div>

              <div
                class="text-xs px-2 py-1 rounded-full shrink-0"
                :class="p.isReady
                  ? 'bg-candle-300/30 text-parchment-50'
                  : 'bg-tavern-800/50 text-parchment-50/60'"
              >
                {{ p.isReady ? 'Ready' : 'Waiting' }}
              </div>
            </div>

            <div
              v-for="i in emptySlots"
              :key="'empty-' + i"
              class="flex items-center justify-center rounded-xl border border-dashed border-wood-700/35
                     bg-tavern-900/30 px-4 py-3 text-parchment-50/35"
            >
              Empty slot
            </div>
          </div>

          <div class="mt-4 text-sm text-ink-900/60">
            To start the game: минимум 2 гравці і всі мають бути <span class="font-semibold">Ready</span>.
          </div>
        </section>

        <!-- RIGHT: CONTROL -->
        <aside class="lg:col-span-5">
          <div class="font-display text-lg">Controls</div>

          <div class="mt-3 rounded-[var(--radius-tavern)] border border-wood-700/30 bg-parchment-50/60 p-4">
            <div class="text-sm font-semibold">Room status</div>

            <div class="mt-3 space-y-2">
              <div class="flex items-center justify-between">
                <span class="text-ink-900/70">All ready</span>
                <span class="font-semibold">{{ allReady ? 'Yes' : 'No' }}</span>
              </div>

              <div class="flex items-center justify-between">
                <span class="text-ink-900/70">You</span>
                <span class="font-semibold">{{ me?.isReady ? 'Ready' : 'Not ready' }}</span>
              </div>

              <div class="flex items-center justify-between">
                <span class="text-ink-900/70">Host</span>
                <span class="font-semibold">{{ isHost ? 'You' : 'Other' }}</span>
              </div>
            </div>

            <div class="mt-4 flex flex-col gap-2">
              <!-- HOST: START -->
              <UiButton
                v-if="isHost"
                variant="primary"
                :disabled="!canStart"
                @click="startGame"
              >
                Start game
              </UiButton>

              <!-- GUEST: READY -->
              <UiButton
                v-else
                :variant="me?.isReady ? 'ghost' : 'primary'"
                @click="toggleReady"
              >
                {{ me?.isReady ? 'Unready' : 'Ready' }}
              </UiButton>

              <UiButton variant="ghost" @click="copyInviteLink">
                Invite
              </UiButton>
            </div>
          </div>

          <div class="mt-3 text-sm text-ink-900/60">
            Пізніше тут будуть: правила кімнати, таймер старту, чат/смайли.
          </div>
        </aside>
      </div>
    </div>
  </TavernShell>
</template>
