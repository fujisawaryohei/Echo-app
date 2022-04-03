import { get } from "../utils/httpClient"
import { Post } from "../types"

export const getPosts = async(path: string, token?: string): Promise<Post[]> => {
  return await get<Post[]>(path, token)
}