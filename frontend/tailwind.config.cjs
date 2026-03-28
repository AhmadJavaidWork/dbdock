/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}',
    './src/components/**/*.{vue,js,ts}' // in case your components are nested
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Montserrat', 'sans-serif'], // default UI font
        mono: ['JetBrains Mono', 'monospace'], // code blocks
        roboto: ['Roboto', 'sans-serif'] // optional secondary
      }
    }
  },
  plugins: []
}
