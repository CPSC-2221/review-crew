/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.{go,js,templ,html}"],
  theme: {
    colors: {
      primary: "rgb(var(--color-primary))",
      secondary: "rgb(var(--color-secondary))",
      background: "rgb(var(--color-background))",
      supporting: "rgb(var(--color-supporting))",
      tertiary: "rgb(var(--color-tertiary))",
    },
    fontFamily: {
      display: ["Kelsi", "display"],
    },
    extend: {},
  },
  plugins: [],
};
