/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      spacing: {
        '104': '26rem',
        '128': '32rem',
        '144': '36rem',
        '160': '40rem',
      },
      keyframes: {
        'ripple': {
          '0%': { width: '0', height: '0', opacity: '.6' },
          '50%': { height: '200px', width: '200px', opacity: '0' },
        },
      },
    },
  },
  plugins: [],
}

