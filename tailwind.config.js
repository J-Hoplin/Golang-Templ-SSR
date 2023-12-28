/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "media",
  content: ["templates/**/*"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["light", "dark"],
  },
  plugins: [require("daisyui")],
};
