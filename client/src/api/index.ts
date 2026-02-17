// src/services/gameApi.ts
export type User = {
  id: string
  username: string
  avatar?: string
}

export type Room = {
  code: string
  hostName: string
  players: number
  maxPlayers: number
  status: 'open' | 'in_game'
  createdAtIso: string
}

export type LeaderboardRow = {
  rank: number
  username: string
  games: number
  wins: number
  winRate: number // 0..100
  rating: number
}

export type Session = {
  token: string
  user: User
}

export type JoinRoomInput = {
  code: string
}

export type LoginInput = {
  username: string
  password: string
}

export type RegisterInput = {
  username: string
  password: string
}

export interface GameApi {
  joinRoom(input: JoinRoomInput): Promise<{ room: Room }>
  listPublicRooms(): Promise<Room[]>

  getLeaderboard(): Promise<LeaderboardRow[]>
}

// ---- helpers ----
const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms))
const nowIso = () => new Date().toISOString()

const rand = (min: number, max: number) =>
  Math.floor(Math.random() * (max - min + 1)) + min

// ---- mock data ----
let mockRooms: Room[] = [
  { code: 'TAV-AA22', hostName: 'Bohdan', players: 1, maxPlayers: 4, status: 'open', createdAtIso: nowIso() },
  { code: 'TAV-QW77', hostName: 'Yulia', players: 3, maxPlayers: 6, status: 'open', createdAtIso: nowIso() },
  { code: 'TAV-ZK33', hostName: 'Max', players: 2, maxPlayers: 2, status: 'in_game', createdAtIso: nowIso() },
]

const mockLeaderboard: LeaderboardRow[] = Array.from({ length: 12 }).map((_, i) => {
  const games = rand(10, 120)
  const wins = rand(0, games)
  const winRate = Math.round((wins / games) * 100)
  return {
    rank: i + 1,
    username: ['Hermann', 'Scipio', 'Joan', 'Pelagius', 'SunTzu', 'Aethel', 'Kusu', 'Baibars', 'CaoCao', 'Bjorn'][i % 10] + `#${rand(10, 99)}`,
    games,
    wins,
    winRate,
    rating: 1200 + (12 - i) * 37 + rand(-20, 20),
  }
})

// ---- mock api ----
export const gameApi: GameApi = {
  async joinRoom(input) {
    await sleep(450)
    const code = input.code.trim().toUpperCase()
    const room = mockRooms.find((r) => r.code === code)
    if (!room) throw new Error('Room not found')
    if (room.status !== 'open') throw new Error('Room is already in game')
    if (room.players >= room.maxPlayers) throw new Error('Room is full')

    room.players += 1
    return { room }
  },

  async listPublicRooms() {
    await sleep(350)
    return mockRooms
  },

  async getLeaderboard() {
    await sleep(450)
    return mockLeaderboard
  },
}
