'use server'

import { revalidatePath } from 'next/cache'

export async function addNote(formData: FormData) {

  console.log(formData)
  revalidatePath('/about')
}