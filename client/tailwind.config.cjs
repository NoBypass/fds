module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    fontFamily: {
      title: ['GoodTiming', 'sans-serif'],
      default: ['Merriweather', 'sans-serif'],
    },
    extend: {
      skew: {
        '20': '20deg',
      },
      spacing: {
        '128': '32rem',
      },
      width: {
        '128': '32rem',
      }
    }
  },
  plugins: [],
}