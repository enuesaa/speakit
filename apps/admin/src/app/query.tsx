'use client'
import { ReactNode } from 'react'
import { QueryClientProvider, QueryClient } from '@tanstack/react-query'

const client = new QueryClient()

type Props = {
  children: ReactNode,
}
export const QueryProvider = ({ children }: Props) => {
  return (
    <QueryClientProvider client={client}>
      {children}
    </QueryClientProvider>
  )
}
