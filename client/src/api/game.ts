import fetchApi from "@/packages/fetchApi.ts";
import type {Currency} from "@/api/common.ts";

export enum JoinType {
  ANYONE = 'anyone',
  FRIENDS = 'friends',
  LINK = 'link'
}

export interface CreateGameInput {
  currency_id: number
  bet: number
  winning_points: number
  join_type: JoinType
}

export interface Game {
  id: number
  code: string
  bet: number
  winning_points: number
  link: string
  currency: Currency
}

export const createGame = async (input: CreateGameInput): Promise<Game> => {
  const data = await fetchApi.post('/games', input)

  return data.data;
}
