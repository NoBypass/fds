import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
      port: 5000
  },
  define: {
    'import.meta.env': {
      VITE_GH_TOKEN: process.env.VITE_GH_TOKEN,
      VITE_API_KEY: process.env.VITE_API_KEY,
    },
  },
})
