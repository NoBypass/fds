/** @types {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      backgroundImage: (theme) => ({
        'gradient-primary': `linear-gradient(120deg, ${theme('colors.purple.500')}, ${theme('colors.indigo.500')})`,
        'gradient-neutral': `linear-gradient(to bottom, ${theme('colors.gray.800/10%')}, ${theme('colors.gray.950/50%')})`,
      }),
      boxShadow: (theme) => ({
        'inset-primary': `inset 0 0 12px ${theme('colors.purple.400/25%')}`,
        'inset-neutral': `inset 0 0 12px ${theme('colors.gray.400/25%')}`,
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
      },
    },
  },
  plugins: [],
}

