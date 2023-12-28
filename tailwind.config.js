/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "class",
  content: ["templates/**/*"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["light", "dark"],
  },
  plugins: [require("daisyui")],
};
