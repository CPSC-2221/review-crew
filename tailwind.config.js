//@type {import('tailwindcss').Config}

module.exports = {
  content: ["./views/**/*.{go,js,templ,html}"],
  theme: {
    colors: {
      primary: "var(--color-primary)",
      secondary: "var(--color-secondary)",
      background: "var(--color-background)",
      supporting: "var(--color-supporting)",
      tertiary: "var(--color-tertiary)",
      tertiaryhover: "var(--color-tertiary-hover)",
      tertiaryclick: "var(--color-tertiary-click)",
    },
    //fontFamily: {
    //  display: ["Kelsi"],
    //},
    extend: {},
  },
  plugins: [],
};
