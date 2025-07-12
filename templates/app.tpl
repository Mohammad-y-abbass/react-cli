import React from 'react'
{{if .UseTailwind}}
import './index.css'
{{end}}

function App() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
      <div className="text-center">
        <h1 className="text-4xl font-bold text-gray-900 mb-4">
          Welcome to React
        </h1>
        <p className="text-lg text-gray-600 mb-8">
          Start building something amazing!
        </p>
        <div className="space-x-4">
          <a
            href="https://react.dev"
            target="_blank"
            rel="noopener noreferrer"
            className="inline-block bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded-lg transition-colors duration-200"
          >
            Learn React
          </a>
          <a
            href="https://vitejs.dev"
            target="_blank"
            rel="noopener noreferrer"
            className="inline-block bg-purple-600 hover:bg-purple-700 text-white font-semibold py-2 px-4 rounded-lg transition-colors duration-200"
          >
            Learn Vite
          </a>
        </div>
        {{if .UseTailwind}}
        <div className="mt-8 p-4 bg-white rounded-lg shadow-lg">
          <p className="text-sm text-gray-500">
            🎨 Tailwind CSS is configured and ready to use!
          </p>
        </div>
        {{end}}
        {{if .UseTypeScript}}
        <div className="mt-4 p-4 bg-green-50 rounded-lg">
          <p className="text-sm text-green-700">
            📝 TypeScript is configured and ready to use!
          </p>
        </div>
        {{end}}
      </div>
    </div>
  )
}

export default App
