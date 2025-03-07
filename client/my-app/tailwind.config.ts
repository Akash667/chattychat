import type { Config } from "tailwindcss";

export default {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        'dark-primary': '#131a1c',
        'dark-secondary': '#1b2224',
        red: '#e74c4c',
        green: '#6bb05d',
        blue: '#0183ff',
        grey: '#dddfe2',
        white: '#fff',
      },
    },
  },
  plugins: [],
} satisfies Config;
