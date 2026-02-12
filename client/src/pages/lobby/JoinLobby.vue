<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import TavernShell from '../../components/TavernShell.vue'
import UiButton from '../../components/UiButton.vue'
import { gameApi, type Room } from '../../api/'
import { useRouter } from 'vue-router'

const router = useRouter()

const code = ref('')
const loading = ref(false)
const error = ref('')

const rooms = ref<Room[]>([])
const roomsLoading = ref(false)

/** пагінація */
const page = ref(1)
const pageSize = ref(5) // можеш поставити 6/8 — як тобі комфортніше

const normalizedCode = computed(() => code.value.trim().toUpperCase())

const totalPages = computed(() => {
  const total = rooms.value.length
  return Math.max(1, Math.ceil(total / pageSize.value))
})

const safePage = computed(() => Math.min(Math.max(page.value, 1), totalPages.value))

const pagedRooms = computed(() => {
  const start = (safePage.value - 1) * pageSize.value
  return rooms.value.slice(start, start + pageSize.value)
})

const showingFrom = computed(() => {
  if (rooms.value.length === 0) return 0
  return (safePage.value - 1) * pageSize.value + 1
})

const showingTo = computed(() => {
  if (rooms.value.length === 0) return 0
  return Math.min(safePage.value * pageSize.value, rooms.value.length)
})

const join = async (roomCode?: string) => {
  error.value = ''
  loading.value = true
  try {
    const c = (roomCode ?? normalizedCode.value).trim()
    if (!c) throw new Error('Enter room code')

    await gameApi.joinRoom({ code: c })

    // Join → lobby screen
    await router.push(`/lobby/${c}`)  }
  catch (e: any) {
    error.value = e?.message ?? 'Failed to join'
  } finally {
    loading.value = false
  }
}

const loadRooms = async () => {
  roomsLoading.value = true
  try {
    rooms.value = await gameApi.listPublicRooms()
    page.value = 1 // при оновленні завжди на першу сторінку
  } finally {
    roomsLoading.value = false
  }
}

const prevPage = () => {
  page.value = Math.max(1, safePage.value - 1)
}

const nextPage = () => {
  page.value = Math.min(totalPages.value, safePage.value + 1)
}

onMounted(loadRooms)
</script>

<template>
  <TavernShell>
    <div class="panel">
      <div class="font-display text-xl">Join room</div>

      <div v-if="error" class="mt-4 rounded-xl border border-danger-500/40 bg-danger-500/10 p-3 text-sm">
        {{ error }}
      </div>

      <label class="mt-4 block text-sm font-semibold">Room code</label>

      <input
        v-model="code"
        class="mt-1 w-full rounded-xl border border-wood-700/35 bg-parchment-50 px-4 py-2 text-ink-900 outline-none
               focus:ring-2 focus:ring-candle-400/60"
        placeholder="e.g. TAV-9X2K"
        @keydown.enter="join()"
      />

      <div class="mt-4 flex gap-2">
        <UiButton variant="primary" :disabled="loading" @click="join()">
          {{ loading ? 'Joining…' : 'Join' }}
        </UiButton>

        <RouterLink to="/">
          <UiButton variant="ghost">Back</UiButton>
        </RouterLink>
      </div>

      <!-- PUBLIC ROOMS -->
      <div class="mt-8">
        <div class="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <div class="font-display text-lg">Public rooms</div>
            <div class="text-sm text-ink-900/70">Pick one to join (mock). Later — REST.</div>
          </div>

          <div class="flex items-center gap-2">
            <UiButton variant="ghost" :disabled="roomsLoading" @click="loadRooms">
              {{ roomsLoading ? 'Refreshing…' : 'Refresh' }}
            </UiButton>
          </div>
        </div>

        <div v-if="roomsLoading" class="mt-4 text-ink-900/60">Loading…</div>

        <div v-else class="mt-4 space-y-2">
          <button
            v-for="r in pagedRooms"
            :key="r.code"
            class="w-full text-left flex items-center justify-between rounded-xl border border-wood-700/35
                   bg-parchment-50/60 px-4 py-3 transition hover:bg-parchment-50/80 disabled:opacity-60"
            :disabled="r.status !== 'open' || r.players >= r.maxPlayers || loading"
            @click="join(r.code)"
          >
            <div>
              <div class="font-semibold text-ink-900">{{ r.code }}</div>
              <div class="text-sm text-ink-900/65">
                Host: {{ r.hostName }} • {{ r.players }}/{{ r.maxPlayers }} • {{ r.status }}
              </div>
            </div>

            <div class="text-ink-900/70">
              <span v-if="r.status !== 'open'">In game</span>
              <span v-else-if="r.players >= r.maxPlayers">Full</span>
              <span v-else class="text-candle-400 font-semibold">Join →</span>
            </div>
          </button>
        </div>

        <!-- PAGINATION -->
        <div class="mt-4 flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
          <div class="text-sm text-ink-900/65">
            Showing <span class="font-semibold">{{ showingFrom }}</span>–<span class="font-semibold">{{ showingTo }}</span>
            of <span class="font-semibold">{{ rooms.length }}</span>
          </div>

          <div class="flex items-center gap-2">
            <UiButton variant="ghost" :disabled="safePage === 1 || roomsLoading" @click="prevPage">
              ← Prev
            </UiButton>

            <div class="text-sm text-ink-900/70">
              Page <span class="font-semibold">{{ safePage }}</span> / <span class="font-semibold">{{ totalPages }}</span>
            </div>

            <UiButton variant="ghost" :disabled="safePage === totalPages || roomsLoading" @click="nextPage">
              Next →
            </UiButton>
          </div>
        </div>
      </div>
    </div>
  </TavernShell>
</template>
