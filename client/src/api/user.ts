import fetchApi from '../packages/fetchApi.ts';

interface User {
  id: number
  username: number
}

export const getProfile = async () => {
  const { data } = await fetchApi.get('/profile');

  return data.data as User;
}
