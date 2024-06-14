module.exports = {
  content: ['./internal/frontend/**/*.{gohtml, js, go}'],
  darkMode: 'selector',
  theme: {
    extend: {
      gridTemplateRows: {
        'sidebar': 'auto 1fr',
      },
      gridTemplateColumns: {
        'topbar': 'auto 1fr',
      },
      boxShadow: {
        '2xl': '0 4px 25px 5px rgba(255, 255, 255, 1) inset',
      },
      keyframes: {
        'loader-spin': {
          '0%': { transform: 'rotate(0deg)' },
          '100%': { transform: 'rotate(1turn)' },
        },
        bump: {
          '0%, 100%': { transform: 'scale(1)' },
          '50%': { transform: 'scale(0.985)' },
        },
        blink: {
          '0%, 100%': { opacity: '1' },
          '50%': { opacity: '0' },
        },
      },
      animation: {
        bump: 'bump 0.35s ease-in-out',
        blink: 'blink 1s linear infinite',
        'loader-linear-spin': 'loader-spin .8s linear infinite',
        'loader-ease-spin': 'loader-spin .8s ease infinite',
      },
      borderWidth: {
        '3': '3px',
      },
      colors: {
        neutral: {
          750: '#343434',
        }
      },
    },
  },
  plugins: [],
}