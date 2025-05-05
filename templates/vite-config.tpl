import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
{{if .UseTailwind}}
import tailwindcss from '@tailwindcss/vite'
{{end}}

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), {{if .UseTailwind}} tailwindcss() {{end}}],
})
