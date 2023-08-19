// see https://github.com/anymaniax/orval/blob/master/samples/react-query/custom-client/src/api/mutator/custom-client.ts

type Args = {
  url: string
  method: 'get' | 'post' | 'put' | 'delete'
  params?: Record<string, any>
  headers?: Record<string, any>
  data?: any
  signal?: any
}
export const useClient = <T>() => {
  return async ({ url, method, params, headers, data }: Args): Promise<T> => {
    const res = await fetch(url + new URLSearchParams(params), {
      method,
      headers,
      ...(data ? { body: JSON.stringify(data) } : {}),
    })

    return res.json()
  }
}
