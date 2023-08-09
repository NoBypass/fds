/** @types {import('tailwindcss').Config} */
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
      animation: {
        'spinner-linear': 'spinner-spin .8s linear infinite',
        'spinner-ease': 'spinner-spin .8s ease infinite',
      },
      keyframes: {
        'ripple': {
          '0%': { width: '0', height: '0', opacity: '.6' },
          '50%': { height: '200px', width: '200px', opacity: '0' },
        },
        'spinner-spin': {
          '0%': { transform: 'rotate(0deg)' },
          '100%': { transform: 'rotate(360deg)' },
        },
      },
      borderWidth: {
        '3': '3px',
      },
    },
  },
  plugins: [],
}

