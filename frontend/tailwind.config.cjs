/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}',
    './src/components/**/*.{vue,js,ts}'
  ],
  darkMode: 'class',
  theme: {
    extend: {
      fontFamily: {
        sans: ['Montserrat', 'sans-serif'],
        mono: ['JetBrains Mono', 'monospace'],
        roboto: ['Roboto', 'sans-serif']
      },
      colors: {
        primary: '#6F5AF8',
        'primary-light': '#E7E3FF',
        'primary-dark-hover': '#7F69FF',
        secondary: '#4BCFFA',
        danger: '#E42933',
        'danger-light': '#FCE7E7',
        'danger-dark-hover': '#D95353',
        'accent-dark': '#4BCFFA',
        surface: {
          light: '#FFFFFF',
          dark: '#2A2B38'
        },
        background: {
          light: '#F8F9FA',
          dark: '#1F2029'
        },
        text: {
          light: '#1F2937',
          dark: '#F5F5F5',
          'secondary-light': '#4B5563',
          'secondary-dark': '#D1D5DB',
          'accent-light': '#6F5AF8',
          'accent-dark': '#4BCFFA'
        }
      },
      borderRadius: {
        lg: '0.75rem'
      },
      boxShadow: {
        card: '0 4px 12px rgba(0,0,0,0.1)'
      }
    }
  },
  plugins: []
}
