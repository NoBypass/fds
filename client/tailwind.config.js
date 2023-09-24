/** @types {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      backgroundImage: (theme) => ({
        'gradient-glass': 'linear-gradient(to top, rgba(255, 255, 255, 0.2), rgba(255, 255, 255, 0.05))',
        'gradient-primary-glass': 'linear-gradient(to top, rgba(168, 85, 247, 0.3), rgba(168, 85, 247, 0.08))',
        'gradient-primary': `linear-gradient(120deg, ${theme('colors.purple.500')}, ${theme('colors.indigo.500')})`,
      }),
      spacing: {
        '104': '26rem',
        '128': '32rem',
        '144': '36rem',
        '160': '40rem',
      },
      animation: {
        'spinner-linear': 'spinner-spin .8s linear infinite',
        'spinner-ease': 'spinner-spin .8s ease infinite',
        'resize': 'resize .2s ease-in-out forwards',
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
        'resize': {
            '0%': { transform: 'scale(1)' },
            '50%': { transform: 'scale(0.98)' },
            '100%': { transform: 'scale(1)' },
        }
      },
      borderWidth: {
        '3': '3px',
      },
      screens: {
        '2xl': '1536px',
        '3xl': '1920px',
      }
    },
  },
  plugins: [],
}

