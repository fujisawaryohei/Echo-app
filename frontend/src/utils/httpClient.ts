import axios, { AxiosResponse } from "axios"

// envによってURLを変更する処理を追加
// const BASE_URL = 'http://127.0.0.1:4010'
const BASE_URL = 'http://localhost:8000'

export const get = async<T>(path: string, token?: string): Promise<T> => {
  return await axios.get<T>(BASE_URL + path, Headers(token)).then((response: AxiosResponse<T>) => {
    return responseHandler<T>(response)
  })
}

const Headers = (token?: string) => {
  return {
    headers: {
      Authorization: `Bearer ${token}`
    }
  }
}

const responseHandler = <T>(response: AxiosResponse<T>): Promise<T> => {
  switch (response.status) {
    case 200:
      return Promise.resolve(response.data)
    case 201:
      return Promise.resolve(response.data)
    case 400:
      return Promise.reject(response.data)
    case 401:
      return Promise.reject(new Error('セッションが切れました。ログインしてください。'))
    case 500:
      return Promise.reject(new Error('サーバーでエラーが発生しました。再度お試しください。'))
    default:
      return Promise.reject(new Error('予期せぬエラーが発生しました。'))
  }
}