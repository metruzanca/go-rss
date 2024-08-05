// This file is needed for the vscode tailwind extension to provide autocomplete
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./lib/views/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [],
};
