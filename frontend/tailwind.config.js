/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      'pinkCustom': '#FC5185',
      'grayCustom': '#FBFBFB'

    },
  },
  plugins: [],
}