<script setup lang="ts">
import TavernShell from '../../components/TavernShell.vue'
import MenuLink from '../../components/MenuLink.vue'
</script>

<template>
  <TavernShell>
    <div class="mx-auto max-w-2xl space-y-3">
      <MenuLink
        title="Create game"
        description="Host a room and share the code with friends."
        to="/create"
      />

      <MenuLink
        title="Join room"
        description="Enter a room code or pick from public rooms."
        to="/join"
      />

      <MenuLink
        title="Leaderboard"
        description="Top players, wins and rating."
        to="/leaderboard"
      />
    </div>
  </TavernShell>
</template>
