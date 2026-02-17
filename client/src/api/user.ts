import fetchApi from '../packages/fetchApi.ts';
import type {Currency} from "@/api/common.ts";

export interface Balance {
  amount: number
  currency: Currency
}

export interface User {
  id: number
  username: number
  balance: Balance
}

export const getProfile = async () => {
  const { data } = await fetchApi.get('/profile');

  return data.data as User;
}
