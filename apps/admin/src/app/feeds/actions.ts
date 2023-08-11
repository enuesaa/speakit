'use server'

export type FeedRequestData = {
  name: string
  url: string
}
export async function addFeed(formData: FeedRequestData) {
  console.log(formData)
}
