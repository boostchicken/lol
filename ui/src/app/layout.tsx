import './globals.css'
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  manifest: "manifest.json",
  title: "Boostchicken LoL",
  description: "Admin interface for Boostchicken LoL"
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html data-bs-theme="dark" lang="en">
      <body className={inter.className}>{children}</body>
    </html>
  )
}
