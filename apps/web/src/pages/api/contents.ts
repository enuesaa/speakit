import type { NextApiRequest, NextApiResponse } from 'next'
import { z } from 'zod'

const validator = z.object({
  name: z.string().min(1),
})
type RequestSchema = z.infer<typeof validator>
type ResponseSchema = {
  message: string
}
 
export default function handler(req: NextApiRequest, res: NextApiResponse<ResponseSchema>) {
  // see https://stackoverflow.com/questions/66739797/how-to-handle-a-post-request-in-next-js
  if (req.method !== 'POST') {
    res.status(404).send({ message: 'Not found.' })
    return
  }
  const result = validator.safeParse(req.body)
  if (result.success) {
    const reqbody = result.data
  } else {
    res.status(401).json({ message: result.error.toString() })
  }

  res.status(200).json({ message: 'a' })
}
